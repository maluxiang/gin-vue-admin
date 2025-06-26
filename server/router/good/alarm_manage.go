package good

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AlarmManageRouter struct{}

// InitAlarmManageRouter 初始化 alarmManage表 路由信息
func (s *AlarmManageRouter) InitAlarmManageRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	alarmManageRouter := Router.Group("alarmManage").Use(middleware.OperationRecord())
	alarmManageRouterWithoutRecord := Router.Group("alarmManage")
	alarmManageRouterWithoutAuth := PublicRouter.Group("alarmManage")
	{
		alarmManageRouter.POST("createAlarmManage", alarmManageApi.CreateAlarmManage)             // 新建alarmManage表
		alarmManageRouter.DELETE("deleteAlarmManage", alarmManageApi.DeleteAlarmManage)           // 删除alarmManage表
		alarmManageRouter.DELETE("deleteAlarmManageByIds", alarmManageApi.DeleteAlarmManageByIds) // 批量删除alarmManage表
		alarmManageRouter.PUT("updateAlarmManage", alarmManageApi.UpdateAlarmManage)              // 更新alarmManage表
	}
	{
		alarmManageRouterWithoutRecord.GET("findAlarmManage", alarmManageApi.FindAlarmManage)       // 根据ID获取alarmManage表
		alarmManageRouterWithoutRecord.GET("getAlarmManageList", alarmManageApi.GetAlarmManageList) // 获取alarmManage表列表
	}
	{
		alarmManageRouterWithoutAuth.GET("getAlarmManagePublic", alarmManageApi.GetAlarmManagePublic) // alarmManage表开放接口
	}
}
