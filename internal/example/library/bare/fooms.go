// Code generated by "protoc-gen-orm-ent-grpc". DO NOT EDIT

package bare

import (
	context "context"
	library "github.com/lesomnus/proto-orm/internal/example/library"
	ent "github.com/lesomnus/proto-orm/internal/example/library/ent"
	fooms "github.com/lesomnus/proto-orm/internal/example/library/ent/fooms"
	predicate "github.com/lesomnus/proto-orm/internal/example/library/ent/predicate"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type FooMsServiceServer struct {
	db *ent.Client
	library.UnimplementedFooMsServiceServer
}

func NewFooMsServiceServer(db *ent.Client) *FooMsServiceServer {
	return &FooMsServiceServer{db: db}
}

func (s *FooMsServiceServer) Add(ctx context.Context, req *library.FooMsAddRequest) (*library.FooMs, error) {
	q := s.db.FooMs.Create()
	q.SetID(req.GetId())
	if v := req.GetVmsDouble(); v != nil {
		q.SetVmsDouble(v)
	}
	if v := req.GetVmsInt64(); v != nil {
		q.SetVmsInt64(v)
	}
	if v := req.GetVmsUint64(); v != nil {
		q.SetVmsUint64(v)
	}
	if v := req.GetVmsBool(); v != nil {
		q.SetVmsBool(v)
	}
	if v := req.GetVmsString(); v != nil {
		q.SetVmsString(v)
	}
	if v := req.GetVmsBytes(); v != nil {
		q.SetVmsBytes(v)
	}

	v, err := q.Save(ctx)
	if err != nil {
		return nil, StatusFromEntError(err)
	}

	return v.Proto(), nil
}

func (s *FooMsServiceServer) Get(ctx context.Context, req *library.FooMsGetRequest) (*library.FooMs, error) {
	q := s.db.FooMs.Query()
	if p, err := FooMsPick(req); err != nil {
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

func (s *FooMsServiceServer) Patch(ctx context.Context, req *library.FooMsPatchRequest) (*emptypb.Empty, error) {
	id, err := FooMsGetId(ctx, s.db, req.GetKey())
	if err != nil {
		return nil, err
	}

	q := s.db.FooMs.UpdateOneID(id)
	if v := req.GetVmsDouble(); v != nil {
		q.SetVmsDouble(v)
	}
	if v := req.GetVmsInt64(); v != nil {
		q.SetVmsInt64(v)
	}
	if v := req.GetVmsUint64(); v != nil {
		q.SetVmsUint64(v)
	}
	if v := req.GetVmsBool(); v != nil {
		q.SetVmsBool(v)
	}
	if v := req.GetVmsString(); v != nil {
		q.SetVmsString(v)
	}
	if v := req.GetVmsBytes(); v != nil {
		q.SetVmsBytes(v)
	}

	if _, err := q.Save(ctx); err != nil {
		return nil, StatusFromEntError(err)
	}

	return nil, nil
}

func (s *FooMsServiceServer) Erase(ctx context.Context, req *library.FooMsGetRequest) (*emptypb.Empty, error) {
	p, err := FooMsPick(req)
	if err != nil {
		return nil, err
	}
	if _, err := s.db.FooMs.Delete().Where(p).Exec(ctx); err != nil {
		return nil, StatusFromEntError(err)
	}

	return nil, nil
}

func FooMsPick(req *library.FooMsGetRequest) (predicate.FooMs, error) {
	switch req.WhichKey() {
	case library.FooMsGetRequest_Id_case:
		return fooms.IDEQ(req.GetId()), nil
	case library.FooMsGetRequest_Key_not_set_case:
		return nil, status.Errorf(codes.InvalidArgument, "key not provided")
	default:
		return nil, status.Errorf(codes.Unimplemented, "unknown type of key")
	}
}

func FooMsGetId(ctx context.Context, db *ent.Client, req *library.FooMsGetRequest) (int64, error) {
	var z int64
	if req.HasId() {
		return req.GetId(), nil
	}

	p, err := FooMsPick(req)
	if err != nil {
		return z, err
	}

	v, err := db.FooMs.Query().Where(p).OnlyID(ctx)
	if err != nil {
		return z, status.Errorf(codes.Internal, "query: %s", err)
	}

	return v, nil
}