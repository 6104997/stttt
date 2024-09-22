package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type UserGoldDataSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	UserID1        string     `json:"userID1" form:"userID1" `
	UserID2        string     `json:"userID2" form:"userID2" `
	UserID3        string     `json:"userID3" form:"userID3" `
	request.PageInfo
}
