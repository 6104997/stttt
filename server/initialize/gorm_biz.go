package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/st"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(st.AppUser{}, st.ArouselImage{})
	if err != nil {
		return err
	}
	return nil
}
