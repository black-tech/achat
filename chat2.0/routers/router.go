package routers

import (
	"funx/controllers"
	// "funx/models"
	"github.com/astaxie/beego"
	// "net/websocket"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/home", &controllers.MainController{})

	beego.Router("/info", &controllers.InfoController{})
	beego.Router("/chat", &controllers.ChatController{})

	// beego.Router("/chat_websocket", websocket.Handler(models.ChatroomServer))
	// beego.Router("/b", &controllers.BController{})
}
