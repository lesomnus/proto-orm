// Code generated by "protoc-gen-orm-ent-grpc". DO NOT EDIT

package bare

import (
	context "context"
	uuid "github.com/google/uuid"
	library "github.com/lesomnus/proto-orm/internal/example/library"
	ent "github.com/lesomnus/proto-orm/internal/example/library/ent"
	loan "github.com/lesomnus/proto-orm/internal/example/library/ent/loan"
	predicate "github.com/lesomnus/proto-orm/internal/example/library/ent/predicate"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type LoanServiceServer struct {
	db *ent.Client
	library.UnimplementedLoanServiceServer
}

func NewLoanServiceServer(db *ent.Client) *LoanServiceServer {
	return &LoanServiceServer{db: db}
}

func (s *LoanServiceServer) Add(ctx context.Context, req *library.LoanAddRequest) (*library.Loan, error) {
	q := s.db.Loan.Create()
	if req.HasId() {
		if v, err := uuid.FromBytes(req.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "id: %s", err)
		} else {
			q.SetID(v)
		}
	}
	if id, err := BookGetId(ctx, s.db, req.GetSubject()); err != nil {
		return nil, err
	} else {
		q.SetSubjectID(id)
	}
	if id, err := MemberGetId(ctx, s.db, req.GetBorrower()); err != nil {
		return nil, err
	} else {
		q.SetBorrowerID(id)
	}
	q.SetDateDue(req.GetDateDue().AsTime())
	if req.HasDateReturn() {
		q.SetDateReturn(req.GetDateReturn().AsTime())
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

func (s *LoanServiceServer) Get(ctx context.Context, req *library.LoanGetRequest) (*library.Loan, error) {
	q := s.db.Loan.Query()
	if p, err := LoanPick(req); err != nil {
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

func (s *LoanServiceServer) Patch(ctx context.Context, req *library.LoanPatchRequest) (*emptypb.Empty, error) {
	id, err := LoanGetId(ctx, s.db, req.GetKey())
	if err != nil {
		return nil, err
	}

	q := s.db.Loan.UpdateOneID(id)
	if req.HasDateDue() {
		q.SetDateDue(req.GetDateDue().AsTime())
	}
	if req.GetDateReturnNull() {
		q.ClearDateReturn()
	} else if req.HasDateReturn() {
		q.SetDateReturn(req.GetDateReturn().AsTime())
	}

	if _, err := q.Save(ctx); err != nil {
		return nil, StatusFromEntError(err)
	}

	return nil, nil
}

func (s *LoanServiceServer) Erase(ctx context.Context, req *library.LoanGetRequest) (*emptypb.Empty, error) {
	p, err := LoanPick(req)
	if err != nil {
		return nil, err
	}
	if _, err := s.db.Loan.Delete().Where(p).Exec(ctx); err != nil {
		return nil, StatusFromEntError(err)
	}

	return nil, nil
}

func LoanPick(req *library.LoanGetRequest) (predicate.Loan, error) {
	switch req.WhichKey() {
	case library.LoanGetRequest_Id_case:
		if v, err := uuid.FromBytes(req.GetId()); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "id: %s", "err")
		} else {
			return loan.IDEQ(v), nil
		}
	case library.LoanGetRequest_Key_not_set_case:
		return nil, status.Errorf(codes.InvalidArgument, "key not provided")
	default:
		return nil, status.Errorf(codes.Unimplemented, "unknown type of key")
	}
}

func LoanGetId(ctx context.Context, db *ent.Client, req *library.LoanGetRequest) (uuid.UUID, error) {
	var z uuid.UUID
	if req.HasId() {
		if v, err := uuid.FromBytes(req.GetId()); err != nil {
			return z, status.Errorf(codes.InvalidArgument, "key.id: %s", err)
		} else {
			return v, nil
		}
	}

	p, err := LoanPick(req)
	if err != nil {
		return z, err
	}

	v, err := db.Loan.Query().Where(p).OnlyID(ctx)
	if err != nil {
		return z, status.Errorf(codes.Internal, "query: %s", err)
	}

	return v, nil
}
