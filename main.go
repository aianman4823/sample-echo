package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"local.packages/handler"
)

func main() {
	// Echoのインスタンスを作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア(ログとか)はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/", handler.Hello())

	// サーバー起動
	e.Start(":1323")
}

