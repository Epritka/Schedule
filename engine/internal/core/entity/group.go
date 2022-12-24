package entity

type Group struct {
	Id                       int    `json:"id,omitempty"`
	Name                     string `json:"name"`
	FacultyId                int    `json:"facultyId,omitempty"`
	YearId                   int    `json:"yearId,omitempty"`
	EducationalInstitutionId int    `json:"educationalInstitutionId,omitempty"`
}
