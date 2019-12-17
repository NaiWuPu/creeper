package service

import (
	"creeper/app"
	"creeper/db_conn"
	"creeper/orm_model"
	"errors"
)

func CreateApp(name string) *OutPut {
	appModel := new(orm_model.App)
	appModel.Name = name
	appModel.Secret = app.UniqueId()
	db, err := db_conn.GetGormDB()
	if err != nil {
		return &OutPut{Error: err}
	}
	db.Create(appModel)
	if appModel.ID <= 0 {
		return &OutPut{Error: errors.New("创建失败")}
	}
	return &OutPut{Data: appModel}
}
