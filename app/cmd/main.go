package main

import (
	"sabj/app/controller"
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

	
    diaryRepo := db.NewDiaryRepository()      // ← DB実装
    diaryUC := &usecase.DiaryUsecase{Repo: diaryRepo}
    diaryCtrl := &controller.DiaryController{UC: diaryUC}

    e.POST("/api/diary", diaryCtrl.CreateDiary)

    e.Logger.Fatal(e.Start(":8080"))
}


