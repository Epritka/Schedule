package entity

type Day struct {
	Number  *int     `json:"number,omitempty"`
	Lessons []Lesson `json:"lessons"`
}

type DayFilter struct {
	GroupId int    `json:"groupId"`
	Date    string `json:"date"`
}
