package testgrpc_test

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"testing"

	"entgo.io/ent/dialect/sql"
	"github.com/lesomnus/proto-orm/internal/example/library"
	"github.com/lesomnus/proto-orm/internal/example/library/bare"
	"github.com/lesomnus/proto-orm/internal/example/library/ent"
	"github.com/lesomnus/proto-orm/internal/example/library/ent/locker"
	"github.com/lesomnus/proto-orm/internal/example/library/ent/member"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

type Server struct {
	t  *testing.T
	Db *ent.Client

	Author bare.AuthorServiceServer
	Book   bare.BookServiceServer
	Member bare.MemberServiceServer
	Locker library.LockerServiceServer
}

func NewServer(t *testing.T) *Server {
	ctx := context.TODO()

	driver, err := sql.Open("sqlite3", ":memory:?cache=shared&_fk=1")
	require.NoError(t, err)

	driver.DB().SetMaxOpenConns(1)

	db := ent.NewClient(ent.Driver(driver))
	err = db.Schema.Create(ctx)
	require.NoError(t, err)

	return &Server{
		t:  t,
		Db: ent.NewClient(ent.Driver(driver)),

		Author: bare.NewAuthorServiceServer(db),
		Book:   bare.NewBookServiceServer(db),
		Member: bare.NewMemberServiceServer(db),
		Locker: bare.NewLockerServiceServer(db),
	}
}

func (s *Server) Close() error {
	return s.Db.Close()
}

type Client struct {
	Server   *Server
	Listener *bufconn.Listener
	Conn     grpc.ClientConnInterface

	wg sync.WaitGroup

	grpc_server *grpc.Server

	Author library.AuthorServiceClient
	Book   library.BookServiceClient
	Member library.MemberServiceClient
	Locker library.LockerServiceClient

	v_author *library.Author
	v_book   *library.Book
	v_member *library.Member
	v_locker *library.Locker
}

func NewClient(s *Server) *Client {
	ctx := context.TODO()
	db := s.Db

	v_author, err := db.Author.Create().
		SetAlias("dostoevsky").
		SetName("Fyodor Dostoevsky").
		Save(ctx)
	require.NoError(s.t, err)

	v_book, err := db.Book.Create().
		SetTitle("Crime and Punishment").
		AddAuthors(v_author).
		Save(ctx)
	require.NoError(s.t, err)

	v_member, err := db.Member.Create().
		SetName("lesomnus").
		SetLabels(map[string]string{
			"foo": "bar",
		}).
		SetProfile(library.Profile_builder{
			Email: "lesomnus@gmail.com",
		}.Build()).
		SetLevel(int(library.Member_LEVEL_BRONZE)).
		Save(ctx)
	require.NoError(s.t, err)

	v_locker, err := db.Locker.Create().
		SetOwner(v_member).
		SetName("Davy Jones's").
		SetNumber(42).
		Save(ctx)
	require.NoError(s.t, err)

	v_locker, err = db.Locker.Query().
		Where(locker.ID(v_locker.ID)).
		WithOwner().
		Only(ctx)
	require.NoError(s.t, err)

	v_author, err = db.Author.Get(ctx, v_author.ID)
	require.NoError(s.t, err)

	v_member, err = db.Member.Query().
		Where(member.ID(v_member.ID)).
		WithLocker().
		Only(ctx)
	require.NoError(s.t, err)

	listener := bufconn.Listen(1 << 20)
	conn, err := grpc.NewClient("passthrough://bufnet",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(ctx context.Context, _ string) (net.Conn, error) {
			return listener.DialContext(ctx)
		}),
	)
	require.NoError(s.t, err)

	grpc_server := grpc.NewServer()
	library.RegisterAuthorServiceServer(grpc_server, s.Author)
	library.RegisterBookServiceServer(grpc_server, s.Book)
	library.RegisterMemberServiceServer(grpc_server, s.Member)
	library.RegisterLockerServiceServer(grpc_server, s.Locker)

	v := &Client{
		Server:   s,
		Listener: listener,
		Conn:     conn,

		grpc_server: grpc_server,

		Author: library.NewAuthorServiceClient(conn),
		Book:   library.NewBookServiceClient(conn),
		Member: library.NewMemberServiceClient(conn),
		Locker: library.NewLockerServiceClient(conn),

		v_author: v_author.Proto(),
		v_book:   v_book.Proto(),
		v_member: v_member.Proto(),
		v_locker: v_locker.Proto(),
	}

	v.wg.Add(1)
	go func() {
		defer v.wg.Done()
		if err := v.grpc_server.Serve(listener); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			require.Fail(s.t, fmt.Sprintf("server failed: %s", err.Error()))
		}
	}()

	return v
}

func (c *Client) Close() error {
	c.grpc_server.GracefulStop()
	c.wg.Wait()
	return nil
}

func T(run func(ctx context.Context, x *require.Assertions, c *Client)) func(t *testing.T) {
	return func(t *testing.T) {
		ctx := context.TODO()

		s := NewServer(t)
		defer s.Close()

		c := NewClient(s)
		defer c.Close()

		x := require.New(t)
		run(ctx, x, c)
	}
}
