package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/goldenHouseManagement"
	"github.com/flipped-aurora/gin-vue-admin/server/model/st"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(st.AppUser{}, st.ArouselImage{}, st.AnnouncementManagement{}, st.ArticleManagement{}, goldenHouseManagement.GoldenRoomForm{}, goldenHouseManagement.UserGoldData{}, st.CategoricalNavigationManagement{})
	if err != nil {
		return err
	}
	return nil
}
