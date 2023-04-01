package main

// 注意API格式
import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

const ApiKey = "1ac7e50c9e8245f5aa8846800e117502"

type Daily struct {
	FxDate   string `json:"fxDate"`
	TempMax  string `json:"tempMax"`
	TempMin  string `json:"tempMin"`
	Humidity string `json:"humidity"`
}
type Weather struct {
	Today []Daily `json:"daily"`
}

func main() {
	router := gin.Default()

	router.GET("/weather/:city", func(c *gin.Context) { // 城市搜索也可以通过请求一个API来找到
		city := c.Param("city")
		city = city[1:] // bug来源: 此处返回的值会包含冒号而原来请求的API是不包含冒号的
		apiUrl := fmt.Sprintf("https://devapi.qweather.com/v7/weather/3d?location=%v&key=%v", city, ApiKey)

		response, err := http.Get(apiUrl)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer response.Body.Close()

		// 和风天气使用了gzip压缩
		// 检查响应头中是否包含Content-Encoding为gzip的信息，
		// 如果包含，则使用gzip.NewReader函数创建一个gzip.Reader对象来解析响应体数据
		var reader io.ReadCloser = response.Body
		if response.Header.Get("Content-Encoding") == "gzip" {
			reader, err = gzip.NewReader(response.Body)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			defer reader.Close()
		}

		bytesData, err := io.ReadAll(reader)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var weatherData Weather

		// 反序列化
		err = json.Unmarshal(bytesData, &weatherData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"city": 101010100, "tempMax": weatherData.Today[0].TempMax, "tempMin": weatherData.Today[0].TempMin, "humidity": weatherData.Today[0].Humidity})
	})

	router.Run(":8080")
}
