package st

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	ArouselImageRouter
	AppUserRouter
}

var (
	arouselImageApi = api.ApiGroupApp.StApiGroup.ArouselImageApi
	appUserApi      = api.ApiGroupApp.StApiGroup.AppUserApi
)
