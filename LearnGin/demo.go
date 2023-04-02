package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"net/http"
	"path"
)

type UserInfo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Userid   string `form:"userid" json:"userid"`
}

func main() {
	// 创建一个服务
	r := gin.Default()
	// 添加一个icon
	r.Use(favicon.New("./favicon.ico"))
	// 修改文件上传的最大内存限制
	r.MaxMultipartMemory = 8 << 40
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

	// 法二：user/info/userid/username 获取URL参数
	r.GET("/user/info/:userid/:username", func(context *gin.Context) {
		userid := context.Param("userid")
		username := context.Param("username")
		context.JSON(http.StatusOK, gin.H{
			"userid from frontend":   userid,
			"username from frontend": username,
		})
	})

	//参数绑定JSON
	r.POST("/userVal", func(context *gin.Context) {
		var userInfo UserInfo
		if err := context.ShouldBind(&userInfo); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v", userInfo) // %#v返回字面量，包括变量名和其值
			context.JSON(http.StatusOK, userInfo)
		}
	})

	// 参数绑定form
	r.POST("/userForm", func(context *gin.Context) {
		var userInfo UserInfo
		if err := context.ShouldBind(&userInfo); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v", userInfo)
			context.JSON(http.StatusOK, userInfo)
		}
	})

	// 参数绑定请求参数
	r.POST("/userinfo", func(context *gin.Context) {
		var userinfo UserInfo
		if err := context.ShouldBind(&userinfo); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v", userinfo)
			context.JSON(http.StatusOK, userinfo)
		}

	})

	// 上传单个文件
	r.GET("/upload", func(context *gin.Context) {
		context.HTML(http.StatusOK, "Upload.html", gin.H{
			"status": "ok",
		})
	})
	// POST方法将上传的文件保存在服务器本地
	r.POST("/uploadFile", func(context *gin.Context) {
		if file, err := context.FormFile("myFile"); err != nil {
			context.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
		} else {
			filePath := path.Join("./", file.Filename)
			if err := context.SaveUploadedFile(file, filePath); err != nil {
				context.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			}
		}

	})

	// 上传多个文件, HTML表单中input属性改成multiple,同时要改变action的路径
	r.POST("/uploadMultiFile", func(context *gin.Context) {
		// 先获取表单对象
		if form, err := context.MultipartForm(); err != nil {
			context.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
		} else {
			// 再从表单对象中取出文件
			files := form.File["myFile"]
			// 将每个文件都保存到本地
			for index, file := range files {
				filepath := fmt.Sprintf("./%d_%s", index, file.Filename) // 拼接路径
				context.SaveUploadedFile(file, filepath)
			}
			context.JSON(http.StatusOK, gin.H{
				"uploadFileNum": len(files),
			})
		}

	})

	// 路由
	r.GET("/bilibili", func(context *gin.Context) {
		// 外部重定向，会改变地址
		context.Redirect(http.StatusMovedPermanently, "https://www.bilibili.com/")
	})

	r.GET("/a", func(context *gin.Context) {
		// 实现内部重定向，不会改变地址
		context.Request.URL.Path = "/b" // 将路径修改
		r.HandleContext(context)        // 通过引擎处理当前上下文
	})

	r.GET("/b", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "This is msg from b",
		})
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

	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/check", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"method": "get",
			})
		})

		shopGroup.POST("/check", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"method": "post",
			})
		})

		shopGroup.PUT("/check", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"method": "put",
			})
		})

		shopGroup.DELETE("/check", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"method": "delete",
			})
		})
	}
	//服务器端口,注意冒号
	r.Run(":8000") // listen and serve on "localhost:8000"
}
