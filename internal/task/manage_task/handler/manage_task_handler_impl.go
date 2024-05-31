package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/internal/task/manage_task/dto"
	"github.com/sawalreverr/recything/internal/task/manage_task/usecase"
	"github.com/sawalreverr/recything/pkg"
)

type ManageTaskHandlerImpl struct {
	Usecase usecase.ManageTaskUsecase
}

func NewManageTaskHandler(usecase usecase.ManageTaskUsecase) *ManageTaskHandlerImpl {
	return &ManageTaskHandlerImpl{Usecase: usecase}
}

func (handler *ManageTaskHandlerImpl) CreateTaskHandler(c echo.Context) error {
	claims := c.Get("user").(*helper.JwtCustomClaims)
	var request dto.CreateTaskResquest
	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid request body, detail : "+err.Error())
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}
	taskChallange, err := handler.Usecase.CreateTaskUsecase(&request, claims.UserID)

	if err != nil {
		if errors.Is(err, pkg.ErrTaskStepsNull) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrTaskStepsNull.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, detail : "+err.Error())
	}

	taskStep := []dto.TaskSteps{}

	data := dto.CreateTaskResponse{
		Id:          taskChallange.ID,
		Title:       taskChallange.Title,
		Description: taskChallange.Description,
		Thumbnail:   taskChallange.Thumbnail,
		StartDate:   taskChallange.StartDate,
		EndDate:     taskChallange.EndDate,
		Steps:       taskStep,
	}
	for _, step := range taskChallange.TaskSteps {
		taskSteps := dto.TaskSteps{
			Title:       step.Title,
			Description: step.Description,
		}
		taskStep = append(taskStep, taskSteps)
	}
	data.Steps = taskStep

	responseData := helper.ResponseData(http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, responseData)

}

func (handler *ManageTaskHandlerImpl) GetTaskChallengePaginationHandler(c echo.Context) error {
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "10"
	}
	limitInt, errLimit := strconv.Atoi(limit)
	if errLimit != nil || limitInt <= 0 {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid limit parameter")
	}
	pageInt, errPage := strconv.Atoi(page)
	if errPage != nil || pageInt <= 0 {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid page parameter")
	}

	tasks, totalData, err := handler.Usecase.GetTaskChallengePagination(pageInt, limitInt)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, detail: "+err.Error())
	}

	var data []dto.DataTasks
	for _, task := range tasks {
		var taskSteps []dto.TaskSteps
		for _, step := range task.TaskSteps {
			taskSteps = append(taskSteps, dto.TaskSteps{
				Id:          step.ID,
				Title:       step.Title,
				Description: step.Description,
			})
		}
		data = append(data, dto.DataTasks{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Thumbnail:   task.Thumbnail,
			StartDate:   task.StartDate,
			EndDate:     task.EndDate,
			Steps:       taskSteps,
			TaskCreator: dto.TaskCreatorAdmin{
				Id:   task.AdminId,
				Name: task.Admin.Name,
			},
		})
	}

	totalPage := totalData / limitInt
	if totalData%limitInt != 0 {
		totalPage++
	}

	responseDataPagination := dto.GetTaskPagination{
		Code:      http.StatusOK,
		Message:   "success",
		Data:      data,
		Page:      pageInt,
		Limit:     limitInt,
		TotalData: totalData,
		TotalPage: totalPage,
	}

	return c.JSON(http.StatusOK, responseDataPagination)
}

func (handler *ManageTaskHandlerImpl) UploadThumbnailHandler(c echo.Context) error {
	file, errFile := c.FormFile("thumbnail")
	if errFile != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "thumbnail is required")
	}

	if file.Size > 2*1024*1024 {
		return helper.ErrorHandler(c, http.StatusBadRequest, "file is too large")
	}

	if !strings.HasPrefix(file.Header.Get("Content-Type"), "image") {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid file type")
	}

	src, errOpen := file.Open()
	if errOpen != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, "failed to open file: "+errOpen.Error())
	}
	defer src.Close()

	imageUrl, err := helper.UploadToCloudinary(src, "task_thumbnail")
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, pkg.ErrUploadCloudinary.Error())
	}

	data := dto.TaskUploadThumbnailResponse{
		Thumbnail: imageUrl,
	}

	responseData := helper.ResponseData(http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, responseData)

}

func (handler *ManageTaskHandlerImpl) GetTaskByIdHandler(c echo.Context) error {
	id := c.Param("taskId")
	task, err := handler.Usecase.GetTaskByIdUsecase(id)
	if err != nil {
		if errors.Is(err, pkg.ErrTaskNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrTaskNotFound.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, detail: "+err.Error())
	}

	var taskSteps []dto.TaskSteps
	for _, step := range task.TaskSteps {
		taskSteps = append(taskSteps, dto.TaskSteps{
			Id:          step.ID,
			Title:       step.Title,
			Description: step.Description,
		})
	}
	data := dto.TaskGetByIdResponse{
		Id:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Thumbnail:   task.Thumbnail,
		StartDate:   task.StartDate,
		EndDate:     task.EndDate,
		Steps:       taskSteps,
		TaskCreator: dto.TaskCreatorAdmin{
			Id:   task.AdminId,
			Name: task.Admin.Name,
		},
	}
	responseData := helper.ResponseData(http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, responseData)
}

func (handler *ManageTaskHandlerImpl) UpdateTaskHandler(c echo.Context) error {
	var request dto.UpdateTaskRequest
	id := c.Param("taskId")
	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid request body, detail: "+err.Error())
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if len(request.Steps) == 0 {
		return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrTaskStepsNull.Error())
	}

	task, err := handler.Usecase.UpdateTaskChallengeUsecase(&request, id)
	if err != nil {
		if errors.Is(err, pkg.ErrTaskNotFound) {
			return helper.ErrorHandler(c, http.StatusNotFound, pkg.ErrTaskNotFound.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error, detail: "+err.Error())
	}
	var taskSteps []dto.TaskSteps
	for _, step := range task.TaskSteps {
		taskSteps = append(taskSteps, dto.TaskSteps{
			Id:          step.ID,
			Title:       step.Title,
			Description: step.Description,
		})
	}
	data := dto.UpdateTaskResponse{
		Id:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Thumbnail:   task.Thumbnail,
		StartDate:   task.StartDate,
		EndDate:     task.EndDate,
		Steps:       taskSteps,
	}
	responseData := helper.ResponseData(http.StatusOK, "success", data)
	return c.JSON(http.StatusOK, responseData)
}
