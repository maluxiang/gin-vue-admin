
package good

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/good"
    goodReq "github.com/flipped-aurora/gin-vue-admin/server/model/good/request"
)

type CertificateService struct {}
// CreateCertificate 创建certificate表记录
// Author [yourname](https://github.com/yourname)
func (certificateService *CertificateService) CreateCertificate(ctx context.Context, certificate *good.Certificate) (err error) {
	err = global.GVA_DB.Create(certificate).Error
	return err
}

// DeleteCertificate 删除certificate表记录
// Author [yourname](https://github.com/yourname)
func (certificateService *CertificateService)DeleteCertificate(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&good.Certificate{},"id = ?",id).Error
	return err
}

// DeleteCertificateByIds 批量删除certificate表记录
// Author [yourname](https://github.com/yourname)
func (certificateService *CertificateService)DeleteCertificateByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]good.Certificate{},"id in ?",ids).Error
	return err
}

// UpdateCertificate 更新certificate表记录
// Author [yourname](https://github.com/yourname)
func (certificateService *CertificateService)UpdateCertificate(ctx context.Context, certificate good.Certificate) (err error) {
	err = global.GVA_DB.Model(&good.Certificate{}).Where("id = ?",certificate.Id).Updates(&certificate).Error
	return err
}

// GetCertificate 根据id获取certificate表记录
// Author [yourname](https://github.com/yourname)
func (certificateService *CertificateService)GetCertificate(ctx context.Context, id string) (certificate good.Certificate, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&certificate).Error
	return
}
// GetCertificateInfoList 分页获取certificate表记录
// Author [yourname](https://github.com/yourname)
func (certificateService *CertificateService)GetCertificateInfoList(ctx context.Context, info goodReq.CertificateSearch) (list []good.Certificate, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&good.Certificate{})
    var certificates []good.Certificate
    // 如果有条件搜索 下方会自动创建搜索语句
    
    if info.CertificateName != nil && *info.CertificateName != "" {
        db = db.Where("certificate_name = ?", *info.CertificateName)
    }
    if info.CertificateStatus != nil && *info.CertificateStatus != "" {
        db = db.Where("certificate_status = ?", *info.CertificateStatus)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&certificates).Error
	return  certificates, total, err
}
func (certificateService *CertificateService)GetCertificatePublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}

// Equcount 设备统计
// Author [yourname](https://github.com/yourname)
func (certificateService *CertificateService)Equcount(ctx context.Context) (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&good.Certificate{})
    return db.Error
}


