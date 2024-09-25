package st

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/st"
	stReq "github.com/flipped-aurora/gin-vue-admin/server/model/st/request"
	"github.com/gofrs/uuid/v5"
)

type AppUserService struct{}

// CreateAppUser 创建appUser表记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) CreateAppUser(appUser *st.AppUser) (err error) {
	err = global.GVA_DB.Create(appUser).Error
	return err
}

// DeleteAppUser 删除appUser表记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) DeleteAppUser(ID string) (err error) {
	err = global.GVA_DB.Delete(&st.AppUser{}, "id = ?", ID).Error
	return err
}

// DeleteAppUserByIds 批量删除appUser表记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) DeleteAppUserByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]st.AppUser{}, "id in ?", IDs).Error
	return err
}

// UpdateAppUser 更新appUser表记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) UpdateAppUser(appUser st.AppUser) (err error) {
	err = global.GVA_DB.Model(&st.AppUser{}).Where("id = ?", appUser.ID).Updates(&appUser).Error
	return err
}

// GetAppUser 根据ID获取appUser表记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) GetAppUser(ID string) (appUser st.AppUser, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&appUser).Error
	return
}

// GetAppUserInfoList 分页获取appUser表记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) GetAppUserInfoList(info stReq.AppUserSearch) (list []st.AppUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&st.AppUser{})
	var appUsers []st.AppUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Nickname != "" {
		db = db.Where("nickname LIKE ?", "%"+info.Nickname+"%")
	}
	if info.CurrentPoints != nil {
		db = db.Where("current_points <> ?", info.CurrentPoints)
	}
	if info.Phone != "" {
		db = db.Where("phone = ?", info.Phone)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&appUsers).Error
	return appUsers, total, err
}
func (appUserService *AppUserService) GetAppUserPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// Login 方法介绍
// Author [yourname](https://github.com/yourname)
func (appUserService *AppUserService) Login(openpid string) (appUser st.AppUser, err error) {
	// 查询数据库中是否存在openpid,如果存在则返回数据，如果不存在则创建数据并返回数据
	err = global.GVA_DB.Where("openpid = ?", openpid).First(&appUser).Error
	if err != nil {

		appUser.Openpid = openpid
		appUser.Uuid = uuid.Must(uuid.NewV4())
		err = global.GVA_DB.Create(&appUser).Error
		if err != nil {
			return
		}

	}

	fmt.Println("Login", err)
	//db := global.GVA_DB.Model(&st.AppUser{})
	return
}

// GetUserinfo 用户数据表单
// Author [yourname](https://github.com/yourname)
func (appUserService *AppUserService) GetUserinfo(id uint) (user st.AppUser, err error) {
	// 请在这里实现自己的业务逻辑
	err = global.GVA_DB.First(&user, id).Error
	return
}

// UpdateTheImage 更新头像
// Author [yourname](https://github.com/yourname)
func (appUserService *AppUserService) UpdateTheImage(id uint, image string) (user st.AppUser, err error) {
	// 请在这里实现自己的业务逻辑

	user.HeadPortrait = image
	err = global.GVA_DB.Where("id = ?", id).Updates(&user).Error

	return
}

// Upnickname 方法介绍
// Author [yourname](https://github.com/yourname)
func (appUserService *AppUserService) Upnickname(id uint, nickname string) (user st.AppUser, err error) {
	// 请在这里实现自己的业务逻辑
	user.Nickname = nickname
	err = global.GVA_DB.Where("id = ?", id).Updates(&user).Error
	return
}
