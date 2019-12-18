package orm_model

import (
	"creeper/cache"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"time"
)

type AccessToken struct {
	gorm.Model
	AppID  uint   `gorm:"column:app_id"`          //app id
	Secret string `gorm:"column:secret;size:255"` //32位秘钥
	Token  string `gorm:"column:token"`           //AccessToken
}

func (AccessToken) TableName() string {
	return "c_access_token"
}

func (m *AccessToken) AfterCreate(scope *gorm.Scope) error {
	//设置token到缓存变量中
	err := cache.SetAccessToken(
		m.Token,
		m.AppID,
		m.CreatedAt.Add(time.Duration(cache.AccessTokenExpires*int64(time.Second))))
	if err != nil {
		logrus.Error("AccessToken AfterCreate : ", err)
	}
	return nil
}
