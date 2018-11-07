package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"resource-backend/pkg/file"
	"runtime"
	"time"
)

var (
	logSavePath = "runtime/logs/"
	logSaveName = "log"
	LogFileExt = "log"
	TimeFormat = "20060102"
)

type Level int

var (
	F *os.File
	DefaultPrefix = ""
	DefaultCallerDepth = 2

	logger *log.Logger
	logPrefix = ""
	levelFlags = []string{"DEBUG", "INFO", "WARNING", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init()  {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func getLogFilePath() string {
	return fmt.Sprintf("%s", logSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s%s", logSaveName, time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		dir, _ := os.Getwd()
		err := file.MkDir(dir + "/" + getLogFilePath())
		if err != nil {
			panic(err)
		}
	case os.IsPermission(err):
		log.Fatalf("Permission:%v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile:%v", err)
	}
	return handle
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}

func setPrefix(level Level) {
	_, logFile, line, ok := runtime.Caller(DefaultCallerDepth)
	fmt.Println(ok)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(logFile), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	fmt.Println(logPrefix)
	logger.SetPrefix(logPrefix)
}
