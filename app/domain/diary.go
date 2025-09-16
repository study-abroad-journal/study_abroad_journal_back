package domain

import "time"

type Diary struct {
    ID            int64     `json:"id"`
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
    Save(diary *Diary) error
    // FindByID(id int64) (*Diary, error)
    // FindAll(userID int64) ([]*Diary, error)
	Create(d *Diary) error
    // Update(diary *Diary) error
    // Delete(id int64) error
}


