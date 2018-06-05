package log

import (
	"fmt"
	"log"
	"os"
	"strings"
	//"sync"
	"runtime"
	"time"
)

const (
	info = iota
	warn
	error
	fatal
)

type Log struct {
	logLevel int

	logChannel chan string
}

var (
	logObj             = Log{logLevel: error, logChannel: make(chan string, 1024)}
	levelMap           = make(map[string]int)
	levelStrMap        = make(map[int]string)
	LogPath     string = "../log/" //windows 下调试模式
)

func initLevelMap() {
	levelMap["fatal"] = fatal
	levelMap["error"] = error
	levelMap["warn"] = warn
	levelMap["info"] = info

	levelStrMap[fatal] = "fatal"
	levelStrMap[error] = "error"
	levelStrMap[warn] = "warn"
	levelStrMap[info] = "info"

	if strings.Contains(runtime.GOOS, "linux") {
		//linux 发布模式
		LogPath = `/home/parameterCheckww/parameterCheckService2/src/log/`
	}
}

func init() {
	initLevelMap()
}

func Fatal(v ...interface{}) {
	logObj.Fatal(v...)
}

func Error(v ...interface{}) {
	logObj.Error(v...)
}

func Warn(v ...interface{}) {
	logObj.Warn(v...)
}

func Info(v ...interface{}) {
	logObj.Info(v...)
}

func SetLogLevel(level string) {
	logObj.SetLogLevel(level)

}
func GetLogLevel() int {
	return logObj.GetLogLevel()
}

func RunLogFileThread() {

	logObj.RunLogFileThread()
}

func (self *Log) SetLogLevel(level string) {
	_, ok := levelMap[level]
	if ok {
		self.logLevel = levelMap[level]
		log.Println("logLevel has changed:" + level)
	} else {
		log.Println("loglevel is error")
	}

}
func (self *Log) GetLogLevel() int {
	return self.logLevel
}

func (self *Log) Fatal(v ...interface{}) {
	self.printWithCheckLevel(fatal, v...)
}

func (self *Log) Error(v ...interface{}) {
	self.printWithCheckLevel(error, v...)
}

func (self *Log) Warn(v ...interface{}) {
	self.printWithCheckLevel(warn, v...)
}

func (self *Log) Info(v ...interface{}) {
	self.printWithCheckLevel(info, v...)
}

func (self *Log) printWithCheckLevel(logLevel int, v ...interface{}) {
	if self.logLevel <= logLevel {
		log.Println(v...)
	}
	str := strings.TrimRight(strings.TrimPrefix(strings.TrimSpace(fmt.Sprintln(v)), "["), "]")
	self.logChannel <- "[" + levelStrMap[logLevel] + "] " + strings.TrimSpace(str) + "\n"
}

func (self *Log) WriteLogFile(log string) {
	//	logFile, err := os.OpenFile("../log/log_"+time.Now().Format("2006-01-02")+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	logFile, err := os.OpenFile(LogPath+"log_"+time.Now().Format("2006-01-02")+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open log file failed!!!")
		return
	}
	logFile.WriteString(time.Now().Format("2006-01-02 15:04:05") + " " + log)
	logFile.Close()
}

func (self *Log) RunLogFileThread() {
	for {
		strlog := <-self.logChannel
		self.WriteLogFile(strlog)

	}

}
