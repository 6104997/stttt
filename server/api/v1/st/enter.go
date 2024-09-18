package st

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	AppUserApi
	ArouselImageApi
}

var (
	appUserService      = service.ServiceGroupApp.StServiceGroup.AppUserService
	arouselImageService = service.ServiceGroupApp.StServiceGroup.ArouselImageService
)
