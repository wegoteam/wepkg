package config

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"testing"
)

func TestConfig(t *testing.T) {
	viper.AddRemoteProvider("etcd3", "http://127.0.0.1:4001", "./config.json")
	viper.SetConfigType("json") // because there is no file extension in a stream of bytes, supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
	err := viper.ReadRemoteConfig()
	if err != nil { // 处理错误
		fmt.Errorf("Fatal error config file: %s \n", err)
	}
	keys := viper.AllKeys()
	fmt.Println(keys)
}

func TestViper(t *testing.T) {
	viper := viper.New()
	viper.SetConfigName("config") // 配置文件的文件名，没有扩展名，如 .yaml, .toml 这样的扩展名
	viper.SetConfigType("yaml")   // 设置扩展名。在这里设置文件的扩展名。另外，如果配置文件的名称没有扩展名，则需要配置这个选项
	//viper.AddConfigPath("/etc/appname/")  // 查找配置文件所在路径
	//viper.AddConfigPath("$HOME/.appname") // 多次调用AddConfigPath，可以添加多个搜索路径
	viper.AddConfigPath(".")    // 还可以在工作目录中搜索配置文件
	err := viper.ReadInConfig() // 搜索并读取配置文件
	if err != nil {             // 处理错误
		//if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		//	// Config file not found; ignore error if desired
		//} else {
		//	// Config file was found but another error was produced
		//}
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	var mysql = &MySQL{}
	err = viper.UnmarshalKey("mysql", mysql)
	if err != nil {
		fmt.Println(err)
	}

	var red = &redis.Options{}

	err = viper.UnmarshalKey("redis", red)
	if err != nil {
		fmt.Println(err)
	}

	val := redis.NewClient(red).Get(context.Background(), "test").Val()
	fmt.Println(val)
	fmt.Println(red)
	fmt.Println(mysql)
	keys := viper.AllKeys()
	fmt.Println(keys)
}
