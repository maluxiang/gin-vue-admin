package good

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/good"
    goodReq "github.com/flipped-aurora/gin-vue-admin/server/model/good/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type AlarmManageApi struct {}



// CreateAlarmManage 创建alarmManage表
// @Tags AlarmManage
// @Summary 创建alarmManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body good.AlarmManage true "创建alarmManage表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /alarmManage/createAlarmManage [post]
func (alarmManageApi *AlarmManageApi) CreateAlarmManage(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var alarmManage good.AlarmManage
	err := c.ShouldBindJSON(&alarmManage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = alarmManageService.CreateAlarmManage(ctx,&alarmManage)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteAlarmManage 删除alarmManage表
// @Tags AlarmManage
// @Summary 删除alarmManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body good.AlarmManage true "删除alarmManage表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /alarmManage/deleteAlarmManage [delete]
func (alarmManageApi *AlarmManageApi) DeleteAlarmManage(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := alarmManageService.DeleteAlarmManage(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAlarmManageByIds 批量删除alarmManage表
// @Tags AlarmManage
// @Summary 批量删除alarmManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /alarmManage/deleteAlarmManageByIds [delete]
func (alarmManageApi *AlarmManageApi) DeleteAlarmManageByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := alarmManageService.DeleteAlarmManageByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAlarmManage 更新alarmManage表
// @Tags AlarmManage
// @Summary 更新alarmManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body good.AlarmManage true "更新alarmManage表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /alarmManage/updateAlarmManage [put]
func (alarmManageApi *AlarmManageApi) UpdateAlarmManage(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var alarmManage good.AlarmManage
	err := c.ShouldBindJSON(&alarmManage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = alarmManageService.UpdateAlarmManage(ctx,alarmManage)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAlarmManage 用id查询alarmManage表
// @Tags AlarmManage
// @Summary 用id查询alarmManage表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询alarmManage表"
// @Success 200 {object} response.Response{data=good.AlarmManage,msg=string} "查询成功"
// @Router /alarmManage/findAlarmManage [get]
func (alarmManageApi *AlarmManageApi) FindAlarmManage(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	realarmManage, err := alarmManageService.GetAlarmManage(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(realarmManage, c)
}
// GetAlarmManageList 分页获取alarmManage表列表
// @Tags AlarmManage
// @Summary 分页获取alarmManage表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query goodReq.AlarmManageSearch true "分页获取alarmManage表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /alarmManage/getAlarmManageList [get]
func (alarmManageApi *AlarmManageApi) GetAlarmManageList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo goodReq.AlarmManageSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := alarmManageService.GetAlarmManageInfoList(ctx,pageInfo)
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

// GetAlarmManagePublic 不需要鉴权的alarmManage表接口
// @Tags AlarmManage
// @Summary 不需要鉴权的alarmManage表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /alarmManage/getAlarmManagePublic [get]
func (alarmManageApi *AlarmManageApi) GetAlarmManagePublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    alarmManageService.GetAlarmManagePublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的alarmManage表接口信息",
    }, "获取成功", c)
}
