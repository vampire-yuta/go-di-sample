package entities

import (
	"time"

	"github.com/google/uuid"
)

// User はユーザーエンティティを表します
type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewUser は新しいユーザーを作成します
func NewUser(name, email string) *User {
	now := time.Now()
	return &User{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
