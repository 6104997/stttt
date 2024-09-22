package goldenHouseManagement

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/goldenHouseManagement"
	goldenHouseManagementReq "github.com/flipped-aurora/gin-vue-admin/server/model/goldenHouseManagement/request"
)

type UserGoldDataService struct{}

// CreateUserGoldData 创建用户打金数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (userGoldDataService *UserGoldDataService) CreateUserGoldData(userGoldData *goldenHouseManagement.UserGoldData) (err error) {
	err = global.GVA_DB.Create(userGoldData).Error
	return err
}

// DeleteUserGoldData 删除用户打金数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (userGoldDataService *UserGoldDataService) DeleteUserGoldData(ID string) (err error) {
	err = global.GVA_DB.Delete(&goldenHouseManagement.UserGoldData{}, "id = ?", ID).Error
	return err
}

// DeleteUserGoldDataByIds 批量删除用户打金数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (userGoldDataService *UserGoldDataService) DeleteUserGoldDataByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]goldenHouseManagement.UserGoldData{}, "id in ?", IDs).Error
	return err
}

// UpdateUserGoldData 更新用户打金数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (userGoldDataService *UserGoldDataService) UpdateUserGoldData(userGoldData goldenHouseManagement.UserGoldData) (err error) {
	err = global.GVA_DB.Model(&goldenHouseManagement.UserGoldData{}).Where("id = ?", userGoldData.ID).Updates(&userGoldData).Error
	return err
}

// GetUserGoldData 根据ID获取用户打金数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (userGoldDataService *UserGoldDataService) GetUserGoldData(ID string) (userGoldData goldenHouseManagement.UserGoldData, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&userGoldData).Error
	return
}

// GetUserGoldDataInfoList 分页获取用户打金数据记录
// Author [piexlmax](https://github.com/piexlmax)
func (userGoldDataService *UserGoldDataService) GetUserGoldDataInfoList(info goldenHouseManagementReq.UserGoldDataSearch) (list []goldenHouseManagement.UserGoldData, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&goldenHouseManagement.UserGoldData{})
	var userGoldDatas []goldenHouseManagement.UserGoldData
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.UserID1 != "" {
		db = db.Where("user_id1 LIKE ?", "%"+info.UserID1+"%")
	}
	if info.UserID2 != "" {
		db = db.Where("user_id2 LIKE ?", "%"+info.UserID2+"%")
	}
	if info.UserID3 != "" {
		db = db.Where("user_id3 LIKE ?", "%"+info.UserID3+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&userGoldDatas).Error
	return userGoldDatas, total, err
}
func (userGoldDataService *UserGoldDataService) GetUserGoldDataPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// AddTheUserSGoldData 添加打金用户数据
// Author [yourname](https://github.com/yourname)
func (userGoldDataService *UserGoldDataService) AddTheUserSGoldData() (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&goldenHouseManagement.UserGoldData{})
	return db.Error
}
