
package good

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/good"
    goodReq "github.com/flipped-aurora/gin-vue-admin/server/model/good/request"
)

type GoodTypeService struct {}
// CreateGoodType 创建goodType表记录
// Author [yourname](https://github.com/yourname)
func (goodTypeService *GoodTypeService) CreateGoodType(ctx context.Context, goodType *good.GoodType) (err error) {
	err = global.GVA_DB.Create(goodType).Error
	return err
}

// DeleteGoodType 删除goodType表记录
// Author [yourname](https://github.com/yourname)
func (goodTypeService *GoodTypeService)DeleteGoodType(ctx context.Context, typeId string) (err error) {
	err = global.GVA_DB.Delete(&good.GoodType{},"type_id = ?",typeId).Error
	return err
}

// DeleteGoodTypeByIds 批量删除goodType表记录
// Author [yourname](https://github.com/yourname)
func (goodTypeService *GoodTypeService)DeleteGoodTypeByIds(ctx context.Context, typeIds []string) (err error) {
	err = global.GVA_DB.Delete(&[]good.GoodType{},"type_id in ?",typeIds).Error
	return err
}

// UpdateGoodType 更新goodType表记录
// Author [yourname](https://github.com/yourname)
func (goodTypeService *GoodTypeService)UpdateGoodType(ctx context.Context, goodType good.GoodType) (err error) {
	err = global.GVA_DB.Model(&good.GoodType{}).Where("type_id = ?",goodType.TypeId).Updates(&goodType).Error
	return err
}

// GetGoodType 根据typeId获取goodType表记录
// Author [yourname](https://github.com/yourname)
func (goodTypeService *GoodTypeService)GetGoodType(ctx context.Context, typeId string) (goodType good.GoodType, err error) {
	err = global.GVA_DB.Where("type_id = ?", typeId).First(&goodType).Error
	return
}
// GetGoodTypeInfoList 分页获取goodType表记录
// Author [yourname](https://github.com/yourname)
func (goodTypeService *GoodTypeService)GetGoodTypeInfoList(ctx context.Context, info goodReq.GoodTypeSearch) (list []good.GoodType, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&good.GoodType{})
    var goodTypes []good.GoodType
    // 如果有条件搜索 下方会自动创建搜索语句
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&goodTypes).Error
	return  goodTypes, total, err
}
func (goodTypeService *GoodTypeService)GetGoodTypePublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
