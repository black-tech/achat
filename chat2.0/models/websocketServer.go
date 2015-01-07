package models

import (
	"bufio"
	"container/list"
	"io"
	"log"
	"net/websocket"
)

var WSChatClients list.List

type WSChatClient struct {
	ws      *websocket.Conn
	item_ws *list.Element
}

func NewWSChatClient(ws *websocket.Conn) (w *WSChatClient) {
	w = &WSChatClient{}
	w.ws = ws
	w.item_ws = WSChatClients.PushBack(w)
	return
}
func (w *WSChatClient) Close() {
	WSChatClients.PushBack(w.item_ws)
	w.ws.Close()
}
func (w *WSChatClient) SendMessage(self *list.Element, data string) {
	for item := WSChatClients.Front(); item != nil; item = item.Next() {
		wc, ok := item.Value.(*WSChatClient)
		if !ok {
			panic("item not *websocket.Conn")
		}
		if item == self {
			continue
		}
		io.WriteString(wc.ws, data)
	}
}

// func SendMessage(self *list.Element, data string) {
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

func ChatroomServer(ws *websocket.Conn) {
	w := NewWSChatClient(ws)
	defer w.Close()
	// auth := false
	// ch := NewChatHome()
	// user.DoNothing()
	// item_ws := WSConns.PushBack(ws)
	// defer WSConns.Remove(item_ws)

	// SendMessage(nil, fmt.Sprintf("welcome %s join\n", "name"))

	r := bufio.NewReader(ws)
	log.Println("Connected ")
	for {
		data, err := r.ReadBytes('\n')
		if err != nil {
			// if b, _ := ch.IsOnline("user.Name"); b {
			// 	ch.RmOnlineUser("user.Name")
			// 	ch.WaittingToOnline()
			// } else {
			// 	ch.RmWaittingUser("user.Name")
			// }
			break
		}
		log.Println("Received: " + string(data))
		// log.Println(`{"Code":"online_user_list","Data": "` + ch.OnlineListToBase64() + `"}`)
		// var msg Msg
		// err = json.Unmarshal(data, &msg)
		// if err != nil {
		// 	SendMessage(item_ws, `{"Code":"error","Data":"JSON Error","Desc":"Nothing"}`)
		// }
	}
	log.Println("Listenning Over")
}
