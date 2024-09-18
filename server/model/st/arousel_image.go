// 自动生成模板ArouselImage
package st

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// arouselImage表 结构体  ArouselImage
type ArouselImage struct {
	global.GVA_MODEL
	Image                              string `json:"image" form:"image" gorm:"column:image;comment:;size:191;" binding:"required"`                                                                        //轮播图
	IntroductionToRotatingBroadcasting string `json:"introductionToRotatingBroadcasting" form:"introductionToRotatingBroadcasting" gorm:"column:introduction_to_rotating_broadcasting;comment:;size:191;"` //轮播介绍
	Path                               string `json:"path" form:"path" gorm:"column:path;comment:;size:191;"`                                                                                              //跳转地址
}

// TableName arouselImage表 ArouselImage自定义表名 arousel_image
func (ArouselImage) TableName() string {
	return "arousel_image"
}
