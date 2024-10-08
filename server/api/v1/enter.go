package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/goldenHouseManagement"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/st"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup                system.ApiGroup
	ExampleApiGroup               example.ApiGroup
	StApiGroup                    st.ApiGroup
	GoldenHouseManagementApiGroup goldenHouseManagement.ApiGroup
}
