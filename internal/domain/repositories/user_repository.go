package repositories

import "go-di/internal/domain/entities"

// UserRepository はユーザーリポジトリのインターフェースです
type UserRepository interface {
	Create(user *entities.User) error
	GetByID(id string) (*entities.User, error)
	GetAll() ([]*entities.User, error)
	Update(user *entities.User) error
	Delete(id string) error
}
