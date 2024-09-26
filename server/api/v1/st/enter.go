package st

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	AppUserApi
	ArouselImageApi
	AnnouncementManagementApi
	ArticleManagementApi
}

var (
	appUserService                = service.ServiceGroupApp.StServiceGroup.AppUserService
	arouselImageService           = service.ServiceGroupApp.StServiceGroup.ArouselImageService
	announcementManagementService = service.ServiceGroupApp.StServiceGroup.AnnouncementManagementService
	articleManagementService      = service.ServiceGroupApp.StServiceGroup.ArticleManagementService
)
