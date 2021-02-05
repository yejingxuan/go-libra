package main

import (
	"encoding/base64"
	"github.com/yejingxuan/go-libra/pkg/utils"
)

func main() {
	key := "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCHoRwcuCKprPUhZh3IG6+NxfHiIgXp24aMJ6I69iJMKtInUgymmdB4RcZ7FfX2yRUj/aiXzGPYTyVErLo2fb88Yi/YOse3S/j31OjswYe/1X4PsH5Jo52PKNtjs151nVc8UzQ8mHRKZzKLH+ySRsZLTWVs7nIwDlvcqydj/NI4ZwIDAQAB"
	toString := base64.StdEncoding.EncodeToString([]byte(key))

	utils.RSAEncrypt("111111", toString)

	/*password := "yjx123456"
	encryptPassword := "$2a$10$fCNZTZniMUZqd6ctSzx.h.sly8bRIqMejvVbLAtNpuo2LB88nOvT."

	err := bcrypt.CompareHashAndPassword([]byte(encryptPassword), []byte(password))
	if err != nil{
		println("error")
		return
	}
	println("success")

	md5Text := utils.MD5Encrypt("123456")
	println(md5Text)*/

	/*decrypt, _ := utils.AesDecrypt("4Mo9jgiv9FD7vQrdb7Zhgg==", "e1d941aaf4eab500a134d75k", "Tk7XMAIqC9IfG6De")
	println(string(decrypt))*/


	/*encrypt, _ := utils.BcryptEncrypt("111111", 10)
	println(encrypt)*/
	/*res := "01000"
	i := res[1:]
	println(i)
	*/
	decrypt, _ := utils.AesDecrypt("4NbQe6JsHKp+e4QstOBZbA==", "e1d941aaf4eab500", "Tk7XMAIqC9IfG6De")
	println(string(decrypt))

	en, _ := utils.AesEncrypt("Fxpc@100200", "e1d941aaf4eab500", "Tk7XMAIqC9IfG6De")
	println(string(en))

	en2, _ := utils.AesEncrypt("111111", "e1d941aaf4eab500", "Tk7XMAIqC9IfG6De")
	println(string(en2))

	/*tracer := "http://81.70.90.102:9994/#/single-login?uuid=5e2aa79cd14b4c5dbf4c81c7a95123ba"
	comma := strings.Index(tracer, "uuid=")

	fmt.Println(tracer[comma+len("uuid="):])*/

	/*v := float64(1234)
	//s1 := strconv.FormatFloat(v, 'E', -1, 64)//float32s2 := strconv.FormatFloat(v, 'E', -1, 64)
	s1 := strconv.FormatFloat(v, 'f', -1, 64)
	println(s1)*/
}
