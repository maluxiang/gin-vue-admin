
package good

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/good"
    goodReq "github.com/flipped-aurora/gin-vue-admin/server/model/good/request"
)

type GoodAccountService struct {}
// CreateGoodAccount 创建goodAccount表记录
// Author [yourname](https://github.com/yourname)
func (goodAccountService *GoodAccountService) CreateGoodAccount(ctx context.Context, goodAccount *good.GoodAccount) (err error) {
	err = global.GVA_DB.Create(goodAccount).Error
	return err
}

// DeleteGoodAccount 删除goodAccount表记录
// Author [yourname](https://github.com/yourname)
func (goodAccountService *GoodAccountService)DeleteGoodAccount(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&good.GoodAccount{},"id = ?",id).Error
	return err
}

// DeleteGoodAccountByIds 批量删除goodAccount表记录
// Author [yourname](https://github.com/yourname)
func (goodAccountService *GoodAccountService)DeleteGoodAccountByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]good.GoodAccount{},"id in ?",ids).Error
	return err
}

// UpdateGoodAccount 更新goodAccount表记录
// Author [yourname](https://github.com/yourname)
func (goodAccountService *GoodAccountService)UpdateGoodAccount(ctx context.Context, goodAccount good.GoodAccount) (err error) {
	err = global.GVA_DB.Model(&good.GoodAccount{}).Where("id = ?",goodAccount.Id).Updates(&goodAccount).Error
	return err
}

// GetGoodAccount 根据id获取goodAccount表记录
// Author [yourname](https://github.com/yourname)
func (goodAccountService *GoodAccountService)GetGoodAccount(ctx context.Context, id string) (goodAccount good.GoodAccount, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&goodAccount).Error
	return
}
// GetGoodAccountInfoList 分页获取goodAccount表记录
// Author [yourname](https://github.com/yourname)
func (goodAccountService *GoodAccountService)GetGoodAccountInfoList(ctx context.Context, info goodReq.GoodAccountSearch) (list []good.GoodAccount, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&good.GoodAccount{})
    var goodAccounts []good.GoodAccount
    // 如果有条件搜索 下方会自动创建搜索语句
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&goodAccounts).Error
	return  goodAccounts, total, err
}
func (goodAccountService *GoodAccountService)GetGoodAccountPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
