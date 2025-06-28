package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/good"
	"github.com/flipped-aurora/gin-vue-admin/server/model/product"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(good.Product{}, good.GoodType{}, good.GoodAccount{}, good.EquManage{}, good.EquGroup{}, good.AlarmManage{}, good.Certificate{}, product.EquAddress{})
	if err != nil {
		return err
	}
	return nil
}
