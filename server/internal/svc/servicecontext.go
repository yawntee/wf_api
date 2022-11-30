package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"wf_api/server/internal/config"
	"wf_api/server/model"
)

type Models struct {
	model.UserModel
	model.GameUserModel
}

type ServiceContext struct {
	Config config.Config
	Models
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config: c,
		Models: Models{
			UserModel:     model.NewUserModel(conn),
			GameUserModel: model.NewGameUserModel(conn),
		},
	}
}
