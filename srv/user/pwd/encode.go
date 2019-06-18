package pwd

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

func EncodePwd(code string, pasword string) string {
	plainPwd := fmt.Sprintf("%s:%s", code, pasword)
	data := sha256.Sum256([]byte(plainPwd))
	encodedPwd := md5.Sum(data[:])
	return fmt.Sprintf("%x", encodedPwd)
}
