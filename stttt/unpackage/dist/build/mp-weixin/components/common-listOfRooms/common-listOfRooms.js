"use strict";const t=require("../../common/vendor.js");if(!Array){t.resolveComponent("uni-icons")()}Math;const e={__name:"common-listOfRooms",props:{roomNumber:{type:String,default:"房间号"},roomPassword:{type:String,default:"房间密码"},currentStatus:{type:String,default:"状态"},CreatedAt:{type:String,default:"创建时间"}},setup:e=>(o,r)=>({a:t.p({type:"calendar",size:"20"}),b:t.t(e.roomNumber),c:t.t(e.roomPassword),d:t.p({type:"loop",size:"20"}),e:t.t(e.CreatedAt),f:t.p({type:"eye",size:"20"}),g:t.t(e.currentStatus),h:t.o((o=>{return r=e.roomNumber,void t.index.setClipboardData({data:r,success:function(){t.index.showToast({title:"复制成功",icon:"none"})},fail:function(e){t.index.showToast({title:"复制失败",icon:"none"})}});var r}))})},o=t._export_sfc(e,[["__scopeId","data-v-bf9861f6"]]);wx.createComponent(o);
