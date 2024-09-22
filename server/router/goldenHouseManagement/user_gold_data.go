package goldenHouseManagement

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserGoldDataRouter struct{}

func (s *UserGoldDataRouter) InitUserGoldDataRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	userGoldDataRouter := Router.Group("userGoldData").Use(middleware.OperationRecord())
	userGoldDataRouterWithoutRecord := Router.Group("userGoldData")
	userGoldDataRouterWithoutAuth := PublicRouter.Group("userGoldData")
	{
		userGoldDataRouter.POST("createUserGoldData", userGoldDataApi.CreateUserGoldData)
		userGoldDataRouter.DELETE("deleteUserGoldData", userGoldDataApi.DeleteUserGoldData)
		userGoldDataRouter.DELETE("deleteUserGoldDataByIds", userGoldDataApi.DeleteUserGoldDataByIds)
		userGoldDataRouter.PUT("updateUserGoldData", userGoldDataApi.UpdateUserGoldData)
	}
	{
		userGoldDataRouterWithoutRecord.GET("findUserGoldData", userGoldDataApi.FindUserGoldData)
		userGoldDataRouterWithoutRecord.GET("getUserGoldDataList", userGoldDataApi.GetUserGoldDataList)
	}
	{
		userGoldDataRouterWithoutAuth.GET("getUserGoldDataPublic", userGoldDataApi.GetUserGoldDataPublic)
		userGoldDataRouterWithoutAuth.GET("addTheUserSGoldData", userGoldDataApi.AddTheUserSGoldData)
	}
}
