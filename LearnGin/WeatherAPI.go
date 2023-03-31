package main

// 不知道为什么会返回空数据...
import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Daily struct {
	FxDate   string `json:"fxDate"`
	TempMax  string `json:"tempMax"`
	TempMin  string `json:"tempMin"`
	Humidity string `json:"humidity"`
}
type Weather struct {
	Today [1]Daily `json:"daily"`
}

func main() {
	router := gin.Default()

	router.GET("/weather/:city", func(c *gin.Context) {
		city := c.Param("city")
		apiKey := "1ac7e50c9e8245f5aa8846800e117502"
		apiUrl := fmt.Sprintf("https://devapi.qweather.com/v7/weather/3d?location=%s&key=%s", city, apiKey)

		response, err := http.Get(apiUrl)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer response.Body.Close()
		fmt.Println("getDone")

		var reader io.ReadCloser

		// 检查响应头中是否包含Content-Encoding为gzip的信息，
		// 如果包含，则使用gzip.NewReader函数创建一个gzip.Reader对象来解析响应体数据
		if response.Header.Get("Content-Encoding") == "gzip" {
			reader, err = gzip.NewReader(response.Body)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		} else {
			reader = response.Body
		}

		defer reader.Close()
		fmt.Println("gzipDone")

		bodyBytes, err := io.ReadAll(reader)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("readAllDone")

		var weatherData Weather

		err = json.Unmarshal(bodyBytes, &weatherData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(weatherData.Today[0])
		c.JSON(http.StatusOK, gin.H{"city": city, "tempMax": weatherData.Today[0].TempMax, "tempMin": weatherData.Today[0].TempMin, "humidity": weatherData.Today[0].Humidity})
	})

	router.Run(":8080")
}
