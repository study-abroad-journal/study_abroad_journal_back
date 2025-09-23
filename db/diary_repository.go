package db

import (
	"database/sql"
	"fmt"
	"log"
	"sabj/internal/domain"
	"time"

	_ "github.com/lib/pq"
)

type DiaryRepository struct {
    DB *sql.DB
}

func NewDiaryRepository()(*DiaryRepository, error) {
    dsn := "postgres://user:pass@db:5432/sabj_db?sslmode=disable"   // data source name
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
    return &DiaryRepository{DB: db}, nil
}

func (r *DiaryRepository) Create(d *domain.Diary) (*domain.Diary, error) {
	query := `
		INSERT INTO diaries
			(user_id, title, category_id, latitude, longitude, text, corrected_text, created_at)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING diary_id, created_at
	`

	createdAt := time.Now()
	err := r.DB.QueryRow(
		query,
		d.UserID,
		d.Title,
		d.CategoryID,
		d.Latitude,
		d.Longitude,
		d.Text,
		d.CorrectedText,
		createdAt,
	).Scan(&d.DiaryID, &d.CreatedAt)

	if err != nil {
		return nil, err
	}

	return d, nil
}

func (r *DiaryRepository) GetByID(id int64) (*domain.Diary, error) {
    d := &domain.Diary{}
    query := 
        `SELECT diary_id, user_id, title, category_id, latitude, longitude, text, corrected_text, created_at 
        FROM diaries 
        WHERE diary_id=$1`

    err := r.DB.QueryRow(query, id).Scan(
        &d.DiaryID, &d.UserID, &d.Title, &d.CategoryID, &d.Latitude, &d.Longitude, &d.Text, &d.CorrectedText, &d.CreatedAt,
    )
    if err != nil {
        return nil, err
    }
    return d, nil
}

func (r *DiaryRepository) GetAll(userID int64) ([]*domain.Diary, error) {
    query := `
        SELECT diary_id, user_id, title, category_id, latitude, longitude, text, corrected_text, created_at 
        FROM diaries 
        WHERE user_id = $1 
        ORDER BY created_at DESC
    `
    rows, err := r.DB.Query(query, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var diaries []*domain.Diary
    for rows.Next() {
        d := &domain.Diary{}
        err := rows.Scan(
            &d.DiaryID, &d.UserID, &d.Title, &d.CategoryID, &d.Latitude, &d.Longitude, &d.Text, &d.CorrectedText, &d.CreatedAt,
        )
        if err != nil {
            return nil, err
        }
        diaries = append(diaries, d)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return diaries, nil
}
// Delete
func (r *DiaryRepository) Delete(id int64) error {
    _, err := r.DB.Exec(`DELETE FROM diaries WHERE diary_id=$1`, id)
    return err
}

// Update
func (r *DiaryRepository) Update(d *domain.Diary) error {
    query := `
        UPDATE diaries
        SET title=$1, category_id=$2, latitude=$3, longitude=$4, text=$5, corrected_text=$6
        WHERE diary_id=$7
    `
    _, err := r.DB.Exec(query, d.Title, d.CategoryID, d.Latitude, d.Longitude, d.Text, d.CorrectedText, d.DiaryID)
    return err
}
