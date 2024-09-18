package st

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppUserRouter struct{}

// InitAppUserRouter 初始化 appUser表 路由信息
func (s *AppUserRouter) InitAppUserRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	appUserRouter := Router.Group("appUser").Use(middleware.OperationRecord())
	appUserRouterWithoutRecord := Router.Group("appUser")
	appUserRouterWithoutAuth := PublicRouter.Group("appUser")
	{
		appUserRouter.POST("createAppUser", appUserApi.CreateAppUser)             // 新建appUser表
		appUserRouter.DELETE("deleteAppUser", appUserApi.DeleteAppUser)           // 删除appUser表
		appUserRouter.DELETE("deleteAppUserByIds", appUserApi.DeleteAppUserByIds) // 批量删除appUser表
		appUserRouter.PUT("updateAppUser", appUserApi.UpdateAppUser)              // 更新appUser表
	}
	{
		appUserRouterWithoutRecord.GET("findAppUser", appUserApi.FindAppUser)       // 根据ID获取appUser表
		appUserRouterWithoutRecord.GET("getAppUserList", appUserApi.GetAppUserList) // 获取appUser表列表
	}
	{
		appUserRouterWithoutAuth.GET("getAppUserPublic", appUserApi.GetAppUserPublic) // appUser表开放接口
	}
}
