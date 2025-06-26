import service from '@/utils/request'
// @Tags Certificate
// @Summary 创建certificate表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Certificate true "创建certificate表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /certificate/createCertificate [post]
export const createCertificate = (data) => {
  return service({
    url: '/certificate/createCertificate',
    method: 'post',
    data
  })
}

// @Tags Certificate
// @Summary 删除certificate表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Certificate true "删除certificate表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /certificate/deleteCertificate [delete]
export const deleteCertificate = (params) => {
  return service({
    url: '/certificate/deleteCertificate',
    method: 'delete',
    params
  })
}

// @Tags Certificate
// @Summary 批量删除certificate表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除certificate表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /certificate/deleteCertificate [delete]
export const deleteCertificateByIds = (params) => {
  return service({
    url: '/certificate/deleteCertificateByIds',
    method: 'delete',
    params
  })
}

// @Tags Certificate
// @Summary 更新certificate表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Certificate true "更新certificate表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /certificate/updateCertificate [put]
export const updateCertificate = (data) => {
  return service({
    url: '/certificate/updateCertificate',
    method: 'put',
    data
  })
}

// @Tags Certificate
// @Summary 用id查询certificate表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Certificate true "用id查询certificate表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /certificate/findCertificate [get]
export const findCertificate = (params) => {
  return service({
    url: '/certificate/findCertificate',
    method: 'get',
    params
  })
}

// @Tags Certificate
// @Summary 分页获取certificate表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取certificate表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /certificate/getCertificateList [get]
export const getCertificateList = (params) => {
  return service({
    url: '/certificate/getCertificateList',
    method: 'get',
    params
  })
}

// @Tags Certificate
// @Summary 不需要鉴权的certificate表接口
// @Accept application/json
// @Produce application/json
// @Param data query goodReq.CertificateSearch true "分页获取certificate表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /certificate/getCertificatePublic [get]
export const getCertificatePublic = () => {
  return service({
    url: '/certificate/getCertificatePublic',
    method: 'get',
  })
}
// Equcount 设备统计
// @Tags Certificate
// @Summary 设备统计
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /certificate/count [POST]
export const count = () => {
  return service({
    url: '/certificate/count',
    method: 'POST'
  })
}
