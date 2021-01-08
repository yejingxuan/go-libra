package main

import (
	"github.com/yejingxuan/go-libra/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "yjx123456"
	encryptPassword := "$2a$10$fCNZTZniMUZqd6ctSzx.h.sly8bRIqMejvVbLAtNpuo2LB88nOvT."

	err := bcrypt.CompareHashAndPassword([]byte(encryptPassword), []byte(password))
	if err != nil{
		println("error")
		return
	}
	println("success")

	md5Text := utils.MD5Encrypt("123456")
	println(md5Text)
}
