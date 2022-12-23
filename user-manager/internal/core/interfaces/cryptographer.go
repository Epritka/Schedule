package interfaces

type Cryptographer interface {
	Encrypt(password string) (string, error)
	Match(hashPassword, password string) bool
}
