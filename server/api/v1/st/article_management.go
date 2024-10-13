package st

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/st"
	stReq "github.com/flipped-aurora/gin-vue-admin/server/model/st/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"go.uber.org/zap"
)

type ArticleManagementApi struct{}

// CreateArticleManagement 创建文章
// @Tags ArticleManagement
// @Summary 创建文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body st.ArticleManagement true "创建文章"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /articleManagement/createArticleManagement [post]
func (articleManagementApi *ArticleManagementApi) CreateArticleManagement(c *gin.Context) {
	var articleManagement st.ArticleManagement
	err := c.ShouldBindJSON(&articleManagement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	articleManagement.Uuid = uuid.Must(uuid.NewV4())
	err = articleManagementService.CreateArticleManagement(&articleManagement)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteArticleManagement 删除文章
// @Tags ArticleManagement
// @Summary 删除文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body st.ArticleManagement true "删除文章"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /articleManagement/deleteArticleManagement [delete]
func (articleManagementApi *ArticleManagementApi) DeleteArticleManagement(c *gin.Context) {
	ID := c.Query("ID")
	err := articleManagementService.DeleteArticleManagement(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteArticleManagementByIds 批量删除文章
// @Tags ArticleManagement
// @Summary 批量删除文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /articleManagement/deleteArticleManagementByIds [delete]
func (articleManagementApi *ArticleManagementApi) DeleteArticleManagementByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := articleManagementService.DeleteArticleManagementByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateArticleManagement 更新文章
// @Tags ArticleManagement
// @Summary 更新文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body st.ArticleManagement true "更新文章"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /articleManagement/updateArticleManagement [put]
func (articleManagementApi *ArticleManagementApi) UpdateArticleManagement(c *gin.Context) {
	var articleManagement st.ArticleManagement
	err := c.ShouldBindJSON(&articleManagement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = articleManagementService.UpdateArticleManagement(articleManagement)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindArticleManagement 用id查询文章
// @Tags ArticleManagement
// @Summary 用id查询文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query st.ArticleManagement true "用id查询文章"
// @Success 200 {object} response.Response{data=st.ArticleManagement,msg=string} "查询成功"
// @Router /articleManagement/findArticleManagement [get]
func (articleManagementApi *ArticleManagementApi) FindArticleManagement(c *gin.Context) {
	ID := c.Query("ID")
	rearticleManagement, err := articleManagementService.GetArticleManagement(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rearticleManagement, c)
}

// GetArticleManagementList 分页获取文章列表
// @Tags ArticleManagement
// @Summary 分页获取文章列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query stReq.ArticleManagementSearch true "分页获取文章列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /articleManagement/getArticleManagementList [get]
func (articleManagementApi *ArticleManagementApi) GetArticleManagementList(c *gin.Context) {
	var pageInfo stReq.ArticleManagementSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := articleManagementService.GetArticleManagementInfoList(pageInfo)
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

// GetAListOfArticles 获取文章列表
// @Tags ArticleManagement
// @Summary 获取文章列表
// @accept application/json
// @Produce application/json
// @Param data query stReq.ArticleManagementSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /articleManagement/getAListOfArticles [GET]
func (articleManagementApi *ArticleManagementApi) GetAListOfArticles(c *gin.Context) {
	// 请添加自己的业务逻辑
	articleClassificationId := c.Query("articleClassificationId")
	var pageInfo stReq.ArticleManagementSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := articleManagementService.GetAListOfArticles(pageInfo, articleClassificationId)

	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)

	id := utils.GetUserID(c)
	clientIP := c.ClientIP()
	if clientIP == "::1" {
		clientIP = "127.0.0.1"
	}
	err = interfaceToAccessTheDataService.AddUserAccessRecords(id, clientIP, "获取文章列表", "/articleManagement/getAListOfArticles")
}

// GetArticleManagementByUUID 根据UUID获取文章记录
// @Tags ArticleManagement
// @Summary 根据UUID获取文章记录
// @accept application/json
// @Produce application/json
// @Param data query stReq.ArticleManagementSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /articleManagement/getArticleManagementByUUID [GET]
func (articleManagementApi *ArticleManagementApi) GetArticleManagementByUUID(c *gin.Context) {
	// 请添加自己的业务逻辑
	UUID := c.Query("uuid")
	rearticleManagement, err := articleManagementService.GetArticleManagementByUUID(UUID)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData(rearticleManagement, c)
	id := utils.GetUserID(c)
	clientIP := c.ClientIP()
	if clientIP == "::1" {
		clientIP = "127.0.0.1"
	}
	err = interfaceToAccessTheDataService.AddUserAccessRecords(id, clientIP, "根据UUID获取文章记录", "/articleManagement/getArticleManagementByUUID")
}
