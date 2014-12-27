package models

import (
	// "container/list"
	"errors"
	"github.com/astaxie/beego"
	"github.com/hoisie/redis"
	_ "log"
)

type ChatHome struct {
	MAX_ONLINE_COUNT int
	rls_online       string
	rls_waitting     string
	// waitting_count   int
	// online_users     *list.List
	// waitting_users   *list.List
	rediscli redis.Client
}

func NewChatHome() (ch *ChatHome) {
	ch = &ChatHome{}
	ch.MAX_ONLINE_COUNT = 10
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

// func (ch *ChatHome) WaittingToOnline() (string, error) {
// 	v, err := ch.rediscli.Lpop(ch.rls_waitting)
// 	if err == nil {
// 		n, _ := strconv.Atoi(string(v))
// 		return n
// 	}
// 	return '0'
// }
