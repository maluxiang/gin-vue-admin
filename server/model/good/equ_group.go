
// 自动生成模板EquGroup
package good
import (
)

// equGroup表 结构体  EquGroup
type EquGroup struct {
  Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"`  //id字段
  GroupId  *string `json:"groupId" form:"groupId" gorm:"column:group_id;size:255;"`  //groupId字段
  GroupName  *string `json:"groupName" form:"groupName" gorm:"column:group_name;size:255;"`  //groupName字段
  Remark  *string `json:"remark" form:"remark" gorm:"column:remark;size:255;"`  //remark字段
}


// TableName equGroup表 EquGroup自定义表名 equ_group
func (EquGroup) TableName() string {
    return "equ_group"
}





