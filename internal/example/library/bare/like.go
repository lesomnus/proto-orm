// Code generated by "protoc-gen-orm-ent-grpc". DO NOT EDIT

package bare

import (
	context "context"
	uuid "github.com/google/uuid"
	library "github.com/lesomnus/proto-orm/internal/example/library"
	ent "github.com/lesomnus/proto-orm/internal/example/library/ent"
	like "github.com/lesomnus/proto-orm/internal/example/library/ent/like"
	predicate "github.com/lesomnus/proto-orm/internal/example/library/ent/predicate"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type LikeServiceServer struct {
	db *ent.Client
	library.UnimplementedLikeServiceServer
}

func NewLikeServiceServer(db *ent.Client) *LikeServiceServer {
	return &LikeServiceServer{db: db}
}

func (s *LikeServiceServer) Add(ctx context.Context, req *library.LikeAddRequest) (*library.Like, error) {
	q := s.db.Like.Create()
	if v := req.Id; v != nil {
		if w, err := uuid.FromBytes(v); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "id: %s", err)
		} else {
			q.SetID(w)
		}
	}
	if id, err := BookGetId(ctx, s.db, req.GetSubject()); err != nil {
		return nil, err
	} else {
		q.SetSubjectID(id)
	}
	if id, err := MemberGetId(ctx, s.db, req.GetActor()); err != nil {
		return nil, err
	} else {
		q.SetActorID(id)
	}
	if v := req.DateCreated; v != nil {
		q.SetDateCreated(v.AsTime())
	}

	v, err := q.Save(ctx)
	if err != nil {
		return nil, StatusFromEntError(err)
	}

	return v.Proto(), nil
}

func (s *LikeServiceServer) Get(ctx context.Context, req *library.LikeGetRequest) (*library.Like, error) {
	q := s.db.Like.Query()
	if p, err := LikePick(req); err != nil {
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

func (s *LikeServiceServer) Patch(ctx context.Context, req *library.LikePatchRequest) (*emptypb.Empty, error) {
	id, err := LikeGetId(ctx, s.db, req.GetKey())
	if err != nil {
		return nil, err
	}

	q := s.db.Like.UpdateOneID(id)

	if _, err := q.Save(ctx); err != nil {
		return nil, StatusFromEntError(err)
	}

	return nil, nil
}

func (s *LikeServiceServer) Erase(ctx context.Context, req *library.LikeGetRequest) (*emptypb.Empty, error) {
	p, err := LikePick(req)
	if err != nil {
		return nil, err
	}
	if _, err := s.db.Like.Delete().Where(p).Exec(ctx); err != nil {
		return nil, StatusFromEntError(err)
	}

	return nil, nil
}

func LikePick(req *library.LikeGetRequest) (predicate.Like, error) {
	switch k := req.GetKey().(type) {
	case *library.LikeGetRequest_Id:
		if v, err := uuid.FromBytes(k.Id); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "id: %s", "err")
		} else {
			return like.IDEQ(v), nil
		}
	case *library.LikeGetRequest_Holders:
		ps := make([]predicate.Like, 0, 2)
		if v, err := uuid.FromBytes(k.Holders.SubjectId); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "holders.subject_id: %s", "err")
		} else {
			ps = append(ps, like.SubjectIDEQ(v))
		}
		if p, err := MemberPick(k.Holders.GetActor()); err != nil {
			s, _ := status.FromError(err)
			return nil, status.Errorf(codes.InvalidArgument, "holders.actor: %s", s.Message())
		} else {
			ps = append(ps, like.HasActorWith(p))
		}
		return like.And(ps...), nil
	case nil:
		return nil, status.Errorf(codes.InvalidArgument, "key not provided")
	default:
		return nil, status.Errorf(codes.Unimplemented, "unknown type of key")
	}
}

func LikeGetId(ctx context.Context, db *ent.Client, req *library.LikeGetRequest) (uuid.UUID, error) {
	var z uuid.UUID
	k := req.GetKey()
	if r, ok := k.(*library.LikeGetRequest_Id); ok {
		if v, err := uuid.FromBytes(r.Id); err != nil {
			return z, status.Errorf(codes.InvalidArgument, "key.id: %s", err)
		} else {
			return v, nil
		}
	}

	p, err := LikePick(req)
	if err != nil {
		return z, err
	}

	v, err := db.Like.Query().Where(p).OnlyID(ctx)
	if err != nil {
		return z, status.Errorf(codes.Internal, "query: %s", err)
	}

	return v, nil
}
