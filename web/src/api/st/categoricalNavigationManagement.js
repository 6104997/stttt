import service from '@/utils/request'
// @Tags CategoricalNavigationManagement
// @Summary 创建分类导航管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CategoricalNavigationManagement true "创建分类导航管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /categoricalNavigationManagement/createCategoricalNavigationManagement [post]
export const createCategoricalNavigationManagement = (data) => {
  return service({
    url: '/categoricalNavigationManagement/createCategoricalNavigationManagement',
    method: 'post',
    data
  })
}

// @Tags CategoricalNavigationManagement
// @Summary 删除分类导航管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CategoricalNavigationManagement true "删除分类导航管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /categoricalNavigationManagement/deleteCategoricalNavigationManagement [delete]
export const deleteCategoricalNavigationManagement = (params) => {
  return service({
    url: '/categoricalNavigationManagement/deleteCategoricalNavigationManagement',
    method: 'delete',
    params
  })
}

// @Tags CategoricalNavigationManagement
// @Summary 批量删除分类导航管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除分类导航管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /categoricalNavigationManagement/deleteCategoricalNavigationManagement [delete]
export const deleteCategoricalNavigationManagementByIds = (params) => {
  return service({
    url: '/categoricalNavigationManagement/deleteCategoricalNavigationManagementByIds',
    method: 'delete',
    params
  })
}

// @Tags CategoricalNavigationManagement
// @Summary 更新分类导航管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CategoricalNavigationManagement true "更新分类导航管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /categoricalNavigationManagement/updateCategoricalNavigationManagement [put]
export const updateCategoricalNavigationManagement = (data) => {
  return service({
    url: '/categoricalNavigationManagement/updateCategoricalNavigationManagement',
    method: 'put',
    data
  })
}

// @Tags CategoricalNavigationManagement
// @Summary 用id查询分类导航管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CategoricalNavigationManagement true "用id查询分类导航管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /categoricalNavigationManagement/findCategoricalNavigationManagement [get]
export const findCategoricalNavigationManagement = (params) => {
  return service({
    url: '/categoricalNavigationManagement/findCategoricalNavigationManagement',
    method: 'get',
    params
  })
}

// @Tags CategoricalNavigationManagement
// @Summary 分页获取分类导航管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取分类导航管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /categoricalNavigationManagement/getCategoricalNavigationManagementList [get]
export const getCategoricalNavigationManagementList = (params) => {
  return service({
    url: '/categoricalNavigationManagement/getCategoricalNavigationManagementList',
    method: 'get',
    params
  })
}

// @Tags CategoricalNavigationManagement
// @Summary 不需要鉴权的分类导航管理接口
// @accept application/json
// @Produce application/json
// @Param data query stReq.CategoricalNavigationManagementSearch true "分页获取分类导航管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /categoricalNavigationManagement/getCategoricalNavigationManagementPublic [get]
export const getCategoricalNavigationManagementPublic = () => {
  return service({
    url: '/categoricalNavigationManagement/getCategoricalNavigationManagementPublic',
    method: 'get',
  })
}
// GetAListOfNavigationCategories 获取导航数据
// @Tags CategoricalNavigationManagement
// @Summary 获取导航数据
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /categoricalNavigationManagement/getAListOfNavigationCategories [GET]
export const getAListOfNavigationCategories = () => {
  return service({
    url: '/categoricalNavigationManagement/getAListOfNavigationCategories',
    method: 'GET'
  })
}
