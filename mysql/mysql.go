package mysql

import (
	"database/sql"

	"github.com/flymzero/homespider"
	"github.com/flymzero/homespider/logs"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const (
	mysqlType   = homespider.MYSQL_TYEP
	mysqlDSName = homespider.MYSQL_USER + ":" + homespider.MYSQL_PWD + "@tcp(" + homespider.MYSQL_IP + ":" + homespider.MYSQL_PORT + ")/" + homespider.MYSQL_DATABASE + "?charset=utf8"
)

func init() {
	var err error
	db, err = sql.Open(mysqlType, mysqlDSName)
	if err != nil {
		logs.OtherLog(logs.Fatal, err)
	}
	err = db.Ping()
	if err != nil {
		logs.OtherLog(logs.Fatal, err)
	}

}

func Query(query string) (*sql.Rows, error) {
	return db.Query(query)
}

func CreateTable(query string) error {
	_, err := db.Exec(query)
	return err
}

func Exec(query string) (sql.Result, error) {
	return db.Exec(query)
}

func Begin(f func(t *sql.Tx)) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	f(tx)
	return tx.Commit()

}

// func Temp() {
// 	_, err := db.Exec(`CREATE TABLE job51_city (
// 		id int(20) NOT NULL,
// 		city_id varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
// 		city_name varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
// 		PRIMARY KEY (id)
// 		);`)

// }
