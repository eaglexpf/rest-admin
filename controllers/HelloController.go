package controllers

import (
	//	"fmt"

	"github.com/eaglexpf/rest-admin/pkg"
	"github.com/eaglexpf/rest-admin/service"
	"github.com/gin-gonic/gin"
)

type HelloController struct {
	pkg.Controller
}

func (this *HelloController) RegisterRouter(router *gin.Engine) {
	r := router.Group("/user")
	r.GET("/", this.hello)
}

func (this *HelloController) hello(c *gin.Context) {
	var userService service.UserService
	count := userService.Count()
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": count,
	})
}
