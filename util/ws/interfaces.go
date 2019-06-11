package ws

import "github.com/gorilla/websocket"

// 在线服务接口  websocket
type OnlineService interface {
	UserOnline(uid string, socket *websocket.Conn) //用户上线
	UserOffline(uid string)                        //用户下线
	IsUserOnline(uid string) bool                  //用户是否在线
	Push(uid string, content string)               //信息单发
	PushAll(content string)                        //信息群发
	UserOnlineList() map[string]bool               //所有在线用户
}
