package main

import (
	"io/ioutil"
	"net/http"
	//导入echo包
	"github.com/labstack/echo"
)

func main() {
	//实例化echo对象。
	e := echo.New()

	//注册一个Get请求, 路由地址为: /tizi365  并且绑定一个控制器函数, 这里使用的是闭包函数。
	e.POST("/tizi365", func(c echo.Context) error {
		//控制器函数直接返回一个字符串，http响应状态为http.StatusOK，就是200状态。
		h := c.Request().Header
		c.JSONPretty(http.StatusOK, h, "\t")
		b, _ := ioutil.ReadAll(c.Request().Body)
		return c.String(200, string(b))
	})

	//启动http server, 并监听8080端口，冒号（:）前面为空的意思就是绑定网卡所有Ip地址，本机支持的所有ip地址都可以访问。
	e.Start(":18080")
}
