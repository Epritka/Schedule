package entity

const (
	EvenWeek string = "even"
	OddWeek  string = "odd"
)

type Day struct {
	Id       int      `json:"id,omitempty"`
	GroupId  int      `json:"groupId,omitempty"`
	Number   int      `json:"number,omitempty"`
	WeekType string   `json:"weekType,omitempty"`
	Lessons  []Lesson `json:"lessons"`
}

type DayFilter struct {
	GroupId int    `json:"groupId"`
	Date    string `json:"date"`
}
