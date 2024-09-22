// 自动生成模板AnnouncementManagement
package st

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 公告管理 结构体  AnnouncementManagement
type AnnouncementManagement struct {
	global.GVA_MODEL
	TheContentOfTheAnnouncement string `json:"theContentOfTheAnnouncement" form:"theContentOfTheAnnouncement" gorm:"column:the_content_of_the_announcement;comment:;" binding:"required"` //公告内容
}

// TableName 公告管理 AnnouncementManagement自定义表名 announcement_management
func (AnnouncementManagement) TableName() string {
	return "announcement_management"
}
