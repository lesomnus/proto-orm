// Code generated by "protoc-gen-orm-ent-grpc". DO NOT EDIT

package bare

import (
	context "context"
	uuid "github.com/google/uuid"
	library "github.com/lesomnus/proto-orm/internal/example/library"
	ent "github.com/lesomnus/proto-orm/internal/example/library/ent"
	locker "github.com/lesomnus/proto-orm/internal/example/library/ent/locker"
	predicate "github.com/lesomnus/proto-orm/internal/example/library/ent/predicate"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type LockerServiceServer struct {
	db *ent.Client
	library.UnimplementedLockerServiceServer
}

func NewLockerServiceServer(db *ent.Client) *LockerServiceServer {
	return &LockerServiceServer{db: db}
}

func (s *LockerServiceServer) Add(ctx context.Context, req *library.LockerAddRequest) (*library.Locker, error) {
	q := s.db.Locker.Create()
	if v := req.Id; v != nil {
		if w, err := uuid.FromBytes(v); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "id: %s", err)
		} else {
			q.SetID(w)
		}
	}
	if id, err := MemberGetId(ctx, s.db, req.GetOwner()); err != nil {
		return nil, err
	} else {
		q.SetOwnerID(id)
	}

	v, err := q.Save(ctx)
	if err != nil {
		return nil, StatusFromEntError(err)
	}

	return v.Proto(), nil
}

func (s *LockerServiceServer) Get(ctx context.Context, req *library.LockerGetRequest) (*library.Locker, error) {
	q := s.db.Locker.Query()
	if p, err := LockerPick(req); err != nil {
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

func (s *LockerServiceServer) Patch(ctx context.Context, req *library.LockerPatchRequest) (*emptypb.Empty, error) {
	id, err := LockerGetId(ctx, s.db, req.GetKey())
	if err != nil {
		return nil, err
	}

	q := s.db.Locker.UpdateOneID(id)
	if req.OwnerNull {
		q.ClearOwner()
	} else if id, err := LockerGetId(ctx, s.db, req.GetKey()); err != nil {
		return nil, err
	} else {
		q.SetOwnerID(id)
	}

	if _, err := q.Save(ctx); err != nil {
		return nil, StatusFromEntError(err)
	}

	return nil, nil
}

func (s *LockerServiceServer) Erase(ctx context.Context, req *library.LockerGetRequest) (*emptypb.Empty, error) {
	p, err := LockerPick(req)
	if err != nil {
		return nil, err
	}
	if _, err := s.db.Locker.Delete().Where(p).Exec(ctx); err != nil {
		return nil, StatusFromEntError(err)
	}

	return nil, nil
}

func LockerPick(req *library.LockerGetRequest) (predicate.Locker, error) {
	switch k := req.GetKey().(type) {
	case *library.LockerGetRequest_Id:
		if v, err := uuid.FromBytes(k.Id); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "id: %s", "err")
		} else {
			return locker.IDEQ(v), nil
		}
	case nil:
		return nil, status.Errorf(codes.InvalidArgument, "key not provided")
	default:
		return nil, status.Errorf(codes.Unimplemented, "unknown type of key")
	}
}

func LockerGetId(ctx context.Context, db *ent.Client, req *library.LockerGetRequest) (uuid.UUID, error) {
	var z uuid.UUID
	k := req.GetKey()
	if r, ok := k.(*library.LockerGetRequest_Id); ok {
		if v, err := uuid.FromBytes(r.Id); err != nil {
			return z, status.Errorf(codes.InvalidArgument, "key.id: %s", err)
		} else {
			return v, nil
		}
	}

	p, err := LockerPick(req)
	if err != nil {
		return z, err
	}

	v, err := db.Locker.Query().Where(p).OnlyID(ctx)
	if err != nil {
		return z, status.Errorf(codes.Internal, "query: %s", err)
	}

	return v, nil
}
