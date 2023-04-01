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

	// 加载静态资源文件如：.css, .js 以及图片
	r.Static("/static", "./static") // HTML中所有以/static开头的文件都去demo.go同级目录下的static去找
	// 加载静态页面
	r.LoadHTMLGlob("templates/*")

	// 响应HTML页面
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "This is the data from backend!",
		})
	})

	// 访问地址，处理请求，包括GET,POST，PUT和DELETE
	r.GET("/get", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "GET",
		})
	})

	r.POST("/post", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{ //处理JSON可用map（也即gin.H // ）或者结构体
			"message": "POST",
		})
	})

	r.DELETE("/delete", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "DELETE",
		})
	})

	r.PUT("/put", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "PUT",
		})
	})

	// 获取请求参数
	// 法一：url?userid=xxx&username=xxx
	r.GET("/user/info", func(context *gin.Context) {
		userid := context.DefaultQuery("userid", "000000") // 默认Query,如果没有请求userid则会自动返回其值会默认为000000
		username := context.Query("username")
		context.JSON(http.StatusOK, gin.H{
			"userid from frontend":   userid,
			"username from frontend": username,
		})
	})
	// 表单验证
	r.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "getForm.html", nil)
	})
	// 一次请求一个响应，验证登录时应该再发起一个服务
	r.POST("/login", func(context *gin.Context) {
		password := context.PostForm("password")
		username := context.PostForm("username")

		if password == "123456" && username == "kingkia" {
			context.HTML(http.StatusOK, "index.html", gin.H{ // 返回的是HTML,
				"msg": "Success!",
			})
		}
	})

	// 法二：user/info/userid/username 推荐使用
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
