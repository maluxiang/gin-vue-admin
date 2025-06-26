package good

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EquManageRouter struct{}

func (s *EquManageRouter) InitEquManageRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	equManageRouter := Router.Group("equManage").Use(middleware.OperationRecord())
	equManageRouterWithoutRecord := Router.Group("equManage")
	equManageRouterWithoutAuth := PublicRouter.Group("equManage")
	{
		equManageRouter.POST("createEquManage", equManageApi.CreateEquManage)
		equManageRouter.DELETE("deleteEquManage", equManageApi.DeleteEquManage)
		equManageRouter.DELETE("deleteEquManageByIds", equManageApi.DeleteEquManageByIds)
		equManageRouter.PUT("updateEquManage", equManageApi.UpdateEquManage)
	}
	{
		equManageRouterWithoutRecord.GET("findEquManage", equManageApi.FindEquManage)
		equManageRouterWithoutRecord.GET("getEquManageList", equManageApi.GetEquManageList)
	}
	{
		equManageRouterWithoutAuth.GET("getEquManagePublic", equManageApi.GetEquManagePublic)
		equManageRouterWithoutAuth.POST("counts", equManageApi.EquCount)
	}
}
