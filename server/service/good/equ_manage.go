package good

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/good"
	goodReq "github.com/flipped-aurora/gin-vue-admin/server/model/good/request"
	"gorm.io/gorm"
	"time"
)

type EquManageService struct {
	DB *gorm.DB
}

type EquManageStats struct {
	Total      int64  `json:"total"`      // 全部设备数量
	Online     int64  `json:"online"`     // 在线设备数量
	Offline    int64  `json:"offline"`    // 离线设备数量
	Unused     int64  `json:"unused"`     // 未启用设备数量
	UpdateTime string `json:"updateTime"` // 数据更新时间
}

func NewEquManageService(db *gorm.DB) *EquManageService {
	return &EquManageService{DB: db}
}

func (equManageService *EquManageService) GetEquManageStats(ctx context.Context) (EquManageStats, error) {
	var stats EquManageStats
	// 获取全部设备数量
	if err := equManageService.DB.Model(&good.EquManage{}).Count(&stats.Total).Error; err != nil {
		return stats, err
	}
	// 获取在线设备数量（假设 equStatus 字段值为 "在线" 代表在线，根据实际枚举调整）
	if err := equManageService.DB.Model(&good.EquManage{}).Where("equ_status =?", "在线").Count(&stats.Online).Error; err != nil {
		return stats, err
	}
	// 获取离线设备数量（假设 equStatus 字段值为 "离线" 代表离线，根据实际枚举调整）
	if err := equManageService.DB.Model(&good.EquManage{}).Where("equ_status =?", "离线").Count(&stats.Offline).Error; err != nil {
		return stats, err
	}
	// 获取未启用设备数量（假设 equStatus 字段值为 "未启用" 代表未启用，根据实际枚举调整）
	if err := equManageService.DB.Model(&good.EquManage{}).Where("equ_status =?", "未启用").Count(&stats.Unused).Error; err != nil {
		return stats, err
	}
	// 设置更新时间（可根据实际需求，比如从数据库记录取或直接用当前时间）
	stats.UpdateTime = time.Now().Format("2006/01/02 15:04:05")
	return stats, nil
}

// CreateEquManage 创建equManage表记录
// Author [yourname](https://github.com/yourname)
func (equManageService *EquManageService) CreateEquManage(ctx context.Context, equManage *good.EquManage) (err error) {
	err = global.GVA_DB.Create(equManage).Error
	return err
}

// DeleteEquManage 删除equManage表记录
// Author [yourname](https://github.com/yourname)
func (equManageService *EquManageService) DeleteEquManage(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&good.EquManage{}, "id = ?", id).Error
	return err
}

// DeleteEquManageByIds 批量删除equManage表记录
// Author [yourname](https://github.com/yourname)
func (equManageService *EquManageService) DeleteEquManageByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]good.EquManage{}, "id in ?", ids).Error
	return err
}

// UpdateEquManage 更新equManage表记录
// Author [yourname](https://github.com/yourname)
func (equManageService *EquManageService) UpdateEquManage(ctx context.Context, equManage good.EquManage) (err error) {
	err = global.GVA_DB.Model(&good.EquManage{}).Where("id = ?", equManage.Id).Updates(&equManage).Error
	return err
}

// GetEquManage 根据id获取equManage表记录
// Author [yourname](https://github.com/yourname)
func (equManageService *EquManageService) GetEquManage(ctx context.Context, id string) (equManage good.EquManage, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&equManage).Error
	return
}

// GetEquManageInfoList 分页获取equManage表记录
// Author [yourname](https://github.com/yourname)
func (equManageService *EquManageService) GetEquManageInfoList(ctx context.Context, info goodReq.EquManageSearch) (list []good.EquManage, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&good.EquManage{})
	var equManages []good.EquManage
	// 如果有条件搜索 下方会自动创建搜索语句

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&equManages).Error
	return equManages, total, err
}

func (equManageService *EquManageService) GetEquManagePublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// EquCount 设备统计
// Author [yourname](https://github.com/yourname)
func (equManageService *EquManageService) EquCount(ctx context.Context) (stats map[string]int64, err error) {
	// 初始化统计结果
	stats = make(map[string]int64)

	// 统计设备总数
	var total int64
	err = global.GVA_DB.Model(&good.EquManage{}).Count(&total).Error
	if err != nil {
		return nil, err
	}
	stats["total"] = total

	// 统计不同状态的设备数量
	// 定义一个结构体来匹配查询结果
	type StatusCount struct {
		EquStatus string `gorm:"column:equ_status"`
		Count     int64  `gorm:"column:count"`
	}
	var statusCounts []StatusCount

	err = global.GVA_DB.Model(&good.EquManage{}).Select("equ_status, COUNT(*) as count").Group("equ_status").Scan(&statusCounts).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(statusCounts)
	// 定义状态映射表
	statusMap := map[string]string{
		"在线":  "online",
		"离线":  "offline",
		"未启用": "disabled",
		// 可以根据需要添加更多映射
	}

	// 将查询结果转换为 map，同时转换状态名称
	for _, sc := range statusCounts {
		// 查找状态映射，如果没有匹配则使用原始值
		englishStatus, exists := statusMap[sc.EquStatus]
		if !exists {
			englishStatus = sc.EquStatus // 没有匹配时保留原始值
		}
		stats[englishStatus] = sc.Count
	}
	fmt.Println(stats)
	return stats, nil
}
