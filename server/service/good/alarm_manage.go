
package good

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/good"
    goodReq "github.com/flipped-aurora/gin-vue-admin/server/model/good/request"
)

type AlarmManageService struct {}
// CreateAlarmManage 创建alarmManage表记录
// Author [yourname](https://github.com/yourname)
func (alarmManageService *AlarmManageService) CreateAlarmManage(ctx context.Context, alarmManage *good.AlarmManage) (err error) {
	err = global.GVA_DB.Create(alarmManage).Error
	return err
}

// DeleteAlarmManage 删除alarmManage表记录
// Author [yourname](https://github.com/yourname)
func (alarmManageService *AlarmManageService)DeleteAlarmManage(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&good.AlarmManage{},"id = ?",id).Error
	return err
}

// DeleteAlarmManageByIds 批量删除alarmManage表记录
// Author [yourname](https://github.com/yourname)
func (alarmManageService *AlarmManageService)DeleteAlarmManageByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]good.AlarmManage{},"id in ?",ids).Error
	return err
}

// UpdateAlarmManage 更新alarmManage表记录
// Author [yourname](https://github.com/yourname)
func (alarmManageService *AlarmManageService)UpdateAlarmManage(ctx context.Context, alarmManage good.AlarmManage) (err error) {
	err = global.GVA_DB.Model(&good.AlarmManage{}).Where("id = ?",alarmManage.Id).Updates(&alarmManage).Error
	return err
}

// GetAlarmManage 根据id获取alarmManage表记录
// Author [yourname](https://github.com/yourname)
func (alarmManageService *AlarmManageService)GetAlarmManage(ctx context.Context, id string) (alarmManage good.AlarmManage, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&alarmManage).Error
	return
}
// GetAlarmManageInfoList 分页获取alarmManage表记录
// Author [yourname](https://github.com/yourname)
func (alarmManageService *AlarmManageService)GetAlarmManageInfoList(ctx context.Context, info goodReq.AlarmManageSearch) (list []good.AlarmManage, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&good.AlarmManage{})
    var alarmManages []good.AlarmManage
    // 如果有条件搜索 下方会自动创建搜索语句
    
    if info.EquId != nil && *info.EquId != "" {
        db = db.Where("equ_id = ?", *info.EquId)
    }
    if info.EquName != nil && *info.EquName != "" {
        db = db.Where("equ_name = ?", *info.EquName)
    }
			if len(info.AlarmTimeRange) == 2 {
				db = db.Where("alarm_time BETWEEN ? AND ? ", info.AlarmTimeRange[0], info.AlarmTimeRange[1])
			}
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&alarmManages).Error
	return  alarmManages, total, err
}
func (alarmManageService *AlarmManageService)GetAlarmManagePublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
