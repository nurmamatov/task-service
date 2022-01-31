package postgres

import (
	pb "khusniddin/task-servise/genproto"
	"time"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

// NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(task *pb.Task) (*pb.Task, error) {
	query := `INSERT INTO tasks(assignee,title,deadline,status,created_at,updated_at,deleted_at) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id,assignee,title`
	CreatedAt := time.Now().Format(time.RFC3339)

	err := r.db.QueryRow(query, task.Assignee, task.Title, task.Deadline, task.Status, CreatedAt, task.UpdatedAt, task.DeletedAt).Scan(
		&task.Id,
		&task.Assignee,
		&task.Title,
	)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *userRepo) Get(task *pb.Task) (*pb.Task, error) {
	query := `SELECT assignee,title,deadline,status,created_at,updated_at,deleted_at FROM tasks WHERE id=$1`
	err := r.db.QueryRow(query, task.Id).Scan(
		&task.Assignee,
		&task.Title,
		&task.Deadline,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt,
		&task.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *userRepo) List(task *pb.Empty) (*pb.Tasks, error) {
	query := `SELECT id,assignee,title,deadline,status,created_at,updated_at,deleted_at FROM tasks`
	row, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	tasks := pb.Tasks{}
	for row.Next() {
		task_one := pb.Task{}
		err := row.Scan(
			&task_one.Id,
			&task_one.Assignee,
			&task_one.Title,
			&task_one.Deadline,
			&task_one.Status,
			&task_one.CreatedAt,
			&task_one.UpdatedAt,
			&task_one.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks.Tasks = append(tasks.Tasks, &task_one)
	}
	return &tasks, nil
}

func (r *userRepo) Delete(task *pb.Task) (*pb.Message, error) {
	query := `UPDATE tasks SET deleted_at=$1 WHERE id=$2`
	_, err := r.db.Exec(query, time.Now().Format(time.RFC3339), task.Id)
	if err != nil {
		return &pb.Message{Message: "Can't delete"}, err
	}
	return &pb.Message{Message: "Ok!"}, nil
}

func (r *userRepo) Update(task *pb.Task) (*pb.Task, error) {
	query := `UPDATE tasks SET assignee=$1, title=$2, deadline=$3, status=$4, updated_at=$5 WHERE id=$6 RETURNING assignee,title,deadline,status,updated_at`
	err := r.db.QueryRow(query, task.Id).Scan(
		&task.Assignee,
		&task.Title,
		&task.Deadline,
		&task.Status,
		&task.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *userRepo) ListOverdue(empty *pb.Empty) (*pb.Tasks, error) {
	query := `SELECT id,assignee,title,deadline,status,created_at,updated_at,deleted_at FROM tasks WHERE deadline<$1`
	now_time := time.Now().Format(time.RFC3339)
	rows, err := r.db.Query(query, now_time)
	if err != nil {
		return nil, err
	}
	tasks := pb.Tasks{}
	for rows.Next() {
		task := pb.Task{}
		err = rows.Scan(
			&task.Id,
			&task.Assignee,
			&task.Title,
			&task.Deadline,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
			&task.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks.Tasks = append(tasks.Tasks, &task)
	}
	return &tasks, nil
}
