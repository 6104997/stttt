package goldenHouseManagement

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/goldenHouseManagement"
	goldenHouseManagementReq "github.com/flipped-aurora/gin-vue-admin/server/model/goldenHouseManagement/request"
)

type GoldenRoomFormService struct{}

// CreateGoldenRoomForm 创建金房表单记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldenRoomFormService *GoldenRoomFormService) CreateGoldenRoomForm(goldenRoomForm *goldenHouseManagement.GoldenRoomForm) (err error) {
	err = global.GVA_DB.Create(goldenRoomForm).Error
	return err
}

// DeleteGoldenRoomForm 删除金房表单记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldenRoomFormService *GoldenRoomFormService) DeleteGoldenRoomForm(ID string) (err error) {
	err = global.GVA_DB.Delete(&goldenHouseManagement.GoldenRoomForm{}, "id = ?", ID).Error
	return err
}

// DeleteGoldenRoomFormByIds 批量删除金房表单记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldenRoomFormService *GoldenRoomFormService) DeleteGoldenRoomFormByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]goldenHouseManagement.GoldenRoomForm{}, "id in ?", IDs).Error
	return err
}

// UpdateGoldenRoomForm 更新金房表单记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldenRoomFormService *GoldenRoomFormService) UpdateGoldenRoomForm(goldenRoomForm goldenHouseManagement.GoldenRoomForm) (err error) {
	err = global.GVA_DB.Model(&goldenHouseManagement.GoldenRoomForm{}).Where("id = ?", goldenRoomForm.ID).Updates(&goldenRoomForm).Error
	return err
}

// GetGoldenRoomForm 根据ID获取金房表单记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldenRoomFormService *GoldenRoomFormService) GetGoldenRoomForm(ID string) (goldenRoomForm goldenHouseManagement.GoldenRoomForm, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&goldenRoomForm).Error
	return
}

// GetGoldenRoomFormInfoList 分页获取金房表单记录
// Author [piexlmax](https://github.com/piexlmax)
func (goldenRoomFormService *GoldenRoomFormService) GetGoldenRoomFormInfoList(info goldenHouseManagementReq.GoldenRoomFormSearch) (list []goldenHouseManagement.GoldenRoomForm, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&goldenHouseManagement.GoldenRoomForm{})
	var goldenRoomForms []goldenHouseManagement.GoldenRoomForm
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.RoomNumber != "" {
		db = db.Where("room_number = ?", info.RoomNumber)
	}
	if info.RoomPassword != "" {
		db = db.Where("room_password = ?", info.RoomPassword)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&goldenRoomForms).Error
	return goldenRoomForms, total, err
}
func (goldenRoomFormService *GoldenRoomFormService) GetGoldenRoomFormPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// GetAListOfGoldenHouses 获取金房列表
// Author [yourname](https://github.com/yourname)
func (goldenRoomFormService *GoldenRoomFormService) GetAListOfGoldenHouses() (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&goldenHouseManagement.GoldenRoomForm{})
	return db.Error
}

// AddTheGoldenRoomData 添加金房数据
// Author [yourname](https://github.com/yourname)
func (goldenRoomFormService *GoldenRoomFormService) AddTheGoldenRoomData() (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&goldenHouseManagement.GoldenRoomForm{})
	return db.Error
}

// DeleteTheGoldenRoomData 删除金房数据
// Author [yourname](https://github.com/yourname)
func (goldenRoomFormService *GoldenRoomFormService) DeleteTheGoldenRoomData() (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&goldenHouseManagement.GoldenRoomForm{})
	return db.Error
}

// UpdateTheGoldenRoomData 更新金房数据
// Author [yourname](https://github.com/yourname)
func (goldenRoomFormService *GoldenRoomFormService) UpdateTheGoldenRoomData() (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&goldenHouseManagement.GoldenRoomForm{})
	return db.Error
}
