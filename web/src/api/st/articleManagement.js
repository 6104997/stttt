import service from '@/utils/request'
// @Tags ArticleManagement
// @Summary 创建文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ArticleManagement true "创建文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /articleManagement/createArticleManagement [post]
export const createArticleManagement = (data) => {
  return service({
    url: '/articleManagement/createArticleManagement',
    method: 'post',
    data
  })
}

// @Tags ArticleManagement
// @Summary 删除文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ArticleManagement true "删除文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /articleManagement/deleteArticleManagement [delete]
export const deleteArticleManagement = (params) => {
  return service({
    url: '/articleManagement/deleteArticleManagement',
    method: 'delete',
    params
  })
}

// @Tags ArticleManagement
// @Summary 批量删除文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /articleManagement/deleteArticleManagement [delete]
export const deleteArticleManagementByIds = (params) => {
  return service({
    url: '/articleManagement/deleteArticleManagementByIds',
    method: 'delete',
    params
  })
}

// @Tags ArticleManagement
// @Summary 更新文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ArticleManagement true "更新文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /articleManagement/updateArticleManagement [put]
export const updateArticleManagement = (data) => {
  return service({
    url: '/articleManagement/updateArticleManagement',
    method: 'put',
    data
  })
}

// @Tags ArticleManagement
// @Summary 用id查询文章
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ArticleManagement true "用id查询文章"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /articleManagement/findArticleManagement [get]
export const findArticleManagement = (params) => {
  return service({
    url: '/articleManagement/findArticleManagement',
    method: 'get',
    params
  })
}

// @Tags ArticleManagement
// @Summary 分页获取文章列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取文章列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /articleManagement/getArticleManagementList [get]
export const getArticleManagementList = (params) => {
  return service({
    url: '/articleManagement/getArticleManagementList',
    method: 'get',
    params
  })
}

// @Tags ArticleManagement
// @Summary 不需要鉴权的文章接口
// @accept application/json
// @Produce application/json
// @Param data query stReq.ArticleManagementSearch true "分页获取文章列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /articleManagement/getArticleManagementPublic [get]
export const getArticleManagementPublic = () => {
  return service({
    url: '/articleManagement/getArticleManagementPublic',
    method: 'get',
  })
}
// GetAListOfArticles 获取文章列表
// @Tags ArticleManagement
// @Summary 获取文章列表
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /articleManagement/getAListOfArticles [GET]
export const getAListOfArticles = () => {
  return service({
    url: '/articleManagement/getAListOfArticles',
    method: 'GET'
  })
}
// GetArticleManagementByUUID 根据UUID获取文章记录
// @Tags ArticleManagement
// @Summary 根据UUID获取文章记录
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /articleManagement/getArticleManagementByUUID [GET]
export const getArticleManagementByUUID = () => {
  return service({
    url: '/articleManagement/getArticleManagementByUUID',
    method: 'GET'
  })
}
