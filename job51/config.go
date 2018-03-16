package job51

/************mysql*************/
const (
	JOB51_TABLE_CITY = "job51_city"
)

type Job51TableCity struct {
	Id       string `json:"id"`
	CityId   string `json:"city_id"`
	CityName string `json:"city_name"`
}

const (
	JOB51_AREA_SITE = "http://js.51jobcdn.com/in/js/h5/dd/jobarea.js?"
)
