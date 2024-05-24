package user

import (
	u "github.com/sawalreverr/recything/internal/user"
	"github.com/sawalreverr/recything/pkg"
)

type userUsecase struct {
	userRepository u.UserRepository
}

func NewUserUsecase(userRepo u.UserRepository) u.UserUsecase {
	return &userUsecase{userRepository: userRepo}
}

func (uc *userUsecase) UpdateUserDetail(userID string, user u.UserDetail) error {
	userFound, err := uc.userRepository.FindByID(userID)
	if err != nil {
		return pkg.ErrUserNotFound
	}

	userFound.Name = user.Name
	userFound.Email = user.Email
	userFound.PhoneNumber = user.PhoneNumber
	userFound.Gender = user.Gender
	userFound.BirthDate = user.ParsedBirthDate
	userFound.Address = user.Address

	if err := uc.userRepository.Update(*userFound); err != nil {
		return pkg.ErrStatusInternalError
	}

	return nil
}

func (uc *userUsecase) UpdateUserPicture(userID string, picture_url string) error {
	userFound, err := uc.userRepository.FindByID(userID)
	if err != nil {
		return pkg.ErrUserNotFound
	}

	userFound.PictureURL = picture_url
	if err := uc.userRepository.Update(*userFound); err != nil {
		return pkg.ErrStatusInternalError
	}

	return nil
}

func (uc *userUsecase) FindUserByID(userID string) (*u.UserResponse, error) {
	userFound, err := uc.userRepository.FindByID(userID)
	if err != nil {
		return nil, pkg.ErrUserNotFound
	}

	response := u.UserResponse{
		ID:          userFound.ID,
		Name:        userFound.Name,
		Email:       userFound.Email,
		PhoneNumber: userFound.PhoneNumber,
		Point:       userFound.Point,
		Gender:      userFound.Gender,
		BirthDate:   userFound.BirthDate,
		Address:     userFound.Address,
		PictureURL:  userFound.PictureURL,
	}

	return &response, nil
}

func (uc *userUsecase) FindAllUser(page int, limit int, sortBy string, sortType string) (*u.UserPaginationResponse, error) {
	var usersResponse []u.UserResponse
	users, err := uc.userRepository.FindAll(page, limit, sortBy, sortType)
	if err != nil {
		return nil, pkg.ErrStatusInternalError
	}

	totalCount, err := uc.userRepository.CountAllUser()
	if err != nil {
		return nil, pkg.ErrStatusInternalError
	}

	for _, user := range *users {
		response := u.UserResponse{
			ID:          user.ID,
			Name:        user.Name,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			Point:       user.Point,
			Gender:      user.Gender,
			BirthDate:   user.BirthDate,
			Address:     user.Address,
			PictureURL:  user.PictureURL,
		}

		usersResponse = append(usersResponse, response)
	}

	paginationResponse := u.UserPaginationResponse{
		TotalCount: totalCount,
		Page:       page,
		Limit:      limit,
		Users:      usersResponse,
	}

	return &paginationResponse, nil
}

func (uc *userUsecase) DeleteUser(userID string) error {
	userFound, err := uc.userRepository.FindByID(userID)
	if err != nil {
		return pkg.ErrUserNotFound
	}

	if err := uc.userRepository.Delete(userFound.ID); err != nil {
		return pkg.ErrStatusInternalError
	}

	return nil
}
