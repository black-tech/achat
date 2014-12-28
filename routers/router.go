package routers

import (
	"achat/controllers"
	"achat/models"
	"github.com/astaxie/beego"
	"log"
	"net/websocket"
)

func init() {

	ch := models.NewChatHome()
	ch.InitGlobal()
	// connid = 0

	beego.Handler("/ws_chat", websocket.Handler(models.ChatroomServer))

	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/chat_home", &controllers.ChatController{})

	log.Println("Router init Over")
}

// var connid int
// var WSConns *list.List

// func SendMessage(self *list.Element, data string) {
// 	log.Println("WS:STart")
// 	for item := WSConns.Front(); item != nil; item = item.Next() {
// 		ws, ok := item.Value.(*websocket.Conn)
// 		if !ok {
// 			panic("item not *websocket.Conn")
// 		}
// 		if item == self {
// 			continue
// 		}
// 		io.WriteString(ws, data)
// 	}
// }

// func ChatroomServer(ws *websocket.Conn) {
// 	defer ws.Close()

// 	connid++
// 	fmt.Println(connid)

// 	item := WSConns.PushBack(ws)
// 	defer WSConns.Remove(item)

// 	name := fmt.Sprintf("user%d", connid)
// 	SendMessage(nil, fmt.Sprintf("welcome %s join\n", name))

// 	r := bufio.NewReader(ws)

// 	for {
// 		data, err := r.ReadBytes('\n')
// 		if err != nil {
// 			fmt.Printf("disconnected id: %d\n", connid)
// 			SendMessage(item, fmt.Sprintf("%s offline\n", name))
// 			break
// 		}
// 		fmt.Printf("%s: %s", name, data)
// 		SendMessage(nil, fmt.Sprintf("%s\t> %s", name, data))
// 	}
// }
