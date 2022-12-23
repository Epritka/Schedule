package entity

type Permission string

type Role struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions"`
}
