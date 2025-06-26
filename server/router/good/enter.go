package good

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	ProductRouter
	GoodTypeRouter
	GoodAccountRouter
	EquManageRouter
	EquGroupRouter
	AlarmManageRouter
	CertificateRouter
}

var (
	productApi     = api.ApiGroupApp.GoodApiGroup.ProductApi
	goodTypeApi    = api.ApiGroupApp.GoodApiGroup.GoodTypeApi
	goodAccountApi = api.ApiGroupApp.GoodApiGroup.GoodAccountApi
	equManageApi   = api.ApiGroupApp.GoodApiGroup.EquManageApi
	equGroupApi    = api.ApiGroupApp.GoodApiGroup.EquGroupApi
	alarmManageApi = api.ApiGroupApp.GoodApiGroup.AlarmManageApi
	certificateApi = api.ApiGroupApp.GoodApiGroup.CertificateApi
)
