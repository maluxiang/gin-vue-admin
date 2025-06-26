package good

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/good"
	goodReq "github.com/flipped-aurora/gin-vue-admin/server/model/good/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EquManageApi struct{}

// CreateEquManage 创建equManage表
// @Tags EquManage
// @Summary 创建equManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body good.EquManage true "创建equManage表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /equManage/createEquManage [post]

func (equManageApi *EquManageApi) CreateEquManage(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var equManage good.EquManage
	err := c.ShouldBindJSON(&equManage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = equManageService.CreateEquManage(ctx, &equManage)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteEquManage 删除equManage表
// @Tags EquManage
// @Summary 删除equManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body good.EquManage true "删除equManage表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /equManage/deleteEquManage [delete]
func (equManageApi *EquManageApi) DeleteEquManage(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	id := c.Query("id")
	err := equManageService.DeleteEquManage(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteEquManageByIds 批量删除equManage表
// @Tags EquManage
// @Summary 批量删除equManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /equManage/deleteEquManageByIds [delete]
func (equManageApi *EquManageApi) DeleteEquManageByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := equManageService.DeleteEquManageByIds(ctx, ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateEquManage 更新equManage表
// @Tags EquManage
// @Summary 更新equManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body good.EquManage true "更新equManage表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /equManage/updateEquManage [put]
func (equManageApi *EquManageApi) UpdateEquManage(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var equManage good.EquManage
	err := c.ShouldBindJSON(&equManage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = equManageService.UpdateEquManage(ctx, equManage)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindEquManage 用id查询equManage表
// @Tags EquManage
// @Summary 用id查询equManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询equManage表"
// @Success 200 {object} response.Response{data=good.EquManage,msg=string} "查询成功"
// @Router /equManage/findEquManage [get]
func (equManageApi *EquManageApi) FindEquManage(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	id := c.Query("id")
	reequManage, err := equManageService.GetEquManage(ctx, id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reequManage, c)
}

// GetEquManageList 分页获取equManage表列表
// @Tags EquManage
// @Summary 分页获取equManage表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query goodReq.EquManageSearch true "分页获取equManage表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /equManage/getEquManageList [get]
func (equManageApi *EquManageApi) GetEquManageList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo goodReq.EquManageSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := equManageService.GetEquManageInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetEquManagePublic 不需要鉴权的equManage表接口
// @Tags EquManage
// @Summary 不需要鉴权的equManage表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /equManage/getEquManagePublic [get]
func (equManageApi *EquManageApi) GetEquManagePublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	equManageService.GetEquManagePublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的equManage表接口信息",
	}, "获取成功", c)
}

// EquCount 设备统计
// @Tags EquManage
// @Summary 设备统计
// @Accept application/json
// @Produce application/json
// @Param data query goodReq.EquManageSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /equManage/counts [POST]
func (equManageApi *EquManageApi) EquCount(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()
	// 请添加自己的业务逻辑
	status, err := equManageService.EquCount(ctx)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.EquCount(status, c)
}
