package service

import (
	"creeper/app"
	"creeper/db_conn"
	"creeper/orm_model"
	"errors"
)

//生成access_token
func CreateAccessToken(appID uint, secret string) *OutPut {
	db, err := db_conn.GetGormDB()
	defer db.Close()
	if err != nil {
		return &OutPut{Error: err}
	}
	//验证appID与secret
	appOutPut := FindAppBySecret(appID, secret)
	if appOutPut.Error != nil {
		return appOutPut
	}
	//access_token model
	accessTokenModel := new(orm_model.AccessToken)
	accessTokenModel.AppID = appID
	//32位token
	accessTokenModel.Token = app.UniqueId()
	accessTokenModel.Secret = secret
	db.Create(accessTokenModel)
	if accessTokenModel.ID <= 0 {
		return &OutPut{Error: errors.New("创建access_token失败")}
	}

	return &OutPut{Data: accessTokenModel}
}
