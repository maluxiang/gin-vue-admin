import service from '@/utils/request'
// @Tags EquGroup
// @Summary 创建equGroup表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EquGroup true "创建equGroup表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /equGroup/createEquGroup [post]
export const createEquGroup = (data) => {
  return service({
    url: '/equGroup/createEquGroup',
    method: 'post',
    data
  })
}

// @Tags EquGroup
// @Summary 删除equGroup表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EquGroup true "删除equGroup表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /equGroup/deleteEquGroup [delete]
export const deleteEquGroup = (params) => {
  return service({
    url: '/equGroup/deleteEquGroup',
    method: 'delete',
    params
  })
}

// @Tags EquGroup
// @Summary 批量删除equGroup表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除equGroup表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /equGroup/deleteEquGroup [delete]
export const deleteEquGroupByIds = (params) => {
  return service({
    url: '/equGroup/deleteEquGroupByIds',
    method: 'delete',
    params
  })
}

// @Tags EquGroup
// @Summary 更新equGroup表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.EquGroup true "更新equGroup表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /equGroup/updateEquGroup [put]
export const updateEquGroup = (data) => {
  return service({
    url: '/equGroup/updateEquGroup',
    method: 'put',
    data
  })
}

// @Tags EquGroup
// @Summary 用id查询equGroup表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.EquGroup true "用id查询equGroup表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /equGroup/findEquGroup [get]
export const findEquGroup = (params) => {
  return service({
    url: '/equGroup/findEquGroup',
    method: 'get',
    params
  })
}

// @Tags EquGroup
// @Summary 分页获取equGroup表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取equGroup表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /equGroup/getEquGroupList [get]
export const getEquGroupList = (params) => {
  return service({
    url: '/equGroup/getEquGroupList',
    method: 'get',
    params
  })
}

// @Tags EquGroup
// @Summary 不需要鉴权的equGroup表接口
// @Accept application/json
// @Produce application/json
// @Param data query goodReq.EquGroupSearch true "分页获取equGroup表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /equGroup/getEquGroupPublic [get]
export const getEquGroupPublic = () => {
  return service({
    url: '/equGroup/getEquGroupPublic',
    method: 'get',
  })
}
