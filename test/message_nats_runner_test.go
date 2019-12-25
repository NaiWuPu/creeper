package test

import (
	"creeper/message"
	"testing"
)

func Test_runner(t *testing.T) {
	err := message.SendMessage("test1", []byte("test1"))
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	err = message.SendMessage("test2", []byte("test2"))
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	err = message.SendMessage("test3", []byte("test3"))
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	err = message.SendMessage("test4", []byte("test4"))
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	err = message.SendMessage("test5", []byte("test5"))
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
