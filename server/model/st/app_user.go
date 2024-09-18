// 自动生成模板AppUser
package st

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gofrs/uuid/v5"
)

// appUser表 结构体  AppUser
type AppUser struct {
	global.GVA_MODEL
	Nickname      string    `json:"nickname" form:"nickname" gorm:"column:nickname;comment:;size:191;"`                                                                        //昵称
	HeadPortrait  string    `json:"headPortrait" form:"headPortrait" gorm:"default:https://8.134.253.195:16624/down/v4wAyJrfG3rh.png;column:head_portrait;comment:;size:191;"` //头像
	Openpid       string    `json:"openpid" form:"openpid" gorm:"column:openpid;comment:;size:191;"`                                                                           //openpid
	Uuid          uuid.UUID `json:"uuid" form:"uuid" gorm:"column:uuid;comment:;size:191;"`                                                                                    //uuid
	CurrentPoints *int      `json:"currentPoints" form:"currentPoints" gorm:"column:current_points;comment:;"`                                                                 //当前积分
	SignIn        *int      `json:"signIn" form:"signIn" gorm:"column:sign_in;comment:;"`                                                                                      //签到
	Phone         string    `json:"phone" form:"phone" gorm:"column:phone;comment:;size:11;"`                                                                                  //电话号
}

// TableName appUser表 AppUser自定义表名 app_user
func (AppUser) TableName() string {
	return "app_user"
}
