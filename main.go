package main

import (
	"context"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/wing8169/golang-htmx-chat-app/templates"
)

var (
	upgrader = websocket.Upgrader{}
)

func main() {
	e := echo.New()
	manager := NewManager()
	go manager.HandleClientListEventChannel(context.Background())
	e.GET("/", func(c echo.Context) error {
		component := templates.Index()
		return component.Render(context.Background(), c.Response().Writer)
	})
	e.GET("/ws/chat", manager.Handle)

	e.Logger.Fatal(e.Start(":3000"))
}
