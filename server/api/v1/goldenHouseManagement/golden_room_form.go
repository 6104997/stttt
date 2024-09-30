package goldenHouseManagement

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/goldenHouseManagement"
	goldenHouseManagementReq "github.com/flipped-aurora/gin-vue-admin/server/model/goldenHouseManagement/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GoldenRoomFormApi struct{}

// CreateGoldenRoomForm 创建金房管理
// @Tags GoldenRoomForm
// @Summary 创建金房管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body goldenHouseManagement.GoldenRoomForm true "创建金房管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /goldenRoomForm/createGoldenRoomForm [post]
func (goldenRoomFormApi *GoldenRoomFormApi) CreateGoldenRoomForm(c *gin.Context) {
	var goldenRoomForm goldenHouseManagement.GoldenRoomForm
	err := c.ShouldBindJSON(&goldenRoomForm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = goldenRoomFormService.CreateGoldenRoomForm(&goldenRoomForm)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteGoldenRoomForm 删除金房管理
// @Tags GoldenRoomForm
// @Summary 删除金房管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body goldenHouseManagement.GoldenRoomForm true "删除金房管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /goldenRoomForm/deleteGoldenRoomForm [delete]
func (goldenRoomFormApi *GoldenRoomFormApi) DeleteGoldenRoomForm(c *gin.Context) {
	ID := c.Query("ID")
	err := goldenRoomFormService.DeleteGoldenRoomForm(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteGoldenRoomFormByIds 批量删除金房管理
// @Tags GoldenRoomForm
// @Summary 批量删除金房管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /goldenRoomForm/deleteGoldenRoomFormByIds [delete]
func (goldenRoomFormApi *GoldenRoomFormApi) DeleteGoldenRoomFormByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := goldenRoomFormService.DeleteGoldenRoomFormByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateGoldenRoomForm 更新金房管理
// @Tags GoldenRoomForm
// @Summary 更新金房管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body goldenHouseManagement.GoldenRoomForm true "更新金房管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /goldenRoomForm/updateGoldenRoomForm [put]
func (goldenRoomFormApi *GoldenRoomFormApi) UpdateGoldenRoomForm(c *gin.Context) {
	var goldenRoomForm goldenHouseManagement.GoldenRoomForm
	err := c.ShouldBindJSON(&goldenRoomForm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = goldenRoomFormService.UpdateGoldenRoomForm(goldenRoomForm)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindGoldenRoomForm 用id查询金房管理
// @Tags GoldenRoomForm
// @Summary 用id查询金房管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query goldenHouseManagement.GoldenRoomForm true "用id查询金房管理"
// @Success 200 {object} response.Response{data=goldenHouseManagement.GoldenRoomForm,msg=string} "查询成功"
// @Router /goldenRoomForm/findGoldenRoomForm [get]
func (goldenRoomFormApi *GoldenRoomFormApi) FindGoldenRoomForm(c *gin.Context) {
	ID := c.Query("ID")
	regoldenRoomForm, err := goldenRoomFormService.GetGoldenRoomForm(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(regoldenRoomForm, c)
}

// GetGoldenRoomFormList 分页获取金房管理列表
// @Tags GoldenRoomForm
// @Summary 分页获取金房管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query goldenHouseManagementReq.GoldenRoomFormSearch true "分页获取金房管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /goldenRoomForm/getGoldenRoomFormList [get]
func (goldenRoomFormApi *GoldenRoomFormApi) GetGoldenRoomFormList(c *gin.Context) {
	var pageInfo goldenHouseManagementReq.GoldenRoomFormSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := goldenRoomFormService.GetGoldenRoomFormInfoList(pageInfo)
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

// GetAListOfGoldenHouses 获取金房列表
// @Tags GoldenRoomForm
// @Summary 获取金房列表
// @accept application/json
// @Produce application/json
// @Param data query goldenHouseManagementReq.GoldenRoomFormSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /goldenRoomForm/getAListOfGoldenHouses [GET]
func (goldenRoomFormApi *GoldenRoomFormApi) GetAListOfGoldenHouses(c *gin.Context) {
	// 请添加自己的业务逻辑
	roomdata, err := goldenRoomFormService.GetAListOfGoldenHouses()
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData(roomdata, c)
}

// listOfRooms
// AddTheGoldenRoomData 添加房间数据
// @Tags GoldenRoomForm
// @Summary 添加房间数据
// @accept application/json
// @Produce application/json
// @Param data query goldenHouseManagementReq.GoldenRoomFormSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /goldenRoomForm/addTheGoldenRoomData [GET]
func (goldenRoomFormApi *GoldenRoomFormApi) AddTheGoldenRoomData(c *gin.Context) {
	// 请添加自己的业务逻辑
	var goldenRoomForm goldenHouseManagementReq.GoldenRoomFormSearch
	err := c.ShouldBindQuery(&goldenRoomForm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	roomdata, err := goldenRoomFormService.AddTheGoldenRoomData(goldenRoomForm)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData(roomdata, c)
}

// DeleteTheGoldenRoomData 删除金房数据
// @Tags GoldenRoomForm
// @Summary 删除金房数据
// @accept application/json
// @Produce application/json
// @Param data query goldenHouseManagementReq.GoldenRoomFormSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /goldenRoomForm/deleteTheGoldenRoomData [GET]
func (goldenRoomFormApi *GoldenRoomFormApi) DeleteTheGoldenRoomData(c *gin.Context) {
	// 请添加自己的业务逻辑
	ID := c.Query("id")
	err := goldenRoomFormService.DeleteTheGoldenRoomData(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateTheGoldenRoomData 更新金房数据
// @Tags GoldenRoomForm
// @Summary 更新金房数据
// @accept application/json
// @Produce application/json
// @Param data query goldenHouseManagementReq.GoldenRoomFormSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /goldenRoomForm/updateTheGoldenRoomData [GET]
func (goldenRoomFormApi *GoldenRoomFormApi) UpdateTheGoldenRoomData(c *gin.Context) {
	// 请添加自己的业务逻辑
	var goldenRoomForm goldenHouseManagementReq.GoldenRoomFormSearch
	err := c.ShouldBindQuery(&goldenRoomForm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = goldenRoomFormService.UpdateTheGoldenRoomData(goldenRoomForm)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)

}
