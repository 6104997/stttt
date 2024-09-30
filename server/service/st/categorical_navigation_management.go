package st

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/st"
	stReq "github.com/flipped-aurora/gin-vue-admin/server/model/st/request"
)

type CategoricalNavigationManagementService struct{}

// CreateCategoricalNavigationManagement 创建分类导航管理记录
// Author [yourname](https://github.com/yourname)
func (categoricalNavigationManagementService *CategoricalNavigationManagementService) CreateCategoricalNavigationManagement(categoricalNavigationManagement *st.CategoricalNavigationManagement) (err error) {
	err = global.GVA_DB.Create(categoricalNavigationManagement).Error
	return err
}

// DeleteCategoricalNavigationManagement 删除分类导航管理记录
// Author [yourname](https://github.com/yourname)
func (categoricalNavigationManagementService *CategoricalNavigationManagementService) DeleteCategoricalNavigationManagement(ID string) (err error) {
	err = global.GVA_DB.Delete(&st.CategoricalNavigationManagement{}, "id = ?", ID).Error
	return err
}

// DeleteCategoricalNavigationManagementByIds 批量删除分类导航管理记录
// Author [yourname](https://github.com/yourname)
func (categoricalNavigationManagementService *CategoricalNavigationManagementService) DeleteCategoricalNavigationManagementByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]st.CategoricalNavigationManagement{}, "id in ?", IDs).Error
	return err
}

// UpdateCategoricalNavigationManagement 更新分类导航管理记录
// Author [yourname](https://github.com/yourname)
func (categoricalNavigationManagementService *CategoricalNavigationManagementService) UpdateCategoricalNavigationManagement(categoricalNavigationManagement st.CategoricalNavigationManagement) (err error) {
	err = global.GVA_DB.Model(&st.CategoricalNavigationManagement{}).Where("id = ?", categoricalNavigationManagement.ID).Updates(&categoricalNavigationManagement).Error
	return err
}

// GetCategoricalNavigationManagement 根据ID获取分类导航管理记录
// Author [yourname](https://github.com/yourname)
func (categoricalNavigationManagementService *CategoricalNavigationManagementService) GetCategoricalNavigationManagement(ID string) (categoricalNavigationManagement st.CategoricalNavigationManagement, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&categoricalNavigationManagement).Error
	return
}

// GetCategoricalNavigationManagementInfoList 分页获取分类导航管理记录
// Author [yourname](https://github.com/yourname)
func (categoricalNavigationManagementService *CategoricalNavigationManagementService) GetCategoricalNavigationManagementInfoList(info stReq.CategoricalNavigationManagementSearch) (list []st.CategoricalNavigationManagement, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&st.CategoricalNavigationManagement{})
	var categoricalNavigationManagements []st.CategoricalNavigationManagement
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.NavigationName != "" {
		db = db.Where("navigation_name LIKE ?", "%"+info.NavigationName+"%")
	}
	if info.NavigationUrl != "" {
		db = db.Where("navigation_url LIKE ?", "%"+info.NavigationUrl+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&categoricalNavigationManagements).Error
	return categoricalNavigationManagements, total, err
}

// GetAListOfNavigationCategories 获取导航数据
// Author [yourname](https://github.com/yourname)
func (categoricalNavigationManagementService *CategoricalNavigationManagementService) GetAListOfNavigationCategories() (list []st.CategoricalNavigationManagement, err error) {
	// 请在这里实现自己的业务逻辑
	err = global.GVA_DB.Find(&list).Error
	return
}
