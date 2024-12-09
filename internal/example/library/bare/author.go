// Code generated by "protoc-gen-orm-ent-grpc". DO NOT EDIT

package bare

import (
	context "context"
	uuid "github.com/google/uuid"
	library "github.com/lesomnus/proto-orm/internal/example/library"
	ent "github.com/lesomnus/proto-orm/internal/example/library/ent"
	author "github.com/lesomnus/proto-orm/internal/example/library/ent/author"
	predicate "github.com/lesomnus/proto-orm/internal/example/library/ent/predicate"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type AuthorServiceServer struct {
	db *ent.Client
	library.UnimplementedAuthorServiceServer
}

func NewAuthorServiceServer(db *ent.Client) *AuthorServiceServer {
	return &AuthorServiceServer{db: db}
}

func (s *AuthorServiceServer) Add(ctx context.Context, req *library.AuthorAddRequest) (*library.Author, error) {
	q := s.db.Author.Create()
	if v, err := uuid.FromBytes(req.Id); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id: %s", err)
	} else {
		q.SetID(v)
	}
	q.SetAlias(req.Alias)
	q.SetName(req.Name)
	q.SetDateCreated(req.DateCreated.AsTime())

	v, err := q.Save(ctx)
	if err != nil {
		return nil, StatusFromEntError(err)
	}

	return v.Proto(), nil
}

func (s *AuthorServiceServer) Get(ctx context.Context, req *library.AuthorGetRequest) (*library.Author, error) {
	q := s.db.Author.Query()
	if p, err := AuthorPick(req); err != nil {
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

func (s *AuthorServiceServer) Patch(ctx context.Context, req *library.AuthorPatchRequest) (*library.Author, error) {
	id, err := AuthorGetId(ctx, s.db, req.GetKey())
	if err != nil {
		return nil, err
	}

	q := s.db.Author.UpdateOneID(id)
	if req.Alias != nil {
		q.SetAlias(*req.Alias)
	}
	if req.Name != nil {
		q.SetName(*req.Name)
	}

	v, err := q.Save(ctx)
	if err != nil {
		return nil, StatusFromEntError(err)
	}

	return v.Proto(), nil
}

func (s *AuthorServiceServer) Erase(ctx context.Context, req *library.AuthorGetRequest) (*library.Author, error) {
	p, err := AuthorPick(req)
	if err != nil {
		return nil, err
	}
	if _, err := s.db.Author.Delete().Where(p).Exec(ctx); err != nil {
		return nil, StatusFromEntError(err)
	}

	return nil, nil
}

func AuthorPick(req *library.AuthorGetRequest) (predicate.Author, error) {
	switch k := req.GetKey().(type) {
	case *library.AuthorGetRequest_Id:
		if v, err := uuid.FromBytes(k.Id); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "id: %s", "err")
		} else {
			return author.IDEQ(v), nil
		}
	case *library.AuthorGetRequest_Alias:
		return author.AliasEQ(k.Alias), nil
	case nil:
		return nil, status.Errorf(codes.InvalidArgument, "key not provided")
	default:
		return nil, status.Errorf(codes.Unimplemented, "unknown type of key")
	}
}

func AuthorGetId(ctx context.Context, db *ent.Client, req *library.AuthorGetRequest) (uuid.UUID, error) {
	var z uuid.UUID
	k := req.GetKey()
	if r, ok := k.(*library.AuthorGetRequest_Id); ok {
		if v, err := uuid.FromBytes(r.Id); err != nil {
			return z, status.Errorf(codes.InvalidArgument, "Id: %s", "err")
		} else {
			return v, nil
		}
	}

	p, err := AuthorPick(req)
	if err != nil {
		return z, err
	}

	v, err := db.Author.Query().Where(p).OnlyID(ctx)
	if err != nil {
		return z, status.Errorf(codes.Internal, "query: %s", err)
	}

	return v, nil
}
