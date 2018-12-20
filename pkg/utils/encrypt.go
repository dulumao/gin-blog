package utils

import "golang.org/x/crypto/bcrypt"

// admin 123456
func Encrypt(str string) (pwd string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPassword(hashPwd, pwd string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd))
	return
}
