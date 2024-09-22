package goldenHouseManagement

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GoldenRoomFormRouter struct{}

func (s *GoldenRoomFormRouter) InitGoldenRoomFormRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	goldenRoomFormRouter := Router.Group("goldenRoomForm").Use(middleware.OperationRecord())
	goldenRoomFormRouterWithoutRecord := Router.Group("goldenRoomForm")
	goldenRoomFormRouterWithoutAuth := PublicRouter.Group("goldenRoomForm")
	{
		goldenRoomFormRouter.POST("createGoldenRoomForm", goldenRoomFormApi.CreateGoldenRoomForm)
		goldenRoomFormRouter.DELETE("deleteGoldenRoomForm", goldenRoomFormApi.DeleteGoldenRoomForm)
		goldenRoomFormRouter.DELETE("deleteGoldenRoomFormByIds", goldenRoomFormApi.DeleteGoldenRoomFormByIds)
		goldenRoomFormRouter.PUT("updateGoldenRoomForm", goldenRoomFormApi.UpdateGoldenRoomForm)
	}
	{
		goldenRoomFormRouterWithoutRecord.GET("findGoldenRoomForm", goldenRoomFormApi.FindGoldenRoomForm)
		goldenRoomFormRouterWithoutRecord.GET("getGoldenRoomFormList", goldenRoomFormApi.GetGoldenRoomFormList)
	}
	{
		goldenRoomFormRouterWithoutAuth.GET("getGoldenRoomFormPublic", goldenRoomFormApi.GetGoldenRoomFormPublic)
		goldenRoomFormRouterWithoutAuth.GET("getAListOfGoldenHouses", goldenRoomFormApi.GetAListOfGoldenHouses)
		goldenRoomFormRouterWithoutAuth.GET("addTheGoldenRoomData", goldenRoomFormApi.AddTheGoldenRoomData)
		goldenRoomFormRouterWithoutAuth.GET("deleteTheGoldenRoomData", goldenRoomFormApi.DeleteTheGoldenRoomData)
		goldenRoomFormRouterWithoutAuth.GET("updateTheGoldenRoomData", goldenRoomFormApi.UpdateTheGoldenRoomData)
	}
}
