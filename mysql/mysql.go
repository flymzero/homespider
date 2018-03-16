package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/flymzero/homespider"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const (
	mysqlType   = homespider.MYSQL_TYEP
	mysqlDSName = homespider.MYSQL_USER + ":" + homespider.MYSQL_PWD + "@tcp(" + homespider.MYSQL_IP + ":" + homespider.MYSQL_PORT + ")/" + homespider.MYSQL_DATABASE + "?charset=utf8"
)

func init() {
	var err error
	db, err = sql.Open(mysqlType, mysqlDSName) //sql.Open(homespider.MYSQL_TYEP, "root:5513505@tcp(localhost:3306)/job_spider?charset=utf8")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Print("连接成功")
	}
}

func Temp() {
	fmt.Print(db)

}
