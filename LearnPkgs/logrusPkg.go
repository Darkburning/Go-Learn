package main

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

//// init函数在go程序启动时自动调用
//func init() {
//	// 设置日志格式为json格式
//	log.SetFormatter(&log.JSONFormatter{})
//
//	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
//	// 日志消息输出可以是任意的io.writer类型
//	log.SetOutput(os.Stdout)
//
//	// 设置日志级别为warn以上,若级别低于warn则不会显示
//	log.SetLevel(log.WarnLevel)
//}

// logrus提供了New()函数来创建一个logrus的实例
// 项目中，可以创建任意数量的logrus实例
// 一般创建一个logrus实例即可
var log = logrus.New()

// DefaultFieldHook 实现hook接口
type DefaultFieldHook struct {
}

func (hook *DefaultFieldHook) Fire(entry *logrus.Entry) error {
	entry.Data["request_id"] = "hookDefine!"
	return nil
}

func (hook *DefaultFieldHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// DefaultFormatter 实现Formatter接口
type DefaultFormatter struct {
}

// 颜色
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

func (f *DefaultFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同level去显示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出
		// 利用了ANSI ESCAPE code控制终端字体颜色
		// 输出调用者的信息, 哪个文件，第几行 以及 调用函数名
		_, err := fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", timestamp, levelColor, entry.Level, entry.Data["request_id"], entry.Data["user_ip"], entry.Message)
		if err != nil {
			return nil, err
		}
	}
	return b.Bytes(), nil
}

func main() {
	// 为当前logrus实例设置消息的输出，同样地，
	// 可以设置logrus实例的输出到任意io.writer
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)
	// 设置是否记录函数调用者的信息
	log.SetReportCaller(true)

	// 为当前logrus实例设置消息输出格式为自定义格式。
	log.SetFormatter(&DefaultFormatter{})

	// hook注册
	log.AddHook(&DefaultFieldHook{})

	// 在一个应用中、或者应用的一部分中，都有一些固定的Field
	// 我们可以创建一个logrus.Entry实例，为这个实例设置默认Fields，在上下文中使用这个logrus.Entry实例记录日
	requestLogger := log.WithFields(logrus.Fields{
		"request_id": "request_id",
		"user_ip":    "user_ip"})
	requestLogger.Info("something happened on that request") //  will log request_id and user_ip
	requestLogger.Warn("something not great happened")

}
