package st

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/st"
	stReq "github.com/flipped-aurora/gin-vue-admin/server/model/st/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ArouselImageApi struct{}

// CreateArouselImage 创建arouselImage表
// @Tags ArouselImage
// @Summary 创建arouselImage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body st.ArouselImage true "创建arouselImage表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /arouselImage/createArouselImage [post]
func (arouselImageApi *ArouselImageApi) CreateArouselImage(c *gin.Context) {
	var arouselImage st.ArouselImage
	err := c.ShouldBindJSON(&arouselImage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = arouselImageService.CreateArouselImage(&arouselImage)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteArouselImage 删除arouselImage表
// @Tags ArouselImage
// @Summary 删除arouselImage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body st.ArouselImage true "删除arouselImage表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /arouselImage/deleteArouselImage [delete]
func (arouselImageApi *ArouselImageApi) DeleteArouselImage(c *gin.Context) {
	ID := c.Query("ID")
	err := arouselImageService.DeleteArouselImage(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteArouselImageByIds 批量删除arouselImage表
// @Tags ArouselImage
// @Summary 批量删除arouselImage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /arouselImage/deleteArouselImageByIds [delete]
func (arouselImageApi *ArouselImageApi) DeleteArouselImageByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := arouselImageService.DeleteArouselImageByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateArouselImage 更新arouselImage表
// @Tags ArouselImage
// @Summary 更新arouselImage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body st.ArouselImage true "更新arouselImage表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /arouselImage/updateArouselImage [put]
func (arouselImageApi *ArouselImageApi) UpdateArouselImage(c *gin.Context) {
	var arouselImage st.ArouselImage
	err := c.ShouldBindJSON(&arouselImage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = arouselImageService.UpdateArouselImage(arouselImage)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindArouselImage 用id查询arouselImage表
// @Tags ArouselImage
// @Summary 用id查询arouselImage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query st.ArouselImage true "用id查询arouselImage表"
// @Success 200 {object} response.Response{data=st.ArouselImage,msg=string} "查询成功"
// @Router /arouselImage/findArouselImage [get]
func (arouselImageApi *ArouselImageApi) FindArouselImage(c *gin.Context) {
	ID := c.Query("ID")
	rearouselImage, err := arouselImageService.GetArouselImage(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rearouselImage, c)
}

// GetArouselImageList 分页获取arouselImage表列表
// @Tags ArouselImage
// @Summary 分页获取arouselImage表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query stReq.ArouselImageSearch true "分页获取arouselImage表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /arouselImage/getArouselImageList [get]
func (arouselImageApi *ArouselImageApi) GetArouselImageList(c *gin.Context) {
	var pageInfo stReq.ArouselImageSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := arouselImageService.GetArouselImageInfoList(pageInfo)
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

// GetArouselImagePublic 不需要鉴权的arouselImage表接口
// @Tags ArouselImage
// @Summary 不需要鉴权的arouselImage表接口
// @accept application/json
// @Produce application/json
// @Param data query stReq.ArouselImageSearch true "分页获取arouselImage表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /arouselImage/getArouselImagePublic [get]
func (arouselImageApi *ArouselImageApi) GetArouselImagePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	arouselImageService.GetArouselImagePublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的arouselImage表接口信息",
	}, "获取成功", c)
}

// GetImage 获取轮播图图片
// @Tags ArouselImage
// @Summary 方法介绍
// @accept application/json
// @Produce application/json
// @Param data query stReq.ArouselImageSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /arouselImage/GetImage [GET]
func (arouselImageApi *ArouselImageApi) GetImage(c *gin.Context) {
	// 获取轮播图图片
	listImage, err := arouselImageService.GetImage()
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData(listImage, c)
}
