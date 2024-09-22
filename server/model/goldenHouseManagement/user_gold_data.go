// 自动生成模板UserGoldData
package goldenHouseManagement

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 用户打金数据 结构体  UserGoldData
type UserGoldData struct {
	global.GVA_MODEL
	RoomNumber string `json:"roomNumber" form:"roomNumber" gorm:"column:room_number;comment:;"` //房间号
	UserID1    string `json:"userID1" form:"userID1" gorm:"column:user_id1;comment:;"`          //用户id1
	UserID2    string `json:"userID2" form:"userID2" gorm:"column:user_id2;comment:;"`          //用户id2
	UserID3    string `json:"userID3" form:"userID3" gorm:"column:user_id3;comment:;"`          //用户id3
}

// TableName 用户打金数据 UserGoldData自定义表名 userGoldData
func (UserGoldData) TableName() string {
	return "userGoldData"
}
