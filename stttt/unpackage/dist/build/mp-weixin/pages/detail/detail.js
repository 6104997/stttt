"use strict";const e=require("../../common/vendor.js"),t=require("../../api/request.js");if(!Array){(e.resolveComponent("uni-tag")+e.resolveComponent("mp-html"))()}Math||((()=>"../../uni_modules/uni-tag/components/uni-tag/uni-tag.js")+(()=>"../../uni_modules/mp-html/components/mp-html/mp-html.js"))();const a={__name:"detail",setup(a){const n=e.ref({});let o;return e.onLoad((e=>{o=e.uuid,(async()=>{const e=await t.getArticleManagementByUUID(o);n.value=e.data,n.value.CreatedAt=t.convertTimeFormat(e.data.CreatedAt)})()})),(t,a)=>e.e({a:n.value.select},n.value.select?{b:e.p({inverted:!0,text:"置顶",type:"error"})}:{},{c:e.t(n.value.title),d:e.t(n.value.author),e:e.t(n.value.CreatedAt),f:e.p({content:n.value.content}),g:e.t(n.value.view_count)})}},n=e._export_sfc(a,[["__scopeId","data-v-37218d45"]]);a.__runtimeHooks=2,wx.createPage(n);
