package good

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/good"
    goodReq "github.com/flipped-aurora/gin-vue-admin/server/model/good/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type CertificateApi struct {}



// CreateCertificate 创建certificate表
// @Tags Certificate
// @Summary 创建certificate表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body good.Certificate true "创建certificate表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /certificate/createCertificate [post]
func (certificateApi *CertificateApi) CreateCertificate(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var certificate good.Certificate
	err := c.ShouldBindJSON(&certificate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = certificateService.CreateCertificate(ctx,&certificate)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteCertificate 删除certificate表
// @Tags Certificate
// @Summary 删除certificate表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body good.Certificate true "删除certificate表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /certificate/deleteCertificate [delete]
func (certificateApi *CertificateApi) DeleteCertificate(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := certificateService.DeleteCertificate(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCertificateByIds 批量删除certificate表
// @Tags Certificate
// @Summary 批量删除certificate表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /certificate/deleteCertificateByIds [delete]
func (certificateApi *CertificateApi) DeleteCertificateByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := certificateService.DeleteCertificateByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateCertificate 更新certificate表
// @Tags Certificate
// @Summary 更新certificate表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body good.Certificate true "更新certificate表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /certificate/updateCertificate [put]
func (certificateApi *CertificateApi) UpdateCertificate(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var certificate good.Certificate
	err := c.ShouldBindJSON(&certificate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = certificateService.UpdateCertificate(ctx,certificate)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindCertificate 用id查询certificate表
// @Tags Certificate
// @Summary 用id查询certificate表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询certificate表"
// @Success 200 {object} response.Response{data=good.Certificate,msg=string} "查询成功"
// @Router /certificate/findCertificate [get]
func (certificateApi *CertificateApi) FindCertificate(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	recertificate, err := certificateService.GetCertificate(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(recertificate, c)
}
// GetCertificateList 分页获取certificate表列表
// @Tags Certificate
// @Summary 分页获取certificate表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query goodReq.CertificateSearch true "分页获取certificate表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /certificate/getCertificateList [get]
func (certificateApi *CertificateApi) GetCertificateList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo goodReq.CertificateSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := certificateService.GetCertificateInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetCertificatePublic 不需要鉴权的certificate表接口
// @Tags Certificate
// @Summary 不需要鉴权的certificate表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /certificate/getCertificatePublic [get]
func (certificateApi *CertificateApi) GetCertificatePublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    certificateService.GetCertificatePublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的certificate表接口信息",
    }, "获取成功", c)
}
// Equcount 设备统计
// @Tags Certificate
// @Summary 设备统计
// @Accept application/json
// @Produce application/json
// @Param data query goodReq.CertificateSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /certificate/count [POST]
func (certificateApi *CertificateApi)Equcount(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()
    // 请添加自己的业务逻辑
    err := certificateService.Equcount(ctx)
    if err != nil {
        global.GVA_LOG.Error("失败!", zap.Error(err))
   		response.FailWithMessage("失败", c)
   		return
   	}
   	response.OkWithData("返回数据",c)
}


