"use strict";const e=require("../../common/vendor.js"),t=require("../../utils/system.js");if(!Array){e.resolveComponent("uni-icons")()}Math;const r={__name:"custom-bav-bar",props:{title:{type:String,default:"标题"}},setup:r=>(s,a)=>({a:e.unref(t.getStatusBarHeight)()+"px",b:e.t(r.title),c:e.p({type:"search",color:"#888",size:"18"}),d:e.unref(t.getTitbarHeight)()+"px",e:e.unref(t.getNavBarheight)()+"px"})},s=e._export_sfc(r,[["__scopeId","data-v-643d39ae"]]);wx.createComponent(s);