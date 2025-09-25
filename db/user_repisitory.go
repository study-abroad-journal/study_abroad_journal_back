package db

import (
	"database/sql"
	"fmt"
	"log"
	"sabj/internal/domain"

	_ "github.com/lib/pq"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository() (*UserRepository, error) {
	dsn := "postgres://user:pass@db:5432/sabj_db?sslmode=disable" // data source name
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("DB接続失敗:", err)
		return nil, err
	}

	// 接続確認
	if err := db.Ping(); err != nil {
		log.Fatal("DB ping失敗:", err)
		return nil, err
	}
	fmt.Println("DB接続成功！")
	return &UserRepository{DB: db}, nil
}

func (ur *UserRepository) Create(user *domain.User) error {
	// データベースにユーザーを挿入するSQLクエリ
	query := "INSERT INTO users (user_id, email) VALUES ($1, $2)"

	// SQLを実行
	_, err := ur.DB.Exec(query, user.User_id, user.Name)
	if err != nil {
		log.Printf("ユーザー作成エラー: %v", err)
		return err
	}

	fmt.Printf("ユーザー作成成功: %s\n", user.Name)
	return nil
}
