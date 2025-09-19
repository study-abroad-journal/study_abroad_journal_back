package controller

import (
	"fmt"
	"net/http"
	"sabj/internal/domain"
	"sabj/internal/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DiaryController struct {
    UC *usecase.DiaryUsecase
}

// CreateDiary - 日記作成（既存）
func (c *DiaryController) CreateDiary(ctx echo.Context) error {
    var diary domain.Diary
    if err := ctx.Bind(&diary); err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
    }

    if diary.Text == "" {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "text is required"})
    }

    if err := c.UC.CreateDiary(&diary); err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }

    return ctx.JSON(http.StatusCreated, diary)
}

// GetDiary - 特定の日記を取得
func (c *DiaryController) GetDiary(ctx echo.Context) error {
    // パスパラメータからIDを取得
    idParam := ctx.Param("id")
    id, err := strconv.ParseInt(idParam, 10, 64)
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid diary id"})
    }

    // TODO: 認証実装後はJWTからuserIDを取得
    // 現在は仮でuser_id=1とする
    userID := int64(1)

    diary, err := c.UC.GetDiary(id, userID)
    if err != nil {
        if err.Error() == "access denied: you can only view your own diaries" {
            return ctx.JSON(http.StatusForbidden, map[string]string{"error": err.Error()})
        }
        return ctx.JSON(http.StatusNotFound, map[string]string{"error": "diary not found"})
    }

    return ctx.JSON(http.StatusOK, diary)
}

// GetAllDiaries - ユーザーの全日記を取得
func (c *DiaryController) GetAllDiaries(ctx echo.Context) error {
    // TODO: 認証実装後はJWTからuserIDを取得
    // 現在は仮でuser_id=1とする
    userID := int64(1)

    diaries, err := c.UC.GetAllDiaries(userID)
    if err != nil {
        return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to get diaries"})
    }

    return ctx.JSON(http.StatusOK, map[string]interface{}{
        "diaries": diaries,
        "count":   len(diaries),
    })
}

// UpdateDiary - 日記を更新
func (c *DiaryController) UpdateDiary(ctx echo.Context) error {
    // デバッグ用
    fmt.Println("UpdateDiary called")
    
    // パスパラメータからIDを取得
    idParam := ctx.Param("id")
    id, err := strconv.ParseInt(idParam, 10, 64)
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid diary id"})
    }

    var updatedDiary domain.Diary
    if err := ctx.Bind(&updatedDiary); err != nil {
        fmt.Println("Bind error:", err)  // ← これが原因の可能性
        return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
    }

    fmt.Printf("Bound diary: %+v\n", updatedDiary)

    if updatedDiary.Text == "" {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "text is required"})
    }

    // TODO: 認証実装後はJWTからuserIDを取得
    userID := int64(1)

    err = c.UC.UpdateDiary(id, &updatedDiary, userID)
    if err != nil {
        if err.Error() == "access denied: you can only edit your own diaries" {
            return ctx.JSON(http.StatusForbidden, map[string]string{"error": err.Error()})
        }
        return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update diary"})
    }

    return ctx.JSON(http.StatusOK, updatedDiary)
}

// DeleteDiary - 日記を削除
func (c *DiaryController) DeleteDiary(ctx echo.Context) error {
    // パスパラメータからIDを取得
    idParam := ctx.Param("id")
    id, err := strconv.ParseInt(idParam, 10, 64)
    if err != nil {
        return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid diary id"})
    }

    // TODO: 認証実装後はJWTからuserIDを取得
    userID := int64(1)

    err = c.UC.DeleteDiary(id, userID)
    if err != nil {
        if err.Error() == "access denied: you can only delete your own diaries" {
            return ctx.JSON(http.StatusForbidden, map[string]string{"error": err.Error()})
        }
        return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete diary"})
    }

    return ctx.JSON(http.StatusOK, map[string]string{"message": "diary deleted successfully"})
}