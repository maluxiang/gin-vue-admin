
// 自动生成模板GoodType
package good
import (
)

// goodType表 结构体  GoodType
type GoodType struct {
  TypeId  *int `json:"typeId" form:"typeId" gorm:"primarykey;column:type_id;size:10;"`  //typeId字段
  EqumentType  *string `json:"equmentType" form:"equmentType" gorm:"column:equment_type;size:255;"`  //equmentType字段
  WorkType  *string `json:"workType" form:"workType" gorm:"column:work_type;size:255;"`  //workType字段
  BusinessType  *string `json:"businessType" form:"businessType" gorm:"column:business_type;size:255;"`  //businessType字段
  Remark  *string `json:"remark" form:"remark" gorm:"column:remark;size:255;"`  //remark字段
}


// TableName goodType表 GoodType自定义表名 good_type
func (GoodType) TableName() string {
    return "good_type"
}





