package routers

import (
	"github.com/gin-gonic/gin"
	v1 "homework/internal/apiserver/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	requestTest := v1.NewTestRequest()

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/SendRequestHeader", requestTest.SendRequestHeader)
		apiv1.GET("/healthz", requestTest.CheckHealth)
	}
	return r
}
