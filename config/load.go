// Package config 为您执行以下操作：
//查找、加载和解组 JSON、TOML、YAML、YML、HCL、INI、envfile 或 Java 属性格式的配置文件。
//提供一种机制来为不同的配置选项设置默认值。
//提供一种机制来为通过命令行标志指定的选项设置覆盖值。
//提供别名系统以轻松重命名参数而不破坏现有代码。
//可以轻松区分用户提供的命令行或配置文件与默认值相同的时间。

package config

import (
	"errors"
	"fmt"
	flagUtil "github.com/spf13/pflag"
	configUtil "github.com/spf13/viper"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

var (
	once       sync.Once
	systemProp = make(map[string]interface{})
	config     *Config
	isLoad     = false
)

// Config
// @Description: 配置对象
type Config struct {
	Config   *configUtil.Viper // 配置对象
	Name     string            // 配置文件名称
	Type     string            // 配置文件类型
	Path     []string          // 配置文件路径
	Profiles string            // 配置文件环境
}

// init
// @Description: 初始化
func init() {
	once.Do(func() {
		initConfig()
	})
}

// GetConfig
// @Description: 获取配置对象
// @return *viper.Viper
func GetConfig() *Config {
	return config
}

// SetConfig
// @Description: 设置配置对象
// @param: configName
// @param: configType
// @param: profiles
// @param: confPaths
func SetConfig(configName, configType, profiles string, confPaths []string) *Config {
	c := configUtil.New()
	c.SetConfigName(configName)
	c.SetConfigType(configType)
	c.SetEnvPrefix(profiles)
	fmt.Printf("set config AddDefault ConfigName=%s,ConfigType=%s \n", configName, configType)
	for _, path := range confPaths {
		c.AddConfigPath(path)
		fmt.Printf("set config AddConfigPath %s \n", path)
	}
	isLoad = false
	config = &Config{
		Config:   c,
		Name:     configName,
		Type:     configType,
		Path:     confPaths,
		Profiles: profiles,
	}
	return config
}

// NewConfig
// @Description: 创建配置对象
// @param: configName 配置文件名称
// @param: configType 配置文件类型
// @param: profiles 配置文件环境
// @param: confPaths 配置文件路径
// @return *Config
func NewConfig(configName, configType, profiles string, confPaths []string) *Config {
	c := configUtil.New()
	c.SetConfigName(configName)
	c.SetConfigType(configType)
	c.SetEnvPrefix(profiles)
	fmt.Printf("new config AddDefault ConfigName=%s,ConfigType=%s \n", configName, configType)
	for _, path := range confPaths {
		c.AddConfigPath(path)
		fmt.Printf("new config AddConfigPath %s \n", path)
	}
	isLoad = false
	return &Config{
		Config:   c,
		Name:     configName,
		Type:     configType,
		Path:     confPaths,
		Profiles: profiles,
	}
}

// initConfig
// @Description: 初始化系统配置参数
func initConfig() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("initConfig panic", err)
		}
	}()
	var profiles, configPath string
	flagUtil.StringVar(&configPath, "config", "./config/config.yaml", "help message for config path")
	flagUtil.StringVar(&profiles, "server.profiles", "dev", "help message for server.profiles")
	flagUtil.StringVar(&profiles, "profiles", "dev", "help message for profiles")
	flagUtil.Parse()
	flagUtil.Visit(func(f *flagUtil.Flag) {
		systemProp[f.Name] = f.Value
		fmt.Printf("initConfig flag name=%s,value=%s \n", f.Name, f.Value)
	})
	c := configUtil.New()
	// 配置文件的文件名，没有扩展名，如 .yaml, .toml 这样的扩展名
	// 设置扩展名。在这里设置文件的扩展名。另外，如果配置文件的名称没有扩展名，则需要配置这个选项
	paths, fileName, fileType := parseFilePath(configPath)
	c.SetConfigName(fileName)
	c.SetConfigType(fileType)
	var confPathList = make([]string, 0)
	confPathList = append(confPathList, paths)
	fmt.Printf("initConfig AddDefault ConfigName=%s,ConfigType=%s \n", fileName, fileType)
	c.AddConfigPath(paths)
	fmt.Printf("initConfig AddConfigPath %s \n", paths)
	config = &Config{
		Config:   c,
		Name:     fileName,
		Type:     fileType,
		Path:     confPathList,
		Profiles: profiles,
	}
}

// parseFilePath
// @Description: 解析文件路径
// @param: files 文件路径
// @return paths 文件路径
// @return fileName 文件名称
// @return fileType 文件类型
func parseFilePath(files string) (paths, fileName, fileType string) {
	if files == "" || files == "." {
		return "./", "config", ""
	}
	paths, fileName = filepath.Split(files)
	if paths == "" {
		paths = "./"
	}
	fileType = path.Ext(files)
	fileName = strings.TrimSuffix(fileName, fileType)
	fileType = strings.TrimPrefix(fileType, ".")
	if fileName == "" && fileType == "" {
		fileName = "config"
	}
	return
}

// Load
// @Description: 加载配置文件
// @param: prefix
// @param: data
// @return error
func (config *Config) Load(prefix string, data interface{}) error {
	c := config.Config
	// 搜索并读取配置文件
	if !isLoad {
		readErr := c.ReadInConfig()
		readErr = c.MergeInConfig()
		if readErr != nil {
			fmt.Errorf("fatal error config file: %s \n", readErr)
			return errors.New("fatal error config file")
		}
		isLoad = true
	}
	// 解析
	parseErr := c.UnmarshalKey(prefix, &data)
	if parseErr != nil {
		fmt.Errorf("fatal error load config prop: %s err=%s \n", prefix, parseErr)
		return errors.New("fatal error load config prop")
	}
	return nil
}

// Read
// @Description: 读取配置文件
// @receiver: config
// @param: prefix
// @param: data
// @return error
func Read(config *Config) error {
	c := config.Config
	readErr := c.ReadInConfig()
	readErr = c.MergeInConfig()
	if readErr != nil {
		fmt.Errorf("fatal error config file: %s \n", readErr)
		return errors.New("fatal error config file")
	}
	return nil
}

// Parse
// @Description: 解析配置文件
// @receiver: config *Config
// @param: prefix
// @param: data
// @return error
func (config *Config) Parse(prefix string, data interface{}) error {
	c := config.Config
	// 解析
	parseErr := c.UnmarshalKey(prefix, &data)
	if parseErr != nil {
		fmt.Errorf("fatal error load config prop: %s err=%s \n", prefix, parseErr)
		return errors.New("fatal error load config prop")
	}
	return nil
}
