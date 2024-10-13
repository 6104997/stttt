package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type InterfaceToAccessTheDataSearch struct {
	StartCreatedAt        *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt          *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Nickname              string     `json:"nickname" form:"nickname" `
	TheNameOfTheInterface string     `json:"theNameOfTheInterface" form:"theNameOfTheInterface" `
	Ip                    string     `json:"ip" form:"ip" `
	Url                   string     `json:"url" form:"url" `
	request.PageInfo
}
