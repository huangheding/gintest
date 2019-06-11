package ws

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

// User Online
type user struct {
	sockets         map[string]*websocket.Conn // 在线信息
	chanUserOnline  chan *online               // 用户上线
	chanUserOffline chan string                // 用户下线
	chanPush        chan *contents             // 推消息
	chanIsOnline    chan *isOnline             // 是否在线
}

type online struct {
	UID    string
	Socket *websocket.Conn
}
type isOnline struct {
	UIDs []string
	Resp chan map[string]bool
}
type contents struct {
	UID string `json:"uid"`
	// Event   string `json:"event"`
	Content string `json:"content"`
}

//ws client
type Client struct {
	conn *websocket.Conn
}

var (
	once          sync.Once
	singleton     *user
	ServiceOnline OnlineService
)

func InitWs() {
	once.Do(func() {
		singleton = &user{
			sockets:         make(map[string]*websocket.Conn),
			chanUserOnline:  make(chan *online),
			chanUserOffline: make(chan string),
			chanPush:        make(chan *contents),
			chanIsOnline:    make(chan *isOnline),
		}
		singleton.Run()
	})
	//全局变量
	ServiceOnline = singleton
}

func (s *user) Run() {
	go func() {
		for {
			select {
			case onlineUser := <-s.chanUserOnline:
				s.processUserOnline(onlineUser)
			case uid := <-s.chanUserOffline:
				s.processUserOffline(uid)
			case req := <-s.chanIsOnline:
				s.processIsOnline(req)
			case msg := <-s.chanPush:
				s.processPush(msg)
			}
		}
	}()
}

func ServeWs(c *gin.Context) {
	conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if error != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	client := &Client{conn: conn}
	go client.connPump()
	go client.readPump()
}

func (c *Client) readPump() {
	var uid string
	defer func() {
		fmt.Println("readPump stop,", c.conn)
		ServiceOnline.UserOffline(uid) //注销
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait)) //读取连接有效期延时，超时失效返回err
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	b := true
	for {
		_, message, err := c.conn.ReadMessage() //读取到信息
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("client:readPump:error: ", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, []byte{'\n'}, []byte{' '}, -1))
		msg := contents{}
		err2 := json.Unmarshal(message, &msg)
		if err2 != nil {
			log.Println(err2)
			continue
		}
		if b { //注册
			uid = msg.UID
			ServiceOnline.UserOnline(uid, c.conn)
			b = false
		}
	}
}

func (c *Client) connPump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		fmt.Println("connPump stop", c.conn)
		ticker.Stop()
		c.conn.Close() //下线关闭连接
	}()
	for {
		select {
		case <-ticker.C: //定时器启动
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

/******** ws channel 编辑********/
// 上线
func (s *user) UserOnline(uid string, socket *websocket.Conn) {
	s.chanUserOnline <- &online{UID: uid, Socket: socket}
}

// 下线
func (s *user) UserOffline(uid string) {
	s.chanUserOffline <- uid
}

// 全部在线人员
func (s *user) UserOnlineList() map[string]bool {
	req := &isOnline{UIDs: []string{}, Resp: make(chan map[string]bool)}
	s.chanIsOnline <- req
	return <-req.Resp
}

// IsUserOnline 在线判断
func (s *user) IsUserOnline(uid string) bool {
	req := &isOnline{UIDs: []string{uid}, Resp: make(chan map[string]bool)}
	s.chanIsOnline <- req
	res := <-req.Resp
	return res[uid]
}

// 推送消息-单发
func (s *user) Push(uid string, content string) {
	s.chanPush <- &contents{UID: uid, Content: content}
}

// 推送消息-群发
func (s *user) PushAll(content string) {
	s.chanPush <- &contents{Content: content}
}

/*******底层ws操作*********/
//添加上线连接
func (s *user) processUserOnline(onlineUser *online) {
	s.sockets[onlineUser.UID] = onlineUser.Socket
}

//移除上线连接
func (s *user) processUserOffline(uid string) {
	delete(s.sockets, uid)
}

//判断连接是否存在
func (s *user) processIsOnline(req *isOnline) {
	result := make(map[string]bool)
	if len(req.UIDs) > 0 {
		for _, uid := range req.UIDs {
			_, ok := s.sockets[uid]
			result[uid] = ok
		}
	} else {
		for k := range s.sockets {
			result[k] = true
		}
	}

	req.Resp <- result
}

//推送消息
func (s *user) processPush(msg *contents) {
	b, e := json.Marshal(msg)
	if e != nil {
		return
	}
	if len(msg.UID) > 0 {
		socket, ok := s.sockets[msg.UID]
		if !ok {
			return
		}
		if err := sent(socket, b); err != nil {
			socket.Close()
			delete(s.sockets, msg.UID)
		}
	} else {
		for _, socket := range s.sockets {
			if err := sent(socket, b); err != nil {
				socket.Close()
				delete(s.sockets, msg.UID)
			}
		}
	}
}
func sent(c *websocket.Conn, b []byte) error {
	c.SetWriteDeadline(time.Now().Add(writeWait))
	if len(b) == 0 {
		c.WriteMessage(websocket.CloseMessage, []byte{})
		return nil
	}
	w, err := c.NextWriter(websocket.TextMessage)
	if err != nil {
		return err
	}
	w.Write(b)
	if err := w.Close(); err != nil {
		return err
	}
	return nil
}
