package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Mysql Mysql `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
}
type Redis struct {
	Addr     string `yaml:"addr"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Db       string `yaml:"db"`
	PoolSize int    `yaml:"poolSize"`
}

type Mysql struct {
	Host          string `yaml:"host"`
	Port          int    `yaml:"port"`
	Db            string `yaml:"db"`
	User          string `yaml:"user"`
	PassWord      string `yaml:"password"`
	LogLevel      string `yaml:"log_level"`
	Configuration string `yaml:"configuration"`
}

var config *Config

func Configuration() *Config {
	if config != nil {
		return config
	}

	// 设置需要读取的文件路径
	viper.AddConfigPath("./")
	// 设置需要读取的文件名
	viper.SetConfigName("config")
	// 设置需要读取的文件类型
	viper.SetConfigType("yaml")
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("config file error: %s", err)
		os.Exit(1)
	}
	// 将读取到的信息反序列化绑定到指定类型的结构体上
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("config file error: %s", err)
		os.Exit(1)
	}

	fmt.Println("Configuration file :", config)
	return config
}

func main() {
	config = Configuration()
}
