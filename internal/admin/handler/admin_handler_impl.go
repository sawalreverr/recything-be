package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sawalreverr/recything/internal/admin/dto"
	"github.com/sawalreverr/recything/internal/admin/usecase"
	"github.com/sawalreverr/recything/internal/helper"
	"github.com/sawalreverr/recything/pkg"
)

type adminHandlerImpl struct {
	Usecase usecase.AdminUsecase
}

func NewAdminHandler(adminUsecase usecase.AdminUsecase) *adminHandlerImpl {
	return &adminHandlerImpl{Usecase: adminUsecase}
}

func (handler *adminHandlerImpl) AddAdminHandler(c echo.Context) error {
	var request dto.AdminRequestCreate
	if err := c.Bind(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "invalid request body")
	}

	if err := c.Validate(&request); err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	if request.Role != "admin" && request.Role != "super admin" {
		return helper.ErrorHandler(c, http.StatusBadRequest, "role must be admin or super admin")
	}

	file, errFile := c.FormFile("profile_photo")
	if errFile != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, "profile_photo not found")
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

	admin, errUc := handler.Usecase.AddAdminUsecase(request, src)
	if errUc != nil {
		if errors.Is(errUc, pkg.ErrEmailAlreadyExist) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrEmailAlreadyExist.Error())
		}

		if errors.Is(errUc, pkg.ErrUploadCloudinary) {
			return helper.ErrorHandler(c, http.StatusBadRequest, pkg.ErrUploadCloudinary.Error())
		}
		return helper.ErrorHandler(c, http.StatusInternalServerError, "internal server error")
	}

	data := dto.AdminResponseRegister{
		Id:           admin.ID,
		Name:         admin.Name,
		Email:        admin.Email,
		Role:         admin.Role,
		ProfilePhoto: admin.ImageUrl,
	}
	responseData := helper.ResponseData(http.StatusCreated, "success", data)
	return c.JSON(http.StatusCreated, responseData)
}

func (handler *adminHandlerImpl) GetDataAllAdminHandler(c echo.Context) error {
	limit := c.QueryParam("limit")

	if limit == "" {
		limit = "10"
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusBadRequest, err.Error())
	}

	admins, err := handler.Usecase.GetDataAllAdminUsecase(limitInt)
	if err != nil {
		return helper.ErrorHandler(c, http.StatusInternalServerError, err.Error())
	}

	// Create the slice of AdminDataGetAll directly
	data := []dto.AdminDataGetAll{}

	for _, admin := range admins {
		data = append(data, dto.AdminDataGetAll{
			Id:    admin.ID,
			Name:  admin.Name,
			Email: admin.Email,
			Role:  admin.Role,
		})
	}

	dataRes := dto.AdminResponseGetDataAll{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
		Limit:   limitInt,
		Total:   len(admins),
	}

	return c.JSON(http.StatusOK, dataRes)
}

func (handler *adminHandlerImpl) UpdateAdminHandler(c echo.Context) error {
	return nil
}
