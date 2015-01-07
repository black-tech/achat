package models

import (
	"container/list"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/hoisie/redis"
	"log"
	"math/big"
	// "net/websocket"
	"strconv"
	"strings"
	"time"
)

type User struct {
	Id       int
	Name     string
	Pw       string
	Salt     string
	Cookie   string
	rediscli redis.Client
	// Ws       websocket.Conn
}

var ConnUsers list.List

func NewUser(username string, password string) (u *User, err error) {

	u = &User{}
	u.Name = username
	u.Pw = password
	u.rediscli.Addr = beego.AppConfig.String("redis_addr")
	if authPW := beego.AppConfig.String("redis_auth"); "" != u.Pw {
		u.rediscli.Auth(authPW)
	}
	var userps []byte
	userps, err = u.rediscli.Hget("userlist", u.Name)
	if err != nil || u.Pw != string(userps) {
		log.Println(u.Pw)
		err = errors.New("Username Or Password Error")
	}
	return
}
func (u *User) LoginSucc() {
	u.GetNewSalt()
	vs := base64.URLEncoding.EncodeToString([]byte(u.Pw))
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	h := hmac.New(sha1.New, []byte(u.Salt))
	fmt.Fprintf(h, "%s%s", vs, timestamp)
	sig := fmt.Sprintf("%02x", h.Sum(nil))
	u.Cookie = strings.Join([]string{vs, timestamp, sig}, "|")

	ConnUsers.PushBack(u)

}
func (u *User) GetNewSalt() {
	rnd, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		log.Printf("GetNewSalt:rand.Int() error : %v \n", err)
		rnd = big.NewInt(4369)
	}
	u.Salt = fmt.Sprintln(rnd)
}
func (u *User) DoNothing() {

}
