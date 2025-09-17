package main

import (
	"fmt"
	"log"
	"sabj/db"
	"sabj/internal/controller"
	"sabj/internal/domain"
	"sabj/internal/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("🚀 アプリケーション開始")
	
	e := echo.New()
	fmt.Println("✅ Echo インスタンス作成完了")

	e.Static("/", "web")
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})
	fmt.Println("✅ 基本ルート設定完了")

	fmt.Println("📊 DB接続を試行中...")
	diaryRepo, err := db.NewDiaryRepository()
	if err != nil {
		fmt.Printf("❌ DB接続エラー: %v\n", err)
		log.Fatal(err)
	}
	fmt.Println("✅ DB接続成功")

	fmt.Println("🔧 UseCase作成中...")
	diaryUC := &usecase.DiaryUsecase{Repo: diaryRepo}
	fmt.Println("✅ UseCase作成完了")

	fmt.Println("🎮 Controller作成中...")
	diaryCtrl := &controller.DiaryController{UC: diaryUC}
	fmt.Println("✅ Controller作成完了")

	fmt.Println("🛣️ APIルート設定中...")
	e.POST("/api/diary", diaryCtrl.CreateDiary)
	fmt.Println("✅ API ルート設定完了")

	fmt.Println("🧪 テストデータ作成中...")
	newDiary := &domain.Diary{
		UserID: 1,
		Title:  "初めての日記",
		Text:   "今日は Go アプリを動かせた！",
	}

	created, err := diaryRepo.Create(newDiary)
	if err != nil {
		fmt.Printf("❌ テストデータ作成エラー: %v\n", err)
		log.Fatal(err)
	}
	fmt.Printf("✅ テストデータ作成成功: %+v\n", created)

	fmt.Println("📋 登録されたルート一覧:")
	for _, r := range e.Routes() {
		fmt.Printf("  %s %s\n", r.Method, r.Path)
	}

	fmt.Println("🎯 サーバーを :8080 で起動します...")
	e.Logger.Fatal(e.Start(":8080"))
}
// package main

// import (
// 	"fmt"
// 	"log"
// 	"sabj/internal/controller"
// 	"sabj/internal/domain"
// 	"sabj/internal/usecase"
// 	"sabj/db"

// 	"github.com/labstack/echo/v4"
// )

// func main() {
//     e := echo.New()

// 	e.Static("/", "web")
// 	e.GET("/health", func(c echo.Context) error {
//     return c.JSON(200, map[string]string{"status": "ok"})
// 	})

// 	diaryRepo, err := db.NewDiaryRepository()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

//     diaryUC := &usecase.DiaryUsecase{Repo: diaryRepo}
//     diaryCtrl := &controller.DiaryController{UC: diaryUC}


// 	e.POST("/api/diary", diaryCtrl.CreateDiary)

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

   


// 	for _, r := range e.Routes() {
//     fmt.Println(r.Method, r.Path)
// }

//     e.Logger.Fatal(e.Start(":8080"))
// }


