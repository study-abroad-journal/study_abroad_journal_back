package main

import (
	"fmt"
	"log"
	"sabj/app/controller"
	"sabj/app/domain"
	"sabj/app/usecase"
	"sabj/db"

	"github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

	e.Static("/", "web")
	e.GET("/health", func(c echo.Context) error {
    return c.JSON(200, map[string]string{"status": "ok"})
	})

	diaryRepo, err := db.NewDiaryRepository()
	if err != nil {
		log.Fatal(err)
	}

    diaryUC := &usecase.DiaryUsecase{Repo: diaryRepo}
    diaryCtrl := &controller.DiaryController{UC: diaryUC}



	newDiary := &domain.Diary{
		UserID: 1,
		Title:  "初めての日記",
		Text:   "今日は Go アプリを動かせた！",
	}

	created, err := diaryRepo.Create(newDiary)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created diary: %+v\n", created)

    e.POST("/api/diary", diaryCtrl.CreateDiary)

    e.Logger.Fatal(e.Start(":8080"))
}


