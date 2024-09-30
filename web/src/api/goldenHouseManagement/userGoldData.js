import service from '@/utils/request'
// @Tags UserGoldData
// @Summary 创建用户打金数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserGoldData true "创建用户打金数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userGoldData/createUserGoldData [post]
export const createUserGoldData = (data) => {
  return service({
    url: '/userGoldData/createUserGoldData',
    method: 'post',
    data
  })
}

// @Tags UserGoldData
// @Summary 删除用户打金数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserGoldData true "删除用户打金数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userGoldData/deleteUserGoldData [delete]
export const deleteUserGoldData = (params) => {
  return service({
    url: '/userGoldData/deleteUserGoldData',
    method: 'delete',
    params
  })
}

// @Tags UserGoldData
// @Summary 批量删除用户打金数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户打金数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userGoldData/deleteUserGoldData [delete]
export const deleteUserGoldDataByIds = (params) => {
  return service({
    url: '/userGoldData/deleteUserGoldDataByIds',
    method: 'delete',
    params
  })
}

// @Tags UserGoldData
// @Summary 更新用户打金数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserGoldData true "更新用户打金数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userGoldData/updateUserGoldData [put]
export const updateUserGoldData = (data) => {
  return service({
    url: '/userGoldData/updateUserGoldData',
    method: 'put',
    data
  })
}

// @Tags UserGoldData
// @Summary 用id查询用户打金数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.UserGoldData true "用id查询用户打金数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userGoldData/findUserGoldData [get]
export const findUserGoldData = (params) => {
  return service({
    url: '/userGoldData/findUserGoldData',
    method: 'get',
    params
  })
}

// @Tags UserGoldData
// @Summary 分页获取用户打金数据列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户打金数据列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userGoldData/getUserGoldDataList [get]
export const getUserGoldDataList = (params) => {
  return service({
    url: '/userGoldData/getUserGoldDataList',
    method: 'get',
    params
  })
}

// @Tags UserGoldData
// @Summary 不需要鉴权的用户打金数据接口
// @accept application/json
// @Produce application/json
// @Param data query goldenHouseManagementReq.UserGoldDataSearch true "分页获取用户打金数据列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /userGoldData/getUserGoldDataPublic [get]
export const getUserGoldDataPublic = () => {
  return service({
    url: '/userGoldData/getUserGoldDataPublic',
    method: 'get',
  })
}
