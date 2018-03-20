package logs

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/flymzero/homespider"
)

var (
	logPath, timeStr, logWebName, logOtherName string
	logWebFile, logOtherFile                   *os.File
	webLog, otherLog                           *log.Logger
)

var (
	Debug = "[debug] "
	Info  = "[info] "
	Warn  = "[warn] "
	Error = "[error] "
	Fatal = "[fatal] "
)

func init() {
	//路径
	logPath = path.Dir(homespider.LOGS_PATH)
	if logPath == "." {
		logPath = path.Dir(os.Args[0])
	}
}

func resetLogs() {
	curTimeStr := time.Now().Format(homespider.LOGS_FILE_PREFIX)
	if curTimeStr != timeStr {
		timeStr = curTimeStr
		logWebName = timeStr + homespider.LOGS_TYPE_WEB
		logOtherName = timeStr + homespider.LOGS_TYPE_OTHER
		//web
		if logWebFile != nil {
			logWebFile.Close()
		}
		logWebFile, err := os.OpenFile(path.Join(logPath, logWebName), os.O_RDWR|os.O_APPEND|os.O_CREATE, 0766)
		if err != nil {
			log.Fatal(err)
		}
		webLog = log.New(logWebFile, "", log.Ldate|log.Ltime|log.Lshortfile)
		//
		if logOtherFile != nil {
			logOtherFile.Close()
		}
		logOtherFile, err = os.OpenFile(path.Join(logPath, logOtherName), os.O_RDWR|os.O_APPEND|os.O_CREATE, 0766)
		if err != nil {
			log.Fatal(err)
		}
		otherLog = log.New(logOtherFile, "", log.Ldate|log.Ltime|log.Lshortfile)
	}
}

func WebLog(level string, v ...interface{}) {
	resetLogs()
	webLog.SetPrefix(level)
	s := fmt.Sprint(v...)
	webLog.Output(2, s)
	log.Printf("%s %s", level, s)
	if level == Fatal {
		os.Exit(1)
	}
}

func OtherLog(level string, v ...interface{}) {
	resetLogs()
	otherLog.SetPrefix(level)
	s := fmt.Sprint(v...)
	otherLog.Output(2, s)
	log.Printf("%s %s", level, s)
	if level == Fatal {
		os.Exit(1)
	}
}
