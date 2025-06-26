package good

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/good"
    goodReq "github.com/flipped-aurora/gin-vue-admin/server/model/good/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type GoodTypeApi struct {}



// CreateGoodType 创建goodType表
// @Tags GoodType
// @Summary 创建goodType表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body good.GoodType true "创建goodType表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /goodType/createGoodType [post]
func (goodTypeApi *GoodTypeApi) CreateGoodType(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var goodType good.GoodType
	err := c.ShouldBindJSON(&goodType)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = goodTypeService.CreateGoodType(ctx,&goodType)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteGoodType 删除goodType表
// @Tags GoodType
// @Summary 删除goodType表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body good.GoodType true "删除goodType表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /goodType/deleteGoodType [delete]
func (goodTypeApi *GoodTypeApi) DeleteGoodType(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	typeId := c.Query("typeId")
	err := goodTypeService.DeleteGoodType(ctx,typeId)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteGoodTypeByIds 批量删除goodType表
// @Tags GoodType
// @Summary 批量删除goodType表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /goodType/deleteGoodTypeByIds [delete]
func (goodTypeApi *GoodTypeApi) DeleteGoodTypeByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	typeIds := c.QueryArray("typeIds[]")
	err := goodTypeService.DeleteGoodTypeByIds(ctx,typeIds)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateGoodType 更新goodType表
// @Tags GoodType
// @Summary 更新goodType表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body good.GoodType true "更新goodType表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /goodType/updateGoodType [put]
func (goodTypeApi *GoodTypeApi) UpdateGoodType(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var goodType good.GoodType
	err := c.ShouldBindJSON(&goodType)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = goodTypeService.UpdateGoodType(ctx,goodType)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindGoodType 用id查询goodType表
// @Tags GoodType
// @Summary 用id查询goodType表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param typeId query int true "用id查询goodType表"
// @Success 200 {object} response.Response{data=good.GoodType,msg=string} "查询成功"
// @Router /goodType/findGoodType [get]
func (goodTypeApi *GoodTypeApi) FindGoodType(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	typeId := c.Query("typeId")
	regoodType, err := goodTypeService.GetGoodType(ctx,typeId)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(regoodType, c)
}
// GetGoodTypeList 分页获取goodType表列表
// @Tags GoodType
// @Summary 分页获取goodType表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query goodReq.GoodTypeSearch true "分页获取goodType表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /goodType/getGoodTypeList [get]
func (goodTypeApi *GoodTypeApi) GetGoodTypeList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo goodReq.GoodTypeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := goodTypeService.GetGoodTypeInfoList(ctx,pageInfo)
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

// GetGoodTypePublic 不需要鉴权的goodType表接口
// @Tags GoodType
// @Summary 不需要鉴权的goodType表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /goodType/getGoodTypePublic [get]
func (goodTypeApi *GoodTypeApi) GetGoodTypePublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    goodTypeService.GetGoodTypePublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的goodType表接口信息",
    }, "获取成功", c)
}
