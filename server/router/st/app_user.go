package st

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppUserRouter struct{}

func (s *AppUserRouter) InitAppUserRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	appUserRouter := Router.Group("appUser").Use(middleware.OperationRecord())
	appUserRouterWithoutRecord := Router.Group("appUser")
	appUserRouterWithoutAuth := PublicRouter.Group("appUser")
	{
		appUserRouter.POST("createAppUser", appUserApi.CreateAppUser)
		appUserRouter.DELETE("deleteAppUser", appUserApi.DeleteAppUser)
		appUserRouter.DELETE("deleteAppUserByIds", appUserApi.DeleteAppUserByIds)
		appUserRouter.PUT("updateAppUser", appUserApi.UpdateAppUser)
	}
	{
		appUserRouterWithoutRecord.GET("findAppUser", appUserApi.FindAppUser)
		appUserRouterWithoutRecord.GET("getAppUserList", appUserApi.GetAppUserList)
	}
	{
		appUserRouterWithoutAuth.GET("getAppUserPublic", appUserApi.GetAppUserPublic)
		appUserRouterWithoutAuth.POST("Login", appUserApi.Login)
	}
}
