package st

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/st"
	stReq "github.com/flipped-aurora/gin-vue-admin/server/model/st/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CategoricalNavigationManagementApi struct{}

// CreateCategoricalNavigationManagement 创建分类导航管理
// @Tags CategoricalNavigationManagement
// @Summary 创建分类导航管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body st.CategoricalNavigationManagement true "创建分类导航管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /categoricalNavigationManagement/createCategoricalNavigationManagement [post]
func (categoricalNavigationManagementApi *CategoricalNavigationManagementApi) CreateCategoricalNavigationManagement(c *gin.Context) {
	var categoricalNavigationManagement st.CategoricalNavigationManagement
	err := c.ShouldBindJSON(&categoricalNavigationManagement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = categoricalNavigationManagementService.CreateCategoricalNavigationManagement(&categoricalNavigationManagement)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteCategoricalNavigationManagement 删除分类导航管理
// @Tags CategoricalNavigationManagement
// @Summary 删除分类导航管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body st.CategoricalNavigationManagement true "删除分类导航管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /categoricalNavigationManagement/deleteCategoricalNavigationManagement [delete]
func (categoricalNavigationManagementApi *CategoricalNavigationManagementApi) DeleteCategoricalNavigationManagement(c *gin.Context) {
	ID := c.Query("ID")
	err := categoricalNavigationManagementService.DeleteCategoricalNavigationManagement(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCategoricalNavigationManagementByIds 批量删除分类导航管理
// @Tags CategoricalNavigationManagement
// @Summary 批量删除分类导航管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /categoricalNavigationManagement/deleteCategoricalNavigationManagementByIds [delete]
func (categoricalNavigationManagementApi *CategoricalNavigationManagementApi) DeleteCategoricalNavigationManagementByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := categoricalNavigationManagementService.DeleteCategoricalNavigationManagementByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateCategoricalNavigationManagement 更新分类导航管理
// @Tags CategoricalNavigationManagement
// @Summary 更新分类导航管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body st.CategoricalNavigationManagement true "更新分类导航管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /categoricalNavigationManagement/updateCategoricalNavigationManagement [put]
func (categoricalNavigationManagementApi *CategoricalNavigationManagementApi) UpdateCategoricalNavigationManagement(c *gin.Context) {
	var categoricalNavigationManagement st.CategoricalNavigationManagement
	err := c.ShouldBindJSON(&categoricalNavigationManagement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = categoricalNavigationManagementService.UpdateCategoricalNavigationManagement(categoricalNavigationManagement)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindCategoricalNavigationManagement 用id查询分类导航管理
// @Tags CategoricalNavigationManagement
// @Summary 用id查询分类导航管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query st.CategoricalNavigationManagement true "用id查询分类导航管理"
// @Success 200 {object} response.Response{data=st.CategoricalNavigationManagement,msg=string} "查询成功"
// @Router /categoricalNavigationManagement/findCategoricalNavigationManagement [get]
func (categoricalNavigationManagementApi *CategoricalNavigationManagementApi) FindCategoricalNavigationManagement(c *gin.Context) {
	ID := c.Query("ID")
	recategoricalNavigationManagement, err := categoricalNavigationManagementService.GetCategoricalNavigationManagement(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(recategoricalNavigationManagement, c)
}

// GetCategoricalNavigationManagementList 分页获取分类导航管理列表
// @Tags CategoricalNavigationManagement
// @Summary 分页获取分类导航管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query stReq.CategoricalNavigationManagementSearch true "分页获取分类导航管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /categoricalNavigationManagement/getCategoricalNavigationManagementList [get]
func (categoricalNavigationManagementApi *CategoricalNavigationManagementApi) GetCategoricalNavigationManagementList(c *gin.Context) {
	var pageInfo stReq.CategoricalNavigationManagementSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := categoricalNavigationManagementService.GetCategoricalNavigationManagementInfoList(pageInfo)
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

// getAListOfNavigationCategories
// GetAListOfNavigationCategories 获取导航数据
// @Tags CategoricalNavigationManagement
// @Summary 获取导航数据
// @accept application/json
// @Produce application/json
// @Param data query stReq.CategoricalNavigationManagementSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /categoricalNavigationManagement/getAListOfNavigationCategories [GET]
func (categoricalNavigationManagementApi *CategoricalNavigationManagementApi) GetAListOfNavigationCategories(c *gin.Context) {
	// 请添加自己的业务逻辑
	list, err := categoricalNavigationManagementService.GetAListOfNavigationCategories()
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData(list, c)
}

//customGetArticles
