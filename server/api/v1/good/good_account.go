package good

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/good"
    goodReq "github.com/flipped-aurora/gin-vue-admin/server/model/good/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type GoodAccountApi struct {}



// CreateGoodAccount 创建goodAccount表
// @Tags GoodAccount
// @Summary 创建goodAccount表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body good.GoodAccount true "创建goodAccount表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /goodAccount/createGoodAccount [post]
func (goodAccountApi *GoodAccountApi) CreateGoodAccount(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var goodAccount good.GoodAccount
	err := c.ShouldBindJSON(&goodAccount)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = goodAccountService.CreateGoodAccount(ctx,&goodAccount)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteGoodAccount 删除goodAccount表
// @Tags GoodAccount
// @Summary 删除goodAccount表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body good.GoodAccount true "删除goodAccount表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /goodAccount/deleteGoodAccount [delete]
func (goodAccountApi *GoodAccountApi) DeleteGoodAccount(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := goodAccountService.DeleteGoodAccount(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteGoodAccountByIds 批量删除goodAccount表
// @Tags GoodAccount
// @Summary 批量删除goodAccount表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /goodAccount/deleteGoodAccountByIds [delete]
func (goodAccountApi *GoodAccountApi) DeleteGoodAccountByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := goodAccountService.DeleteGoodAccountByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateGoodAccount 更新goodAccount表
// @Tags GoodAccount
// @Summary 更新goodAccount表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body good.GoodAccount true "更新goodAccount表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /goodAccount/updateGoodAccount [put]
func (goodAccountApi *GoodAccountApi) UpdateGoodAccount(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var goodAccount good.GoodAccount
	err := c.ShouldBindJSON(&goodAccount)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = goodAccountService.UpdateGoodAccount(ctx,goodAccount)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindGoodAccount 用id查询goodAccount表
// @Tags GoodAccount
// @Summary 用id查询goodAccount表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询goodAccount表"
// @Success 200 {object} response.Response{data=good.GoodAccount,msg=string} "查询成功"
// @Router /goodAccount/findGoodAccount [get]
func (goodAccountApi *GoodAccountApi) FindGoodAccount(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	regoodAccount, err := goodAccountService.GetGoodAccount(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(regoodAccount, c)
}
// GetGoodAccountList 分页获取goodAccount表列表
// @Tags GoodAccount
// @Summary 分页获取goodAccount表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query goodReq.GoodAccountSearch true "分页获取goodAccount表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /goodAccount/getGoodAccountList [get]
func (goodAccountApi *GoodAccountApi) GetGoodAccountList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo goodReq.GoodAccountSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := goodAccountService.GetGoodAccountInfoList(ctx,pageInfo)
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

// GetGoodAccountPublic 不需要鉴权的goodAccount表接口
// @Tags GoodAccount
// @Summary 不需要鉴权的goodAccount表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /goodAccount/getGoodAccountPublic [get]
func (goodAccountApi *GoodAccountApi) GetGoodAccountPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    goodAccountService.GetGoodAccountPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的goodAccount表接口信息",
    }, "获取成功", c)
}
