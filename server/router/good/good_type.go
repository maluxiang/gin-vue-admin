package good

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GoodTypeRouter struct {}

// InitGoodTypeRouter 初始化 goodType表 路由信息
func (s *GoodTypeRouter) InitGoodTypeRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	goodTypeRouter := Router.Group("goodType").Use(middleware.OperationRecord())
	goodTypeRouterWithoutRecord := Router.Group("goodType")
	goodTypeRouterWithoutAuth := PublicRouter.Group("goodType")
	{
		goodTypeRouter.POST("createGoodType", goodTypeApi.CreateGoodType)   // 新建goodType表
		goodTypeRouter.DELETE("deleteGoodType", goodTypeApi.DeleteGoodType) // 删除goodType表
		goodTypeRouter.DELETE("deleteGoodTypeByIds", goodTypeApi.DeleteGoodTypeByIds) // 批量删除goodType表
		goodTypeRouter.PUT("updateGoodType", goodTypeApi.UpdateGoodType)    // 更新goodType表
	}
	{
		goodTypeRouterWithoutRecord.GET("findGoodType", goodTypeApi.FindGoodType)        // 根据ID获取goodType表
		goodTypeRouterWithoutRecord.GET("getGoodTypeList", goodTypeApi.GetGoodTypeList)  // 获取goodType表列表
	}
	{
	    goodTypeRouterWithoutAuth.GET("getGoodTypePublic", goodTypeApi.GetGoodTypePublic)  // goodType表开放接口
	}
}
