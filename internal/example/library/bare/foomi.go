// Code generated by "protoc-gen-orm-ent-grpc". DO NOT EDIT

package bare

import (
	context "context"
	library "github.com/lesomnus/proto-orm/internal/example/library"
	ent "github.com/lesomnus/proto-orm/internal/example/library/ent"
	foomi "github.com/lesomnus/proto-orm/internal/example/library/ent/foomi"
	predicate "github.com/lesomnus/proto-orm/internal/example/library/ent/predicate"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type FooMiServiceServer struct {
	db *ent.Client
	library.UnimplementedFooMiServiceServer
}

func NewFooMiServiceServer(db *ent.Client) *FooMiServiceServer {
	return &FooMiServiceServer{db: db}
}

func (s *FooMiServiceServer) Add(ctx context.Context, req *library.FooMiAddRequest) (*library.FooMi, error) {
	q := s.db.FooMi.Create()
	q.SetID(req.GetId())
	if v := req.GetVmiDouble(); v != nil {
		q.SetVmiDouble(v)
	}
	if v := req.GetVmiInt64(); v != nil {
		q.SetVmiInt64(v)
	}
	if v := req.GetVmiUint64(); v != nil {
		q.SetVmiUint64(v)
	}
	if v := req.GetVmiBool(); v != nil {
		q.SetVmiBool(v)
	}
	if v := req.GetVmiString(); v != nil {
		q.SetVmiString(v)
	}
	if v := req.GetVmiBytes(); v != nil {
		q.SetVmiBytes(v)
	}

	v, err := q.Save(ctx)
	if err != nil {
		return nil, StatusFromEntError(err)
	}

	return v.Proto(), nil
}

func (s *FooMiServiceServer) Get(ctx context.Context, req *library.FooMiGetRequest) (*library.FooMi, error) {
	q := s.db.FooMi.Query()
	if p, err := FooMiPick(req); err != nil {
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

func (s *FooMiServiceServer) Patch(ctx context.Context, req *library.FooMiPatchRequest) (*emptypb.Empty, error) {
	id, err := FooMiGetId(ctx, s.db, req.GetKey())
	if err != nil {
		return nil, err
	}

	q := s.db.FooMi.UpdateOneID(id)
	if v := req.GetVmiDouble(); v != nil {
		q.SetVmiDouble(v)
	}
	if v := req.GetVmiInt64(); v != nil {
		q.SetVmiInt64(v)
	}
	if v := req.GetVmiUint64(); v != nil {
		q.SetVmiUint64(v)
	}
	if v := req.GetVmiBool(); v != nil {
		q.SetVmiBool(v)
	}
	if v := req.GetVmiString(); v != nil {
		q.SetVmiString(v)
	}
	if v := req.GetVmiBytes(); v != nil {
		q.SetVmiBytes(v)
	}

	if _, err := q.Save(ctx); err != nil {
		return nil, StatusFromEntError(err)
	}

	return nil, nil
}

func (s *FooMiServiceServer) Erase(ctx context.Context, req *library.FooMiGetRequest) (*emptypb.Empty, error) {
	p, err := FooMiPick(req)
	if err != nil {
		return nil, err
	}
	if _, err := s.db.FooMi.Delete().Where(p).Exec(ctx); err != nil {
		return nil, StatusFromEntError(err)
	}

	return nil, nil
}

func FooMiPick(req *library.FooMiGetRequest) (predicate.FooMi, error) {
	switch req.WhichKey() {
	case library.FooMiGetRequest_Id_case:
		return foomi.IDEQ(req.GetId()), nil
	case library.FooMiGetRequest_Key_not_set_case:
		return nil, status.Errorf(codes.InvalidArgument, "key not provided")
	default:
		return nil, status.Errorf(codes.Unimplemented, "unknown type of key")
	}
}

func FooMiGetId(ctx context.Context, db *ent.Client, req *library.FooMiGetRequest) (int64, error) {
	var z int64
	if req.HasId() {
		return req.GetId(), nil
	}

	p, err := FooMiPick(req)
	if err != nil {
		return z, err
	}

	v, err := db.FooMi.Query().Where(p).OnlyID(ctx)
	if err != nil {
		return z, status.Errorf(codes.Internal, "query: %s", err)
	}

	return v, nil
}
