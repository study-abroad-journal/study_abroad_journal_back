package db

import (
	"database/sql"
	"sabj/app/domain"
)

type DiaryRepository struct {
    Conn *sql.DB
}

func NewDiaryRepository() *DiaryRepository {
    // TODO: 実際はDB接続をここで作る
    return &DiaryRepository{Conn: nil}

}

func (r *DiaryRepository) Create(d *domain.Diary) error {
    // 仮実装: 本当はINSERT文を書く
    d.ID = 1
    return nil
}

func (r *DiaryRepository) Save(diary *domain.Diary) error {
    // 仮実装: 本当はINSERT文を書く
    diary.ID = 1
    return nil
}
