package main

import (
	"log"
	"os"
	//"fmt"
)

var (
	LogThreadHold int = 30 //to add the log if
	loggers       map[int]*log.Logger
)

var LogCategory = map[int]string{
	10: "[Error]    ",
	20: "[Warning]  ",
	30: "[Info]     ",
	40: "[Trace]    ",
}

func init() {
	handler := os.Stdout
	loggers = make(map[int]*log.Logger)
	for k, v := range LogCategory {
		loggers[k] = log.New(handler, v, log.Ldate|log.Ltime|log.Lshortfile)
	}
	loggers[99] = log.New(handler, "[undefine] ", log.Ldate|log.Ltime|log.Lshortfile) //undefined LogCategory
}

/*AddLog: generate log
func AddLog(logCategory int, logString string, logArgs ...string) {
}
*/
func AddLog(logCategory int, logString string, args ...string) {
	if logCategory <= LogThreadHold {
		_, ok := loggers[logCategory]
		if ok {
			loggers[logCategory].Println(logString)
		} else {
			loggers[99].Println(logString)
		}
	}
}
