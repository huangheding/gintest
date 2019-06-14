package model

import (
	"encoding/json"
	"time"

	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

type App_notify struct {
	ID           string    `json:"id"`
	NotifyTypeId string    `json:"notifyTypeId"`
	Category     string    `json:"category"`
	UserId       string    `json:"userId"`
	CreatedTime  time.Time `json:"date"`

	Content *Content `json:"content"`
}
type Content struct {
	Key   string `json:"key"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func (m *App_notify) InsertAppNotify() error {
	uid, _ := uuid.NewV1()
	jsonstr, _ := json.Marshal(m.Content)
	sql := "INSERT into app_notify(id,content,category,notify_type_id,created_time) VALUES(UUID_TO_BIN(?,true),?,?,UUID_TO_BIN(?,true),?)"
	log.Info(sql)
	dbitem := db.Exec(sql, uid.String(), jsonstr, m.Category, m.NotifyTypeId, time.Now())
	if err := dbitem.Error; err != nil {
		return err
	}
	m.ID = uid.String()
	m.InsertAppUserNotify()
	return nil
}
func (m *App_notify) InsertAppUserNotify() error {
	uid, _ := uuid.NewV1()
	sql := "INSERT into app_user_notify(id,is_read,notify_id,user_id,created_time) VALUES(UUID_TO_BIN(?,true),?,UUID_TO_BIN(?,true),UUID_TO_BIN(?,true),?)"
	if err := db.Exec(sql, uid.String(), 0, m.ID, m.UserId, time.Now()).Error; err != nil {
		return err
	}
	return nil

}
