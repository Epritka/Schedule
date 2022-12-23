package entity

type AuthSourceType string

const (
	LdapAuthSourceType     AuthSourceType = "ldap"
	KerberosAuthSourceType AuthSourceType = "kerberos"
)

type AuthSourceFilters struct {
	Id         int            `json:"id"`
	Name       string         `json:"name"`
	SourceType AuthSourceType `json:"sourceType"`
	Url        string         `json:"url"`

	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type AuthSource struct {
	Id         int            `json:"id"`
	Name       string         `json:"name"`
	SourceType AuthSourceType `json:"sourceType"`
	Url        string         `json:"url"`
	IsActive   bool           `json:"isActive"`
	Meta       map[string]any `json:"meta"`
}
