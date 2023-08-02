package log

// Trace
// @Description: 在标准日志记录器的Trace级别记录消息
// @param: args
func Trace(args ...interface{}) {
	logger.Trace(args...)
}

// Debug
// @Description: 在标准日志记录器的Debug级别记录消息
// @param: args
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Info
// @Description: 在标准日志记录器的Info级别记录消息
// @param: args
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Notice
// @Description: 在标准日志记录器的Notice级别记录消息
// @param: args
func Notice(args ...interface{}) {
	logger.Warn(args...)
}

// Warn
// @Description: 在标准日志记录器的Warn级别记录消息
// @param: args
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Error
// @Description: 在标准日志记录器的Error级别记录消息
// @param: args
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Panic
// @Description: 在标准日志记录器的Panic级别记录消息
// @param: args
func Panic(args ...interface{}) {
	logger.Panic(args...)
}

// Fatal
// @Description: Fatal在标准日志记录器上以Fatal级别记录消息，然后进程将退出并将状态设置为1
// @param: args
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Tracef
// @Description: 在标准日志记录器的Trace级别记录格式化消息
// @param: format
// @param: args
func Tracef(format string, args ...interface{}) {
	logger.Tracef(format, args...)
}

// Debugf
// @Description: 在标准日志记录器的Debug级别记录格式化消息
// @param: format
// @param: args
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Infof
// @Description: 在标准日志记录器的Info级别记录格式化消息
// @param: format
// @param: args
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Noticef
// @Description: 在标准日志记录器的Notice级别记录格式化消息
// @param: format
// @param: args
func Noticef(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Warnf
// @Description: 在标准日志记录器的Warn级别记录格式化消息
// @param: format
// @param: args
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Errorf
// @Description: 在标准日志记录器的Error级别记录格式化消息
// @param: format
// @param: args
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// Panicf
// @Description: 在标准日志记录器的Panic级别记录格式化消息
// @param: format
// @param: args
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

// Fatalf
// @Description: Fatal在标准日志记录器上以Fatal级别记录格式化消息，然后进程将退出并将状态设置为1
// @param: format
// @param: args
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

// Traceln
// @Description: 在标准日志记录器的Trace级别记录消息
// @param: args
func Traceln(args ...interface{}) {
	logger.Traceln(args...)
}

// Debugln
// @Description: 在标准日志记录器的Debug级别记录消息
// @param: args
func Debugln(args ...interface{}) {
	logger.Debugln(args...)
}

// Infoln
// @Description: 在标准日志记录器的Info级别记录消息
// @param: args
func Infoln(args ...interface{}) {
	logger.Infoln(args...)
}

// Noticeln
// @Description: 在标准日志记录器的Notice级别记录消息
// @param: args
func Noticeln(args ...interface{}) {
	logger.Warnln(args...)
}

// Warnln
// @Description: 在标准日志记录器的Warn级别记录消息
// @param: args
func Warnln(args ...interface{}) {
	logger.Warnln(args...)
}

// Errorln
// @Description: 在标准日志记录器的Error级别记录消息
// @param: args
func Errorln(args ...interface{}) {
	logger.Errorln(args...)
}

// Panicln
// @Description: 在标准日志记录器的Panic级别记录消息
// @param: args
func Panicln(args ...interface{}) {
	logger.Panicln(args...)
}

// Fatalln
// @Description: Fatal在标准日志记录器上以Fatal级别记录消息，然后进程将退出并将状态设置为1
// @param: args
func Fatalln(args ...interface{}) {
	logger.Fatalln(args...)
}
