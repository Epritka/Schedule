package entity

type (
	WeekValue int
	Weekday   int
)

const (
	EvenWeek WeekValue = iota
	OddWeek
)

const (
	day Weekday = iota
	Monday
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
	Group                  Group                  `json:"groupId"`

	EvenWeek []Day `json:"evenWeek"`
	OddWeek  []Day `json:"oddWeek"`
}

type EducationalInstitution struct {
	Id   *int   `json:"id,omitempty"`
	Name string `json:"name"`
}

type Faculty struct {
	Id   *int   `json:"id,omitempty"`
	Name string `json:"name"`
}

type Year struct {
	Id   *int   `json:"id,omitempty"`
	Name string `json:"name"`
}

type Group struct {
	Id   *int   `json:"id,omitempty"`
	Name string `json:"name"`
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
	SubGroupNumber *int   `json:"subGroupNumber"`
}

type Time struct {
	Start string `json:"start"`
	End   string `json:"end"`
}
