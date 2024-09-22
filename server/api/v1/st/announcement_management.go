package st

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/st"
	stReq "github.com/flipped-aurora/gin-vue-admin/server/model/st/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AnnouncementManagementApi struct{}

// CreateAnnouncementManagement 创建公告管理
// @Tags AnnouncementManagement
// @Summary 创建公告管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body st.AnnouncementManagement true "创建公告管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /announcementManagement/createAnnouncementManagement [post]
func (announcementManagementApi *AnnouncementManagementApi) CreateAnnouncementManagement(c *gin.Context) {
	var announcementManagement st.AnnouncementManagement
	err := c.ShouldBindJSON(&announcementManagement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = announcementManagementService.CreateAnnouncementManagement(&announcementManagement)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteAnnouncementManagement 删除公告管理
// @Tags AnnouncementManagement
// @Summary 删除公告管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body st.AnnouncementManagement true "删除公告管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /announcementManagement/deleteAnnouncementManagement [delete]
func (announcementManagementApi *AnnouncementManagementApi) DeleteAnnouncementManagement(c *gin.Context) {
	ID := c.Query("ID")
	err := announcementManagementService.DeleteAnnouncementManagement(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAnnouncementManagementByIds 批量删除公告管理
// @Tags AnnouncementManagement
// @Summary 批量删除公告管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /announcementManagement/deleteAnnouncementManagementByIds [delete]
func (announcementManagementApi *AnnouncementManagementApi) DeleteAnnouncementManagementByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := announcementManagementService.DeleteAnnouncementManagementByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAnnouncementManagement 更新公告管理
// @Tags AnnouncementManagement
// @Summary 更新公告管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body st.AnnouncementManagement true "更新公告管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /announcementManagement/updateAnnouncementManagement [put]
func (announcementManagementApi *AnnouncementManagementApi) UpdateAnnouncementManagement(c *gin.Context) {
	var announcementManagement st.AnnouncementManagement
	err := c.ShouldBindJSON(&announcementManagement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = announcementManagementService.UpdateAnnouncementManagement(announcementManagement)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAnnouncementManagement 用id查询公告管理
// @Tags AnnouncementManagement
// @Summary 用id查询公告管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query st.AnnouncementManagement true "用id查询公告管理"
// @Success 200 {object} response.Response{data=st.AnnouncementManagement,msg=string} "查询成功"
// @Router /announcementManagement/findAnnouncementManagement [get]
func (announcementManagementApi *AnnouncementManagementApi) FindAnnouncementManagement(c *gin.Context) {
	ID := c.Query("ID")
	reannouncementManagement, err := announcementManagementService.GetAnnouncementManagement(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reannouncementManagement, c)
}

// GetAnnouncementManagementList 分页获取公告管理列表
// @Tags AnnouncementManagement
// @Summary 分页获取公告管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query stReq.AnnouncementManagementSearch true "分页获取公告管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /announcementManagement/getAnnouncementManagementList [get]
func (announcementManagementApi *AnnouncementManagementApi) GetAnnouncementManagementList(c *gin.Context) {
	var pageInfo stReq.AnnouncementManagementSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := announcementManagementService.GetAnnouncementManagementInfoList(pageInfo)
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

// GetAnnouncementManagementPublic 不需要鉴权的公告管理接口
// @Tags AnnouncementManagement
// @Summary 不需要鉴权的公告管理接口
// @accept application/json
// @Produce application/json
// @Param data query stReq.AnnouncementManagementSearch true "分页获取公告管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /announcementManagement/getAnnouncementManagementPublic [get]
func (announcementManagementApi *AnnouncementManagementApi) GetAnnouncementManagementPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	announcementManagementService.GetAnnouncementManagementPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的公告管理接口信息",
	}, "获取成功", c)
}

// GetAListOfAnnouncements 方法介绍
// @Tags AnnouncementManagement
// @Summary 方法介绍
// @accept application/json
// @Produce application/json
// @Param data query stReq.AnnouncementManagementSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /announcementManagement/getAListOfAnnouncements [GET]
func (announcementManagementApi *AnnouncementManagementApi) GetAListOfAnnouncements(c *gin.Context) {
	// 请添加自己的业务逻辑
	announcementList, err := announcementManagementService.GetAListOfAnnouncements()
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData(announcementList, c)
}
