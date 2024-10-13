package st

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	AppUserRouter
	ArouselImageRouter
	AnnouncementManagementRouter
	ArticleManagementRouter
	CategoricalNavigationManagementRouter
	InterfaceToAccessTheDataRouter
}

var (
	appUserApi                         = api.ApiGroupApp.StApiGroup.AppUserApi
	arouselImageApi                    = api.ApiGroupApp.StApiGroup.ArouselImageApi
	announcementManagementApi          = api.ApiGroupApp.StApiGroup.AnnouncementManagementApi
	articleManagementApi               = api.ApiGroupApp.StApiGroup.ArticleManagementApi
	categoricalNavigationManagementApi = api.ApiGroupApp.StApiGroup.CategoricalNavigationManagementApi
	interfaceToAccessTheDataApi        = api.ApiGroupApp.StApiGroup.InterfaceToAccessTheDataApi
)
