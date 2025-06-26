import service from '@/utils/request'
// @Tags AlarmManage
// @Summary 创建alarmManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AlarmManage true "创建alarmManage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /alarmManage/createAlarmManage [post]
export const createAlarmManage = (data) => {
  return service({
    url: '/alarmManage/createAlarmManage',
    method: 'post',
    data
  })
}

// @Tags AlarmManage
// @Summary 删除alarmManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AlarmManage true "删除alarmManage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /alarmManage/deleteAlarmManage [delete]
export const deleteAlarmManage = (params) => {
  return service({
    url: '/alarmManage/deleteAlarmManage',
    method: 'delete',
    params
  })
}

// @Tags AlarmManage
// @Summary 批量删除alarmManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除alarmManage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /alarmManage/deleteAlarmManage [delete]
export const deleteAlarmManageByIds = (params) => {
  return service({
    url: '/alarmManage/deleteAlarmManageByIds',
    method: 'delete',
    params
  })
}

// @Tags AlarmManage
// @Summary 更新alarmManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.AlarmManage true "更新alarmManage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /alarmManage/updateAlarmManage [put]
export const updateAlarmManage = (data) => {
  return service({
    url: '/alarmManage/updateAlarmManage',
    method: 'put',
    data
  })
}

// @Tags AlarmManage
// @Summary 用id查询alarmManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.AlarmManage true "用id查询alarmManage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /alarmManage/findAlarmManage [get]
export const findAlarmManage = (params) => {
  return service({
    url: '/alarmManage/findAlarmManage',
    method: 'get',
    params
  })
}

// @Tags AlarmManage
// @Summary 分页获取alarmManage表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取alarmManage表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /alarmManage/getAlarmManageList [get]
export const getAlarmManageList = (params) => {
  return service({
    url: '/alarmManage/getAlarmManageList',
    method: 'get',
    params
  })
}

// @Tags AlarmManage
// @Summary 不需要鉴权的alarmManage表接口
// @Accept application/json
// @Produce application/json
// @Param data query goodReq.AlarmManageSearch true "分页获取alarmManage表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /alarmManage/getAlarmManagePublic [get]
export const getAlarmManagePublic = () => {
  return service({
    url: '/alarmManage/getAlarmManagePublic',
    method: 'get',
  })
}
