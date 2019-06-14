package app

import (
	"encoding/json"
	"fmt"
	"gintest/model"
	"gintest/util/rs"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	m := &model.App_notify{}

	data, _ := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal(data, &m); err != nil {
		fmt.Println(err)
	}
	b, err := json.Marshal(m.Content)
	if err != nil {
		fmt.Println(err)
		b = nil
	}
	go rs.Produce(string(b))
	m.InsertAppNotify()
	c.JSON(http.StatusOK, "success")

}
