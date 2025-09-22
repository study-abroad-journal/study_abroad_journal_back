package domain

import "time"

type User struct {
	// フロントエンドから送られてくる "id" (Firebase UID) に対応
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
