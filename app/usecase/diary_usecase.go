package usecase

import (
	"sabj/app/domain"
	"time"
)

type DiaryUsecase struct {
    Repo domain.DiaryRepository
}

func (u *DiaryUsecase) CreateDiary(diary *domain.Diary) error {
    diary.CreatedAt = time.Now()  // 保存時に自動付与
    return u.Repo.Save(diary)
}
