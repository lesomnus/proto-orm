// Code generated by "protoc-gen-orm-ent-grpc". DO NOT EDIT

package bare

import (
	context "context"
	library "github.com/lesomnus/proto-orm/internal/example/library"
	ent "github.com/lesomnus/proto-orm/internal/example/library/ent"
	foovo "github.com/lesomnus/proto-orm/internal/example/library/ent/foovo"
	predicate "github.com/lesomnus/proto-orm/internal/example/library/ent/predicate"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type FooVoServiceServer struct {
	db *ent.Client
	library.UnimplementedFooVoServiceServer
}

func NewFooVoServiceServer(db *ent.Client) *FooVoServiceServer {
	return &FooVoServiceServer{db: db}
}

func (s *FooVoServiceServer) Add(ctx context.Context, req *library.FooVoAddRequest) (*library.FooVo, error) {
	q := s.db.FooVo.Create()
	q.SetID(req.GetId())
	q.SetVoDouble(req.GetVoDouble())
	q.SetVoInt64(req.GetVoInt64())
	q.SetVoUint64(req.GetVoUint64())
	q.SetVoBool(req.GetVoBool())
	q.SetVoString(req.GetVoString())
	q.SetVoBytes(req.GetVoBytes())

	v, err := q.Save(ctx)
	if err != nil {
		return nil, StatusFromEntError(err)
	}

	return v.Proto(), nil
}

func (s *FooVoServiceServer) Get(ctx context.Context, req *library.FooVoGetRequest) (*library.FooVo, error) {
	q := s.db.FooVo.Query()
	if p, err := FooVoPick(req); err != nil {
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

func (s *FooVoServiceServer) Patch(ctx context.Context, req *library.FooVoPatchRequest) (*emptypb.Empty, error) {
	id, err := FooVoGetId(ctx, s.db, req.GetKey())
	if err != nil {
		return nil, err
	}

	q := s.db.FooVo.UpdateOneID(id)
	if req.HasVoDouble() {
		q.SetVoDouble(req.GetVoDouble())
	}
	if req.HasVoInt64() {
		q.SetVoInt64(req.GetVoInt64())
	}
	if req.HasVoUint64() {
		q.SetVoUint64(req.GetVoUint64())
	}
	if req.HasVoBool() {
		q.SetVoBool(req.GetVoBool())
	}
	if req.HasVoString() {
		q.SetVoString(req.GetVoString())
	}
	if req.HasVoBytes() {
		q.SetVoBytes(req.GetVoBytes())
	}

	if _, err := q.Save(ctx); err != nil {
		return nil, StatusFromEntError(err)
	}

	return nil, nil
}

func (s *FooVoServiceServer) Erase(ctx context.Context, req *library.FooVoGetRequest) (*emptypb.Empty, error) {
	p, err := FooVoPick(req)
	if err != nil {
		return nil, err
	}
	if _, err := s.db.FooVo.Delete().Where(p).Exec(ctx); err != nil {
		return nil, StatusFromEntError(err)
	}

	return nil, nil
}

func FooVoPick(req *library.FooVoGetRequest) (predicate.FooVo, error) {
	switch req.WhichKey() {
	case library.FooVoGetRequest_Id_case:
		return foovo.IDEQ(req.GetId()), nil
	case library.FooVoGetRequest_Key_not_set_case:
		return nil, status.Errorf(codes.InvalidArgument, "key not provided")
	default:
		return nil, status.Errorf(codes.Unimplemented, "unknown type of key")
	}
}

func FooVoGetId(ctx context.Context, db *ent.Client, req *library.FooVoGetRequest) (int64, error) {
	var z int64
	if req.HasId() {
		return req.GetId(), nil
	}

	p, err := FooVoPick(req)
	if err != nil {
		return z, err
	}

	v, err := db.FooVo.Query().Where(p).OnlyID(ctx)
	if err != nil {
		return z, status.Errorf(codes.Internal, "query: %s", err)
	}

	return v, nil
}
