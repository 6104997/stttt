// 自动生成模板GoldenRoomForm
package goldenHouseManagement

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 金房表单 结构体  GoldenRoomForm
type GoldenRoomForm struct {
	global.GVA_MODEL
	RoomNumber    string `json:"roomNumber" form:"roomNumber" gorm:"column:room_number;comment:;"`          //房间号
	RoomPassword  string `json:"roomPassword" form:"roomPassword" gorm:"column:room_password;comment:;"`    //房间密码
	CurrentStatus string `json:"currentStatus" form:"currentStatus" gorm:"column:current_status;comment:;"` //当前状态
	CreationTime  string `json:"creationTime" form:"creationTime" gorm:"column:creation_time;comment:;"`    //创建时间
}

// TableName 金房表单 GoldenRoomForm自定义表名 goldenRoomForm
func (GoldenRoomForm) TableName() string {
	return "goldenRoomForm"
}
