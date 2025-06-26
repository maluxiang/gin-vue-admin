package good

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EquGroupRouter struct {}

// InitEquGroupRouter 初始化 equGroup表 路由信息
func (s *EquGroupRouter) InitEquGroupRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	equGroupRouter := Router.Group("equGroup").Use(middleware.OperationRecord())
	equGroupRouterWithoutRecord := Router.Group("equGroup")
	equGroupRouterWithoutAuth := PublicRouter.Group("equGroup")
	{
		equGroupRouter.POST("createEquGroup", equGroupApi.CreateEquGroup)   // 新建equGroup表
		equGroupRouter.DELETE("deleteEquGroup", equGroupApi.DeleteEquGroup) // 删除equGroup表
		equGroupRouter.DELETE("deleteEquGroupByIds", equGroupApi.DeleteEquGroupByIds) // 批量删除equGroup表
		equGroupRouter.PUT("updateEquGroup", equGroupApi.UpdateEquGroup)    // 更新equGroup表
	}
	{
		equGroupRouterWithoutRecord.GET("findEquGroup", equGroupApi.FindEquGroup)        // 根据ID获取equGroup表
		equGroupRouterWithoutRecord.GET("getEquGroupList", equGroupApi.GetEquGroupList)  // 获取equGroup表列表
	}
	{
	    equGroupRouterWithoutAuth.GET("getEquGroupPublic", equGroupApi.GetEquGroupPublic)  // equGroup表开放接口
	}
}
