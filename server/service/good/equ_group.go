
package good

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/good"
    goodReq "github.com/flipped-aurora/gin-vue-admin/server/model/good/request"
)

type EquGroupService struct {}
// CreateEquGroup 创建equGroup表记录
// Author [yourname](https://github.com/yourname)
func (equGroupService *EquGroupService) CreateEquGroup(ctx context.Context, equGroup *good.EquGroup) (err error) {
	err = global.GVA_DB.Create(equGroup).Error
	return err
}

// DeleteEquGroup 删除equGroup表记录
// Author [yourname](https://github.com/yourname)
func (equGroupService *EquGroupService)DeleteEquGroup(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&good.EquGroup{},"id = ?",id).Error
	return err
}

// DeleteEquGroupByIds 批量删除equGroup表记录
// Author [yourname](https://github.com/yourname)
func (equGroupService *EquGroupService)DeleteEquGroupByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]good.EquGroup{},"id in ?",ids).Error
	return err
}

// UpdateEquGroup 更新equGroup表记录
// Author [yourname](https://github.com/yourname)
func (equGroupService *EquGroupService)UpdateEquGroup(ctx context.Context, equGroup good.EquGroup) (err error) {
	err = global.GVA_DB.Model(&good.EquGroup{}).Where("id = ?",equGroup.Id).Updates(&equGroup).Error
	return err
}

// GetEquGroup 根据id获取equGroup表记录
// Author [yourname](https://github.com/yourname)
func (equGroupService *EquGroupService)GetEquGroup(ctx context.Context, id string) (equGroup good.EquGroup, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&equGroup).Error
	return
}
// GetEquGroupInfoList 分页获取equGroup表记录
// Author [yourname](https://github.com/yourname)
func (equGroupService *EquGroupService)GetEquGroupInfoList(ctx context.Context, info goodReq.EquGroupSearch) (list []good.EquGroup, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&good.EquGroup{})
    var equGroups []good.EquGroup
    // 如果有条件搜索 下方会自动创建搜索语句
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&equGroups).Error
	return  equGroups, total, err
}
func (equGroupService *EquGroupService)GetEquGroupPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
