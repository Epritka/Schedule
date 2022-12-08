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
