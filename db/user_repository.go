package db

import (
	"database/sql"
	"log"
	"sabj/internal/domain"
	"time"

	_ "github.com/lib/pq"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(u *domain.User) (*domain.User, error) {
	query := `
        INSERT INTO users (user_id, email, created_at, updated_at)
        VALUES ($1, $2, $3, $4)
        RETURNING created_at, updated_at
    `

	now := time.Now()

	// .Scan()で受け取るのはDBが生成するcreated_atとupdated_atのみになる
	err := r.DB.QueryRow(
		query,
		u.ID, // フロントエンドから受け取ったFirebase UIDをそのまま使う
		u.Email,
		now,
		now,
	).Scan(&u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		log.Printf("UserRepository.Createでエラー発生: %v", err)
		return nil, err
	}

	return u, nil
}
