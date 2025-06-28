package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}
func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	holder(publicGroup, privateGroup)
	{
		goodRouter := router.RouterGroupApp.Good
		goodRouter.InitProductRouter(privateGroup, publicGroup)
		goodRouter.InitGoodTypeRouter(privateGroup, publicGroup)
		goodRouter.InitGoodAccountRouter(privateGroup, publicGroup)
		goodRouter.InitEquManageRouter(privateGroup, publicGroup)
		goodRouter.InitEquGroupRouter(privateGroup, publicGroup)
		goodRouter.InitAlarmManageRouter(privateGroup, publicGroup)
		goodRouter.InitCertificateRouter(privateGroup, publicGroup)
	}
	{
		productRouter := router.RouterGroupApp.Product
		productRouter.InitEquAddressRouter(privateGroup, publicGroup)
	}
}

// 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
