package repository

import (
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
	user_task "github.com/sawalreverr/recything/internal/task/user_task/entity"
)

type UserTaskRepository interface {
	GetAllTasks() ([]task.TaskChallenge, error)
	GetTaskById(id string) (*task.TaskChallenge, error)
	FindLastIdTaskChallenge() (string, error)
	FindUserTask(userId string, userTaskId string) (*user_task.UserTaskChallenge, error)
	CreateUserTask(userTask *user_task.UserTaskChallenge) (*user_task.UserTaskChallenge, error)
	FindTask(taskId string) (*task.TaskChallenge, error)
	UploadImageTask(userTask *user_task.UserTaskChallenge, userTaskId string) (*user_task.UserTaskChallenge, error)
	GetUserTaskByUserId(userId string) ([]user_task.UserTaskChallenge, error)
	GetUserTaskDoneByUserId(userId string) ([]user_task.UserTaskChallenge, error)
	FindUserHasSameTask(userId string, taskId string) (*user_task.UserTaskChallenge, error)
}
