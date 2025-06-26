import service from '@/utils/request'
// @Tags GoodType
// @Summary 创建goodType表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GoodType true "创建goodType表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /goodType/createGoodType [post]
export const createGoodType = (data) => {
  return service({
    url: '/goodType/createGoodType',
    method: 'post',
    data
  })
}

// @Tags GoodType
// @Summary 删除goodType表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GoodType true "删除goodType表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goodType/deleteGoodType [delete]
export const deleteGoodType = (params) => {
  return service({
    url: '/goodType/deleteGoodType',
    method: 'delete',
    params
  })
}

// @Tags GoodType
// @Summary 批量删除goodType表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除goodType表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goodType/deleteGoodType [delete]
export const deleteGoodTypeByIds = (params) => {
  return service({
    url: '/goodType/deleteGoodTypeByIds',
    method: 'delete',
    params
  })
}

// @Tags GoodType
// @Summary 更新goodType表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GoodType true "更新goodType表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goodType/updateGoodType [put]
export const updateGoodType = (data) => {
  return service({
    url: '/goodType/updateGoodType',
    method: 'put',
    data
  })
}

// @Tags GoodType
// @Summary 用id查询goodType表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.GoodType true "用id查询goodType表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /goodType/findGoodType [get]
export const findGoodType = (params) => {
  return service({
    url: '/goodType/findGoodType',
    method: 'get',
    params
  })
}

// @Tags GoodType
// @Summary 分页获取goodType表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取goodType表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodType/getGoodTypeList [get]
export const getGoodTypeList = (params) => {
  return service({
    url: '/goodType/getGoodTypeList',
    method: 'get',
    params
  })
}

// @Tags GoodType
// @Summary 不需要鉴权的goodType表接口
// @Accept application/json
// @Produce application/json
// @Param data query goodReq.GoodTypeSearch true "分页获取goodType表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /goodType/getGoodTypePublic [get]
export const getGoodTypePublic = () => {
  return service({
    url: '/goodType/getGoodTypePublic',
    method: 'get',
  })
}
