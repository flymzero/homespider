package job51

import "time"

/************配置*************/
var (
	JOB51_KEYWORDS = []string{
		"go",
		"java",
		"ios",
		"android",
		"php",
	}
)

const (
	//2全文 1公司
	JOB51_KEYWORDTYPE = "2"
	//
	JOB51_AREA = "杭州"
	//发布日期 0=24小时内 1=近三天 2=近一周 3=近一月 空=所有
	JOB51_ISSUE_DATE = "0"
)

/************mysql*************/
const (
	JOB51_TABLE_CITY   = "job51_city"
	JOB51_TABLE_Detail = "job_detail"
)

type Job51TableCity struct {
	Id       int    `json:"id"`
	CityId   string `json:"city_id"`
	CityName string `json:"city_name"`
}

type Job51TableDetail struct {
	Id       int    `json:"id"`
	Keyword  string `json:"keyword"`
	Platform string `json:"platform"`
	JobId    string `json:"job_id"`
	JobName  string `json:"job_name"`
	JobMoney string `json:"job_money"`
	//JobTime  string `json:"job_time"`
	JobAge     string    `json:"job_age"`
	Company    string    `json:"company"`
	Address    string    `json:"address"`
	Detail     string    `json:"detail"`
	Lng        float64   `json:"lng"`
	Lat        float64   `json:"lat"`
	Province   string    `json:"province"`
	UpdateTime time.Time `json:"updat_time"`
}

/************url*************/
const (
	//
	JOB51_PLATFORM = "51job"
	//热门城市
	JOB51_HOTCITY_URL = "http://js.51jobcdn.com/in/js/h5/dd/hotcityall.js?"
	//所有城市
	JOB51_AREA_URL = "http://js.51jobcdn.com/in/js/h5/dd/jobarea.js?"
	//搜索返回列表
	JOB51_LIST_URL = "http://m.51job.com/search/joblist.php?keyword=%s&keywordtype=%s&jobarea=%s&issuedate=%s"
	//job详细信息
	JOB51_DETAIL_URL = "http://m.51job.com/search/jobdetail.php?jobid=%s&jobtype=0"
	//job详细地址
	JOB51_ADRESS_URL = "http://m.51job.com/search/jobmap.php?jobid=%s&t=1"
)
