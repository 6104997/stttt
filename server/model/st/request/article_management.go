package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ArticleManagementSearch struct {
	StartCreatedAt          *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt            *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Title                   string     `json:"title" form:"title" `
	Article_status          string     `json:"article_status" form:"article_status" `
	ArticleClassificationId string     `json:"article_classification_id" form:"article_classification_id" `
	Author                  string     `json:"author" form:"author" `
	View_count              *int       `json:"view_count" form:"view_count" `
	Select                  *bool      `json:"select" form:"select" `
	AccountNumber           *int       `json:"accountNumber" form:"accountNumber" `
	Contact                 string     `json:"contact" form:"contact" `
	Price                   *int       `json:"price" form:"price" `
	TheLatestArticles       string     `json:"theLatestArticles" form:"theLatestArticles" `
	AccountType             string     `json:"account_type" form:"account_type" `
	request.PageInfo
}
