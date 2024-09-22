import service from '@/utils/request'
// @Tags GoldenRoomForm
// @Summary 创建金房表单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoldenRoomForm true "创建金房表单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /goldenRoomForm/createGoldenRoomForm [post]
export const createGoldenRoomForm = (data) => {
  return service({
    url: '/goldenRoomForm/createGoldenRoomForm',
    method: 'post',
    data
  })
}

// @Tags GoldenRoomForm
// @Summary 删除金房表单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoldenRoomForm true "删除金房表单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goldenRoomForm/deleteGoldenRoomForm [delete]
export const deleteGoldenRoomForm = (params) => {
  return service({
    url: '/goldenRoomForm/deleteGoldenRoomForm',
    method: 'delete',
    params
  })
}

// @Tags GoldenRoomForm
// @Summary 批量删除金房表单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除金房表单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goldenRoomForm/deleteGoldenRoomForm [delete]
export const deleteGoldenRoomFormByIds = (params) => {
  return service({
    url: '/goldenRoomForm/deleteGoldenRoomFormByIds',
    method: 'delete',
    params
  })
}

// @Tags GoldenRoomForm
// @Summary 更新金房表单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoldenRoomForm true "更新金房表单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goldenRoomForm/updateGoldenRoomForm [put]
export const updateGoldenRoomForm = (data) => {
  return service({
    url: '/goldenRoomForm/updateGoldenRoomForm',
    method: 'put',
    data
  })
}

// @Tags GoldenRoomForm
// @Summary 用id查询金房表单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.GoldenRoomForm true "用id查询金房表单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /goldenRoomForm/findGoldenRoomForm [get]
export const findGoldenRoomForm = (params) => {
  return service({
    url: '/goldenRoomForm/findGoldenRoomForm',
    method: 'get',
    params
  })
}

// @Tags GoldenRoomForm
// @Summary 分页获取金房表单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取金房表单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goldenRoomForm/getGoldenRoomFormList [get]
export const getGoldenRoomFormList = (params) => {
  return service({
    url: '/goldenRoomForm/getGoldenRoomFormList',
    method: 'get',
    params
  })
}

// @Tags GoldenRoomForm
// @Summary 不需要鉴权的金房表单接口
// @accept application/json
// @Produce application/json
// @Param data query goldenHouseManagementReq.GoldenRoomFormSearch true "分页获取金房表单列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /goldenRoomForm/getGoldenRoomFormPublic [get]
export const getGoldenRoomFormPublic = () => {
  return service({
    url: '/goldenRoomForm/getGoldenRoomFormPublic',
    method: 'get',
  })
}
// GetAListOfGoldenHouses 获取金房列表
// @Tags GoldenRoomForm
// @Summary 获取金房列表
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /goldenRoomForm/getAListOfGoldenHouses [GET]
export const getAListOfGoldenHouses = () => {
  return service({
    url: '/goldenRoomForm/getAListOfGoldenHouses',
    method: 'GET'
  })
}
// AddTheGoldenRoomData 添加金房数据
// @Tags GoldenRoomForm
// @Summary 添加金房数据
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /goldenRoomForm/addTheGoldenRoomData [GET]
export const addTheGoldenRoomData = () => {
  return service({
    url: '/goldenRoomForm/addTheGoldenRoomData',
    method: 'GET'
  })
}
// DeleteTheGoldenRoomData 删除金房数据
// @Tags GoldenRoomForm
// @Summary 删除金房数据
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /goldenRoomForm/deleteTheGoldenRoomData [GET]
export const deleteTheGoldenRoomData = () => {
  return service({
    url: '/goldenRoomForm/deleteTheGoldenRoomData',
    method: 'GET'
  })
}
// UpdateTheGoldenRoomData 更新金房数据
// @Tags GoldenRoomForm
// @Summary 更新金房数据
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /goldenRoomForm/updateTheGoldenRoomData [GET]
export const updateTheGoldenRoomData = () => {
  return service({
    url: '/goldenRoomForm/updateTheGoldenRoomData',
    method: 'GET'
  })
}
