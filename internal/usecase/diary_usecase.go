package usecase

import (
	"errors"
	"fmt"
	"sabj/internal/domain"
	"time"
)

type DiaryUsecase struct {
    Repo domain.DiaryRepository
}

func (u *DiaryUsecase) CreateDiary(diary *domain.Diary) error {
    diary.CreatedAt = time.Now()  // 保存時に自動付与
    _, err := u.Repo.Create(diary)
    return err
}

func (u *DiaryUsecase) GetDiary(id int64, userID int64) (*domain.Diary, error) {
    diary, err := u.Repo.GetByID(id)
    if err != nil {
        return nil, err
    }

    if diary.UserID != userID {
         return nil, errors.New("access denied: you can only view your own diaries")
    }
    
    return diary, nil
}


func (u *DiaryUsecase) GetAllDiaries(userID int64) ([]*domain.Diary, error) {
    return u.Repo.GetAll(userID)
}

func (u *DiaryUsecase) UpdateDiary(id int64, updatedDiary *domain.Diary, userID int64) error {
    // まず既存の日記を取得して所有者チェック
    existingDiary, err := u.Repo.GetByID(id)
    if err != nil {
        fmt.Printf("GetByID error: %v\n", err)  // ← デバッグログ
        return err
    }
     fmt.Printf("Existing diary found: %+v\n", existingDiary)

    // ビジネスロジック：他人の日記は編集できない
    if existingDiary.UserID != userID {
        fmt.Printf("Access denied: existingDiary.UserID=%d, userID=%d\n", existingDiary.UserID, userID)
        return errors.New("access denied: you can only edit your own diaries")
    }
    
    // ビジネスロジック：IDとUserIDは変更させない
    updatedDiary.DiaryID = id
    updatedDiary.UserID = userID
    updatedDiary.CreatedAt = existingDiary.CreatedAt // 作成日時は保持
    
    fmt.Printf("About to update diary: %+v\n", updatedDiary)
    
    return u.Repo.Update(updatedDiary)
}

// DeleteDiary - 日記を削除（ビジネスロジック：所有者チェック）
func (u *DiaryUsecase) DeleteDiary(id int64, userID int64) error {
    // まず既存の日記を取得して所有者チェック
    existingDiary, err := u.Repo.GetByID(id)
    if err != nil {
        return err
    }
    
    // ビジネスロジック：他人の日記は削除できない
    if existingDiary.UserID != userID {
        return errors.New("access denied: you can only delete your own diaries")
    }
    
    return u.Repo.Delete(id)
}      

