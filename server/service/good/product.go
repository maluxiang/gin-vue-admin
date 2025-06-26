
package good

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/good"
    goodReq "github.com/flipped-aurora/gin-vue-admin/server/model/good/request"
)

type ProductService struct {}
// CreateProduct 创建product表记录
// Author [yourname](https://github.com/yourname)
func (productService *ProductService) CreateProduct(ctx context.Context, product *good.Product) (err error) {
	err = global.GVA_DB.Create(product).Error
	return err
}

// DeleteProduct 删除product表记录
// Author [yourname](https://github.com/yourname)
func (productService *ProductService)DeleteProduct(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&good.Product{},"id = ?",id).Error
	return err
}

// DeleteProductByIds 批量删除product表记录
// Author [yourname](https://github.com/yourname)
func (productService *ProductService)DeleteProductByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]good.Product{},"id in ?",ids).Error
	return err
}

// UpdateProduct 更新product表记录
// Author [yourname](https://github.com/yourname)
func (productService *ProductService)UpdateProduct(ctx context.Context, product good.Product) (err error) {
	err = global.GVA_DB.Model(&good.Product{}).Where("id = ?",product.Id).Updates(&product).Error
	return err
}

// GetProduct 根据id获取product表记录
// Author [yourname](https://github.com/yourname)
func (productService *ProductService)GetProduct(ctx context.Context, id string) (product good.Product, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&product).Error
	return
}
// GetProductInfoList 分页获取product表记录
// Author [yourname](https://github.com/yourname)
func (productService *ProductService)GetProductInfoList(ctx context.Context, info goodReq.ProductSearch) (list []good.Product, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&good.Product{})
    var products []good.Product
    // 如果有条件搜索 下方会自动创建搜索语句
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&products).Error
	return  products, total, err
}
func (productService *ProductService)GetProductPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
