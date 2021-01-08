package main

import "golang.org/x/crypto/bcrypt"

func main() {
	password := "yjx123456"
	encryptPassword := "$2a$10$fCNZTZniMUZqd6ctSzx.h.sly8bRIqMejvVbLAtNpuo2LB88nOvT."

	err := bcrypt.CompareHashAndPassword([]byte(encryptPassword), []byte(password))
	if err != nil{
		println("error")
		return
	}
	println("success")
}
