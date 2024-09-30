package goldenHouseManagement

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/goldenHouseManagement"
	goldenHouseManagementReq "github.com/flipped-aurora/gin-vue-admin/server/model/goldenHouseManagement/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserGoldDataApi struct{}

// CreateUserGoldData 创建用户打金数据
// @Tags UserGoldData
// @Summary 创建用户打金数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body goldenHouseManagement.UserGoldData true "创建用户打金数据"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /userGoldData/createUserGoldData [post]
func (userGoldDataApi *UserGoldDataApi) CreateUserGoldData(c *gin.Context) {
	var userGoldData goldenHouseManagement.UserGoldData
	err := c.ShouldBindJSON(&userGoldData)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userGoldDataService.CreateUserGoldData(&userGoldData)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteUserGoldData 删除用户打金数据
// @Tags UserGoldData
// @Summary 删除用户打金数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body goldenHouseManagement.UserGoldData true "删除用户打金数据"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /userGoldData/deleteUserGoldData [delete]
func (userGoldDataApi *UserGoldDataApi) DeleteUserGoldData(c *gin.Context) {
	ID := c.Query("ID")
	err := userGoldDataService.DeleteUserGoldData(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteUserGoldDataByIds 批量删除用户打金数据
// @Tags UserGoldData
// @Summary 批量删除用户打金数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /userGoldData/deleteUserGoldDataByIds [delete]
func (userGoldDataApi *UserGoldDataApi) DeleteUserGoldDataByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := userGoldDataService.DeleteUserGoldDataByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateUserGoldData 更新用户打金数据
// @Tags UserGoldData
// @Summary 更新用户打金数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body goldenHouseManagement.UserGoldData true "更新用户打金数据"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /userGoldData/updateUserGoldData [put]
func (userGoldDataApi *UserGoldDataApi) UpdateUserGoldData(c *gin.Context) {
	var userGoldData goldenHouseManagement.UserGoldData
	err := c.ShouldBindJSON(&userGoldData)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userGoldDataService.UpdateUserGoldData(userGoldData)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindUserGoldData 用id查询用户打金数据
// @Tags UserGoldData
// @Summary 用id查询用户打金数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query goldenHouseManagement.UserGoldData true "用id查询用户打金数据"
// @Success 200 {object} response.Response{data=goldenHouseManagement.UserGoldData,msg=string} "查询成功"
// @Router /userGoldData/findUserGoldData [get]
func (userGoldDataApi *UserGoldDataApi) FindUserGoldData(c *gin.Context) {
	ID := c.Query("ID")
	reuserGoldData, err := userGoldDataService.GetUserGoldData(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reuserGoldData, c)
}

// GetUserGoldDataList 分页获取用户打金数据列表
// @Tags UserGoldData
// @Summary 分页获取用户打金数据列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query goldenHouseManagementReq.UserGoldDataSearch true "分页获取用户打金数据列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /userGoldData/getUserGoldDataList [get]
func (userGoldDataApi *UserGoldDataApi) GetUserGoldDataList(c *gin.Context) {
	var pageInfo goldenHouseManagementReq.UserGoldDataSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userGoldDataService.GetUserGoldDataInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetUserGoldDataPublic 不需要鉴权的用户打金数据接口
// @Tags UserGoldData
// @Summary 不需要鉴权的用户打金数据接口
// @accept application/json
// @Produce application/json
// @Param data query goldenHouseManagementReq.UserGoldDataSearch true "分页获取用户打金数据列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userGoldData/getUserGoldDataPublic [get]
func (userGoldDataApi *UserGoldDataApi) GetUserGoldDataPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	userGoldDataService.GetUserGoldDataPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的用户打金数据接口信息",
	}, "获取成功", c)
}
