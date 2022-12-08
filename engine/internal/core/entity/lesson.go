package entity

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
