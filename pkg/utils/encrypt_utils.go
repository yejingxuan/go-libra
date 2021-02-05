package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/robertkrimen/otto"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
)

// Bcrypt 算法加密
func BcryptEncrypt(plaintext string, cost int) (string, error) {
	if cost == 0 {
		cost = bcrypt.DefaultCost
	}
	generatePass, err := bcrypt.GenerateFromPassword([]byte(plaintext), cost)
	return string(generatePass), err
}

// AES加密-CBC方式
func AesEncrypt(plaintext string, key string, iv string) (string, error) {
	encodeBytes := []byte(plaintext)
	//根据key 生成密文
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	encodeBytes = PKCS5Padding(encodeBytes, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	crypted := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(crypted, encodeBytes)

	return base64.StdEncoding.EncodeToString(crypted), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//填充
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// MD5加密
func MD5Encrypt(plaintext string) string {
	text := []byte(plaintext)
	ciphertext := md5.Sum(text)
	return fmt.Sprintf("%x", ciphertext)
}

//RSA加密
func RSAEncrypt(plaintext, publicKey string) {
	var publicKey_by_java = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCHoRwcuCKprPUhZh3IG6+NxfHiIgXp24aMJ6I6
9iJMKtInUgymmdB4RcZ7FfX2yRUj/aiXzGPYTyVErLo2fb88Yi/YOse3S/j31OjswYe/1X4PsH5J
o52PKNtjs151nVc8UzQ8mHRKZzKLH+ySRsZLTWVs7nIwDlvcqydj/NI4ZwIDAQAB
-----END PUBLIC KEY-----
`)
	data, _ := RsaEncrypt([]byte(plaintext), publicKey_by_java)
	fmt.Println(base64.StdEncoding.EncodeToString(data))
}

//利用客户端传来的公钥(java语言产生)加密
func RsaEncrypt(origData []byte, pubKey []byte) ([]byte, error) {
	block, _ := pem.Decode(pubKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

//读取js加密
func RSAWeb() {
	filePath := "./RSA.js"
	//先读入文件内容
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	vm := otto.New()

	_, err = vm.Run(string(bytes))
	if err != nil {
		panic(err)
	}

	//data := "你需要传给JS函数的参数"
	//encodeInp是JS函数的函数名
	value, err := vm.Call("encryptPass", nil, "Fxpc@100200")
	if err != nil {
		panic(err)
	}
	fmt.Println(value.String())
}
