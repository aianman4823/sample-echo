package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/aianman4823/sample-echo/handler/handler.go"
)

func main() {
	// Echoのインスタンスを作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア(ログとか)はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/", handler.hello)

	// サーバー起動
	e.Start(":1323")
}

