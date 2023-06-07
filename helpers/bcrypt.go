package helpers

import "golang.org/x/crypto/bcrypt"

func HashPass(password string) string {
	salt := 12
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), salt)

	return string(hashedPassword)
}

func ComparePass(hashPass, password []byte) bool {
	hash, pass := []byte(hashPass), []byte(password)

	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err == nil
}
