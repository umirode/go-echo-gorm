package v1

import (
	"github.com/googollee/go-socket.io"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/umirode/go-rest/Module/Http/Controller"
)

type SocketIOController struct {
	Controller.BaseController
	Controller.SocketIOControllerInterface
}

func NewSocketIOController() *SocketIOController {
	controller := &SocketIOController{}

	return controller
}

func (*SocketIOController) OnConnect(context echo.Context, s socketio.Conn) error {
	logrus.Println("connected:", s.ID())
	return nil
}

func (*SocketIOController) OnError(context echo.Context, e error) {
	logrus.Println("meet error:", e)
}

func (*SocketIOController) OnDisconnect(context echo.Context, s socketio.Conn, msg string) {
	logrus.Println("closed", msg)
}

func (*SocketIOController) OnEventTest(context echo.Context, s socketio.Conn, msg string) {
	logrus.Println("notice:", msg)
	s.Emit("test", msg)
}
