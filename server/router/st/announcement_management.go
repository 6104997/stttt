package st

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AnnouncementManagementRouter struct{}

func (s *AnnouncementManagementRouter) InitAnnouncementManagementRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	announcementManagementRouter := Router.Group("announcementManagement").Use(middleware.OperationRecord())
	announcementManagementRouterWithoutRecord := Router.Group("announcementManagement")
	announcementManagementRouterWithoutAuth := PublicRouter.Group("announcementManagement")
	{
		announcementManagementRouter.POST("createAnnouncementManagement", announcementManagementApi.CreateAnnouncementManagement)
		announcementManagementRouter.DELETE("deleteAnnouncementManagement", announcementManagementApi.DeleteAnnouncementManagement)
		announcementManagementRouter.DELETE("deleteAnnouncementManagementByIds", announcementManagementApi.DeleteAnnouncementManagementByIds)
		announcementManagementRouter.PUT("updateAnnouncementManagement", announcementManagementApi.UpdateAnnouncementManagement)
	}
	{
		announcementManagementRouterWithoutRecord.GET("findAnnouncementManagement", announcementManagementApi.FindAnnouncementManagement)
		announcementManagementRouterWithoutRecord.GET("getAnnouncementManagementList", announcementManagementApi.GetAnnouncementManagementList)
	}
	{
		announcementManagementRouterWithoutAuth.GET("getAnnouncementManagementPublic", announcementManagementApi.GetAnnouncementManagementPublic)
		announcementManagementRouterWithoutAuth.GET("getAListOfAnnouncements", announcementManagementApi.GetAListOfAnnouncements)
	}
}
