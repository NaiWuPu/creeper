package service

import (
	"creeper/app"
	"creeper/cache"
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

//验证access_token(使用缓存处理)
func CheckAccessToken(appID uint, accessToken string) *OutPut {
	cacheAppId, _, err := cache.GetAccessToken(accessToken)
	if err != nil {
		return &OutPut{Error: err}
	}
	if cacheAppId != appID {
		return &OutPut{Error: errors.New("access_token 无效")}
	}
	return &OutPut{}
}
