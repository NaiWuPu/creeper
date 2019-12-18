package service

import (
	"creeper/app"
	"creeper/db_conn"
	"creeper/orm_model"
	"errors"
	"github.com/sirupsen/logrus"
)

//创建应用
func CreateApp(name string) *OutPut {
	appModel := new(orm_model.App)
	appModel.Name = name
	appModel.Secret = app.UniqueId()
	db, err := db_conn.GetGormDB()
	defer db.Close()
	if err != nil {
		return &OutPut{Error: err}
	}
	db.Create(appModel)
	if appModel.ID <= 0 {
		return &OutPut{Error: errors.New("创建失败")}
	}
	return &OutPut{Data: appModel}
}

//重置应用secret
func ResetSecret(appId uint) *OutPut {
	db, err := db_conn.GetGormDB()
	defer db.Close()
	if err != nil {
		return &OutPut{Error: err}
	}
	appModel := new(orm_model.App)
	//确认id对应app信息
	db.First(appModel, map[string]interface{}{"id": appId})
	if appModel.ID <= 0 {
		logrus.Info(appModel)
		return &OutPut{Error: errors.New("未找到对应app")}
	}
	//记录原来的secret
	appResetSecret := new(orm_model.AppResetSecret)
	appResetSecret.AppID = appModel.ID
	appResetSecret.Secret = appModel.Secret
	//开启事务处理
	tx := db.Begin()
	defer tx.Close()
	//记录旧secret
	tx.Create(appResetSecret)
	if appResetSecret.ID <= 0 {
		logrus.Info(appResetSecret)
		tx.Rollback()
		return &OutPut{Error: errors.New("记录secret失败")}
	}
	//保存新secret
	logrus.Info("修改前secret：", appModel.Secret)
	tx.Model(appModel).Updates(map[string]interface{}{"secret": app.UniqueId()})
	logrus.Info("修改后secret：", appModel.Secret)
	if appModel.Secret == appResetSecret.Secret {
		tx.Rollback()
		return &OutPut{Error: errors.New("修改secret失败")}
	}
	tx.Commit()
	return &OutPut{Data: appModel}
}

//通过appId & secret 获取app
func FindAppBySecret(appId uint, secret string) *OutPut {
	db, err := db_conn.GetGormDB()
	defer db.Close()
	if err != nil {
		return &OutPut{Error: err}
	}
	appModel := new(orm_model.App)
	db.Find(appModel, map[string]interface{}{
		"id":     appId,
		"secret": secret})
	if appModel.ID <= 0 {
		return &OutPut{Error: errors.New("未找到对应应用")}
	}
	return &OutPut{Data: appModel}
}
