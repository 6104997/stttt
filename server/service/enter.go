package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/goldenHouseManagement"
	"github.com/flipped-aurora/gin-vue-admin/server/service/st"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup                system.ServiceGroup
	ExampleServiceGroup               example.ServiceGroup
	StServiceGroup                    st.ServiceGroup
	GoldenHouseManagementServiceGroup goldenHouseManagement.ServiceGroup
}
