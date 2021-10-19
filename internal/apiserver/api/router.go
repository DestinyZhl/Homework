package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	v1 "homework/internal/apiserver/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	requestTest := v1.NewTestRequest()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/SendRequestHeader", requestTest.SendRequestHeader)
	}

	r.GET("/healthz", requestTest.CheckHealth)

	return r
}
