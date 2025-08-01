// Code generated by "protoc-gen-orm-ent-grpc", DO NOT EDIT.

package bare

import (
	context "context"
	library "github.com/lesomnus/proto-orm/internal/example/library"
	ent "github.com/lesomnus/proto-orm/internal/example/library/ent"
	foovd "github.com/lesomnus/proto-orm/internal/example/library/ent/foovd"
	predicate "github.com/lesomnus/proto-orm/internal/example/library/ent/predicate"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type FooVdServiceServer struct {
	Db *ent.Client
	library.UnimplementedFooVdServiceServer
}

func NewFooVdServiceServer(db *ent.Client) FooVdServiceServer {
	return FooVdServiceServer{Db: db}
}

func (s FooVdServiceServer) Add(ctx context.Context, req *library.FooVdAddRequest) (*library.FooVd, error) {
	q := s.Db.FooVd.Create()
	q.SetID(req.GetId())
	if req.HasVdDouble() {
		q.SetVdDouble(req.GetVdDouble())
	}
	if req.HasVdInt64() {
		q.SetVdInt64(req.GetVdInt64())
	}
	if req.HasVdUint64() {
		q.SetVdUint64(req.GetVdUint64())
	}
	if req.HasVdBool() {
		q.SetVdBool(req.GetVdBool())
	}
	if req.HasVdString() {
		q.SetVdString(req.GetVdString())
	}
	if req.HasVdBytes() {
		q.SetVdBytes(req.GetVdBytes())
	}

	v, err := q.Save(ctx)
	if err != nil {
		return nil, StatusFromEntError(err)
	}

	return v.Proto(), nil
}

func (s FooVdServiceServer) Get(ctx context.Context, req *library.FooVdGetRequest) (*library.FooVd, error) {
	q := s.Db.FooVd.Query()
	if req.HasSelect() {
		FooVdSelect(q, req.GetSelect())
	} else {
		FooVdSelectEdges(q)
	}

	if p, err := FooVdPick(req); err != nil {
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

func (s FooVdServiceServer) Patch(ctx context.Context, req *library.FooVdPatchRequest) (*emptypb.Empty, error) {
	id, err := FooVdGetId(ctx, s.Db, req.GetKey())
	if err != nil {
		return nil, err
	}

	q := s.Db.FooVd.UpdateOneID(id)
	if req.HasVdDouble() {
		q.SetVdDouble(req.GetVdDouble())
	}
	if req.HasVdInt64() {
		q.SetVdInt64(req.GetVdInt64())
	}
	if req.HasVdUint64() {
		q.SetVdUint64(req.GetVdUint64())
	}
	if req.HasVdBool() {
		q.SetVdBool(req.GetVdBool())
	}
	if req.HasVdString() {
		q.SetVdString(req.GetVdString())
	}
	if req.HasVdBytes() {
		q.SetVdBytes(req.GetVdBytes())
	}

	if _, err := q.Save(ctx); err != nil {
		return nil, StatusFromEntError(err)
	}

	return nil, nil
}

func (s FooVdServiceServer) Erase(ctx context.Context, req *library.FooVdGetRequest) (*emptypb.Empty, error) {
	p, err := FooVdPick(req)
	if err != nil {
		return nil, err
	}
	if _, err := s.Db.FooVd.Delete().Where(p).Exec(ctx); err != nil {
		return nil, StatusFromEntError(err)
	}

	return nil, nil
}

func FooVdPick(req *library.FooVdGetRequest) (predicate.FooVd, error) {
	switch req.WhichKey() {
	case library.FooVdGetRequest_Id_case:
		return foovd.IDEQ(req.GetId()), nil
	case library.FooVdGetRequest_Key_not_set_case:
		return nil, status.Errorf(codes.InvalidArgument, "key not provided")
	default:
		return nil, status.Errorf(codes.Unimplemented, "unknown type of key")
	}
}

func FooVdGetId(ctx context.Context, db *ent.Client, req *library.FooVdGetRequest) (int64, error) {
	var z int64
	if req.HasId() {
		return req.GetId(), nil
	}

	p, err := FooVdPick(req)
	if err != nil {
		return z, err
	}

	v, err := db.FooVd.Query().Where(p).OnlyID(ctx)
	if err != nil {
		return z, status.Errorf(codes.Internal, "query: %s", err)
	}

	return v, nil
}

func FooVdSelectedFields(m *library.FooVdSelect) []string {
	if m.GetAll() {
		return foovd.Columns
	}

	vs := []string{foovd.FieldID}
	if m.GetVdDouble() {
		vs = append(vs, foovd.FieldVdDouble)
	}
	if m.GetVdInt64() {
		vs = append(vs, foovd.FieldVdInt64)
	}
	if m.GetVdUint64() {
		vs = append(vs, foovd.FieldVdUint64)
	}
	if m.GetVdBool() {
		vs = append(vs, foovd.FieldVdBool)
	}
	if m.GetVdString() {
		vs = append(vs, foovd.FieldVdString)
	}
	if m.GetVdBytes() {
		vs = append(vs, foovd.FieldVdBytes)
	}

	return vs
}

func FooVdSelectEdges(q *ent.FooVdQuery) {
}

func FooVdSelect(q *ent.FooVdQuery, m *library.FooVdSelect) {
	if !m.GetAll() {
		fields := FooVdSelectedFields(m)
		q.Select(fields...)
	}
}
