package job51

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/flymzero/homespider"
	"github.com/flymzero/homespider/logs"
	"github.com/flymzero/homespider/mysql"
	"github.com/gocolly/colly"
)

func createTableJob51City() {
	//创建job51_city表
	jtc := Job51TableCity{}
	jt := reflect.TypeOf(jtc)
	err := mysql.CreateTable(fmt.Sprintf(`CREATE TABLE %s (
		%s int(20) NOT NULL AUTO_INCREMENT,
		%s varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
		%s varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
		PRIMARY KEY (%s)
		);`,
		JOB51_TABLE_CITY,
		jt.Field(0).Tag.Get("json"),
		jt.Field(1).Tag.Get("json"),
		jt.Field(2).Tag.Get("json"),
		jt.Field(0).Tag.Get("json")))
	if err == nil {
		logs.OtherLog(logs.Info, fmt.Sprintf("创建【%s】表成功", JOB51_TABLE_CITY))
	}
}

func createTablejobDetail() {
	jd := Job51TableDetail{}
	jt := reflect.TypeOf(jd)
	err := mysql.CreateTable(fmt.Sprintf(`CREATE TABLE %s  (
		%s int(20) NOT NULL AUTO_INCREMENT,
		%s varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
		%s varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
		%s varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
		%s varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
		%s varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
		%s varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
		%s varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
		%s varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
		%s longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
		%s double NOT NULL,
		%s double NOT NULL,
		%s varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
		%s timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0),
		PRIMARY KEY (%s)
	  );`,
		JOB51_TABLE_Detail,
		jt.Field(0).Tag.Get("json"),
		jt.Field(1).Tag.Get("json"),
		jt.Field(2).Tag.Get("json"),
		jt.Field(3).Tag.Get("json"),
		jt.Field(4).Tag.Get("json"),
		jt.Field(5).Tag.Get("json"),
		jt.Field(6).Tag.Get("json"),
		jt.Field(7).Tag.Get("json"),
		jt.Field(8).Tag.Get("json"),
		jt.Field(9).Tag.Get("json"),
		jt.Field(10).Tag.Get("json"),
		jt.Field(11).Tag.Get("json"),
		jt.Field(12).Tag.Get("json"),
		jt.Field(13).Tag.Get("json"),
		jt.Field(0).Tag.Get("json")))

	if err == nil {
		logs.OtherLog(logs.Info, fmt.Sprintf("创建【%s】表成功", JOB51_TABLE_Detail))
	}
}

func GetJob51City() {
	createTableJob51City()
	//
	jtc := Job51TableCity{}
	jt := reflect.TypeOf(jtc)
	//获取表内的数据量
	var count int
	row := mysql.QueryRow(fmt.Sprintf("SELECT COUNT(%s) FROM %s", jt.Field(0).Tag.Get("json"), JOB51_TABLE_CITY))
	if err := row.Scan(&count); err != nil {
		logs.OtherLog(logs.Fatal, err)
	}
	if count > 0 {
		return
	}
	//临时存储变量
	var adrMap = make(map[string]string)
	//网络获取51job热门城市数据
	hotResp, err := http.Get(JOB51_HOTCITY_URL)
	if err != nil {
		logs.OtherLog(logs.Fatal, err)
	}
	defer hotResp.Body.Close()
	hotBody, err := ioutil.ReadAll(hotResp.Body)
	if err != nil {
		logs.OtherLog(logs.Fatal, err)
	}
	//fmt.Print(string(hotBody))
	reg := regexp.MustCompile(`;hotareaall_[cv]\[\d{1,}\]="[\p{Han}a-zA-Z0-9]+`)
	list := reg.FindAllString(string(hotBody), -1)
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
	//网络获取51job全部城市数据
	resp, err := http.Get(JOB51_AREA_URL)
	if err != nil {
		logs.OtherLog(logs.Fatal, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.OtherLog(logs.Fatal, err)
	}
	//fmt.Print(string(body))
	reg = regexp.MustCompile(`;b?area_[cv]\[\d{1,}\]="[\p{Han}a-zA-Z0-9]+`)
	list = reg.FindAllString(string(body), -1)
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

func GetJob51List() {
	createTablejobDetail()
	//获取城市代码
	jtc := Job51TableCity{}
	jt := reflect.TypeOf(jtc)
	var city_id string
	rows, err := mysql.Query(fmt.Sprintf("SELECT %s FROM %s WHERE %s = \"%s\"",
		jt.Field(1).Tag.Get("json"),
		JOB51_TABLE_CITY,
		jt.Field(2).Tag.Get("json"),
		JOB51_AREA))
	if err != nil {
		logs.OtherLog(logs.Fatal, err)
	}
	for rows.Next() {
		if err := rows.Scan(&city_id); err != nil {
			logs.OtherLog(logs.Fatal, err)
		}
	}
	logs.OtherLog(logs.Info, fmt.Sprintf("获取城市 %s id为 【%s】", JOB51_AREA, city_id))

	//
	for _, v := range JOB51_KEYWORDS {
		logs.OtherLog(logs.Info, fmt.Sprintf("准备获取关键词【%s】的数据", v))
		requestUrl := fmt.Sprintf(JOB51_LIST_URL, v, JOB51_KEYWORDTYPE, city_id, JOB51_ISSUE_DATE)
		//
		c := colly.NewCollector(
			colly.MaxDepth(1),
		)
		c.OnHTML("div[class=items]", func(e *colly.HTMLElement) {
			//
			var jobList = make(map[string]string)
			e.ForEach("a[href]", func(_ int, el *colly.HTMLElement) {
				jobId := el.ChildAttr("b[class=jobid]", "value")
				aside := el.ChildText("aside")
				jobList[jobId] = aside
			})
			//fmt.Print(jobList)
			logs.OtherLog(logs.Info, fmt.Sprintf("当前页获取到【%d】条数据", len(jobList)))
			//对每条数据继续查找详情
			getJob51Detail(v, JOB51_PLATFORM, jobList)
		})

		c.OnHTML("a[class=next]", func(e *colly.HTMLElement) {
			//休息一下以防被封
			time.Sleep(time.Duration(rand.Intn(homespider.SLEEPTIMEMIN)+homespider.SLEEPTIMEMAX-homespider.SLEEPTIMEMIN) * time.Second)
			c.Visit(e.Attr("href"))
		})
		c.OnRequest(func(r *colly.Request) {
			logs.OtherLog(logs.Info, fmt.Sprintf("请求【%s】", r.URL.String()))
		})

		c.Visit(requestUrl)
	}
}

func getJob51Detail(keyword string, platform string, jobIds map[string]string) {
	for id, _ := range jobIds {
		//休息一下以防被封
		time.Sleep(time.Duration(rand.Intn(homespider.SLEEPTIMEMIN)+homespider.SLEEPTIMEMAX-homespider.SLEEPTIMEMIN) * time.Second)
		//
		var jobDetail = Job51TableDetail{
			Keyword:  keyword,
			Platform: platform,
			JobId:    id,
		}
		//
		c := colly.NewCollector(
			colly.MaxDepth(1),
		)
		c.OnHTML("div[class=jt]", func(e *colly.HTMLElement) {
			//
			jobDetail.JobName = e.ChildText("p")
		})
		c.OnHTML("body", func(e *colly.HTMLElement) {
			//
			jobDetail.JobMoney = e.ChildText("p[class=jp]")
			jobDetail.Company = e.ChildText("p[class=c_444]")
		})
		c.OnHTML("div[class=jd]", func(e *colly.HTMLElement) {
			//
			jobDetail.JobAge = e.ChildText("span[class=s_n]")
		})
		// c.OnHTML("a[class=arr a2]", func(e *colly.HTMLElement) {
		// 	//
		// 	jobDetail.Address = e.ChildText("span")
		// })
		c.OnHTML("div[class=ain]", func(e *colly.HTMLElement) {
			//
			jobDetail.Detail = e.ChildText("article")
			//获取详细地址
			getJob51Adress(&jobDetail)
			//插入数据库
			insertDetailIntoSql(&jobDetail)
		})

		c.OnRequest(func(r *colly.Request) {
			logs.OtherLog(logs.Info, fmt.Sprintf("请求jobid为【%s】的详细内容", id))
		})

		c.Visit(fmt.Sprintf(JOB51_DETAIL_URL, id))

	}

}

func getJob51Adress(jobdetail *Job51TableDetail) {
	//休息一下以防被封
	time.Sleep(time.Duration(rand.Intn(homespider.SLEEPTIMEMIN)+homespider.SLEEPTIMEMAX-homespider.SLEEPTIMEMIN) * time.Second)
	logs.OtherLog(logs.Info, fmt.Sprintf("请求jobid为【%s】的详细地址", jobdetail.JobId))
	//网络获取job的位置
	resp, err := http.Get(fmt.Sprintf(JOB51_ADRESS_URL, jobdetail.JobId))
	if err != nil {
		logs.OtherLog(logs.Fatal, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.OtherLog(logs.Fatal, err)
	}
	reg, err := regexp.Compile(`var G = [^;]+`)
	if err != nil {
		logs.OtherLog(logs.Fatal, err)
	}
	reg2 := regexp.MustCompile(`'[^']*'`)
	list := reg2.FindAllString(reg.FindString(string(body)), -1)
	//fmt.Print(list)
	//
	if len(list) != 4 {
		logs.OtherLog(logs.Warn, fmt.Sprintf("jobid为【%s】获取位置接口错误", jobdetail.JobId))
	} else {
		jobdetail.Address = strings.Replace(list[2], "'", "", -1)
		jobdetail.Province = strings.Replace(list[3], "'", "", -1)
		lngAndLat := strings.Split(strings.Replace(list[1], "'", "", -1), ",")
		if len(lngAndLat) != 2 {
			logs.OtherLog(logs.Warn, fmt.Sprintf("jobid为【%s】无法获取经纬度", jobdetail.JobId))
		} else {
			lng, _ := strconv.ParseFloat(lngAndLat[0], 64)
			jobdetail.Lng = lng
			lat, _ := strconv.ParseFloat(lngAndLat[1], 64)
			jobdetail.Lat = lat
		}
	}
}

func insertDetailIntoSql(jobdetail *Job51TableDetail) {
	jd := Job51TableDetail{}
	jt := reflect.TypeOf(jd)
	//判断是否已经存在
	var count int
	row := mysql.QueryRow(fmt.Sprintf("SELECT COUNT(%s) FROM %s WHERE %s = %s", jt.Field(3).Tag.Get("json"), JOB51_TABLE_Detail, jt.Field(3).Tag.Get("json"), jobdetail.JobId))
	if err := row.Scan(&count); err != nil {
		logs.OtherLog(logs.Fatal, err)
	}
	if count == 0 {
		//插入
		_, err := mysql.Exec(fmt.Sprintf(`INSERT INTO %s 
			(%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)
		 VALUES 
		 ("%s","%s","%s","%s","%s","%s","%s","%s","%s",%f,%f,"%s")`,
			JOB51_TABLE_Detail,
			jt.Field(1).Tag.Get("json"),
			jt.Field(2).Tag.Get("json"),
			jt.Field(3).Tag.Get("json"),
			jt.Field(4).Tag.Get("json"),
			jt.Field(5).Tag.Get("json"),
			jt.Field(6).Tag.Get("json"),
			jt.Field(7).Tag.Get("json"),
			jt.Field(8).Tag.Get("json"),
			jt.Field(9).Tag.Get("json"),
			jt.Field(10).Tag.Get("json"),
			jt.Field(11).Tag.Get("json"),
			jt.Field(12).Tag.Get("json"),
			jobdetail.Keyword,
			jobdetail.Platform,
			jobdetail.JobId,
			jobdetail.JobName,
			jobdetail.JobMoney,
			jobdetail.JobAge,
			jobdetail.Company,
			jobdetail.Address,
			jobdetail.Detail,
			jobdetail.Lng,
			jobdetail.Lat,
			jobdetail.Province))
		if err != nil {
			logs.OtherLog(logs.Error, err)
		} else {
			logs.OtherLog(logs.Info, fmt.Sprintf("请求jobid为【%s】插入数据库成功", jobdetail.JobId))
		}
	} else {
		//更新
		_, err := mysql.Exec(fmt.Sprintf(`UPDATE %s set 
			%s = "%s", 
			%s = "%s",
			%s = "%s",
			%s = "%s",
			%s = "%s",
			%s = "%s",
			%s = "%s",
			%s = "%s",
			%s = "%s",
			%s = %f,
			%s = %f,
			%s = "%s",
			%s = CURRENT_TIMESTAMP
			 WHERE %s = "%s"`,
			JOB51_TABLE_Detail,
			jt.Field(1).Tag.Get("json"),
			jobdetail.Keyword,
			jt.Field(2).Tag.Get("json"),
			jobdetail.Platform,
			jt.Field(3).Tag.Get("json"),
			jobdetail.JobId,
			jt.Field(4).Tag.Get("json"),
			jobdetail.JobName,
			jt.Field(5).Tag.Get("json"),
			jobdetail.JobMoney,
			jt.Field(6).Tag.Get("json"),
			jobdetail.JobAge,
			jt.Field(7).Tag.Get("json"),
			jobdetail.Company,
			jt.Field(8).Tag.Get("json"),
			jobdetail.Address,
			jt.Field(9).Tag.Get("json"),
			jobdetail.Detail,
			jt.Field(10).Tag.Get("json"),
			jobdetail.Lng,
			jt.Field(11).Tag.Get("json"),
			jobdetail.Lat,
			jt.Field(12).Tag.Get("json"),
			jobdetail.Province,
			jt.Field(12).Tag.Get("json"),

			jt.Field(3).Tag.Get("json"),
			jobdetail.JobId))
		if err != nil {
			logs.OtherLog(logs.Error, err)
		} else {
			logs.OtherLog(logs.Info, fmt.Sprintf("请求jobid为【%s】更新数据库成功", jobdetail.JobId))
		}
	}
}

func GetJobDetail() []map[string]interface{} {
	jd := Job51TableDetail{}
	jt := reflect.TypeOf(jd)

	rows, err := mysql.Query(fmt.Sprintf("SELECT %s, %s, %s, %s, %s FROM %s WHERE %s != 0 AND %s != 0",
		jt.Field(4).Tag.Get("json"),
		jt.Field(5).Tag.Get("json"),
		jt.Field(7).Tag.Get("json"),
		jt.Field(10).Tag.Get("json"),
		jt.Field(11).Tag.Get("json"),
		JOB51_TABLE_Detail,
		jt.Field(10).Tag.Get("json"),
		jt.Field(11).Tag.Get("json")))
	if err != nil {
		logs.OtherLog(logs.Fatal, err)
	}
	var jobName, jobMoney, company string
	var lng, lat float64
	var result []map[string]interface{}
	for rows.Next() {
		if err := rows.Scan(&jobName, &jobMoney, &company, &lng, &lat); err != nil {
			continue
			// logs.OtherLog(logs.Fatal, err)
		} else {
			var temp = map[string]interface{}{}
			//fmt.Printf("%s %s %s %f %f\n", jobName, jobMoney, company, lng, lat)
			temp["name"] = jobName + "\n" + jobMoney + "\n" + company
			temp["value"] = []float64{lng, lat}
			result = append(result, temp)
		}
	}
	return result
}
