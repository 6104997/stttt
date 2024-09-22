import service from '@/utils/request'
// @Tags ArouselImage
// @Summary 创建arouselImage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ArouselImage true "创建arouselImage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /arouselImage/createArouselImage [post]
export const createArouselImage = (data) => {
  return service({
    url: '/arouselImage/createArouselImage',
    method: 'post',
    data
  })
}

// @Tags ArouselImage
// @Summary 删除arouselImage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ArouselImage true "删除arouselImage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /arouselImage/deleteArouselImage [delete]
export const deleteArouselImage = (params) => {
  return service({
    url: '/arouselImage/deleteArouselImage',
    method: 'delete',
    params
  })
}

// @Tags ArouselImage
// @Summary 批量删除arouselImage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除arouselImage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /arouselImage/deleteArouselImage [delete]
export const deleteArouselImageByIds = (params) => {
  return service({
    url: '/arouselImage/deleteArouselImageByIds',
    method: 'delete',
    params
  })
}

// @Tags ArouselImage
// @Summary 更新arouselImage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ArouselImage true "更新arouselImage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /arouselImage/updateArouselImage [put]
export const updateArouselImage = (data) => {
  return service({
    url: '/arouselImage/updateArouselImage',
    method: 'put',
    data
  })
}

// @Tags ArouselImage
// @Summary 用id查询arouselImage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ArouselImage true "用id查询arouselImage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /arouselImage/findArouselImage [get]
export const findArouselImage = (params) => {
  return service({
    url: '/arouselImage/findArouselImage',
    method: 'get',
    params
  })
}

// @Tags ArouselImage
// @Summary 分页获取arouselImage表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取arouselImage表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /arouselImage/getArouselImageList [get]
export const getArouselImageList = (params) => {
  return service({
    url: '/arouselImage/getArouselImageList',
    method: 'get',
    params
  })
}

// @Tags ArouselImage
// @Summary 不需要鉴权的arouselImage表接口
// @accept application/json
// @Produce application/json
// @Param data query stReq.ArouselImageSearch true "分页获取arouselImage表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /arouselImage/getArouselImagePublic [get]
export const getArouselImagePublic = () => {
  return service({
    url: '/arouselImage/getArouselImagePublic',
    method: 'get',
  })
}
// GetImage 方法介绍
// @Tags ArouselImage
// @Summary 方法介绍
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /arouselImage/GetImage [GET]
export const GetImage = () => {
  return service({
    url: '/arouselImage/GetImage',
    method: 'GET'
  })
}
