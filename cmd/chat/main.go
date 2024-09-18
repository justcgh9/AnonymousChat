package main

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/justcgh9/AnonymousChat/database/sqlite"
	"github.com/justcgh9/AnonymousChat/internal/handler/ws"
	httpH "github.com/justcgh9/AnonymousChat/internal/handler/http"
	messageRepo "github.com/justcgh9/AnonymousChat/internal/repo/message"
	messageService "github.com/justcgh9/AnonymousChat/internal/service/message"
)

func main() {
	e := echo.New()
	db := sqlite.NewConn()
	msgR := messageRepo.NewRepo(db)
	msgS := messageService.NewService(msgR)

	msgH := httpH.NewHandler(msgS)
	wsMsgH := ws.NewHandler(msgS)

	e.GET("/ping", func(c echo.Context) error {
		c.Response().Write([]byte("asdasdadasd"))
		return nil
	})

	fs := http.FileServer(http.Dir("web"))

	e.GET("/", func(c echo.Context) error {
		fs.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/messages", func(c echo.Context) error {
		return msgH.HandleGetMessages(c.Response(), c.Request())
	})

	e.GET("/messages/count", func(c echo.Context) error {
		return msgH.HandleGetMessagesCount(c.Response(), c.Request())
	})

	e.GET("/ws/chat", func(c echo.Context) error {
		return wsMsgH.HandleConnection(c.Response(), c.Request())
	})

	err := e.Start("0.0.0.0:8000")
	if err != nil {
		slog.Error(err.Error())
	}
}
