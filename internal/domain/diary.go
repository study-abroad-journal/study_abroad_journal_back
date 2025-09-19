package domain

import "time"

type Diary struct {
    DiaryID            int64     `json:"diary_id"`
    UserID        int64     `json:"user_id"`
    Title         string    `json:"title"`
    CategoryID    int64     `json:"category_id"`
    Latitude      float64   `json:"latitude"`
    Longitude     float64   `json:"longitude"`
    Text          string    `json:"text"`
    CorrectedText string    `json:"corrected_text"`
    CreatedAt     time.Time `json:"created_at"`
}

type DiaryRepository interface {
    Create(d *Diary) (*Diary, error)
    GetByID(id int64) (*Diary, error)   //特定の日記の取得
    GetAll(userID int64) ([]*Diary, error)  //ユーザーの日記一覧取得
    Update(diary *Diary) error
    Delete(id int64) error
}


