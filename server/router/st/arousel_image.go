package st

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ArouselImageRouter struct{}

func (s *ArouselImageRouter) InitArouselImageRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	arouselImageRouter := Router.Group("arouselImage").Use(middleware.OperationRecord())
	arouselImageRouterWithoutRecord := Router.Group("arouselImage")
	arouselImageRouterWithoutAuth := PublicRouter.Group("arouselImage")
	{
		arouselImageRouter.POST("createArouselImage", arouselImageApi.CreateArouselImage)
		arouselImageRouter.DELETE("deleteArouselImage", arouselImageApi.DeleteArouselImage)
		arouselImageRouter.DELETE("deleteArouselImageByIds", arouselImageApi.DeleteArouselImageByIds)
		arouselImageRouter.PUT("updateArouselImage", arouselImageApi.UpdateArouselImage)
	}
	{
		arouselImageRouterWithoutRecord.GET("findArouselImage", arouselImageApi.FindArouselImage)
		arouselImageRouterWithoutRecord.GET("getArouselImageList", arouselImageApi.GetArouselImageList)
	}
	{
		arouselImageRouterWithoutAuth.GET("getArouselImagePublic", arouselImageApi.GetArouselImagePublic)
		arouselImageRouterWithoutAuth.GET("GetImage", arouselImageApi.GetImage)
	}
}
