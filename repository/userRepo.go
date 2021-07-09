package repository

import (
	"context"
	"gorm.io/gorm"

	"task_go/model"
)

type psqlUserRepository struct {
	Conn *gorm.DB
}

func NewPsqlUserRepository(Conn *gorm.DB) model.UserRepository {
	return &psqlUserRepository{Conn}
}

func (m *psqlUserRepository) AddUser(ctx context.Context, user model.User) (res model.User, err error) {
	m.Conn.Create(&user)
	return user, nil
}

func (m *psqlUserRepository) DeleteUser(ctx context.Context, id string) (err error) {
	var user model.User
	m.Conn.Where("id = ?", id).First(&user)
	m.Conn.Delete(&user)
	return nil
}

func (m *psqlUserRepository) UpdateUser(ctx context.Context, id string, user model.User) (res model.User, err error) {
	var userUp model.User
	
	m.Conn.Where("id = ?", id).First(&userUp)

	userUp.Name = user.Name
	userUp.Job = user.Job
	userUp.UpdatedAt = user.UpdatedAt

	m.Conn.Save(&userUp)

	return userUp, nil
}