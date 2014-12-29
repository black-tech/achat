package models

import (
	"bufio"
	"container/list"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"github.com/hoisie/redis"
	"io"
	"log"
	"net/websocket"
	"strconv"
)

var WSConns list.List

type ChatHome struct {
	MAX_ONLINE_COUNT int
	rls_online       string
	rls_waitting     string
	// waitting_count   int
	// online_users     *list.List
	// waitting_users   *list.List
	rediscli redis.Client
}
type Msg struct {
	Code string
	Data string
	Desc string
}

func SendMessage(self *list.Element, data string) {
	for item := WSConns.Front(); item != nil; item = item.Next() {
		ws, ok := item.Value.(*websocket.Conn)
		if !ok {
			panic("item not *websocket.Conn")
		}
		if item == self {
			continue
		}
		io.WriteString(ws, data)
	}
}

func ChatroomServer(ws *websocket.Conn) {
	defer ws.Close()
	auth := false
	var user *User
	ch := NewChatHome()
	// user.DoNothing()

	item_ws := WSConns.PushBack(ws)
	defer WSConns.Remove(item_ws)

	// SendMessage(nil, fmt.Sprintf("welcome %s join\n", "name"))

	r := bufio.NewReader(ws)
	log.Println("Connected ")
	for {
		data, err := r.ReadBytes('\n')
		if err != nil {
			SendMessage(item_ws, `{"Code":"msg","Data":"`+user.Name+` offline"}`)
			if b, _ := ch.IsOnline(user.Name); b {
				ch.RmOnlineUser(user.Name)
				ch.WaittingToOnline()
				SendMessage(item_ws, `{"Code":"OK","Data":"Welcome `+user.Name+`"}`)
			} else {
				ch.RmWaittingUser(user.Name)
			}
			SendMessage(nil, `{"Code":"online_user_count","Data": `+strconv.Itoa(ch.GetOnlineCount())+`}`)
			SendMessage(nil, `{"Code":"waitting_user_count","Data": `+strconv.Itoa(ch.GetWaittingCount())+`}`)
			// SendMessage(nil, `{"Code":"online_user_list","Data": "`+ch.OnlineListToBase64()+`"}`)
			// SendMessage(nil, `{"Code":"waitting_user_list","Data": "`+ch.WaittingListToBase64()+`"}`)
			break
		}
		log.Println("Received: " + string(data))
		log.Println(`{"Code":"online_user_list","Data": "` + ch.OnlineListToBase64() + `"}`)
		var msg Msg
		err = json.Unmarshal(data, &msg)
		if err != nil {
			SendMessage(item_ws, `{"Code":"error","Data":"JSON Error","Desc":"Nothing"}`)
		}
		if msg.Code == "cookie" {
			for item_user := ConnUsers.Front(); item_user != nil; item_user = item_user.Next() {
				user, _ = item_user.Value.(*User)
				if user.Cookie == msg.Data {
					auth = true
					if ch.GetOnlineCount() < ch.MAX_ONLINE_COUNT && ch.GetWaittingCount() == 0 {
						ch.AddOnlineUser(user.Name)
						SendMessage(item_ws, `{"Code":"OK","Data":"Welcome `+user.Name+`"}`)
					} else {
						ch.AddWaittingUser(user.Name)
						SendMessage(nil, `{"Code":"wait","Data": `+strconv.Itoa(ch.GetWaittingCount())+`}`)
					}
					SendMessage(nil, `{"Code":"online_user_count","Data": `+strconv.Itoa(ch.GetOnlineCount())+`}`)
					SendMessage(nil, `{"Code":"waitting_user_count","Data": `+strconv.Itoa(ch.GetWaittingCount())+`}`)
					// SendMessage(nil, `{"Code":"online_user_list","Data": "`+ch.OnlineListToBase64()+`"}`)
					// SendMessage(nil, `{"Code":"waitting_user_list","Data": "`+ch.WaittingListToBase64()+`"}`)

					break
				}
			}
		} else {
			if !auth {
				SendMessage(item_ws, `{"Code":"error","Data":"Auth Error","Desc":"RELOGIN"}`)
			} else {
				// SendMessage(item_ws, `{"Code":"OK","Data":"Auth Error","Desc":"RELOGIN"}`)
				ControlMsg(item_ws, msg, user)
			}
		}
	}
	log.Println("Listenning Over")
}
func ControlMsg(item *list.Element, msg Msg, user *User) {
	switch msg.Code {
	case "msg":
		SendMessage(item, `{"Code":"msg","Data":"`+user.Name+"> "+msg.Data+`"}`)
	case "":
	}
}

func NewChatHome() (ch *ChatHome) {
	ch = &ChatHome{}
	ch.MAX_ONLINE_COUNT, _ = strconv.Atoi(beego.AppConfig.String("max_online"))
	ch.rls_online = "online_list"
	ch.rls_waitting = "waitting_list"
	ch.rediscli.Addr = beego.AppConfig.String("redis_addr")
	if pw := beego.AppConfig.String("redis_auth"); "" != pw {
		ch.rediscli.Auth(pw)
	}
	return
}
func (ch *ChatHome) InitGlobal() {
	ch.rediscli.Del(ch.rls_online)
	ch.rediscli.Del(ch.rls_waitting)
}
func (ch *ChatHome) GetOnlineCount() int {
	v, err := ch.rediscli.Llen(ch.rls_online)
	if err == nil {
		return v
	}
	return 0
}
func (ch *ChatHome) GetWaittingCount() int {
	v, err := ch.rediscli.Llen(ch.rls_waitting)
	if err == nil {
		return v
	}
	return 0
}
func (ch *ChatHome) RmOnlineUser(uname string) (err error) {
	_, err = ch.rediscli.Lrem(ch.rls_online, 0, []byte(uname))
	return
}
func (ch *ChatHome) RmWaittingUser(uname string) (err error) {
	_, err = ch.rediscli.Lrem(ch.rls_waitting, 0, []byte(uname))
	return
}
func (ch *ChatHome) IsLogin(uname string) (bool, error) {
	v, err := ch.rediscli.Lrange(ch.rls_online, 0, -1)
	for _, k := range v {
		if uname == string(k) {
			return true, err
		}
	}
	v, err = ch.rediscli.Lrange(ch.rls_waitting, 0, -1)
	for _, k := range v {
		if uname == string(k) {
			return true, err
		}
	}
	return false, err
}
func (ch *ChatHome) IsOnline(uname string) (bool, error) {
	v, err := ch.rediscli.Lrange(ch.rls_online, 0, -1)
	for _, k := range v {
		if uname == string(k) {
			return true, err
		}
	}
	return false, err
}
func (ch *ChatHome) IsWaitting(uname string) (bool, error) {
	v, err := ch.rediscli.Lrange(ch.rls_waitting, 0, -1)
	for _, k := range v {
		if uname == string(k) {
			return true, err
		}
	}
	return false, err
}
func (ch *ChatHome) AddWaittingUser(uname string) error {
	if b, err := ch.IsLogin(uname); b {
		if err == nil {
			err = errors.New(uname + " logined already")
		}
		return err
	}
	err := ch.rediscli.Lpush(ch.rls_waitting, []byte(uname))
	return err
}
func (ch *ChatHome) AddOnlineUser(uname string) error {
	if b, err := ch.IsOnline(uname); b {
		if err == nil {
			err = errors.New(uname + " is ONLINE already")
		}
		return err
	}
	err := ch.rediscli.Lpush(ch.rls_online, []byte(uname))
	return err
}

func (ch *ChatHome) WaittingToOnline() (string, error) {
	v, err := ch.rediscli.Rpop(ch.rls_waitting)
	if err == nil {
		ch.AddOnlineUser(string(v))
	}
	return string(v), err
}

func (ch *ChatHome) GetOnlineList() []string {
	bl, _ := ch.rediscli.Lrange(ch.rls_online, 0, -1)
	var sl []string
	for i, v := range bl {
		sl[i] = string(v)
	}
	return sl
}

func (ch *ChatHome) GetWaittingList() []string {
	bl, _ := ch.rediscli.Lrange(ch.rls_waitting, 0, -1)
	var sl []string
	for i, v := range bl {
		sl[i] = string(v)
	}
	return sl
}

func (ch *ChatHome) OnlineListToBase64() string {
	sl := ch.GetOnlineList()
	s := ""
	for i, v := range sl {
		s += v + "(" + strconv.Itoa(i) + "); "
	}
	return base64.URLEncoding.EncodeToString([]byte(s))
}

func (ch *ChatHome) WaittingListToBase64() string {
	sl := ch.GetWaittingList()
	s := ""
	for i, v := range sl {
		s += v + "(" + strconv.Itoa(i) + "); "
	}
	return base64.URLEncoding.EncodeToString([]byte(s))
}
