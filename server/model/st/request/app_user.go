package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type AppUserSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Nickname       string     `json:"nickname" form:"nickname" `
	CurrentPoints  *int       `json:"currentPoints" form:"currentPoints" `
	Phone          string     `json:"phone" form:"phone" `
	request.PageInfo
}
