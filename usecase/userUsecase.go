package usecase

import (
	"context"
	"time"
	
	"task_go/model"
)

type userUsecase struct {
	userRepo model.UserRepository
}

func NewUserUsecase(a model.UserRepository) model.UserUsecase {
	return &userUsecase{
		userRepo: a,
	}
}

func (a *userUsecase) AddUser(c context.Context, user model.User) (res model.User, err error) {
	res, err = a.userRepo.AddUser(c, user)
	return
}

func (a *userUsecase) DeleteUser(c context.Context, id string) (err error) {
	err = a.userRepo.DeleteUser(c, id)
	return
}

func (a *userUsecase) UpdateUser(c context.Context, id string, user model.User) (res model.User, err error) {
	user.UpdatedAt = time.Now()
	res, err = a.userRepo.UpdateUser(c, id, user)
	return
}