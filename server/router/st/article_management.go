package st

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ArticleManagementRouter struct{}

func (s *ArticleManagementRouter) InitArticleManagementRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	articleManagementRouter := Router.Group("articleManagement").Use(middleware.OperationRecord())
	articleManagementRouterWithoutRecord := Router.Group("articleManagement")
	articleManagementRouterWithoutAuth := PublicRouter.Group("articleManagement")

	{
		articleManagementRouter.POST("createArticleManagement", articleManagementApi.CreateArticleManagement)
		articleManagementRouter.DELETE("deleteArticleManagement", articleManagementApi.DeleteArticleManagement)
		articleManagementRouter.DELETE("deleteArticleManagementByIds", articleManagementApi.DeleteArticleManagementByIds)
		articleManagementRouter.PUT("updateArticleManagement", articleManagementApi.UpdateArticleManagement)
	}
	{
		articleManagementRouterWithoutRecord.GET("findArticleManagement", articleManagementApi.FindArticleManagement)
		articleManagementRouterWithoutRecord.GET("getArticleManagementList", articleManagementApi.GetArticleManagementList)
	}
	{
		articleManagementRouterWithoutAuth.GET("getAListOfArticles", articleManagementApi.GetAListOfArticles)
		articleManagementRouterWithoutAuth.GET("getArticleManagementByUUID", articleManagementApi.GetArticleManagementByUUID)

	}
}
