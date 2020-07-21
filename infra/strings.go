package infra

import (
	"crypto/md5"
	"fmt"
	"github.com/gofrs/uuid"
)

// 密码处理=========================
func pwdEncrypt(pwd, salt string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(pwd+":"+salt)))
}

func PwdEncrypt(pwd string) (encryptPwd, salt string) {
	salt = UuidV4()
	encryptPwd = pwdEncrypt(pwd, salt)
	return encryptPwd, salt
}

func PwdEquals(pwd, encryptPwd, salt string) bool {
	pwd = pwdEncrypt(pwd, salt)
	return pwd == encryptPwd
}

func UuidV4() string {
	_uuid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return _uuid.String()
}

// 密码处理=========================
