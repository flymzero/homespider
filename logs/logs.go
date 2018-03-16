package logs

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/flymzero/homespider"
)

var (
	logPath, logWebName, logOtherName string
	webLog                            *log.Logger //, otherLog
)

func init() {
	timeStr := time.Now().Format(homespider.LOGS_FILE_PREFIX)
	logWebName = timeStr + homespider.LOGS_TYPE_WEB
	logOtherName = timeStr + homespider.LOGS_TYPE_OTHER

	logPath = path.Dir(homespider.LOGS_PATH)
	if logPath == "." {
		logPath = path.Dir(os.Args[0])
	}

	logWebFile, err := os.OpenFile(path.Join(logPath, logWebName), os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		log.Fatal(err)
	}
	defer logWebFile.Close()
	webLog = log.New(logWebFile, "", log.Ldate|log.Ltime|log.Lshortfile)
	webLog.Print("test")
}

func Temp() {

}
