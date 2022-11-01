package bd

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(pass string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(pass), 8)
	return string(password), err
}
