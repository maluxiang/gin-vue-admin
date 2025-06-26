
// 自动生成模板Certificate
package good
import (
	"time"
)

// certificate表 结构体  Certificate
type Certificate struct {
  Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;size:10;"`  //id字段
  CertificateName  *string `json:"certificateName" form:"certificateName" gorm:"column:certificate_name;size:255;"`  //certificateName字段
  VersionNumber  *string `json:"versionNumber" form:"versionNumber" gorm:"column:version_number;size:255;"`  //versionNumber字段
  CertificateInfo  *string `json:"certificateInfo" form:"certificateInfo" gorm:"column:certificate_info;size:255;"`  //certificateInfo字段
  CertificateStatus  *string `json:"certificateStatus" form:"certificateStatus" gorm:"column:certificate_status;size:255;"`  //certificateStatus字段
  CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;"`  //createdAt字段
}


// TableName certificate表 Certificate自定义表名 certificate
func (Certificate) TableName() string {
    return "certificate"
}





