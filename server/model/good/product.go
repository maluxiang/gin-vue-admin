
// 自动生成模板Product
package good
import (
	"time"
)

// product表 结构体  Product
type Product struct {
  Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"`  //id字段
  Name  *string `json:"name" form:"name" gorm:"column:name;size:255;"`  //name字段
  Status  *string `json:"status" form:"status" gorm:"column:status;size:255;"`  //status字段
  OnlineStatus  *string `json:"onlineStatus" form:"onlineStatus" gorm:"column:online_status;size:255;"`  //onlineStatus字段
  Remark  *string `json:"remark" form:"remark" gorm:"column:remark;size:255;"`  //remark字段
  CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`  //createdAt字段
  UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;"`  //updatedAt字段
  DeletedAt  *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;"`  //deletedAt字段
}


// TableName product表 Product自定义表名 product
func (Product) TableName() string {
    return "product"
}





