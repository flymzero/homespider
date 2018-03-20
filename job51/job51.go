package job51

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
	"strings"

	"github.com/flymzero/homespider/logs"
	"github.com/flymzero/homespider/mysql"
)

func GetJob51City() {
	//创建job51_city表
	jtc := Job51TableCity{}
	jt := reflect.TypeOf(jtc)
	field := jt.Field(0)
	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))
	err := mysql.CreateTable(fmt.Sprintf(`CREATE TABLE %s (
		%s int(20) NOT NULL AUTO_INCREMENT,
		%s varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
		%s varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
		PRIMARY KEY (id)
		);`,
		JOB51_TABLE_CITY, jt.Field(0).Tag.Get("json"), jt.Field(1).Tag.Get("json"), jt.Field(2).Tag.Get("json")))
	if err == nil {
		logs.OtherLog(logs.Info, fmt.Sprintf("创建【%s】表成功", JOB51_TABLE_CITY))
	}
	//获取表内的数据量
	var count int
	rows, err := mysql.Query(fmt.Sprintf("SELECT COUNT(%s) FROM %s", jt.Field(0).Tag.Get("json"), JOB51_TABLE_CITY))
	if err != nil {
		logs.OtherLog(logs.Fatal, err)
	}
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			logs.OtherLog(logs.Fatal, err)
		}
	}
	if count > 0 {
		return
	}
	//网络获取51job城市数据
	resp, err := http.Get(JOB51_AREA_SITE)
	if err != nil {
		logs.OtherLog(logs.Fatal, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.OtherLog(logs.Fatal, err)
	}
	reg := regexp.MustCompile(`;area_[cv]\[\d{1,}\]="[\p{Han}a-zA-Z0-9]+`)
	list := reg.FindAllString(string(body), -1)

	var adrMap = make(map[string]string)
	if len(list)%2 == 0 {
		var adrNum string
		for i, v := range list {
			if i%2 == 0 {
				adrNum = strings.Split(v, "\"")[1]
			} else {
				adrMap[strings.Split(v, "\"")[1]] = adrNum
			}
		}
	}
	//保存到数据库中
	err = mysql.Begin(func(t *sql.Tx) {
		for k, v := range adrMap {
			t.Exec(fmt.Sprintf("INSERT INTO %s (%s, %s) values(?,?)", JOB51_TABLE_CITY, jt.Field(2).Tag.Get("json"), jt.Field(1).Tag.Get("json")), k, v)
		}
	})
	if err != nil {
		logs.OtherLog(logs.Fatal, "城市数据存入数据库失败")
	}
	logs.OtherLog(logs.Info, fmt.Sprintf("共存入数据库【%d】条51job城市数据", len(adrMap)))
}
