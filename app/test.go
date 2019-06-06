package app

import (
	"encoding/json"
	"fmt"
	"gintest/model"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	p := model.Person{}
	result, _ := p.FindPerson()

	c.JSON(http.StatusOK, result)
}

//post 解析几种格式参数方式
func AddPerson(c *gin.Context) {
	code := 0

	//post body form
	// for k, v := range c.Request.PostForm {
	// 	fmt.Printf("k:%v\n", k)
	// 	fmt.Printf("v:%v\n", v)
	// }

	//post body json
	p := model.Person{}
	p.Date = time.Now()
	data, _ := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal(data, &p); err != nil {
		code = -1
		fmt.Println(err)
	}

	//post param
	// p := model.Person{
	// 	Name: c.PostForm("name"),
	// 	Age:  c.DefaultPostForm("age", "22"), //默认值
	// 	Date: time.Now(),
	// }
	err := p.AddPerson()
	if err != nil {
		code = -1
	}

	c.JSON(http.StatusOK, code)
}

//get解析参数
func DeletePerson(c *gin.Context) {

	code := 0
	// c.DefaultQuery("id", "1")  //默认值
	p := model.Person{}
	if err := p.DeletePerson(c.Query("id")); err != nil {
		code = -1
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, code)
}

func UpdatePerson(c *gin.Context) {
	code := 0
	p := model.Person{
		ID:   c.Query("id"),
		Name: c.Query("name"),
		Age:  c.Query("age"),
		Date: time.Now(),
	}
	if err := p.UpdatePerson(); err != nil {
		code = -1
	}

	c.JSON(http.StatusOK, code)
}

func FindInterest(c *gin.Context) {
	p := model.Cfg_interest{}
	result, _ := p.ArrangeInterest()

	c.JSON(http.StatusOK, result)
}
