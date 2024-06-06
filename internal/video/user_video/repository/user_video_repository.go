package repository

import (
	video "github.com/sawalreverr/recything/internal/video/manage_video/entity"
)

type UserVideoRepository interface {
	GetAllVideo() (*[]video.Video, error)
}
