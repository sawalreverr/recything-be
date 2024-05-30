package repository

import (
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
)

type ManageTaskRepository interface {
	CreateTask(task *task.TaskChallenge) (*task.TaskChallenge, error)
	FindLastIdTaskChallenge() (string, error)
	GetTaskChallengePagination(page int, limit int) ([]task.TaskChallenge, int, error)
	GetTaskById(id string) (*task.TaskChallenge, error)
}
