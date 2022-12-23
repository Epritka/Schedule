package entity

import "time"

type ApiTokenFilters struct {
	UserId int `json:"userId"`
	Id     int `json:"id"`

	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type ApiToken struct {
	Id          int        `json:"id"`
	UserId      int        `json:"userId"`
	Token       string     `json:"token"`
	Description string     `json:"description"`
	Ttl         *time.Time `json:"ttl,omitempty"`
}
