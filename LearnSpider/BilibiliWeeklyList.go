package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type VideoData struct {
	gorm.Model
	Title       string `json:"title"`
	UpName      string `json:"upName"`
	Collection  string `json:"collection"`
	Comment     string `json:"comment"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func main() {
	//db := InitGORM()
	var db *gorm.DB
	wg := new(sync.WaitGroup)
	wg.Add(2)
	for i := 1; i <= 2; i++ {
		go SpiderBilibiliWeeklyList(strconv.Itoa(i), wg, db)
	}
	wg.Wait()
}

func SpiderBilibiliWeeklyList(num string, wg *sync.WaitGroup, db *gorm.DB) {
	defer wg.Done()
	// 新建http客户端
	client := new(http.Client)
	req, err := http.NewRequest("GET", "https://www.bilibili.com/v/popular/weekly?num="+num, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 添加头部
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	//req.Header.Set("Accept-Encoding", "gzip, deflate, br") bug来源
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1823.67")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	//req.Header.Set("Sec-Ch-Ua-Platform", "Windows")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	// 2. 解析网页
	defer resp.Body.Close()
	docDetail, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("test0")
	// #app > div > div.popular-video-container.weekly-list > div:nth-child(2) > div > div:nth-child(1)
	// #app > div > div.popular-video-container.weekly-list > div:nth-child(2) > div > div:nth-child(2)
	// 3. 获得节点信息
	//#app > div > div.popular-video-container.weekly-list > div:nth-child(2) > div
	docDetail.Find("#app > div > div.popular-video-container.weekly-list > div:nth-child(2) > div").
		Each(func(i int, s *goquery.Selection) {
			fmt.Println("test1")
			var videoData VideoData
			// #app > div > div.popular-video-container.weekly-list > div:nth-child(2) > div > div:nth-child(1) > div.video-card__content > a > img
			// #app > div > div.popular-video-container.weekly-list > div:nth-child(2) > div > div:nth-child(2) > div.video-card__content > a > img
			imgLabel := s.Find("div:nth-child(" + strconv.Itoa(i+1) + ") > div.video-card__content > a > img")
			img, ok := imgLabel.Attr("src")
			// #app > div > div.popular-video-container.weekly-list > div:nth-child(2) > div > div:nth-child(1) > div.video-card__info > p
			titleLabel := s.Find("div:nth-child(" + strconv.Itoa(i+1) + ") > div.video-card__info > p")
			title, ok := titleLabel.Attr("title")
			fmt.Println(img, title)
			if ok {
				videoData.Image = img
				videoData.Title = title
			}
		})
}
