import service from '@/utils/request'
// @Tags GoodAccount
// @Summary 创建goodAccount表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GoodAccount true "创建goodAccount表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /goodAccount/createGoodAccount [post]
export const createGoodAccount = (data) => {
  return service({
    url: '/goodAccount/createGoodAccount',
    method: 'post',
    data
  })
}

// @Tags GoodAccount
// @Summary 删除goodAccount表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GoodAccount true "删除goodAccount表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goodAccount/deleteGoodAccount [delete]
export const deleteGoodAccount = (params) => {
  return service({
    url: '/goodAccount/deleteGoodAccount',
    method: 'delete',
    params
  })
}

// @Tags GoodAccount
// @Summary 批量删除goodAccount表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除goodAccount表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /goodAccount/deleteGoodAccount [delete]
export const deleteGoodAccountByIds = (params) => {
  return service({
    url: '/goodAccount/deleteGoodAccountByIds',
    method: 'delete',
    params
  })
}

// @Tags GoodAccount
// @Summary 更新goodAccount表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.GoodAccount true "更新goodAccount表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /goodAccount/updateGoodAccount [put]
export const updateGoodAccount = (data) => {
  return service({
    url: '/goodAccount/updateGoodAccount',
    method: 'put',
    data
  })
}

// @Tags GoodAccount
// @Summary 用id查询goodAccount表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.GoodAccount true "用id查询goodAccount表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /goodAccount/findGoodAccount [get]
export const findGoodAccount = (params) => {
  return service({
    url: '/goodAccount/findGoodAccount',
    method: 'get',
    params
  })
}

// @Tags GoodAccount
// @Summary 分页获取goodAccount表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取goodAccount表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /goodAccount/getGoodAccountList [get]
export const getGoodAccountList = (params) => {
  return service({
    url: '/goodAccount/getGoodAccountList',
    method: 'get',
    params
  })
}

// @Tags GoodAccount
// @Summary 不需要鉴权的goodAccount表接口
// @Accept application/json
// @Produce application/json
// @Param data query goodReq.GoodAccountSearch true "分页获取goodAccount表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /goodAccount/getGoodAccountPublic [get]
export const getGoodAccountPublic = () => {
  return service({
    url: '/goodAccount/getGoodAccountPublic',
    method: 'get',
  })
}
