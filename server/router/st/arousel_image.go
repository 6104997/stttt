package st

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ArouselImageRouter struct{}

// InitArouselImageRouter 初始化 arouselImage表 路由信息
func (s *ArouselImageRouter) InitArouselImageRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	arouselImageRouter := Router.Group("arouselImage").Use(middleware.OperationRecord())
	arouselImageRouterWithoutRecord := Router.Group("arouselImage")
	arouselImageRouterWithoutAuth := PublicRouter.Group("arouselImage")
	{
		arouselImageRouter.POST("createArouselImage", arouselImageApi.CreateArouselImage)             // 新建arouselImage表
		arouselImageRouter.DELETE("deleteArouselImage", arouselImageApi.DeleteArouselImage)           // 删除arouselImage表
		arouselImageRouter.DELETE("deleteArouselImageByIds", arouselImageApi.DeleteArouselImageByIds) // 批量删除arouselImage表
		arouselImageRouter.PUT("updateArouselImage", arouselImageApi.UpdateArouselImage)              // 更新arouselImage表
	}
	{
		arouselImageRouterWithoutRecord.GET("findArouselImage", arouselImageApi.FindArouselImage)       // 根据ID获取arouselImage表
		arouselImageRouterWithoutRecord.GET("getArouselImageList", arouselImageApi.GetArouselImageList) // 获取arouselImage表列表
	}
	{
		arouselImageRouterWithoutAuth.GET("getArouselImagePublic", arouselImageApi.GetArouselImagePublic) // arouselImage表开放接口
	}
}
