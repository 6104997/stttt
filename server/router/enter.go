package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/goldenHouseManagement"
	"github.com/flipped-aurora/gin-vue-admin/server/router/st"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System                system.RouterGroup
	Example               example.RouterGroup
	St                    st.RouterGroup
	GoldenHouseManagement goldenHouseManagement.RouterGroup
}
