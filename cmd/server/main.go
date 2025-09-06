package main

import (
	"log"
	"net/http"

	"go-di/internal/di"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// DIコンテナの初期化
	container := di.NewContainer()

	// Echoインスタンスの作成
	e := echo.New()

	// ミドルウェアの設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// ルーティングの設定
	setupRoutes(e, container)

	// サーバーの起動
	log.Println("Server starting on :8080")
	if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatal("Failed to start server:", err)
	}
}

// setupRoutes はルートを設定します
func setupRoutes(e *echo.Echo, container *di.Container) {
	// ヘルスチェック
	e.GET("/health", container.GetUserHandler().HealthCheck)

	// API v1 グループ
	v1 := e.Group("/api/v1")

	// ユーザー関連のルート
	users := v1.Group("/users")
	users.POST("", container.GetUserHandler().CreateUser)
	users.GET("", container.GetUserHandler().GetAllUsers)
	users.GET("/:id", container.GetUserHandler().GetUser)
	users.PUT("/:id", container.GetUserHandler().UpdateUser)
	users.DELETE("/:id", container.GetUserHandler().DeleteUser)
}
