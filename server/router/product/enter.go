package product

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ EquAddressRouter }

var (
	equAddressApi = api.ApiGroupApp.ProductApiGroup.EquAddressApi
)
