package main

import (
	"fmt"
	"log"
	"sabj/db"
	"sabj/internal/controller"
	"sabj/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// --- CORSミドルウェアの設定を追加 ---
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"}, // フロントエンドのURLを許可
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization, // Authorization ヘッダーを許可
		},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	// ------------------------------------
	// 	fmt.Println("Echo インスタンス作成完了")

	e.Static("/", "web")
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})
	//fmt.Println("基本ルート設定完了")

	diaryRepo, err := db.NewDiaryRepository()
	if err != nil {
		//fmt.Printf("DB接続エラー: %v\n", err)
		log.Fatal(err)
	}
	//fmt.Println("DB接続成功")

	diaryUC := &usecase.DiaryUsecase{Repo: diaryRepo}
	//fmt.Println("UseCase作成完了")

	diaryCtrl := &controller.DiaryController{UC: diaryUC}
	fmt.Println("Controller作成完了")

	e.POST("/api/diary", diaryCtrl.CreateDiary)
	fmt.Println("API ルート設定完了")

	e.GET("/api/diary/:id", diaryCtrl.GetDiary)       // 特定の日記取得
	e.GET("/api/diaries", diaryCtrl.GetAllDiaries)    // 全日記取得
	e.PUT("/api/diary/:id", diaryCtrl.UpdateDiary)    // 日記更新
	e.DELETE("/api/diary/:id", diaryCtrl.DeleteDiary) // 日記削除

	// 	newDiary := &domain.Diary{
	// 		UserID: 1,
	// 		Title:  "初めての日記",
	// 		Text:   "今日は Go アプリを動かせた！",
	// 	}

	// 	created, err := diaryRepo.Create(newDiary)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Printf("Created diary: %+v\n", created)

	//fmt.Println("登録されたルート一覧:")
	for _, r := range e.Routes() {
		fmt.Printf("  %s %s\n", r.Method, r.Path)
	}
	// サーバ起動
	e.Logger.Fatal(e.Start(":8080"))
}
