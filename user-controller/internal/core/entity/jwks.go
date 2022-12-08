package entity

type Jwks struct {
	Keys []Key `json:"keys"`
}

type Key struct {
	Id   string `json:"kid"`
	Type string `json:"kty"`
	K    string `json:"k"`
}
