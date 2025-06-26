
// 自动生成模板EquManage
package good
import (
	"time"
)

// equManage表 结构体  EquManage
type EquManage struct {
  Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"`  //id字段
  EquId  *string `json:"equId" form:"equId" gorm:"column:equ_id;size:255;"`  //equId字段
  EquName  *string `json:"equName" form:"equName" gorm:"column:equ_name;size:255;"`  //equName字段
  EquClass  *string `json:"equClass" form:"equClass" gorm:"column:equ_class;size:255;"`  //equClass字段
  EquStatus  *string `json:"equStatus" form:"equStatus" gorm:"column:equ_status;size:255;"`  //equStatus字段
  Remark  *string `json:"remark" form:"remark" gorm:"column:remark;size:255;"`  //remark字段
  CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`  //createdAt字段
}


// TableName equManage表 EquManage自定义表名 equ_manage
func (EquManage) TableName() string {
    return "equ_manage"
}





