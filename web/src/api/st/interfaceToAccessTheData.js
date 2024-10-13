import service from '@/utils/request'
// @Tags InterfaceToAccessTheData
// @Summary 创建接口访问数据管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InterfaceToAccessTheData true "创建接口访问数据管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /interfaceToAccessTheData/createInterfaceToAccessTheData [post]
export const createInterfaceToAccessTheData = (data) => {
  return service({
    url: '/interfaceToAccessTheData/createInterfaceToAccessTheData',
    method: 'post',
    data
  })
}

// @Tags InterfaceToAccessTheData
// @Summary 删除接口访问数据管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InterfaceToAccessTheData true "删除接口访问数据管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /interfaceToAccessTheData/deleteInterfaceToAccessTheData [delete]
export const deleteInterfaceToAccessTheData = (params) => {
  return service({
    url: '/interfaceToAccessTheData/deleteInterfaceToAccessTheData',
    method: 'delete',
    params
  })
}

// @Tags InterfaceToAccessTheData
// @Summary 批量删除接口访问数据管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除接口访问数据管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /interfaceToAccessTheData/deleteInterfaceToAccessTheData [delete]
export const deleteInterfaceToAccessTheDataByIds = (params) => {
  return service({
    url: '/interfaceToAccessTheData/deleteInterfaceToAccessTheDataByIds',
    method: 'delete',
    params
  })
}

// @Tags InterfaceToAccessTheData
// @Summary 更新接口访问数据管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InterfaceToAccessTheData true "更新接口访问数据管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /interfaceToAccessTheData/updateInterfaceToAccessTheData [put]
export const updateInterfaceToAccessTheData = (data) => {
  return service({
    url: '/interfaceToAccessTheData/updateInterfaceToAccessTheData',
    method: 'put',
    data
  })
}

// @Tags InterfaceToAccessTheData
// @Summary 用id查询接口访问数据管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.InterfaceToAccessTheData true "用id查询接口访问数据管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /interfaceToAccessTheData/findInterfaceToAccessTheData [get]
export const findInterfaceToAccessTheData = (params) => {
  return service({
    url: '/interfaceToAccessTheData/findInterfaceToAccessTheData',
    method: 'get',
    params
  })
}

// @Tags InterfaceToAccessTheData
// @Summary 分页获取接口访问数据管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取接口访问数据管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /interfaceToAccessTheData/getInterfaceToAccessTheDataList [get]
export const getInterfaceToAccessTheDataList = (params) => {
  return service({
    url: '/interfaceToAccessTheData/getInterfaceToAccessTheDataList',
    method: 'get',
    params
  })
}

// @Tags InterfaceToAccessTheData
// @Summary 不需要鉴权的接口访问数据管理接口
// @accept application/json
// @Produce application/json
// @Param data query stReq.InterfaceToAccessTheDataSearch true "分页获取接口访问数据管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /interfaceToAccessTheData/getInterfaceToAccessTheDataPublic [get]
export const getInterfaceToAccessTheDataPublic = () => {
  return service({
    url: '/interfaceToAccessTheData/getInterfaceToAccessTheDataPublic',
    method: 'get',
  })
}
