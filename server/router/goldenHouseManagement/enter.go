package goldenHouseManagement

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	GoldenRoomFormRouter
	UserGoldDataRouter
}

var (
	goldenRoomFormApi = api.ApiGroupApp.GoldenHouseManagementApiGroup.GoldenRoomFormApi
	userGoldDataApi   = api.ApiGroupApp.GoldenHouseManagementApiGroup.UserGoldDataApi
)
