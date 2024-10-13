package st

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/st"
	stReq "github.com/flipped-aurora/gin-vue-admin/server/model/st/request"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"go.uber.org/zap"
)

type InterfaceToAccessTheDataApi struct{}

// CreateInterfaceToAccessTheData 创建接口访问数据管理
// @Tags InterfaceToAccessTheData
// @Summary 创建接口访问数据管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body st.InterfaceToAccessTheData true "创建接口访问数据管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /interfaceToAccessTheData/createInterfaceToAccessTheData [post]
func (interfaceToAccessTheDataApi *InterfaceToAccessTheDataApi) CreateInterfaceToAccessTheData(c *gin.Context) {
	var interfaceToAccessTheData st.InterfaceToAccessTheData
	err := c.ShouldBindJSON(&interfaceToAccessTheData)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	interfaceToAccessTheData.Uuid = uuid.Must(uuid.NewV4())
	err = interfaceToAccessTheDataService.CreateInterfaceToAccessTheData(&interfaceToAccessTheData)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteInterfaceToAccessTheData 删除接口访问数据管理
// @Tags InterfaceToAccessTheData
// @Summary 删除接口访问数据管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body st.InterfaceToAccessTheData true "删除接口访问数据管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /interfaceToAccessTheData/deleteInterfaceToAccessTheData [delete]
func (interfaceToAccessTheDataApi *InterfaceToAccessTheDataApi) DeleteInterfaceToAccessTheData(c *gin.Context) {
	ID := c.Query("ID")
	err := interfaceToAccessTheDataService.DeleteInterfaceToAccessTheData(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteInterfaceToAccessTheDataByIds 批量删除接口访问数据管理
// @Tags InterfaceToAccessTheData
// @Summary 批量删除接口访问数据管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /interfaceToAccessTheData/deleteInterfaceToAccessTheDataByIds [delete]
func (interfaceToAccessTheDataApi *InterfaceToAccessTheDataApi) DeleteInterfaceToAccessTheDataByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := interfaceToAccessTheDataService.DeleteInterfaceToAccessTheDataByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateInterfaceToAccessTheData 更新接口访问数据管理
// @Tags InterfaceToAccessTheData
// @Summary 更新接口访问数据管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body st.InterfaceToAccessTheData true "更新接口访问数据管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /interfaceToAccessTheData/updateInterfaceToAccessTheData [put]
func (interfaceToAccessTheDataApi *InterfaceToAccessTheDataApi) UpdateInterfaceToAccessTheData(c *gin.Context) {
	var interfaceToAccessTheData st.InterfaceToAccessTheData
	err := c.ShouldBindJSON(&interfaceToAccessTheData)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = interfaceToAccessTheDataService.UpdateInterfaceToAccessTheData(interfaceToAccessTheData)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindInterfaceToAccessTheData 用id查询接口访问数据管理
// @Tags InterfaceToAccessTheData
// @Summary 用id查询接口访问数据管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query st.InterfaceToAccessTheData true "用id查询接口访问数据管理"
// @Success 200 {object} response.Response{data=st.InterfaceToAccessTheData,msg=string} "查询成功"
// @Router /interfaceToAccessTheData/findInterfaceToAccessTheData [get]
func (interfaceToAccessTheDataApi *InterfaceToAccessTheDataApi) FindInterfaceToAccessTheData(c *gin.Context) {
	ID := c.Query("ID")
	reinterfaceToAccessTheData, err := interfaceToAccessTheDataService.GetInterfaceToAccessTheData(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reinterfaceToAccessTheData, c)
}

// GetInterfaceToAccessTheDataList 分页获取接口访问数据管理列表
// @Tags InterfaceToAccessTheData
// @Summary 分页获取接口访问数据管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query stReq.InterfaceToAccessTheDataSearch true "分页获取接口访问数据管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /interfaceToAccessTheData/getInterfaceToAccessTheDataList [get]
func (interfaceToAccessTheDataApi *InterfaceToAccessTheDataApi) GetInterfaceToAccessTheDataList(c *gin.Context) {
	var pageInfo stReq.InterfaceToAccessTheDataSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := interfaceToAccessTheDataService.GetInterfaceToAccessTheDataInfoList(pageInfo)
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

// GetInterfaceToAccessTheDataPublic 不需要鉴权的接口访问数据管理接口
// @Tags InterfaceToAccessTheData
// @Summary 不需要鉴权的接口访问数据管理接口
// @accept application/json
// @Produce application/json
// @Param data query stReq.InterfaceToAccessTheDataSearch true "分页获取接口访问数据管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /interfaceToAccessTheData/getInterfaceToAccessTheDataPublic [get]
func (interfaceToAccessTheDataApi *InterfaceToAccessTheDataApi) GetInterfaceToAccessTheDataPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	//interfaceToAccessTheDataService.GetInterfaceToAccessTheDataPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的接口访问数据管理接口信息",
	}, "获取成功", c)
}

//addUserAccessRecords
