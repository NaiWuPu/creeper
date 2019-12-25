package test

import (
	"creeper/app"
	"log"
	"testing"
)

var encryptionAesTestKey = "01234567890123456789012345678901"
var encryptionAesTestOrig = "hello world"

func Test_encrypt(t *testing.T) {
	//加密
	encryptOrig := app.AesEncrypt(encryptionAesTestOrig, encryptionAesTestKey)
	log.Println("加密内容：", encryptOrig)
	//解密
	if app.AesDecrypt(encryptOrig, encryptionAesTestKey) != encryptionAesTestOrig {
		t.Error("加密 != 解密")
		t.Fail()
		return
	}
}
