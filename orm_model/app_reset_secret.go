package orm_model

import "github.com/jinzhu/gorm"

type AppResetSecret struct {
	gorm.Model
	AppID  uint   `gorm:"column:app_id"`          //app id
	Secret string `gorm:"column:secret;size:255"` //32位秘钥
}

func (AppResetSecret) TableName() string {
	return "c_app_reset_secret"
}
