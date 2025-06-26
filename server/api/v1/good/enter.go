package good

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	ProductApi
	GoodTypeApi
	GoodAccountApi
	EquManageApi
	EquGroupApi
	AlarmManageApi
	CertificateApi
}

var (
	productService     = service.ServiceGroupApp.GoodServiceGroup.ProductService
	goodTypeService    = service.ServiceGroupApp.GoodServiceGroup.GoodTypeService
	goodAccountService = service.ServiceGroupApp.GoodServiceGroup.GoodAccountService
	equManageService   = service.ServiceGroupApp.GoodServiceGroup.EquManageService
	equGroupService    = service.ServiceGroupApp.GoodServiceGroup.EquGroupService
	alarmManageService = service.ServiceGroupApp.GoodServiceGroup.AlarmManageService
	certificateService = service.ServiceGroupApp.GoodServiceGroup.CertificateService
)
