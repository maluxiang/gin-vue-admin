
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type CertificateSearch struct{
      CertificateName  *string `json:"certificateName" form:"certificateName"` 
      CertificateStatus  *string `json:"certificateStatus" form:"certificateStatus"` 
    request.PageInfo
}
