package entity

type Lesson struct {
	Id         int    `json:"id,omitempty"`
	DayId      int    `json:"dayId,omitempty"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Teacher    string `json:"teacher"`
	Auditorium string `json:"auditorium"`
	SubGroup   string `json:"subGroup,omitempty"`
}
