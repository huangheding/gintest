package app

import (
	"gintest/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	p := model.Person{}
	result, _ := p.FindPerson()

	c.JSON(http.StatusOK, result)
}

func FindInterest(c *gin.Context) {
	p := model.Interest{}
	result, _ := p.ArrangeInterest()

	c.JSON(http.StatusOK, result)
}
