// Code generated by "protoc-gen-orm-ent-grpc". DO NOT EDIT

package bare

import (
	context "context"
	uuid "github.com/google/uuid"
	library "github.com/lesomnus/proto-orm/_example/library"
	ent "github.com/lesomnus/proto-orm/example/library/ent"
	member "github.com/lesomnus/proto-orm/example/library/ent/member"
	predicate "github.com/lesomnus/proto-orm/example/library/ent/predicate"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type MemberServiceServer struct {
	db *ent.Client
	library.UnimplementedMemberServiceServer
}

func NewMemberServiceServer(db *ent.Client) *MemberServiceServer {
	return &MemberServiceServer{db: db}
}

func (s *MemberServiceServer) Add(ctx context.Context, req *library.MemberAddRequest) (*library.Member, error) {
	q := s.db.Member.Create()
	if v, err := uuid.FromBytes(req.Id); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "id: %s", err)
	} else {
		q.SetID(v)
	}
	q.SetName(req.Name)
	q.SetLabels(req.Labels)
	q.SetProfile(req.Profile)
	q.SetLevel(int(req.Level))
	q.SetDateCreated(req.DateCreated.AsTime())

	v, err := q.Save(ctx)
	if err != nil {
		return nil, StatusFromEntError(err)
	}

	return v.Proto(), nil
}

func (s *MemberServiceServer) Get(ctx context.Context, req *library.MemberGetRequest) (*library.Member, error) {
	q := s.db.Member.Query()
	if p, err := MemberPick(req); err != nil {
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

func (s *MemberServiceServer) Patch(ctx context.Context, req *library.MemberPatchRequest) (*library.Member, error) {
	id, err := MemberGetId(ctx, s.db, req.GetKey())
	if err != nil {
		return nil, err
	}

	q := s.db.Member.UpdateOneID(id)
	if req.Name != nil {
		q.SetName(*req.Name)
	}
	if req.Labels != nil {
		q.SetLabels(req.Labels)
	}
	if req.Profile != nil {
		q.SetProfile(req.Profile)
	}
	if req.Level != nil {
		q.SetLevel(int(*req.Level))
	}

	v, err := q.Save(ctx)
	if err != nil {
		return nil, StatusFromEntError(err)
	}

	return v.Proto(), nil
}

func (s *MemberServiceServer) Erase(ctx context.Context, req *library.MemberGetRequest) (*library.Member, error) {
	p, err := MemberPick(req)
	if err != nil {
		return nil, err
	}
	if _, err := s.db.Member.Delete().Where(p).Exec(ctx); err != nil {
		return nil, StatusFromEntError(err)
	}

	return nil, nil
}

func MemberPick(req *library.MemberGetRequest) (predicate.Member, error) {
	switch k := req.GetKey().(type) {
	case *library.MemberGetRequest_Id:
		if v, err := uuid.FromBytes(k.Id); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "id: %s", "err")
		} else {
			return member.IDEQ(v), nil
		}
	case nil:
		return nil, status.Errorf(codes.InvalidArgument, "key not provided")
	default:
		return nil, status.Errorf(codes.Unimplemented, "unknown type of key")
	}
}

func MemberGetId(ctx context.Context, db *ent.Client, req *library.MemberGetRequest) (uuid.UUID, error) {
	var z uuid.UUID
	k := req.GetKey()
	if r, ok := k.(*library.MemberGetRequest_Id); ok {
		if v, err := uuid.FromBytes(r.Id); err != nil {
			return z, status.Errorf(codes.InvalidArgument, "Id: %s", "err")
		} else {
			return v, nil
		}
	}

	p, err := MemberPick(req)
	if err != nil {
		return z, err
	}

	v, err := db.Member.Query().Where(p).OnlyID(ctx)
	if err != nil {
		return z, status.Errorf(codes.Internal, "query: %s", err)
	}

	return v, nil
}
