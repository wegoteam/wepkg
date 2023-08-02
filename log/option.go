package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/wegoteam/wepkg/bean"
	"github.com/wegoteam/wepkg/config"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"strings"
	"sync"
)

// LoggerConfig
// @Description:
type LoggerConfig struct {
	Format     string `json:"format" yaml:"format"`         //日志输出格式，可选项：json、text
	Level      string `json:"level" yaml:"level"`           //日志级别，可选项：debug、info、warn、error、panic、fatal
	Output     string `json:"output" yaml:"output"`         //日志输出位置，可选项：console、file；多个用逗号分隔
	FileName   string `json:"fileName" json:"fileName"`     //日志文件名
	MaxSize    int    `json:"maxSize" json:"maxSize"`       //日志文件最大大小，单位：MB
	MaxAge     int    `json:"maxAge" yaml:"maxAge"`         //日志文件最大保存时间，单位：天
	MaxBackups int    `json:"maxBackups" yaml:"maxBackups"` //日志文件最大备份数量
}

var (
	logger       *logrus.Logger
	loggerConfig *LoggerConfig
	once         sync.Once
	mutex        sync.Mutex
)

// init
// @Description: 初始化日志
func init() {
	once.Do(func() {
		// 读取配置文件
		loggerConfig = getConfig()
		// 设置日志配置
		SetLoggerConfig(*loggerConfig)
	})

	//logier, err := rotatelogs.New(
	//	// 切割后日志文件名称
	//	"./log/log2.txt",
	//	//rotatelogs.WithLinkName(Current.LogDir),   // 生成软链，指向最新日志文件
	//	rotatelogs.WithMaxAge(30*24*time.Hour),    // 文件最大保存时间
	//	rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	//	//rotatelogs.WithRotationCount(3),
	//	//rotatelogs.WithRotationTime(time.Minute), // 日志切割时间间隔
	//)
	//if err != nil {
	//	fmt.Errorf("config local file system logger error. %v", err)
	//}
	//lfHook := lfshook.NewHook(lfshook.WriterMap{
	//	logrus.InfoLevel:  logier,
	//	logrus.FatalLevel: logier,
	//	logrus.DebugLevel: logier,
	//	logrus.WarnLevel:  logier,
	//	logrus.ErrorLevel: logier,
	//	logrus.PanicLevel: logier,
	//}, formatter)
	//logger.AddHook(lfHook)
}

// SetLoggerConfig
// @Description: 设置日志配置
// @param: config
func SetLoggerConfig(config LoggerConfig) {
	mutex.Lock()
	defer mutex.Unlock()
	loggerConfig = &config
	logger = logrus.New()
	var formatter logrus.Formatter
	if loggerConfig.Format == "json" {
		formatter = &logrus.JSONFormatter{}
	} else {
		formatter = &logrus.TextFormatter{}
	}
	logger.SetFormatter(formatter)

	switch loggerConfig.Level {
	case "trace":
		logger.SetLevel(logrus.TraceLevel)
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	case "panic":
		logger.SetLevel(logrus.PanicLevel)
	case "fatal":
		logger.SetLevel(logrus.FatalLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
	}

	var writers = make([]io.Writer, 0)
	output := loggerConfig.Output
	if strings.Contains(output, "console") {
		writers = append(writers, os.Stdout)
	}
	if strings.Contains(output, "file") {
		writers = append(writers, &lumberjack.Logger{
			Filename:   loggerConfig.FileName,
			MaxSize:    loggerConfig.MaxSize,
			MaxBackups: loggerConfig.MaxBackups,
			MaxAge:     loggerConfig.MaxAge,
			LocalTime:  true,
		})
	}
	if len(writers) == 0 {
		writers = append(writers, os.Stdout)
	}
	//同时写文件和屏幕
	fileAndStdoutWriter := io.MultiWriter(writers...)
	logger.SetOutput(fileAndStdoutWriter)
	logger.SetReportCaller(true)
}

// GetLoggerConfig
// @Description: 获取日志配置
func GetLoggerConfig() *LoggerConfig {
	return loggerConfig
}

// getConfig
// @Description: 获取配置文件的配置
// @return *LoggerConfig
func getConfig() *LoggerConfig {
	var logConfig = &LoggerConfig{}
	c := config.GetConfig()
	err := c.Load("logger", logConfig)
	if err != nil {
		fmt.Errorf("Fatal error config file: %s \n", err)
		return &LoggerConfig{
			Format: "text",
			Level:  "info",
			Output: "console",
		}
	}
	isNotExist := bean.IsZero(loggerConfig)
	if isNotExist {
		fmt.Errorf("Init logger config error use default config")
		return &LoggerConfig{
			Format: "text",
			Level:  "info",
			Output: "console",
		}
	}
	fmt.Printf("Init logger config: %v \n", logConfig)
	return logConfig
}
