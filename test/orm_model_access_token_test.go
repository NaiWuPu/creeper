package test

import (
	"creeper/cache"
	"creeper/db_conn"
	"creeper/orm_model"
	"log"
	"testing"
)

func Test_after_create(t *testing.T) {
	at := new(orm_model.AccessToken)
	at.AppID = 1
	at.Secret = "1"
	at.Token = "token"
	db, err := db_conn.GetGormDB()
	if err != nil {
		log.Println(err)
		t.Fail()
		return
	}
	defer db.Close()
	db.Create(at)
	if at.ID <= 0 {
		log.Println(at)
		t.Fail()
		return
	}
	//验证缓存情况
	appId, exp, err := cache.GetAccessToken(at.Token)
	if err != nil {
		log.Println(err)
		t.Fail()
		return
	}
	if appId != at.AppID {
		log.Println(appId)
		t.Fail()
		return
	}
	log.Println(exp.Unix())
}
