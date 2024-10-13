package st

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type InterfaceToAccessTheDataRouter struct{}

// InitInterfaceToAccessTheDataRouter 初始化 接口访问数据管理 路由信息
func (s *InterfaceToAccessTheDataRouter) InitInterfaceToAccessTheDataRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	interfaceToAccessTheDataRouter := Router.Group("interfaceToAccessTheData").Use(middleware.OperationRecord())
	interfaceToAccessTheDataRouterWithoutRecord := Router.Group("interfaceToAccessTheData")
	interfaceToAccessTheDataRouterWithoutAuth := PublicRouter.Group("interfaceToAccessTheData")
	{
		interfaceToAccessTheDataRouter.POST("createInterfaceToAccessTheData", interfaceToAccessTheDataApi.CreateInterfaceToAccessTheData)             // 新建接口访问数据管理
		interfaceToAccessTheDataRouter.DELETE("deleteInterfaceToAccessTheData", interfaceToAccessTheDataApi.DeleteInterfaceToAccessTheData)           // 删除接口访问数据管理
		interfaceToAccessTheDataRouter.DELETE("deleteInterfaceToAccessTheDataByIds", interfaceToAccessTheDataApi.DeleteInterfaceToAccessTheDataByIds) // 批量删除接口访问数据管理
		interfaceToAccessTheDataRouter.PUT("updateInterfaceToAccessTheData", interfaceToAccessTheDataApi.UpdateInterfaceToAccessTheData)              // 更新接口访问数据管理
	}
	{
		interfaceToAccessTheDataRouterWithoutRecord.GET("findInterfaceToAccessTheData", interfaceToAccessTheDataApi.FindInterfaceToAccessTheData)       // 根据ID获取接口访问数据管理
		interfaceToAccessTheDataRouterWithoutRecord.GET("getInterfaceToAccessTheDataList", interfaceToAccessTheDataApi.GetInterfaceToAccessTheDataList) // 获取接口访问数据管理列表
	}
	{
		interfaceToAccessTheDataRouterWithoutAuth.GET("getInterfaceToAccessTheDataPublic", interfaceToAccessTheDataApi.GetInterfaceToAccessTheDataPublic) // 接口访问数据管理开放接口
	}
}
