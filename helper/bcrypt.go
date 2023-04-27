package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(p string) string {
	salt := 10
	password := []byte(p)
	hashedpassword, _ := bcrypt.GenerateFromPassword(password, salt)
	return string(hashedpassword)
}

func CompareHashedPassword(h, p []byte) bool {
	hash, password := []byte(h), []byte(p)
	err := bcrypt.CompareHashAndPassword(hash, password)
	if err != nil {
		return false
	}
	return err == nil
}
