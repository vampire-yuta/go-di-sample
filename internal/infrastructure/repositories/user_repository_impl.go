package repositories

import (
	"errors"
	"sync"

	"go-di/internal/domain/entities"
	"go-di/internal/domain/repositories"
)

// userRepositoryImpl はメモリベースのユーザーリポジトリ実装です
type userRepositoryImpl struct {
	users map[string]*entities.User
	mu    sync.RWMutex
}

// NewUserRepositoryImpl は新しいuserRepositoryImplを作成します
func NewUserRepositoryImpl() repositories.UserRepository {
	return &userRepositoryImpl{
		users: make(map[string]*entities.User),
	}
}

// Create はユーザーを作成します
func (r *userRepositoryImpl) Create(user *entities.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// 既に存在するかチェック
	if _, exists := r.users[user.ID]; exists {
		return errors.New("user already exists")
	}

	r.users[user.ID] = user
	return nil
}

// GetByID はIDでユーザーを取得します
func (r *userRepositoryImpl) GetByID(id string) (*entities.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// GetAll はすべてのユーザーを取得します
func (r *userRepositoryImpl) GetAll() ([]*entities.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]*entities.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}

	return users, nil
}

// Update はユーザーを更新します
func (r *userRepositoryImpl) Update(user *entities.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// 存在するかチェック
	if _, exists := r.users[user.ID]; !exists {
		return errors.New("user not found")
	}

	r.users[user.ID] = user
	return nil
}

// Delete はユーザーを削除します
func (r *userRepositoryImpl) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// 存在するかチェック
	if _, exists := r.users[id]; !exists {
		return errors.New("user not found")
	}

	delete(r.users, id)
	return nil
}
