package good

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/good"
    goodReq "github.com/flipped-aurora/gin-vue-admin/server/model/good/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type EquGroupApi struct {}



// CreateEquGroup 创建equGroup表
// @Tags EquGroup
// @Summary 创建equGroup表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body good.EquGroup true "创建equGroup表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /equGroup/createEquGroup [post]
func (equGroupApi *EquGroupApi) CreateEquGroup(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var equGroup good.EquGroup
	err := c.ShouldBindJSON(&equGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = equGroupService.CreateEquGroup(ctx,&equGroup)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteEquGroup 删除equGroup表
// @Tags EquGroup
// @Summary 删除equGroup表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body good.EquGroup true "删除equGroup表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /equGroup/deleteEquGroup [delete]
func (equGroupApi *EquGroupApi) DeleteEquGroup(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := equGroupService.DeleteEquGroup(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteEquGroupByIds 批量删除equGroup表
// @Tags EquGroup
// @Summary 批量删除equGroup表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /equGroup/deleteEquGroupByIds [delete]
func (equGroupApi *EquGroupApi) DeleteEquGroupByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := equGroupService.DeleteEquGroupByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateEquGroup 更新equGroup表
// @Tags EquGroup
// @Summary 更新equGroup表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body good.EquGroup true "更新equGroup表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /equGroup/updateEquGroup [put]
func (equGroupApi *EquGroupApi) UpdateEquGroup(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var equGroup good.EquGroup
	err := c.ShouldBindJSON(&equGroup)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = equGroupService.UpdateEquGroup(ctx,equGroup)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindEquGroup 用id查询equGroup表
// @Tags EquGroup
// @Summary 用id查询equGroup表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询equGroup表"
// @Success 200 {object} response.Response{data=good.EquGroup,msg=string} "查询成功"
// @Router /equGroup/findEquGroup [get]
func (equGroupApi *EquGroupApi) FindEquGroup(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	reequGroup, err := equGroupService.GetEquGroup(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reequGroup, c)
}
// GetEquGroupList 分页获取equGroup表列表
// @Tags EquGroup
// @Summary 分页获取equGroup表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query goodReq.EquGroupSearch true "分页获取equGroup表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /equGroup/getEquGroupList [get]
func (equGroupApi *EquGroupApi) GetEquGroupList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo goodReq.EquGroupSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := equGroupService.GetEquGroupInfoList(ctx,pageInfo)
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

// GetEquGroupPublic 不需要鉴权的equGroup表接口
// @Tags EquGroup
// @Summary 不需要鉴权的equGroup表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /equGroup/getEquGroupPublic [get]
func (equGroupApi *EquGroupApi) GetEquGroupPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    equGroupService.GetEquGroupPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的equGroup表接口信息",
    }, "获取成功", c)
}
