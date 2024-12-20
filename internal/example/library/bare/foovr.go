// Code generated by "protoc-gen-orm-ent-grpc". DO NOT EDIT

package bare

import (
	context "context"
	library "github.com/lesomnus/proto-orm/internal/example/library"
	ent "github.com/lesomnus/proto-orm/internal/example/library/ent"
	foovr "github.com/lesomnus/proto-orm/internal/example/library/ent/foovr"
	predicate "github.com/lesomnus/proto-orm/internal/example/library/ent/predicate"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type FooVrServiceServer struct {
	db *ent.Client
	library.UnimplementedFooVrServiceServer
}

func NewFooVrServiceServer(db *ent.Client) *FooVrServiceServer {
	return &FooVrServiceServer{db: db}
}

func (s *FooVrServiceServer) Add(ctx context.Context, req *library.FooVrAddRequest) (*library.FooVr, error) {
	q := s.db.FooVr.Create()
	q.SetID(req.GetId())
	if v := req.GetVrDouble(); v != nil {
		q.SetVrDouble(v)
	}
	if v := req.GetVrInt64(); v != nil {
		q.SetVrInt64(v)
	}
	if v := req.GetVrUint64(); v != nil {
		q.SetVrUint64(v)
	}
	if v := req.GetVrBool(); v != nil {
		q.SetVrBool(v)
	}
	if v := req.GetVrString(); v != nil {
		q.SetVrString(v)
	}
	if v := req.GetVrBytes(); v != nil {
		q.SetVrBytes(v)
	}

	v, err := q.Save(ctx)
	if err != nil {
		return nil, StatusFromEntError(err)
	}

	return v.Proto(), nil
}

func (s *FooVrServiceServer) Get(ctx context.Context, req *library.FooVrGetRequest) (*library.FooVr, error) {
	q := s.db.FooVr.Query()
	if p, err := FooVrPick(req); err != nil {
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

func (s *FooVrServiceServer) Patch(ctx context.Context, req *library.FooVrPatchRequest) (*emptypb.Empty, error) {
	id, err := FooVrGetId(ctx, s.db, req.GetKey())
	if err != nil {
		return nil, err
	}

	q := s.db.FooVr.UpdateOneID(id)
	if v := req.GetVrDouble(); v != nil {
		q.SetVrDouble(v)
	}
	if v := req.GetVrInt64(); v != nil {
		q.SetVrInt64(v)
	}
	if v := req.GetVrUint64(); v != nil {
		q.SetVrUint64(v)
	}
	if v := req.GetVrBool(); v != nil {
		q.SetVrBool(v)
	}
	if v := req.GetVrString(); v != nil {
		q.SetVrString(v)
	}
	if v := req.GetVrBytes(); v != nil {
		q.SetVrBytes(v)
	}

	if _, err := q.Save(ctx); err != nil {
		return nil, StatusFromEntError(err)
	}

	return nil, nil
}

func (s *FooVrServiceServer) Erase(ctx context.Context, req *library.FooVrGetRequest) (*emptypb.Empty, error) {
	p, err := FooVrPick(req)
	if err != nil {
		return nil, err
	}
	if _, err := s.db.FooVr.Delete().Where(p).Exec(ctx); err != nil {
		return nil, StatusFromEntError(err)
	}

	return nil, nil
}

func FooVrPick(req *library.FooVrGetRequest) (predicate.FooVr, error) {
	switch req.WhichKey() {
	case library.FooVrGetRequest_Id_case:
		return foovr.IDEQ(req.GetId()), nil
	case library.FooVrGetRequest_Key_not_set_case:
		return nil, status.Errorf(codes.InvalidArgument, "key not provided")
	default:
		return nil, status.Errorf(codes.Unimplemented, "unknown type of key")
	}
}

func FooVrGetId(ctx context.Context, db *ent.Client, req *library.FooVrGetRequest) (int64, error) {
	var z int64
	if req.HasId() {
		return req.GetId(), nil
	}

	p, err := FooVrPick(req)
	if err != nil {
		return z, err
	}

	v, err := db.FooVr.Query().Where(p).OnlyID(ctx)
	if err != nil {
		return z, status.Errorf(codes.Internal, "query: %s", err)
	}

	return v, nil
}