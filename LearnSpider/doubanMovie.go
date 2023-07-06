package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	USERNAME = "root"
	PASSWORD = "123456"
	HOST     = "127.0.0.1"
	PORT     = "3306"
	DBNAME   = "douban_movie"
)

type MovieData struct {
	gorm.Model
	Title    string `json:"title"`
	Director string `json:"director"`
	Picture  string `json:"picture"`
	Actor    string `json:"actor"`
	Year     string `json:"year"`
	Score    string `json:"score"`
	Quote    string `json:"quote"`
}

func main() {
	db := InitGORM()
	//ch := make(chan struct{})
	// 三种方式的性能测试，总的来看，并发使用WaitGroup的方式速度最快
	WaitGroupStart(db)
	//ChannelStart(ch, db)
	//NormalStart(db)
}

func ChannelStart(ch chan struct{}, db *gorm.DB) {
	// 除同步goroutine外还可以传递数据
	start := time.Now()
	for i := 0; i < 10; i++ {
		go Spider(strconv.Itoa(i*25), ch, db)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
	elapsed := time.Since(start)
	fmt.Printf("ChannelStart cost: %s\n", elapsed)
}

func NormalStart(db *gorm.DB) {
	start := time.Now()
	for i := 0; i < 10; i++ {
		Spider(strconv.Itoa(i*25), nil, db)
	}
	elapsed := time.Since(start)
	fmt.Printf("NormalStart Time %s \n", elapsed)
}

func WaitGroupStart(db *gorm.DB) {
	// 适合简单的同步goroutine而不需要进行传递数据
	start := time.Now()
	wg := new(sync.WaitGroup)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			Spider(strconv.Itoa(i*25), nil, db)
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("WaitGroupStart Time %s\n", elapsed)
}

func Spider(page string, ch chan struct{}, db *gorm.DB) {
	// 1. 新建客户端发起请求
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://movie.douban.com/top250?start="+page, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 处理cookie问题

	// 设置头部混淆，避免爬虫行为被发现
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", "https://movie.douban.com/chart")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")

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
	// 3. 获取节点链接
	//#content > div > div.article > ol > li:nth-child(1)
	//#content > div > div.article > ol > li:nth-child(1) > div > div.info > div.hd > a > span:nth-child(1)
	//#content > div > div.article > ol > li:nth-child(1) > div > div.info > div.bd > p:nth-child(1)
	docDetail.Find("#content > div > div.article > ol > li > div").
		Each(func(i int, s *goquery.Selection) {
			var movieData MovieData
			//#content > div > div.article > ol > li:nth-child(16) > div > div.info > div.hd > a > span:nth-child(1)
			title := s.Find("div.info > div.hd > a > span:nth-child(1)").Text()
			img := s.Find("div.pic > a > img")
			imgTmp, ok := img.Attr("src")
			info := strings.Trim(s.Find("div.info > div.bd > p:nth-child(1)").Text(), " ")
			director, actor, year := InfoSpite(info)
			score := strings.Trim(s.Find("div.info > div.bd > div > span.rating_num").Text(), " ")
			score = strings.Trim(score, "\n")
			quote := strings.Trim(s.Find("div.info > div.bd > p.quote > span").Text(), " ")
			if ok {
				movieData.Title = title
				movieData.Director = director
				movieData.Picture = imgTmp
				movieData.Actor = actor
				movieData.Year = year
				movieData.Score = score
				movieData.Quote = quote
				// 4. 保存信息到数据库中或文件中
				InsertDB(db, movieData)
			}
		})
	if ch != nil {
		ch <- struct{}{}
	}
}

func InfoSpite(info string) (director, actor, year string) {
	directorRe, _ := regexp.Compile(`导演:(.*)主演:`)
	if len(director) < 8 {
		director = string(directorRe.Find([]byte(info)))
	} else {
		director = string(directorRe.Find([]byte(info)))[8:]
	}
	director = strings.Trim(director, "主演:")
	actorRe, _ := regexp.Compile(`主演:(.*)`)
	if len(actor) < 8 {
		actor = string(actorRe.Find([]byte(info)))
	} else {
		actor = string(actorRe.Find([]byte(info)))[8:]
	}
	yearRe, _ := regexp.Compile(`(\d+)`)
	year = string(yearRe.Find([]byte(info)))
	return
}

func InitGORM() *gorm.DB {
	dsn := USERNAME + ":" + PASSWORD + "@tcp(" + HOST + ":" + PORT + ")/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(fmt.Sprintf("[%s] mysql连接失败", dsn))
	} else {
		fmt.Println(fmt.Sprintf("[%s] mysql连接成功", dsn))
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour * 4)

	migrator := db.Migrator()
	if migrator.HasTable(&MovieData{}) {
		fmt.Println("already had table movie_data, so drop it.")
		migrator.DropTable(&MovieData{})
	}
	err = migrator.AutoMigrate(&MovieData{})
	if err != nil {
		fmt.Println("AutoMigrate failed!")
	}
	return db
}

func InsertDB(db *gorm.DB, data MovieData) bool {
	tx := db.Begin()
	if err := tx.Create(&data).Error; err != nil {
		tx.Rollback()
		return false
	} else {
		tx.Commit()
		return true
	}
}
