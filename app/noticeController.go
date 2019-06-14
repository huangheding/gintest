package app

import (
	"encoding/json"
	"gintest/model"
	"gintest/util/rs"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	m := &model.App_notify{}

	data, _ := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal(data, &m); err != nil {
		log.Errorf("post body 解析映射失败，err is：%s", err)
	}
	b, err := json.Marshal(m.Content)
	if err != nil {
		log.Errorf("struct to string err is：%s", err)
		b = nil
	}
	go rs.Produce(string(b))
	m.InsertAppNotify()
	c.JSON(http.StatusOK, "success")

}
