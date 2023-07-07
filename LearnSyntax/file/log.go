package main

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var (
	re       *regexp.Regexp
	calTime  int64
	callTime int64
)

func init() {
	re = regexp.MustCompile(`calculation cost time：(\d+) ms, call func cost time: (\d+) ms`)
	callTime = 0
	calTime = 0
}

// GenerateLogFile 生成格式如：calculation cost time：xx ms, call func cost time: xx ms
func GenerateLogFile(path string, dirNum, fileNum, logLineNum int) {
	start := time.Now()
	wg := new(sync.WaitGroup)
	for i := 0; i < dirNum; i++ {
		// openFile
		wg.Add(1)
		dirName := fmt.Sprintf("%s/log/Day%d", path, i+1)
		err := os.MkdirAll(dirName, 0666)
		if err != nil {
			fmt.Println("mkdir error: ", err.Error())
			continue
		}

		// 每个目录并发写入
		go func(dirName string) {
			defer wg.Done()
			for j := 0; j < fileNum; j++ {
				fileName := fmt.Sprintf("/logFile%d", j+1)

				file, err := os.OpenFile(dirName+fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
				if err != nil {
					fmt.Println("openFile error: ", err.Error())
					continue
				}
				defer file.Close()
				// writeFile
				// 在内存中写完所有内容再一次性写入磁盘
				content := make([]byte, 0)
				for k := 0; k < logLineNum; k++ {
					logLine := fmt.Sprintf("calculation cost time：%d ms, call func cost time: %d ms\n", rand.Intn(200), rand.Intn(200))
					content = append(content, []byte(logLine)...)
				}
				_, err = file.Write(content)
				if err != nil {
					fmt.Println("writeFile error: ", err.Error())
					return
				}
			}
		}(dirName)

	}
	wg.Wait()
	fmt.Printf("GenerateLogFile cost: %v", time.Since(start))

}

// ConcurrentProcessDir 并发处理目录
func ConcurrentProcessDir(dir string) {
	start := time.Now()
	files := ListDir(dir)
	wg := sync.WaitGroup{}
	wg.Add(len(files))
	for _, file := range files {
		go func(file string) {
			defer wg.Done()
			ProcessFile(file)
		}(file) //	协程中使用for循环生成的变量时，务必把变量拷贝到协程里去
	}
	wg.Wait()
	fmt.Printf("ConcurrentProcessDir cost: %v\n", time.Since(start))
}

// ProcessDir 普通处理目录
func ProcessDir(dir string) {
	start := time.Now()
	files := ListDir(dir)
	for _, file := range files {
		ProcessFile(file)
	}
	fmt.Printf("ProcessDir cost: %v\n", time.Since(start))
}

// ListDir 返回目录下所有文件的文件名
func ListDir(dir string) []string {
	files := make([]string, 0, 100)
	// WalkDir内部会深度优先遍历目录,只需要处理
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error { // WalkDir比Walk更高效
		if err != nil {
			return err
		} else if info, err := d.Info(); err == nil {
			// 若没出错则将文件名加入切片
			if info.Mode().IsRegular() {
				files = append(files, path)
			}
			return nil
		} else {
			// 出错返回错误
			return err
		}
	})
	fmt.Printf("%s目录下共%d个文件\n", dir, len(files))
	return files
}

// ProcessFile 处理文件，从每个文件中提取数字累加
func ProcessFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("open file failed:%v", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		// 异常情况
		if log, err := reader.ReadString('\n'); err != nil {
			// 读到文件结尾
			if err == io.EOF {
				if len(log) > 0 {
					n1, n2 := ExtractNumber(log)
					if n1 >= 0 {
						atomic.AddInt64(&calTime, int64(n1)) // 原子操作保证并发安全
						atomic.AddInt64(&callTime, int64(n2))
					}
				}
			} else {
				fmt.Printf("readFile failed: %s", err.Error())
			}
			break
			// 正常情况
		} else {
			// ReadString包括分隔符
			log = strings.TrimRight(log, "\n")
			if len(log) > 0 {
				n1, n2 := ExtractNumber(log)
				if n1 >= 0 {
					atomic.AddInt64(&calTime, int64(n1)) // 原子操作保证并发安全
					atomic.AddInt64(&callTime, int64(n2))
				}
			}
		}
	}
}

// ExtractNumber ParseLog
func ExtractNumber(log string) (int, int) {
	index := re.FindAllSubmatchIndex([]byte(log), -1)
	match1 := log[index[0][2]:index[0][3]]
	match2 := log[index[0][4]:index[0][5]]
	num1, err := strconv.Atoi(match1)
	if err != nil {
		fmt.Println("atoi error: " + err.Error())
		return -1, -1
	}
	num2, err := strconv.Atoi(match2)
	if err != nil {
		fmt.Println("atoi error: " + err.Error())
		return -1, -1
	}

	return num1, num2
}
func main() {
	//GenerateLogFile(".", 10, 100, 10000)
	//fmt.Println(ExtractNumber("calculation cost time：102 ms, call func cost time: 126 ms"))
	//ProcessDir("./log")
	ConcurrentProcessDir("./log")
	fmt.Printf("calTime: %dms, callTime: %dms", calTime, callTime)
}
