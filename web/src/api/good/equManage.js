import service from '@/utils/request'
// @Tags EquManage
// @Summary 创建equManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EquManage true "创建equManage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /equManage/createEquManage [post]
export const createEquManage = (data) => {
  return service({
    url: '/equManage/createEquManage',
    method: 'post',
    data
  })
}

// @Tags EquManage
// @Summary 删除equManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EquManage true "删除equManage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /equManage/deleteEquManage [delete]
export const deleteEquManage = (params) => {
  return service({
    url: '/equManage/deleteEquManage',
    method: 'delete',
    params
  })
}

// @Tags EquManage
// @Summary 批量删除equManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除equManage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /equManage/deleteEquManage [delete]
export const deleteEquManageByIds = (params) => {
  return service({
    url: '/equManage/deleteEquManageByIds',
    method: 'delete',
    params
  })
}

// @Tags EquManage
// @Summary 更新equManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EquManage true "更新equManage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /equManage/updateEquManage [put]
export const updateEquManage = (data) => {
  return service({
    url: '/equManage/updateEquManage',
    method: 'put',
    data
  })
}

// @Tags EquManage
// @Summary 用id查询equManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.EquManage true "用id查询equManage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /equManage/findEquManage [get]
export const findEquManage = (params) => {
  return service({
    url: '/equManage/findEquManage',
    method: 'get',
    params
  })
}

// @Tags EquManage
// @Summary 分页获取equManage表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取equManage表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /equManage/getEquManageList [get]
export const getEquManageList = (params) => {
  return service({
    url: '/equManage/getEquManageList',
    method: 'get',
    params
  })
}

// @Tags EquManage
// @Summary 不需要鉴权的equManage表接口
// @Accept application/json
// @Produce application/json
// @Param data query goodReq.EquManageSearch true "分页获取equManage表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /equManage/getEquManagePublic [get]
export const getEquManagePublic = () => {
  return service({
    url: '/equManage/getEquManagePublic',
    method: 'get',
  })
}
// EquCount 设备统计
// @Tags EquManage
// @Summary 设备统计
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /equManage/counts [POST]
export const counts = (params) => {
  return service({
    url: '/equManage/counts',
    method: 'POST',
    params
  })
}
