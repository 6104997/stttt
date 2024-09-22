package goldenHouseManagement

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	GoldenRoomFormApi
	UserGoldDataApi
}

var (
	goldenRoomFormService = service.ServiceGroupApp.GoldenHouseManagementServiceGroup.GoldenRoomFormService
	userGoldDataService   = service.ServiceGroupApp.GoldenHouseManagementServiceGroup.UserGoldDataService
)
