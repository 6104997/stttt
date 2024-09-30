package goldenHouseManagement

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserGoldDataRouter struct{}

// InitUserGoldDataRouter 初始化 用户打金数据 路由信息
func (s *UserGoldDataRouter) InitUserGoldDataRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	userGoldDataRouter := Router.Group("userGoldData").Use(middleware.OperationRecord())
	userGoldDataRouterWithoutRecord := Router.Group("userGoldData")
	userGoldDataRouterWithoutAuth := PublicRouter.Group("userGoldData")
	{
		userGoldDataRouter.POST("createUserGoldData", userGoldDataApi.CreateUserGoldData)             // 新建用户打金数据
		userGoldDataRouter.DELETE("deleteUserGoldData", userGoldDataApi.DeleteUserGoldData)           // 删除用户打金数据
		userGoldDataRouter.DELETE("deleteUserGoldDataByIds", userGoldDataApi.DeleteUserGoldDataByIds) // 批量删除用户打金数据
		userGoldDataRouter.PUT("updateUserGoldData", userGoldDataApi.UpdateUserGoldData)              // 更新用户打金数据
	}
	{
		userGoldDataRouterWithoutRecord.GET("findUserGoldData", userGoldDataApi.FindUserGoldData)       // 根据ID获取用户打金数据
		userGoldDataRouterWithoutRecord.GET("getUserGoldDataList", userGoldDataApi.GetUserGoldDataList) // 获取用户打金数据列表
	}
	{
		userGoldDataRouterWithoutAuth.GET("getUserGoldDataPublic", userGoldDataApi.GetUserGoldDataPublic) // 用户打金数据开放接口
	}
}
