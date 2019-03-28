package main

import (
	"net/http"

	"fmt"

	"github.com/eaglexpf/rest-admin/controllers"
	"github.com/eaglexpf/rest-admin/pkg"
	"github.com/gin-gonic/gin"
)

func registerRouter(router *gin.Engine) {
	new(controllers.HelloController).RegisterRouter(router)
	new(controllers.CommonController).RegisterRouter(router)
}

func main() {
	//	fmt.Println(pkg.LoadData, "main")
	router := gin.Default()
	registerRouter(router)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", pkg.LoadData.HttpPort),
		Handler:        router,
		ReadTimeout:    pkg.LoadData.ReadTimeout,
		WriteTimeout:   pkg.LoadData.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("run success:")
	}
}
