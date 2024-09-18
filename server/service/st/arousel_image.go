package st

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/st"
	stReq "github.com/flipped-aurora/gin-vue-admin/server/model/st/request"
)

type ArouselImageService struct{}

// CreateArouselImage 创建arouselImage表记录
// Author [piexlmax](https://github.com/piexlmax)
func (arouselImageService *ArouselImageService) CreateArouselImage(arouselImage *st.ArouselImage) (err error) {
	err = global.GVA_DB.Create(arouselImage).Error
	return err
}

// DeleteArouselImage 删除arouselImage表记录
// Author [piexlmax](https://github.com/piexlmax)
func (arouselImageService *ArouselImageService) DeleteArouselImage(ID string) (err error) {
	err = global.GVA_DB.Delete(&st.ArouselImage{}, "id = ?", ID).Error
	return err
}

// DeleteArouselImageByIds 批量删除arouselImage表记录
// Author [piexlmax](https://github.com/piexlmax)
func (arouselImageService *ArouselImageService) DeleteArouselImageByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]st.ArouselImage{}, "id in ?", IDs).Error
	return err
}

// UpdateArouselImage 更新arouselImage表记录
// Author [piexlmax](https://github.com/piexlmax)
func (arouselImageService *ArouselImageService) UpdateArouselImage(arouselImage st.ArouselImage) (err error) {
	err = global.GVA_DB.Model(&st.ArouselImage{}).Where("id = ?", arouselImage.ID).Updates(&arouselImage).Error
	return err
}

// GetArouselImage 根据ID获取arouselImage表记录
// Author [piexlmax](https://github.com/piexlmax)
func (arouselImageService *ArouselImageService) GetArouselImage(ID string) (arouselImage st.ArouselImage, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&arouselImage).Error
	return
}

// GetArouselImageInfoList 分页获取arouselImage表记录
// Author [piexlmax](https://github.com/piexlmax)
func (arouselImageService *ArouselImageService) GetArouselImageInfoList(info stReq.ArouselImageSearch) (list []st.ArouselImage, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&st.ArouselImage{})
	var arouselImages []st.ArouselImage
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

	err = db.Find(&arouselImages).Error
	return arouselImages, total, err
}
func (arouselImageService *ArouselImageService) GetArouselImagePublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
