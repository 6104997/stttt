// 自动生成模板ArticleManagement
package st

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gofrs/uuid/v5"
	"time"
)

// 文章 结构体  ArticleManagement
type ArticleManagement struct {
	global.GVA_MODEL
	Uuid            uuid.UUID `json:"uuid" form:"uuid" gorm:"column:uuid;comment:;"`                                         //uuid
	Title           string    `json:"title" form:"title" gorm:"column:title;comment:;"`                                      //资讯标题
	Article_status  string    `json:"article_status" form:"article_status" gorm:"default:1;column:article_status;comment:;"` //文章状态
	Author          string    `json:"author" form:"author" gorm:"default:大宝;column:author;comment:;"`                        //作者
	View_count      *int      `json:"view_count" form:"view_count" gorm:"default:0;column:view_count;comment:;"`             //阅读次数
	Content         string    `json:"content" form:"content" gorm:"column:content;comment:;type:text;"`                      //文章内容
	Select          *bool     `json:"select" form:"select" gorm:"column:select;comment:;"`                                   //置顶
	AccountNumber   *int      `json:"accountNumber" form:"accountNumber" gorm:"column:account_number;comment:;"`             //账号编号
	Contact         string    `json:"contact" form:"contact" gorm:"column:contact;comment:;"`                                //联系方式
	Price           *int      `json:"price" form:"price" gorm:"column:price;comment:;"`                                      //价格
	PreviewTheImage string    `json:"previewTheImage" form:"previewTheImage" gorm:"column:preview_the_image;comment:;"`      //预览图
}

type ArticleResponse struct {
	CreatedAt       time.Time `json:"created_at"`
	Uuid            string    `json:"uuid"`
	Title           string    `json:"title"`
	PreviewTheImage string    `json:"preview_the_image"`
	Author          string    `json:"author"`
}

// TableName 文章 ArticleManagement自定义表名 article_management
func (ArticleManagement) TableName() string {
	return "article_management"
}
