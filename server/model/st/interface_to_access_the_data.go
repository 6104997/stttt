// 自动生成模板InterfaceToAccessTheData
package st

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/gofrs/uuid/v5"
)

// 接口访问数据管理 结构体  InterfaceToAccessTheData
type InterfaceToAccessTheData struct {
	global.GVA_MODEL
	Nickname              string    `json:"nickname" form:"nickname" gorm:"column:nickname;comment:;"`                                            //用户昵称
	Uuid                  uuid.UUID `json:"uuid" form:"uuid" gorm:"column:uuid;comment:;"`                                                        //uuid
	TheNameOfTheInterface string    `json:"theNameOfTheInterface" form:"theNameOfTheInterface" gorm:"column:the_name_of_the_interface;comment:;"` //接口名称
	Ip                    string    `json:"ip" form:"ip" gorm:"column:ip;comment:;"`                                                              //ip
	Url                   string    `json:"url" form:"url" gorm:"column:url;comment:;"`                                                           //url
}

// TableName 接口访问数据管理 InterfaceToAccessTheData自定义表名 interface_to_access_the_data
func (InterfaceToAccessTheData) TableName() string {
	return "interface_to_access_the_data"
}
