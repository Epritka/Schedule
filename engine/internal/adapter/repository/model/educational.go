package model

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
