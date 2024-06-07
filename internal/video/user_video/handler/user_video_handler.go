package handler

import "github.com/labstack/echo/v4"

type UserVideoHandler interface {
	GetAllVideoHandler(c echo.Context) error
	SearchVideoByTitleHandler(c echo.Context) error
	GetVideoDetailHandler(c echo.Context) error
	AddCommentHandler(c echo.Context) error
}