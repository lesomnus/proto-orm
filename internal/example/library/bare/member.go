// Code generated by "protoc-gen-orm-ent-grpc". DO NOT EDIT

package bare

import (
	context "context"
	uuid "github.com/google/uuid"
	library "github.com/lesomnus/proto-orm/internal/example/library"
	ent "github.com/lesomnus/proto-orm/internal/example/library/ent"
	member "github.com/lesomnus/proto-orm/internal/example/library/ent/member"
	predicate "github.com/lesomnus/proto-orm/internal/example/library/ent/predicate"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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
	if req.HasId() {
		if v, err := uuid.FromBytes(req.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "id: %s", err)
		} else {
			q.SetID(v)
		}
	}
	if req.HasName() {
		q.SetName(req.GetName())
	}
	if v := req.GetLabels(); v != nil {
		q.SetLabels(v)
	}
	q.SetProfile(req.GetProfile())
	q.SetLevel(int(req.GetLevel()))
	if v := req.GetLocker(); v != nil {
		if id, err := LockerGetId(ctx, s.db, v); err != nil {
			return nil, err
		} else {
			q.SetLockerID(id)
		}
	}
	if v := req.GetParent(); v != nil {
		if id, err := MemberGetId(ctx, s.db, v); err != nil {
			return nil, err
		} else {
			q.SetParentID(id)
		}
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

func (s *MemberServiceServer) Patch(ctx context.Context, req *library.MemberPatchRequest) (*emptypb.Empty, error) {
	id, err := MemberGetId(ctx, s.db, req.GetKey())
	if err != nil {
		return nil, err
	}

	q := s.db.Member.UpdateOneID(id)
	if req.HasName() {
		q.SetName(req.GetName())
	}
	if v := req.GetLabels(); v != nil {
		q.SetLabels(v)
	}
	if req.HasProfile() {
		q.SetProfile(req.GetProfile())
	}
	if req.HasLevel() {
		q.SetLevel(int(req.GetLevel()))
	}
	if req.GetLockerNull() {
		q.ClearLocker()
	} else if id, err := MemberGetId(ctx, s.db, req.GetKey()); err != nil {
		return nil, err
	} else {
		q.SetLockerID(id)
	}
	if req.GetParentNull() {
		q.ClearParent()
	} else if id, err := MemberGetId(ctx, s.db, req.GetKey()); err != nil {
		return nil, err
	} else {
		q.SetParentID(id)
	}

	if _, err := q.Save(ctx); err != nil {
		return nil, StatusFromEntError(err)
	}

	return nil, nil
}

func (s *MemberServiceServer) Erase(ctx context.Context, req *library.MemberGetRequest) (*emptypb.Empty, error) {
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
	switch req.WhichKey() {
	case library.MemberGetRequest_Id_case:
		if v, err := uuid.FromBytes(req.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "id: %s", "err")
		} else {
			return member.IDEQ(v), nil
		}
	case library.MemberGetRequest_Name_case:
		ps := make([]predicate.Member, 0, 2)
		if p, err := MemberPick(req.GetName().GetParent()); err != nil {
			s, _ := status.FromError(err)
			return nil, status.Errorf(codes.InvalidArgument, "name.parent: %s", s.Message())
		} else {
			ps = append(ps, member.HasParentWith(p))
		}
		ps = append(ps, member.NameEQ(req.GetName().GetName()))
		return member.And(ps...), nil
	case library.MemberGetRequest_Key_not_set_case:
		return nil, status.Errorf(codes.InvalidArgument, "key not provided")
	default:
		return nil, status.Errorf(codes.Unimplemented, "unknown type of key")
	}
}

func MemberGetId(ctx context.Context, db *ent.Client, req *library.MemberGetRequest) (uuid.UUID, error) {
	var z uuid.UUID
	if req.HasId() {
		if v, err := uuid.FromBytes(req.GetId()); err != nil {
			return z, status.Errorf(codes.InvalidArgument, "key.id: %s", err)
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
