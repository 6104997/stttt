package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type GoldenRoomFormSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Id             int        `json:"id" form:"id"`
	RoomNumber     string     `json:"roomNumber" form:"roomNumber" `
	RoomPassword   string     `json:"roomPassword" form:"roomPassword" `
	CurrentStatus  string     `json:"currentStatus" form:"currentStatus" `
	request.PageInfo
}
