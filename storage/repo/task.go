package repo

import (
	pb "khusniddin/task-servise/genproto"
)

type UserStorageI interface {
	Create(*pb.Task) (*pb.Task, error)
	Get(*pb.Task) (*pb.Task, error)
	List(*pb.Empty) (*pb.Tasks, error)
	Update(*pb.Task) (*pb.Task, error)
	Delete(*pb.Task) (*pb.Message, error)
	ListOverdue(*pb.Empty) (*pb.Tasks, error)
}
