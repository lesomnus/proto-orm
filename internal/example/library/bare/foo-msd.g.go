// Code generated by "protoc-gen-orm-ent-grpc", DO NOT EDIT.

package bare

import (
	context "context"
	library "github.com/lesomnus/proto-orm/internal/example/library"
	ent "github.com/lesomnus/proto-orm/internal/example/library/ent"
	foomsd "github.com/lesomnus/proto-orm/internal/example/library/ent/foomsd"
	predicate "github.com/lesomnus/proto-orm/internal/example/library/ent/predicate"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type FooMsdServiceServer struct {
	Db *ent.Client
	library.UnimplementedFooMsdServiceServer
}

func NewFooMsdServiceServer(db *ent.Client) FooMsdServiceServer {
	return FooMsdServiceServer{Db: db}
}

func (s FooMsdServiceServer) Add(ctx context.Context, req *library.FooMsdAddRequest) (*library.FooMsd, error) {
	q := s.Db.FooMsd.Create()
	q.SetID(req.GetId())
	if v := req.GetMsdDouble(); v != nil {
		q.SetMsdDouble(v)
	}
	if v := req.GetMsdInt64(); v != nil {
		q.SetMsdInt64(v)
	}
	if v := req.GetMsdUint64(); v != nil {
		q.SetMsdUint64(v)
	}
	if v := req.GetMsdBool(); v != nil {
		q.SetMsdBool(v)
	}
	if v := req.GetMsdString(); v != nil {
		q.SetMsdString(v)
	}
	if v := req.GetMsdBytes(); v != nil {
		q.SetMsdBytes(v)
	}

	v, err := q.Save(ctx)
	if err != nil {
		return nil, StatusFromEntError(err)
	}

	return v.Proto(), nil
}

func (s FooMsdServiceServer) Get(ctx context.Context, req *library.FooMsdGetRequest) (*library.FooMsd, error) {
	q := s.Db.FooMsd.Query()
	if req.HasSelect() {
		FooMsdSelect(q, req.GetSelect())
	} else {
		FooMsdSelectEdges(q)
	}

	if p, err := FooMsdPick(req); err != nil {
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

func (s FooMsdServiceServer) Patch(ctx context.Context, req *library.FooMsdPatchRequest) (*emptypb.Empty, error) {
	id, err := FooMsdGetId(ctx, s.Db, req.GetKey())
	if err != nil {
		return nil, err
	}

	q := s.Db.FooMsd.UpdateOneID(id)
	if v := req.GetMsdDouble(); v != nil {
		q.SetMsdDouble(v)
	}
	if v := req.GetMsdInt64(); v != nil {
		q.SetMsdInt64(v)
	}
	if v := req.GetMsdUint64(); v != nil {
		q.SetMsdUint64(v)
	}
	if v := req.GetMsdBool(); v != nil {
		q.SetMsdBool(v)
	}
	if v := req.GetMsdString(); v != nil {
		q.SetMsdString(v)
	}
	if v := req.GetMsdBytes(); v != nil {
		q.SetMsdBytes(v)
	}

	if _, err := q.Save(ctx); err != nil {
		return nil, StatusFromEntError(err)
	}

	return nil, nil
}

func (s FooMsdServiceServer) Erase(ctx context.Context, req *library.FooMsdGetRequest) (*emptypb.Empty, error) {
	p, err := FooMsdPick(req)
	if err != nil {
		return nil, err
	}
	if _, err := s.Db.FooMsd.Delete().Where(p).Exec(ctx); err != nil {
		return nil, StatusFromEntError(err)
	}

	return nil, nil
}

func FooMsdPick(req *library.FooMsdGetRequest) (predicate.FooMsd, error) {
	switch req.WhichKey() {
	case library.FooMsdGetRequest_Id_case:
		return foomsd.IDEQ(req.GetId()), nil
	case library.FooMsdGetRequest_Key_not_set_case:
		return nil, status.Errorf(codes.InvalidArgument, "key not provided")
	default:
		return nil, status.Errorf(codes.Unimplemented, "unknown type of key")
	}
}

func FooMsdGetId(ctx context.Context, db *ent.Client, req *library.FooMsdGetRequest) (int64, error) {
	var z int64
	if req.HasId() {
		return req.GetId(), nil
	}

	p, err := FooMsdPick(req)
	if err != nil {
		return z, err
	}

	v, err := db.FooMsd.Query().Where(p).OnlyID(ctx)
	if err != nil {
		return z, status.Errorf(codes.Internal, "query: %s", err)
	}

	return v, nil
}

func FooMsdSelectedFields(m *library.FooMsdSelect) []string {
	if m.GetAll() {
		return foomsd.Columns
	}

	vs := []string{foomsd.FieldID}
	if m.GetMsdDouble() {
		vs = append(vs, foomsd.FieldMsdDouble)
	}
	if m.GetMsdInt64() {
		vs = append(vs, foomsd.FieldMsdInt64)
	}
	if m.GetMsdUint64() {
		vs = append(vs, foomsd.FieldMsdUint64)
	}
	if m.GetMsdBool() {
		vs = append(vs, foomsd.FieldMsdBool)
	}
	if m.GetMsdString() {
		vs = append(vs, foomsd.FieldMsdString)
	}
	if m.GetMsdBytes() {
		vs = append(vs, foomsd.FieldMsdBytes)
	}

	return vs
}

func FooMsdSelectEdges(q *ent.FooMsdQuery) {
}

func FooMsdSelect(q *ent.FooMsdQuery, m *library.FooMsdSelect) {
	if !m.GetAll() {
		fields := FooMsdSelectedFields(m)
		q.Select(fields...)
	}
}
