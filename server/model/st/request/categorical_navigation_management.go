package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CategoricalNavigationManagementSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	NavigationName string     `json:"navigationName" form:"navigationName" `
	NavigationUrl  string     `json:"navigationUrl" form:"navigationUrl" `
	request.PageInfo
}
