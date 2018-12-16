package utils

import (
	"crypto/md5"
	"encoding/hex"
	//"github.com/golang/crypto"
)

var salt = "6(&^%!@#$%^&*(&*^((6151"

func Encrypt(str string, salts ...string) string {

	hash := md5.New()
	hash.Write([]byte(str))
	for _, s := range salts {
		hash.Write([]byte(s))
	}
	hash.Write([]byte(salt))
	sum := hash.Sum(nil)
	// 16进制转字符串
	return hex.EncodeToString(sum)
}
