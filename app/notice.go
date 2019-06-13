package app

import (
	"encoding/json"
	"fmt"
	"gintest/util/rs"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	notifyTypeTd string
	createdTime  int64
	category     string
	userId       string
	content      *Content
}
type Content struct {
	title string
	desc  string
}

func Test(c *gin.Context) {
	m := Message{}

	data, _ := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal(data, &m); err != nil {
		fmt.Println(err)
	}
	item := m.content.desc
	go rs.Produce(item)
	c.JSON(http.StatusOK, "success")

}
