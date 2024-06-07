package usecase

import (
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
	"github.com/sawalreverr/recything/internal/video/user_video/dto"
)

type UserVideoUsecase interface {
	GetAllVideoUsecase() (*[]video.Video, error)
	SearchVideoByTitleUsecase(title string) (*[]video.Video, error)
	GetVideoDetailUsecase(id int) (*video.Video, *[]video.Comment, error)
	AddCommentUsecase(request *dto.AddCommentRequest, userId string) error
}