import service from '@/utils/request'
// @Tags AnnouncementManagement
// @Summary 创建公告管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AnnouncementManagement true "创建公告管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /announcementManagement/createAnnouncementManagement [post]
export const createAnnouncementManagement = (data) => {
  return service({
    url: '/announcementManagement/createAnnouncementManagement',
    method: 'post',
    data
  })
}

// @Tags AnnouncementManagement
// @Summary 删除公告管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AnnouncementManagement true "删除公告管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /announcementManagement/deleteAnnouncementManagement [delete]
export const deleteAnnouncementManagement = (params) => {
  return service({
    url: '/announcementManagement/deleteAnnouncementManagement',
    method: 'delete',
    params
  })
}

// @Tags AnnouncementManagement
// @Summary 批量删除公告管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除公告管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /announcementManagement/deleteAnnouncementManagement [delete]
export const deleteAnnouncementManagementByIds = (params) => {
  return service({
    url: '/announcementManagement/deleteAnnouncementManagementByIds',
    method: 'delete',
    params
  })
}

// @Tags AnnouncementManagement
// @Summary 更新公告管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AnnouncementManagement true "更新公告管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /announcementManagement/updateAnnouncementManagement [put]
export const updateAnnouncementManagement = (data) => {
  return service({
    url: '/announcementManagement/updateAnnouncementManagement',
    method: 'put',
    data
  })
}

// @Tags AnnouncementManagement
// @Summary 用id查询公告管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AnnouncementManagement true "用id查询公告管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /announcementManagement/findAnnouncementManagement [get]
export const findAnnouncementManagement = (params) => {
  return service({
    url: '/announcementManagement/findAnnouncementManagement',
    method: 'get',
    params
  })
}

// @Tags AnnouncementManagement
// @Summary 分页获取公告管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取公告管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /announcementManagement/getAnnouncementManagementList [get]
export const getAnnouncementManagementList = (params) => {
  return service({
    url: '/announcementManagement/getAnnouncementManagementList',
    method: 'get',
    params
  })
}

// @Tags AnnouncementManagement
// @Summary 不需要鉴权的公告管理接口
// @accept application/json
// @Produce application/json
// @Param data query stReq.AnnouncementManagementSearch true "分页获取公告管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /announcementManagement/getAnnouncementManagementPublic [get]
export const getAnnouncementManagementPublic = () => {
  return service({
    url: '/announcementManagement/getAnnouncementManagementPublic',
    method: 'get',
  })
}
// GetAListOfAnnouncements 方法介绍
// @Tags AnnouncementManagement
// @Summary 方法介绍
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /announcementManagement/getAListOfAnnouncements [GET]
export const getAListOfAnnouncements = () => {
  return service({
    url: '/announcementManagement/getAListOfAnnouncements',
    method: 'GET'
  })
}
