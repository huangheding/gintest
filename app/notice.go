package app

import (
	"gintest/model"
	"gintest/util/rs"
	"gintest/util/ws"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	i := 1
	// for i := 0; i < 100; i++ {
	str := "test"
	str += strconv.Itoa(i)
	rs.Produce(str)

	if ws.ServiceOnline.IsUserOnline(strconv.Itoa(i)) {
		config := model.Config.Tomls
		mess := rs.Custom(config.RedisConf.Address, "redismq")
		ws.ServiceOnline.Push(strconv.Itoa(i), mess)
	}
	// }

}
