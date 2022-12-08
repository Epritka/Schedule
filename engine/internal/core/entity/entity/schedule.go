package entity

type (
	WeekType int
	Weekday  int
)

const (
	EvenWeek WeekType = iota
	OddWeek
)

const (
	Monday Weekday = 1 + iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

type Schedule struct {
	EducationalInstitution EducationalInstitution `json:"educationalInstitution"`
	Faculty                Faculty                `json:"faculty"`
	Year                   Year                   `json:"year"`
	Group                  Group                  `json:"group"`

	EvenWeek []Day `json:"evenWeek"`
	OddWeek  []Day `json:"oddWeek"`
}

type Day struct {
	Number  *int     `json:"number,omitempty"`
	Lessons []Lesson `json:"lessons"`
}

type Lesson struct {
	Id             *int   `json:"id,omitempty"`
	Time           Time   `json:"time"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Teacher        string `json:"teacher"`
	Auditorium     string `json:"auditorium"`
	SubGroupNumber *int   `json:"subGroupNumber,omitempty"`
}

type Time struct {
	Start string `json:"start"`
	End   string `json:"end"`
}