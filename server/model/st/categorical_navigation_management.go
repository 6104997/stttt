// 自动生成模板CategoricalNavigationManagement
package st

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 分类导航管理 结构体  CategoricalNavigationManagement
type CategoricalNavigationManagement struct {
	global.GVA_MODEL
	NavigationName string `json:"navigationName" form:"navigationName" gorm:"column:navigation_name;comment:;"` //导航名
	NavigationUrl  string `json:"navigationUrl" form:"navigationUrl" gorm:"column:navigation_url;comment:;"`    //导航数据url
}

// TableName 分类导航管理 CategoricalNavigationManagement自定义表名 categorical_navigation_management
func (CategoricalNavigationManagement) TableName() string {
	return "categorical_navigation_management"
}
