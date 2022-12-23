package cryptographer

import "golang.org/x/crypto/bcrypt"

func (c *cryptographer) Encrypt(password string) (string, error) {
	passwordBytes := []byte(password)

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, c.Cost)
	if err != nil {
		return "", err
	}

	return string(hashedPasswordBytes), err
}
