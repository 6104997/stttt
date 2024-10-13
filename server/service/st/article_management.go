package st

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/st"
	stReq "github.com/flipped-aurora/gin-vue-admin/server/model/st/request"
)

type ArticleManagementService struct{}

// CreateArticleManagement 创建文章记录
// Author [yourname](https://github.com/yourname)
func (articleManagementService *ArticleManagementService) CreateArticleManagement(articleManagement *st.ArticleManagement) (err error) {
	err = global.GVA_DB.Create(articleManagement).Error
	return err
}

// DeleteArticleManagement 删除文章记录
// Author [yourname](https://github.com/yourname)
func (articleManagementService *ArticleManagementService) DeleteArticleManagement(ID string) (err error) {
	err = global.GVA_DB.Delete(&st.ArticleManagement{}, "id = ?", ID).Error
	return err
}

// DeleteArticleManagementByIds 批量删除文章记录
// Author [yourname](https://github.com/yourname)
func (articleManagementService *ArticleManagementService) DeleteArticleManagementByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]st.ArticleManagement{}, "id in ?", IDs).Error
	return err
}

// UpdateArticleManagement 更新文章记录
// Author [yourname](https://github.com/yourname)
func (articleManagementService *ArticleManagementService) UpdateArticleManagement(articleManagement st.ArticleManagement) (err error) {
	err = global.GVA_DB.Model(&st.ArticleManagement{}).Where("id = ?", articleManagement.ID).Updates(&articleManagement).Error
	return err
}

// GetArticleManagement 根据ID获取文章记录
// Author [yourname](https://github.com/yourname)
func (articleManagementService *ArticleManagementService) GetArticleManagement(ID string) (articleManagement st.ArticleManagement, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&articleManagement).Error
	return
}

// GetArticleManagementInfoList 分页获取文章记录
// Author [yourname](https://github.com/yourname)
func (articleManagementService *ArticleManagementService) GetArticleManagementInfoList(info stReq.ArticleManagementSearch) (list []st.ArticleManagement, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&st.ArticleManagement{})
	var articleManagements []st.ArticleManagement
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}
	if info.Article_status != "" {
		db = db.Where("article_status = ?", info.Article_status)
	}
	if info.AccountType != "" {
		db = db.Where("account_type = ?", info.AccountType)
	}
	if info.ArticleClassificationId != "" {
		db = db.Where("article_classification_id = ?", info.ArticleClassificationId)
	}
	if info.Author != "" {
		db = db.Where("author LIKE ?", "%"+info.Author+"%")
	}
	if info.View_count != nil {
		db = db.Where("view_count <> ?", info.View_count)
	}
	if info.Select != nil {
		db = db.Where("select = ?", info.Select)
	}
	if info.AccountNumber != nil {
		db = db.Where("account_number = ?", info.AccountNumber)
	}
	if info.Contact != "" {
		db = db.Where("contact = ?", info.Contact)
	}
	if info.Price != nil {
		db = db.Where("price <> ?", info.Price)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&articleManagements).Error
	return articleManagements, total, err
}

//getAListOfArticles

// GetAListOfArticles 获取文章列表
// Author [yourname](https://github.com/yourname)
func (articleManagementService *ArticleManagementService) GetAListOfArticles(info stReq.ArticleManagementSearch, articleClassificationId string) (list []st.ArticleResponse, total int64, err error) {
	// 请在这里实现自己的业务逻辑
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&st.ArticleManagement{})

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.TheLatestArticles != "" {
		db = db.Order("created_at DESC")
	}
	if info.AccountType != "" {
		db = db.Where("account_type = ?", info.AccountType)

	}
	db = db.Order("created_at DESC")
	db = db.Debug().Where("article_status = ? AND article_classification_id = ?", "1", articleClassificationId)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	//navigationName
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	//TheLatestArticles
	//accountType

	err = db.Select("created_at, uuid, title, preview_the_image, author").Find(&list).Error
	return list, total, err
}

//convertTimeFormat
//根据UUID获取文章记录

// GetArticleManagementByUUID 根据UUID获取文章记录
// Author [yourname](https://github.com/yourname)
func (articleManagementService *ArticleManagementService) GetArticleManagementByUUID(uuid string) (articleManagement st.ArticleManagement, err error) {
	err = global.GVA_DB.Where("uuid = ?", uuid).First(&articleManagement).Error
	if err != nil {
		return
	}

	// 如果阅读次数的指针为 nil，则将其初始化为 0
	if articleManagement.View_count == nil {
		count := 0
		articleManagement.View_count = &count
	}

	// 更新阅读次数
	err = global.GVA_DB.Model(&articleManagement).Update("view_count", *articleManagement.View_count+1).Error
	return
}

//getTheAnnouncementArticle
