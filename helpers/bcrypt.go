package helpers

import "golang.org/x/crypto/bcrypt"

// GeneratePasswordHash handles generating password hash
// bcrypt hashes password of type byte
func GeneratePasswordHash(password string) string {
	// default cost is 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

func PasswordCompare(password string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err
}
