package st

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CategoricalNavigationManagementRouter struct{}

func (s *CategoricalNavigationManagementRouter) InitCategoricalNavigationManagementRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	categoricalNavigationManagementRouter := Router.Group("categoricalNavigationManagement").Use(middleware.OperationRecord())
	categoricalNavigationManagementRouterWithoutRecord := Router.Group("categoricalNavigationManagement")
	categoricalNavigationManagementRouterWithoutAuth := PublicRouter.Group("categoricalNavigationManagement")
	{
		categoricalNavigationManagementRouter.POST("createCategoricalNavigationManagement", categoricalNavigationManagementApi.CreateCategoricalNavigationManagement)
		categoricalNavigationManagementRouter.DELETE("deleteCategoricalNavigationManagement", categoricalNavigationManagementApi.DeleteCategoricalNavigationManagement)
		categoricalNavigationManagementRouter.DELETE("deleteCategoricalNavigationManagementByIds", categoricalNavigationManagementApi.DeleteCategoricalNavigationManagementByIds)
		categoricalNavigationManagementRouter.PUT("updateCategoricalNavigationManagement", categoricalNavigationManagementApi.UpdateCategoricalNavigationManagement)
	}
	{
		categoricalNavigationManagementRouterWithoutRecord.GET("findCategoricalNavigationManagement", categoricalNavigationManagementApi.FindCategoricalNavigationManagement)
		categoricalNavigationManagementRouterWithoutRecord.GET("getCategoricalNavigationManagementList", categoricalNavigationManagementApi.GetCategoricalNavigationManagementList)
	}
	{

		categoricalNavigationManagementRouterWithoutAuth.GET("getAListOfNavigationCategories", categoricalNavigationManagementApi.GetAListOfNavigationCategories)
	}
}
