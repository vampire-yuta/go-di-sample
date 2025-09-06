package usecase

import (
	"errors"
	"time"

	"go-di/internal/domain/entities"
	"go-di/internal/domain/repositories"
)

// UserService はユーザー関連のビジネスロジックを提供します
type UserService struct {
	userRepo repositories.UserRepository
}

// NewUserService は新しいUserServiceを作成します
func NewUserService(userRepo repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// CreateUserRequest はユーザー作成リクエストを表します
type CreateUserRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

// CreateUser は新しいユーザーを作成します
func (s *UserService) CreateUser(req CreateUserRequest) (*entities.User, error) {
	// バリデーション
	if req.Name == "" {
		return nil, errors.New("name is required")
	}
	if req.Email == "" {
		return nil, errors.New("email is required")
	}

	// ユーザー作成
	user := entities.NewUser(req.Name, req.Email)

	// リポジトリに保存
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUser はIDでユーザーを取得します
func (s *UserService) GetUser(id string) (*entities.User, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}

	return s.userRepo.GetByID(id)
}

// GetAllUsers はすべてのユーザーを取得します
func (s *UserService) GetAllUsers() ([]*entities.User, error) {
	return s.userRepo.GetAll()
}

// UpdateUserRequest はユーザー更新リクエストを表します
type UpdateUserRequest struct {
	ID    string `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

// UpdateUser はユーザーを更新します
func (s *UserService) UpdateUser(req UpdateUserRequest) (*entities.User, error) {
	// 既存ユーザーを取得
	user, err := s.userRepo.GetByID(req.ID)
	if err != nil {
		return nil, err
	}

	// フィールドを更新
	user.Name = req.Name
	user.Email = req.Email
	user.UpdatedAt = time.Now()

	// リポジトリに保存
	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser はユーザーを削除します
func (s *UserService) DeleteUser(id string) error {
	if id == "" {
		return errors.New("id is required")
	}

	return s.userRepo.Delete(id)
}
