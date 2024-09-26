package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ArticleManagementSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Title          string     `json:"title" form:"title" `
	Article_status string     `json:"article_status" form:"article_status" `
	Author         string     `json:"author" form:"author" `
	View_count     *int       `json:"view_count" form:"view_count" `
	Select         *bool      `json:"select" form:"select" `
	AccountNumber  *int       `json:"accountNumber" form:"accountNumber" `
	Contact        string     `json:"contact" form:"contact" `
	Price          *int       `json:"price" form:"price" `
	request.PageInfo
}
