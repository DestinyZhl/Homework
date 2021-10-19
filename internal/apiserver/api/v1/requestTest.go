package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"homework/global"
	app "homework/internal/pkg"
	"os"
)

type TestRequest struct {
}

func NewTestRequest() TestRequest {
	return TestRequest{}
}

//@Summary 接收客户端 request，并将 request 中带的 header 写入 response header
//@Description 将 request 中带的 header 写入 response header
//@Tags 功能测试
//@Success 200 {string} string "{"msg": "Success", "HttpCode": 200}"
//@Router /api/v1/SendRequestHeader [get]
func (t TestRequest) SendRequestHeader(c *gin.Context) {
	response := app.NewResponse(c)

	//1. 接收客户端 request，并将 request 中带的 header 写入 response header
	for k, v := range c.Request.Header {
		response.Ctx.Header(k, v[0])
	}
	//2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	version := os.Getenv("VERSION")
	if version != "" {
		fmt.Printf("服务器端VERSION：%s", version)
		global.Logger.Info("服务器端VERSION：%s", version)
		response.Ctx.Header("VERSION", version)
	}

	//3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	fmt.Printf("客户端 IP:%s,http 返回码：200", c.Request.RemoteAddr)
	global.Logger.Info("客户端 IP:%s,http 返回码：200", c.Request.RemoteAddr)

	response.ToResponse(gin.H{"msg": "Success", "HttpCode": 200})
	return
}

//@Summary 健康检测
//@Description 健康检测
//@Tags 健康检测
//@Success 200 {string} string "{"msg": "Success", "HttpCode": 200}"
//@Router /healthz [get]
func (t TestRequest) CheckHealth(c *gin.Context) {

	response := app.NewResponse(c)

	//3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	fmt.Printf("客户端 IP:%v,http 返回码：200", c.Request.RemoteAddr)
	global.Logger.Info("客户端 IP:%v,http 返回码：200", c.Request.RemoteAddr)

	response.ToResponse(gin.H{"msg": "Success", "HttpCode": 200})
	return
}
