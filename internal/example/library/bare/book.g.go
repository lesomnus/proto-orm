// Code generated by "protoc-gen-orm-ent-grpc", DO NOT EDIT.

package bare

import (
	context "context"
	uuid "github.com/google/uuid"
	library "github.com/lesomnus/proto-orm/internal/example/library"
	ent "github.com/lesomnus/proto-orm/internal/example/library/ent"
	book "github.com/lesomnus/proto-orm/internal/example/library/ent/book"
	predicate "github.com/lesomnus/proto-orm/internal/example/library/ent/predicate"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type BookServiceServer struct {
	Db *ent.Client
	library.UnimplementedBookServiceServer
}

func NewBookServiceServer(db *ent.Client) BookServiceServer {
	return BookServiceServer{Db: db}
}

func (s BookServiceServer) Add(ctx context.Context, req *library.BookAddRequest) (*library.Book, error) {
	q := s.Db.Book.Create()
	if req.HasId() {
		if v, err := uuid.FromBytes(req.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "id: %s", err)
		} else {
			q.SetID(v)
		}
	}
	q.SetTitle(req.GetTitle())
	if v := req.GetIndex(); v != nil {
		q.SetIndex(v)
	}
	if v := req.GetGenres(); v != nil {
		q.SetGenres(v)
	}
	if req.HasDateCreated() {
		q.SetDateCreated(req.GetDateCreated().AsTime())
	}

	v, err := q.Save(ctx)
	if err != nil {
		return nil, StatusFromEntError(err)
	}

	return v.Proto(), nil
}

func (s BookServiceServer) Get(ctx context.Context, req *library.BookGetRequest) (*library.Book, error) {
	q := s.Db.Book.Query()
	if req.HasSelect() {
		BookSelect(q, req.GetSelect())
	} else {
	}

	if p, err := BookPick(req); err != nil {
		return nil, err
	} else {
		q.Where(p)
	}

	v, err := q.Only(ctx)
	if err != nil {
		return nil, StatusFromEntError(err)
	}

	return v.Proto(), nil
}

func (s BookServiceServer) Patch(ctx context.Context, req *library.BookPatchRequest) (*emptypb.Empty, error) {
	id, err := BookGetId(ctx, s.Db, req.GetKey())
	if err != nil {
		return nil, err
	}

	q := s.Db.Book.UpdateOneID(id)
	if req.HasTitle() {
		q.SetTitle(req.GetTitle())
	}
	if v := req.GetIndex(); v != nil {
		q.SetIndex(v)
	}
	if v := req.GetGenres(); v != nil {
		q.SetGenres(v)
	}

	if _, err := q.Save(ctx); err != nil {
		return nil, StatusFromEntError(err)
	}

	return nil, nil
}

func (s BookServiceServer) Erase(ctx context.Context, req *library.BookGetRequest) (*emptypb.Empty, error) {
	p, err := BookPick(req)
	if err != nil {
		return nil, err
	}
	if _, err := s.Db.Book.Delete().Where(p).Exec(ctx); err != nil {
		return nil, StatusFromEntError(err)
	}

	return nil, nil
}

func BookPick(req *library.BookGetRequest) (predicate.Book, error) {
	switch req.WhichKey() {
	case library.BookGetRequest_Id_case:
		if v, err := uuid.FromBytes(req.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "id: %s", "err")
		} else {
			return book.IDEQ(v), nil
		}
	case library.BookGetRequest_Key_not_set_case:
		return nil, status.Errorf(codes.InvalidArgument, "key not provided")
	default:
		return nil, status.Errorf(codes.Unimplemented, "unknown type of key")
	}
}

func BookGetId(ctx context.Context, db *ent.Client, req *library.BookGetRequest) (uuid.UUID, error) {
	var z uuid.UUID
	if req.HasId() {
		if v, err := uuid.FromBytes(req.GetId()); err != nil {
			return z, status.Errorf(codes.InvalidArgument, "key.id: %s", err)
		} else {
			return v, nil
		}
	}

	p, err := BookPick(req)
	if err != nil {
		return z, err
	}

	v, err := db.Book.Query().Where(p).OnlyID(ctx)
	if err != nil {
		return z, status.Errorf(codes.Internal, "query: %s", err)
	}

	return v, nil
}

func BookSelectedFields(m *library.BookSelect) []string {
	if m.GetAll() {
		return book.Columns
	}

	vs := []string{book.FieldID}
	if m.GetTitle() {
		vs = append(vs, book.FieldTitle)
	}
	if m.GetIndex() {
		vs = append(vs, book.FieldIndex)
	}
	if m.GetGenres() {
		vs = append(vs, book.FieldGenres)
	}
	if m.GetDateCreated() {
		vs = append(vs, book.FieldDateCreated)
	}

	return vs
}

func BookSelect(q *ent.BookQuery, m *library.BookSelect) {
	if !m.GetAll() {
		fields := BookSelectedFields(m)
		q.Select(fields...)
	}
	if m.HasAuthors() {
		q.WithAuthors(func(q *ent.AuthorQuery) {
			AuthorSelect(q, m.GetAuthors())
		})
	}
}
