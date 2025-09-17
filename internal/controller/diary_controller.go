package controller

import (
	"net/http"
	"sabj/internal/domain"
	"sabj/internal/usecase"

	"github.com/labstack/echo/v4"
)

type DiaryController struct {
    UC *usecase.DiaryUsecase
}

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
