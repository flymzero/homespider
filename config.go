package homespider

/************mysql*************/
const (
	MYSQL_TYEP = "mysql"
	MYSQL_USER = "root"
	MYSQL_PWD  = "123456"
	MYSQL_IP   = "localhost"
	MYSQL_PORT = "3306"
	//MYSQL_CHARSET  = "utf8"
	MYSQL_DATABASE = "home_spider"
)

const (
	SLEEPTIMEMIN = 5
	SLEEPTIMEMAX = 15
)

/************log*************/
const (
	LOGS_PATH = "/Users/program/go/src/github.com/flymzero/" //存放log日志的绝对路径（空的话存在os.Args[0]所在的文件夹）
	//2006_01-02_hs_web.log  2006_01-02_hs_other.log
	LOGS_TYPE_WEB    = "web.log"        //log类型(文件名，用于记录前端访问日志)
	LOGS_TYPE_OTHER  = "other.log"      //log类型(文件名，用于记录其他访问日志)
	LOGS_FILE_PREFIX = "2006-01-02_hs_" //文件名前缀(需要时间格式化）
)
