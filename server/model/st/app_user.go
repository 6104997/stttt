// 自动生成模板AppUser
package st

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gofrs/uuid/v5"
)

var _ system.Login = (*AppUser)(nil)

type SessionResponse struct {
	SessionKey string `json:"session_key"`
	OpenID     string `json:"openid"`
}

// appUser表 结构体  AppUser
type AppUser struct {
	global.GVA_MODEL
	Nickname      string    `json:"nickname" form:"nickname" gorm:"default:微信用户;column:nickname;comment:昵称;size:191;"`                                                   //昵称
	HeadPortrait  string    `json:"headPortrait" form:"headPortrait" gorm:"default:https://stamdin.gys9.com/w700d1q75cms.jpg;column:head_portrait;comment:头像;size:191;"` //头像
	Openpid       string    `json:"openpid" form:"openpid" gorm:"column:openpid;comment:;size:191;"`                                                                     //openpid
	Uuid          uuid.UUID `json:"uuid" form:"uuid" gorm:"column:uuid;comment:uuid;size:191;"`                                                                          //uuid
	CurrentPoints *int      `json:"currentPoints" form:"currentPoints" gorm:"default:0;column:current_points;comment:当前积分;"`                                             //当前积分
	SignIn        *int      `json:"signIn" form:"signIn" gorm:"default:0;column:sign_in;comment:签到;"`                                                                    //签到
	Phone         string    `json:"phone" form:"phone" gorm:"column:phone;comment:电话号;size:11;"`                                                                         //电话号
	AuthorityId   uint      `json:"authorityId" form:"authorityId" gorm:"default:9999;column:authority_id;comment:角色ID;size:191;"`                                       //角色ID
}

func (u AppUser) GetUsername() string {
	return u.Openpid
}

func (u AppUser) GetNickname() string {
	return u.Nickname
}

func (u AppUser) GetUUID() uuid.UUID {
	return u.Uuid
}

func (u AppUser) GetUserId() uint {
	return u.ID
}

func (u AppUser) GetAuthorityId() uint {
	return u.AuthorityId
}

func (u AppUser) GetUserInfo() any {
	return u
}

// TableName appUser表 AppUser自定义表名 app_user
func (AppUser) TableName() string {
	return "app_user"
}
