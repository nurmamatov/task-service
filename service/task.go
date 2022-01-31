package service

import (
	"context"
	pb "khusniddin/task-servise/genproto"
	l "khusniddin/task-servise/pkg/logger"

	"khusniddin/task-servise/storage"

	"github.com/jmoiron/sqlx"
)

// UserService ...
type UserService struct {
	storage storage.IStorage
	logger  l.Logger
}

// NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *UserService) Create(ctx context.Context, req *pb.Task) (*pb.Task, error) {
	res, err := s.storage.User().Create(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) List(ctx context.Context, req *pb.Empty) (*pb.Tasks, error) {
	res, err := s.storage.User().List(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) Get(ctx context.Context, req *pb.Task) (*pb.Task, error) {
	res, err := s.storage.User().Get(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) ListOverdue(ctx context.Context, req *pb.Empty) (*pb.Tasks, error) {
	res, err := s.storage.User().ListOverdue(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) Delete(ctx context.Context, req *pb.Task) (*pb.Message, error) {
	res, err := s.storage.User().Delete(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserService) Update(ctx context.Context, req *pb.Task) (*pb.Task, error) {
	res, err := s.storage.User().Update(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
