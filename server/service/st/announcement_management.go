package st

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/st"
	stReq "github.com/flipped-aurora/gin-vue-admin/server/model/st/request"
)

type AnnouncementManagementService struct{}

// CreateAnnouncementManagement 创建公告管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (announcementManagementService *AnnouncementManagementService) CreateAnnouncementManagement(announcementManagement *st.AnnouncementManagement) (err error) {
	err = global.GVA_DB.Create(announcementManagement).Error
	return err
}

// DeleteAnnouncementManagement 删除公告管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (announcementManagementService *AnnouncementManagementService) DeleteAnnouncementManagement(ID string) (err error) {
	err = global.GVA_DB.Delete(&st.AnnouncementManagement{}, "id = ?", ID).Error
	return err
}

// DeleteAnnouncementManagementByIds 批量删除公告管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (announcementManagementService *AnnouncementManagementService) DeleteAnnouncementManagementByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]st.AnnouncementManagement{}, "id in ?", IDs).Error
	return err
}

// UpdateAnnouncementManagement 更新公告管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (announcementManagementService *AnnouncementManagementService) UpdateAnnouncementManagement(announcementManagement st.AnnouncementManagement) (err error) {
	err = global.GVA_DB.Model(&st.AnnouncementManagement{}).Where("id = ?", announcementManagement.ID).Updates(&announcementManagement).Error
	return err
}

// GetAnnouncementManagement 根据ID获取公告管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (announcementManagementService *AnnouncementManagementService) GetAnnouncementManagement(ID string) (announcementManagement st.AnnouncementManagement, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&announcementManagement).Error
	return
}

// GetAnnouncementManagementInfoList 分页获取公告管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (announcementManagementService *AnnouncementManagementService) GetAnnouncementManagementInfoList(info stReq.AnnouncementManagementSearch) (list []st.AnnouncementManagement, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&st.AnnouncementManagement{})
	var announcementManagements []st.AnnouncementManagement
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&announcementManagements).Error
	return announcementManagements, total, err
}
func (announcementManagementService *AnnouncementManagementService) GetAnnouncementManagementPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// GetAListOfAnnouncements 方法介绍
// Author [yourname](https://github.com/yourname)
func (announcementManagementService *AnnouncementManagementService) GetAListOfAnnouncements() (AnnouncementManagements []st.AnnouncementManagement, err error) {
	// 获取公告列表
	err = global.GVA_DB.Find(&AnnouncementManagements).Error
	return
}
