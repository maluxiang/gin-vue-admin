package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type AlarmManageSearch struct {
	EquId          *string     `json:"equId" form:"equId"`
	EquName        *string     `json:"equName" form:"equName"`
	AlarmTimeRange []time.Time `json:"alarmTimeRange" form:"alarmTimeRange[]"`
	request.PageInfo
}
