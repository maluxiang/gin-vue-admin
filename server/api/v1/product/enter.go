package product

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ EquAddressApi }

var (
	equAddressService = service.ServiceGroupApp.ProductServiceGroup.EquAddressService
)
