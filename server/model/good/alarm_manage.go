
// 自动生成模板AlarmManage
package good
import (
	"time"
)

// alarmManage表 结构体  AlarmManage
type AlarmManage struct {
  Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;size:20;"`  //id字段
  EquId  *string `json:"equId" form:"equId" gorm:"column:equ_id;size:255;"`  //equId字段
  EquName  *string `json:"equName" form:"equName" gorm:"column:equ_name;size:255;"`  //equName字段
  AlarmStatus  *string `json:"alarmStatus" form:"alarmStatus" gorm:"column:alarm_status;size:255;"`  //alarmStatus字段
  AlarmTime  *time.Time `json:"alarmTime" form:"alarmTime" gorm:"column:alarm_time;"`  //alarmTime字段
  CpuStatus  *string `json:"cpuStatus" form:"cpuStatus" gorm:"column:cpu_status;size:255;"`  //cpuStatus字段
}


// TableName alarmManage表 AlarmManage自定义表名 alarm_manage
func (AlarmManage) TableName() string {
    return "alarm_manage"
}





