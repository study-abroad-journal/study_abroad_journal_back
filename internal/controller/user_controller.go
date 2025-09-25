package controller

import (
	"net/http"
	"sabj/internal/domain"
	"sabj/internal/usecase" // あなたのプロジェクトのusecaseパッケージ

	"github.com/labstack/echo/v4"
)

// UserController はユーザー関連のコントローラーです
type UserController struct {
	UC *usecase.UserUsecase
}

// CreateUser は新しいユーザーを作成するハンドラーです
func (c *UserController) CreateUser(ctx echo.Context) error {
	// 1. リクエストボディを格納するための変数を定義
	var user domain.User

	// 2. リクエストボディのJSONを user 変数にバインド（割り当て）
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	// 3. バリデーション：必須項目が空でないかチェック
	if user.User_id == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "user ID is required"})
	}
	if user.Name == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "name is required"})
	}
	if user.Email == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":         "email is required",
			"received_data": user, // user変数の内容を追加
		})
	}

	// 4. UseCaseのメソッドを呼び出し、ユーザー作成処理を依頼
	if err := c.UC.CreateUser(&user); err != nil {
		// e.g., メールアドレスの重複などでエラーが起きた場合
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// 5. 成功レスポンスを返す
	return ctx.JSON(http.StatusCreated, user)
}
