package test

import (
	"creeper/service"
	"log"
	"testing"
)

func Test_create_app(t *testing.T) {
	ot := service.CreateApp("测试")
	if ot.Error != nil {
		t.Error(ot.Error)
		t.Fail()
		return
	}
	log.Println(ot.Data)
}

func Test_reset_secret_app(t *testing.T) {
	ot := service.ResetSecret(1)
	if ot.Error != nil {
		t.Error(ot.Error)
		t.Fail()
		return
	}
	log.Println(ot.Data)
}
