// package main

// import (
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// func main() {
// 	e := echo.New()

// 	// Health check
// 	e.GET("/health", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "OK from Go backend!")
// 	})

// 	// サーバ起動
// 	e.Logger.Fatal(e.Start(":8080"))
// }
