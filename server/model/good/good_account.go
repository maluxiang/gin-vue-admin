
// 自动生成模板GoodAccount
package good
import (
)

// goodAccount表 结构体  GoodAccount
type GoodAccount struct {
  Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;size:10;"`  //id字段
  Pictrue  string `json:"pictrue" form:"pictrue" gorm:"column:pictrue;size:255;"`  //pictrue字段
  Name  *string `json:"name" form:"name" gorm:"column:name;size:255;"`  //name字段
  Count  *int `json:"count" form:"count" gorm:"column:count;size:10;"`  //count字段
  Status  *string `json:"status" form:"status" gorm:"column:status;size:255;"`  //status字段
  EquStatus  *string `json:"equStatus" form:"equStatus" gorm:"column:equ_status;size:255;"`  //equStatus字段
}


// TableName goodAccount表 GoodAccount自定义表名 good_account
func (GoodAccount) TableName() string {
    return "good_account"
}





