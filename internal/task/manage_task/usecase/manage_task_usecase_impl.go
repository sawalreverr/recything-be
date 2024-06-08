package usecase

import (
	"mime/multipart"
	"time"

	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/internal/task/manage_task/dto"
	task "github.com/sawalreverr/recything/internal/task/manage_task/entity"
	"github.com/sawalreverr/recything/internal/task/manage_task/repository"
	"github.com/sawalreverr/recything/pkg"
	"gorm.io/gorm"
)

type ManageTaskUsecaseImpl struct {
	ManageTaskRepository repository.ManageTaskRepository
}

func NewManageTaskUsecase(repository repository.ManageTaskRepository) *ManageTaskUsecaseImpl {
	return &ManageTaskUsecaseImpl{ManageTaskRepository: repository}
}

func (usecase *ManageTaskUsecaseImpl) CreateTaskUsecase(request *dto.CreateTaskResquest, thumbnail []*multipart.FileHeader, adminId string) (*task.TaskChallenge, error) {
	if len(thumbnail) == 0 {
		return nil, pkg.ErrThumbnail
	}
	if len(thumbnail) > 1 {
		return nil, pkg.ErrThumbnailMaximum
	}
	if len(request.TaskSteps) == 0 {
		return nil, pkg.ErrTaskStepsNull
	}
	validImages, errImages := helper.ImagesValidation(thumbnail)
	if errImages != nil {
		return nil, errImages
	}
	urlThumbnail, errUpload := helper.UploadToCloudinary(validImages[0], "task_thumbnail")
	if errUpload != nil {
		return nil, pkg.ErrUploadCloudinary
	}

	findLastId, _ := usecase.ManageTaskRepository.FindLastIdTaskChallenge()
	id := helper.GenerateCustomID(findLastId, "TM")
	startDateString := request.StartDate
	endDateString := request.EndDate
	parsedStartDate, errParsedStartDate := time.Parse("2006-01-02", startDateString)
	if errParsedStartDate != nil {
		return nil, pkg.ErrParsedTime
	}
	parsedEndDate, errParsedEndDate := time.Parse("2006-01-02", endDateString)
	if errParsedEndDate != nil {
		return nil, pkg.ErrParsedTime
	}

	taskChallange := &task.TaskChallenge{
		ID:          id,
		AdminId:     adminId,
		Title:       request.Title,
		Description: request.Description,
		Thumbnail:   urlThumbnail,
		StartDate:   parsedStartDate,
		EndDate:     parsedEndDate,
		Point:       request.Point,
		TaskSteps:   []task.TaskStep{},
		DeletedAt:   gorm.DeletedAt{},
	}

	for _, step := range request.TaskSteps {
		taskStep := task.TaskStep{
			TaskChallengeId: id,
			Title:           step.Title,
			Description:     step.Description,
		}
		taskChallange.TaskSteps = append(taskChallange.TaskSteps, taskStep)
	}

	if _, err := usecase.ManageTaskRepository.CreateTask(taskChallange); err != nil {
		return nil, err
	}
	return taskChallange, nil
}

func (usecase *ManageTaskUsecaseImpl) GetTaskChallengePagination(page int, limit int) ([]task.TaskChallenge, int, error) {
	tasks, total, err := usecase.ManageTaskRepository.GetTaskChallengePagination(page, limit)
	if err != nil {
		return nil, 0, err
	}
	return tasks, total, nil
}

func (usecase *ManageTaskUsecaseImpl) GetTaskByIdUsecase(id string) (*task.TaskChallenge, error) {
	task, err := usecase.ManageTaskRepository.GetTaskById(id)
	if err != nil {
		return nil, pkg.ErrTaskNotFound
	}
	return task, nil
}

func (usecase *ManageTaskUsecaseImpl) UpdateTaskChallengeUsecase(request *dto.UpdateTaskRequest, id string) (*task.TaskChallenge, error) {
	findTask, _ := usecase.ManageTaskRepository.FindTask(id)

	if findTask == nil {
		return nil, pkg.ErrTaskNotFound
	}
	if len(request.TaskSteps) == 0 {
		return nil, pkg.ErrTaskStepsNull
	}
	startDateString := request.StartDate
	endDateString := request.EndDate
	parsedStartDate, errParsedStartDate := time.Parse("2006-01-02", startDateString)
	if errParsedStartDate != nil {
		return nil, pkg.ErrParsedTime
	}
	parsedEndDate, errParsedEndDate := time.Parse("2006-01-02", endDateString)
	if errParsedEndDate != nil {
		return nil, pkg.ErrParsedTime
	}

	taskChallenge := &task.TaskChallenge{
		ID:          id,
		AdminId:     findTask.AdminId,
		Title:       request.Title,
		Description: request.Description,
		Thumbnail:   request.ThumbnailUrl,
		StartDate:   parsedStartDate,
		EndDate:     parsedEndDate,
		Point:       request.Point,
	}

	// Add new steps
	for _, step := range request.TaskSteps {
		taskStep := task.TaskStep{
			TaskChallengeId: id,
			Title:           step.Title,
			Description:     step.Description,
		}
		taskChallenge.TaskSteps = append(taskChallenge.TaskSteps, taskStep)
	}

	updatedTaskChallenge, err := usecase.ManageTaskRepository.UpdateTaskChallenge(taskChallenge, id)
	if err != nil {
		return nil, err
	}

	return updatedTaskChallenge, nil
}

func (usecase *ManageTaskUsecaseImpl) DeleteTaskChallengeUsecase(id string) error {
	if _, err := usecase.ManageTaskRepository.FindTask(id); err != nil {
		return pkg.ErrTaskNotFound
	}
	if err := usecase.ManageTaskRepository.DeleteTaskChallenge(id); err != nil {
		return err
	}
	return nil
}
