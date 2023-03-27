package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"net/http"
)

func main() {
	// 创建一个服务
	r := gin.Default()
	// 添加一个icon
	r.Use(favicon.New("./favicon.ico"))
	// 访问地址，处理请求，包括POST，PUT和DELETE

	// 返回json
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})

	// 加载静态页面
	r.LoadHTMLGlob("templates/*")
	// 加载资源文件
	r.Static("/static", "./static")
	// 响应页面
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "This is the data from backend!",
		})
	})

	// 获取请求参数
	// 法一：url?userid=xxx&username=xxx
	r.GET("/user/info", func(context *gin.Context) {
		userid := context.Query("userid")
		username := context.Query("username")
		context.JSON(http.StatusOK, gin.H{
			"userid from frontend":   userid,
			"username from frontend": username,
		})
	})

	// 法二：user/info/userid/username
	r.GET("/user/info/:userid/:username", func(context *gin.Context) {
		userid := context.Param("userid")
		username := context.Param("username")
		context.JSON(http.StatusOK, gin.H{
			"userid from frontend":   userid,
			"username from frontend": username,
		})
	})

	// 路由
	r.GET("/bilibili", func(context *gin.Context) {
		// 重定向
		context.Redirect(http.StatusMovedPermanently, "https://www.bilibili.com/")
	})
	// 查看HTML文件可知道，404_1的素材存储在本地，而404_2的素材存储在对应的url
	// 404_1 NOTFOUND页面
	//r.NoRoute(func(context *gin.Context) {
	//	context.HTML(http.StatusNotFound, "404_1.html", nil)
	//})
	// 404_2 NOTFOUND页面
	r.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusNotFound, "404_2.html", nil)
	})

	//服务器端口,注意冒号
	r.Run(":8000") // listen and serve on "localhost:8000"
}
