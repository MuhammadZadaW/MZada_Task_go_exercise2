package model

import (
	"context"
	"time"
)

type User struct {
	ID			uint		`json:"id" gorm:"primary_key;auto_increment;column:id"`
	Name		string		`json:"name" gorm:"column:name"`
	Job			string		`json:"job" gorm:"column:job"`
	CreatedAt	time.Time	`json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt	time.Time	`json:"updatedAt" gorm:"column:updatedAt"`
}

type UserRepository interface {

	AddUser(ctx context.Context, user User) (User, error)
	DeleteUser(ctx context.Context, id string) (error)
	UpdateUser(ctx context.Context, id string, user User) (User, error)
	// Fetch(ctx context.Context) (res []User, err error)
	// GetByID(ctx context.Context, id string) (User, error)

}

type UserUsecase interface {

	AddUser(ctx context.Context, user User) (User, error)
	DeleteUser(ctx context.Context, id string) (error)
	UpdateUser(ctx context.Context, id string, user User) (User, error)
	// Fetch(ctx context.Context) ([]User, error)
	// GetByID(ctx context.Context, id string) (User, error)
}