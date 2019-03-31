package main

import (
	"log"
	"os"
	"runtime"
	"path/filepath"
	//"fmt"
)

var (
	LogThreadHold int = 40 //to add the log if
	logger *log.Logger
	//loggers       map[int]*log.Logger
	logFilePath string = `.\cdll.log`
)

var LogCategory = map[int]string{
	10: "[Error]    ",
	20: "[Warning]  ",
	30: "[Info]     ",
	40: "[Trace]    ",
}

func init() {
	// If the file doesn't exist, create it, or append to the file
	logFile, _ := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	logger = log.New(logFile, "", log.Ldate|log.Ltime)
	//logger.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
	//logger.SetOutput(logFile)
	
	/*
	handler := os.Stdout
	loggers = make(map[int]*log.Logger)
	for k, v := range LogCategory {
		loggers[k] = log.New(handler, v, log.Ldate|log.Ltime|log.Lshortfile)
	}
	loggers[99] = log.New(handler, "[undefine] ", log.Ldate|log.Ltime|log.Lshortfile) //undefined LogCategory
	*/
}

//deLogCategory: decode LogCategory int to string
func deLogCategory(intCategory int) string {
	strCategory, ok := LogCategory[intCategory]
	if !ok {
		strCategory = "[undefine] "
	}
	return strCategory
}

/*AddLog: generate log
func AddLog(logCategory int, logString string, logArgs ...string) {
}
*/
func AddLog(logCategory int, logString string, args ...string) {
	if logCategory <= LogThreadHold {
		skip := 1
		_, path, line, _ := runtime.Caller(skip)
		file := filepath.Base(path)
		strCategory := deLogCategory(logCategory)
		logger.Printf("%s, l:%6d, %s, %s\r\n",file, line, strCategory,logString)
	}
	/*
	if logCategory <= LogThreadHold {
		_, ok := loggers[logCategory]
		if ok {
			loggers[logCategory].Println(logString)
		} else {
			loggers[99].Println(logString)
		}
	}
	*/
}
