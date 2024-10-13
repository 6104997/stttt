package st

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/st"
	stReq "github.com/flipped-aurora/gin-vue-admin/server/model/st/request"
	"log"
)

type InterfaceToAccessTheDataService struct{}

// CreateInterfaceToAccessTheData 创建接口访问数据管理记录
// Author [yourname](https://github.com/yourname)
func (interfaceToAccessTheDataService *InterfaceToAccessTheDataService) CreateInterfaceToAccessTheData(interfaceToAccessTheData *st.InterfaceToAccessTheData) (err error) {
	err = global.GVA_DB.Create(interfaceToAccessTheData).Error
	return err
}

// DeleteInterfaceToAccessTheData 删除接口访问数据管理记录
// Author [yourname](https://github.com/yourname)
func (interfaceToAccessTheDataService *InterfaceToAccessTheDataService) DeleteInterfaceToAccessTheData(ID string) (err error) {
	err = global.GVA_DB.Delete(&st.InterfaceToAccessTheData{}, "id = ?", ID).Error
	return err
}

// DeleteInterfaceToAccessTheDataByIds 批量删除接口访问数据管理记录
// Author [yourname](https://github.com/yourname)
func (interfaceToAccessTheDataService *InterfaceToAccessTheDataService) DeleteInterfaceToAccessTheDataByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]st.InterfaceToAccessTheData{}, "id in ?", IDs).Error
	return err
}

// UpdateInterfaceToAccessTheData 更新接口访问数据管理记录
// Author [yourname](https://github.com/yourname)
func (interfaceToAccessTheDataService *InterfaceToAccessTheDataService) UpdateInterfaceToAccessTheData(interfaceToAccessTheData st.InterfaceToAccessTheData) (err error) {
	err = global.GVA_DB.Model(&st.InterfaceToAccessTheData{}).Where("id = ?", interfaceToAccessTheData.ID).Updates(&interfaceToAccessTheData).Error
	return err
}

// GetInterfaceToAccessTheData 根据ID获取接口访问数据管理记录
// Author [yourname](https://github.com/yourname)
func (interfaceToAccessTheDataService *InterfaceToAccessTheDataService) GetInterfaceToAccessTheData(ID string) (interfaceToAccessTheData st.InterfaceToAccessTheData, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&interfaceToAccessTheData).Error
	return
}

// GetInterfaceToAccessTheDataInfoList 分页获取接口访问数据管理记录
// Author [yourname](https://github.com/yourname)
func (interfaceToAccessTheDataService *InterfaceToAccessTheDataService) GetInterfaceToAccessTheDataInfoList(info stReq.InterfaceToAccessTheDataSearch) (list []st.InterfaceToAccessTheData, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&st.InterfaceToAccessTheData{})
	var interfaceToAccessTheDatas []st.InterfaceToAccessTheData
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Nickname != "" {
		db = db.Where("nickname LIKE ?", "%"+info.Nickname+"%")
	}
	if info.TheNameOfTheInterface != "" {
		db = db.Where("the_name_of_the_interface LIKE ?", "%"+info.TheNameOfTheInterface+"%")
	}
	if info.Ip != "" {
		db = db.Where("ip LIKE ?", "%"+info.Ip+"%")
	}
	if info.Url != "" {
		db = db.Where("url = ?", info.Url)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&interfaceToAccessTheDatas).Error
	return interfaceToAccessTheDatas, total, err
}

func (interfaceToAccessTheDataService *InterfaceToAccessTheDataService) AddUserAccessRecords(id uint, ip string, theNameOfTheInterface string, url string) (err error) {
	user := st.AppUser{}
	log.Printf("Failed to find user with ID %d: %v", id, err)
	if err = global.GVA_DB.First(&user, id).Error; err != nil {
		log.Printf("Failed to find user with ID %d: %v", id, err)
		return err
	}

	interfaceToAccessTheData := st.InterfaceToAccessTheData{
		Nickname:              user.Nickname,
		Uuid:                  user.Uuid,
		TheNameOfTheInterface: theNameOfTheInterface,
		Ip:                    ip,
		Url:                   url,
	}
	if err = global.GVA_DB.Create(&interfaceToAccessTheData).Error; err != nil {
		log.Printf("Failed to create access record for user %d: %v", id, err)
		return err
	}

	// 打印结果
	log.Printf("Successfully added access record for user %d", id)

	return nil
}
