import service from '@/utils/request'
// @Tags AppUser
// @Summary 创建appUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppUser true "创建appUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /appUser/createAppUser [post]
export const createAppUser = (data) => {
  return service({
    url: '/appUser/createAppUser',
    method: 'post',
    data
  })
}

// @Tags AppUser
// @Summary 删除appUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppUser true "删除appUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appUser/deleteAppUser [delete]
export const deleteAppUser = (params) => {
  return service({
    url: '/appUser/deleteAppUser',
    method: 'delete',
    params
  })
}

// @Tags AppUser
// @Summary 批量删除appUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除appUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appUser/deleteAppUser [delete]
export const deleteAppUserByIds = (params) => {
  return service({
    url: '/appUser/deleteAppUserByIds',
    method: 'delete',
    params
  })
}

// @Tags AppUser
// @Summary 更新appUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppUser true "更新appUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appUser/updateAppUser [put]
export const updateAppUser = (data) => {
  return service({
    url: '/appUser/updateAppUser',
    method: 'put',
    data
  })
}

// @Tags AppUser
// @Summary 用id查询appUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppUser true "用id查询appUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appUser/findAppUser [get]
export const findAppUser = (params) => {
  return service({
    url: '/appUser/findAppUser',
    method: 'get',
    params
  })
}

// @Tags AppUser
// @Summary 分页获取appUser表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取appUser表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appUser/getAppUserList [get]
export const getAppUserList = (params) => {
  return service({
    url: '/appUser/getAppUserList',
    method: 'get',
    params
  })
}

// @Tags AppUser
// @Summary 不需要鉴权的appUser表接口
// @accept application/json
// @Produce application/json
// @Param data query stReq.AppUserSearch true "分页获取appUser表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /appUser/getAppUserPublic [get]
export const getAppUserPublic = () => {
  return service({
    url: '/appUser/getAppUserPublic',
    method: 'get',
  })
}
// Login 方法介绍
// @Tags AppUser
// @Summary 方法介绍
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /appUser/Login [POST]
export const Login = () => {
  return service({
    url: '/appUser/Login',
    method: 'POST'
  })
}
// GetUserinfo 用户数据表单
// @Tags AppUser
// @Summary 用户数据表单
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /appUser/getUserinfo [GET]
export const getUserinfo = () => {
  return service({
    url: '/appUser/getUserinfo',
    method: 'GET'
  })
}
// UpdateTheImage 更新头像
// @Tags AppUser
// @Summary 更新头像
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /appUser/updateTheImage [GET]
export const updateTheImage = () => {
  return service({
    url: '/appUser/updateTheImage',
    method: 'GET'
  })
}
// Upnickname 更新昵称
// @Tags AppUser
// @Summary 更新昵称
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /appUser/opnickname [GET]
export const opnickname = () => {
  return service({
    url: '/appUser/opnickname',
    method: 'GET'
  })
}
// Upnickname 方法介绍
// @Tags AppUser
// @Summary 方法介绍
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /appUser/upnickname [GET]
export const upnickname = () => {
  return service({
    url: '/appUser/upnickname',
    method: 'GET'
  })
}
