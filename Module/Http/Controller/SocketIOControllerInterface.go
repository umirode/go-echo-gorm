package Controller

import (
	"github.com/googollee/go-socket.io"
	"github.com/labstack/echo"
)

type SocketIOControllerInterface interface {
	OnConnect(context echo.Context, s socketio.Conn) error
	OnError(context echo.Context, e error)
	OnDisconnect(context echo.Context, s socketio.Conn, msg string)
}
