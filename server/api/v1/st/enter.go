package st

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	ArouselImageApi
	AppUserApi
}

var (
	arouselImageService = service.ServiceGroupApp.StServiceGroup.ArouselImageService
	appUserService      = service.ServiceGroupApp.StServiceGroup.AppUserService
)
