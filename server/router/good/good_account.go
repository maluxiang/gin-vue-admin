package good

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GoodAccountRouter struct {}

// InitGoodAccountRouter 初始化 goodAccount表 路由信息
func (s *GoodAccountRouter) InitGoodAccountRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	goodAccountRouter := Router.Group("goodAccount").Use(middleware.OperationRecord())
	goodAccountRouterWithoutRecord := Router.Group("goodAccount")
	goodAccountRouterWithoutAuth := PublicRouter.Group("goodAccount")
	{
		goodAccountRouter.POST("createGoodAccount", goodAccountApi.CreateGoodAccount)   // 新建goodAccount表
		goodAccountRouter.DELETE("deleteGoodAccount", goodAccountApi.DeleteGoodAccount) // 删除goodAccount表
		goodAccountRouter.DELETE("deleteGoodAccountByIds", goodAccountApi.DeleteGoodAccountByIds) // 批量删除goodAccount表
		goodAccountRouter.PUT("updateGoodAccount", goodAccountApi.UpdateGoodAccount)    // 更新goodAccount表
	}
	{
		goodAccountRouterWithoutRecord.GET("findGoodAccount", goodAccountApi.FindGoodAccount)        // 根据ID获取goodAccount表
		goodAccountRouterWithoutRecord.GET("getGoodAccountList", goodAccountApi.GetGoodAccountList)  // 获取goodAccount表列表
	}
	{
	    goodAccountRouterWithoutAuth.GET("getGoodAccountPublic", goodAccountApi.GetGoodAccountPublic)  // goodAccount表开放接口
	}
}
