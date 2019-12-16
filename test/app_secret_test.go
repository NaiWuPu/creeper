package test

import (
	"creeper/app"
	"log"
	"testing"
)

func Test_secret(t *testing.T) {
	a := app.UniqueId()
	log.Println(len(a))
	//创建1000个
	var secrets []string
	for i := 0; i <= 10000; i++ {
		newSecrets := app.UniqueId()
		if len(newSecrets) == 32 {
			t.Error("长度是：", len(newSecrets))
			t.Fail()
			return
		}
		//判断是否有重复的
		for _, s := range secrets {
			if newSecrets == s {
				t.Error("有重复")
				t.Fail()
				return
			}
		}
		secrets = append(secrets, newSecrets)
	}
	log.Println(len(secrets))
}
