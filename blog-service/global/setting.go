package global

import (
	"go-blog-service/pkg/logger"
	"go-blog-service/pkg/setting"
)

var (
	ServerSetting *setting.ServerSettingS
	AppSetting    *setting.AppSettingS
	Logger        *logger.Logger
)
