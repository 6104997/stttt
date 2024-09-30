"use strict";const t=require("../z-paging-utils.js"),e=require("../z-paging-constant.js"),i=require("../z-paging-enum.js"),l={props:{useVirtualList:{type:Boolean,default:t.u.gc("useVirtualList",!1)},useCompatibilityMode:{type:Boolean,default:t.u.gc("useCompatibilityMode",!1)},extraData:{type:Object,default:t.u.gc("extraData",{})},useInnerList:{type:Boolean,default:t.u.gc("useInnerList",!1)},forceCloseInnerList:{type:Boolean,default:t.u.gc("forceCloseInnerList",!1)},cellKeyName:{type:String,default:t.u.gc("cellKeyName","")},innerListStyle:{type:Object,default:t.u.gc("innerListStyle",{})},innerCellStyle:{type:Object,default:t.u.gc("innerCellStyle",{})},preloadPage:{type:[Number,String],default:t.u.gc("preloadPage",12),validator:e=>(e<=0&&t.u.consoleErr("preload-page必须大于0！"),e>0)},cellHeightMode:{type:String,default:t.u.gc("cellHeightMode",i.Enum.CellHeightMode.Fixed)},fixedCellHeight:{type:[Number,String],default:t.u.gc("fixedCellHeight",0)},virtualListCol:{type:[Number,String],default:t.u.gc("virtualListCol",1)},virtualScrollFps:{type:[Number,String],default:t.u.gc("virtualScrollFps",80)}},data:()=>({virtualListKey:t.u.getInstanceId(),virtualPageHeight:0,virtualCellHeight:0,virtualScrollTimeStamp:0,virtualList:[],virtualPlaceholderTopHeight:0,virtualPlaceholderBottomHeight:0,virtualTopRangeIndex:0,virtualBottomRangeIndex:0,lastVirtualTopRangeIndex:0,lastVirtualBottomRangeIndex:0,virtualItemInsertedCount:0,virtualHeightCacheList:[],getCellHeightRetryCount:{fixed:0,dynamic:0},pagingOrgTop:-1,updateVirtualListFromDataChange:!1}),watch:{realTotalData(){this.updateVirtualListRender()},virtualList(t){this.$emit("update:virtualList",t),this.$emit("virtualListChange",t)},virtualPlaceholderTopHeight(t){this.$emit("virtualTopHeightChange",t)}},computed:{virtualCellIndexKey:()=>e.c.listCellIndexKey,finalUseVirtualList(){return this.useVirtualList&&this.usePageScroll&&t.u.consoleErr("使用页面滚动时，开启虚拟列表无效！"),this.useVirtualList&&!this.usePageScroll},finalUseInnerList(){return this.useInnerList||this.finalUseVirtualList&&!this.forceCloseInnerList},finalCellKeyName(){return this.cellKeyName},finalVirtualPageHeight(){return this.virtualPageHeight>0?this.virtualPageHeight:this.windowHeight},finalFixedCellHeight(){return t.u.convertToPx(this.fixedCellHeight)},virtualRangePageHeight(){return this.finalVirtualPageHeight*this.preloadPage},virtualScrollDisTimeStamp(){return 1e3/this.virtualScrollFps}},methods:{doInsertVirtualListItem(l,a){if(this.cellHeightMode!==i.Enum.CellHeightMode.Dynamic)return;this.virtualItemInsertedCount++,l&&"[object Object]"===Object.prototype.toString.call(l)||(l={item:l});const h=this.virtualCellIndexKey;l[h]=`custom-${this.virtualItemInsertedCount}`,l[e.c.listCellIndexUniqueKey]=`${this.virtualListKey}-${l[h]}`,this.$nextTick((async()=>{let i=0;for(;i<=10;){await t.u.wait(e.c.delayTime);const s=await this._getNodeClientRect(`#zp-id-${l[h]}`,this.finalUseInnerList);if(!s){i++;continue}const r=s?s[0].height:0,n=this.virtualHeightCacheList[a-1],o=n?n.totalHeight:0;this.virtualHeightCacheList.splice(a,0,{height:r,lastTotalHeight:o,totalHeight:o+r});for(let t=a+1;t<this.virtualHeightCacheList.length;t++){const e=this.virtualHeightCacheList[t];e.lastTotalHeight+=r,e.totalHeight+=r}this._updateVirtualScroll(this.oldScrollTop);break}}))},didUpdateVirtualListCell(t){if(this.cellHeightMode!==i.Enum.CellHeightMode.Dynamic)return;const e=this.virtualHeightCacheList[t];this.$nextTick((()=>{this._getNodeClientRect(`#zp-id-${t}`,this.finalUseInnerList).then((i=>{const l=i?i[0].height:0,a=l-e.height;e.height=l,e.totalHeight=e.lastTotalHeight+l;for(let e=t+1;e<this.virtualHeightCacheList.length;e++){const t=this.virtualHeightCacheList[e];t.totalHeight+=a,t.lastTotalHeight+=a}}))}))},didDeleteVirtualListCell(t){if(this.cellHeightMode!==i.Enum.CellHeightMode.Dynamic)return;const e=this.virtualHeightCacheList[t];for(let i=t+1;i<this.virtualHeightCacheList.length;i++){const t=this.virtualHeightCacheList[i];t.totalHeight-=e.height,t.lastTotalHeight-=e.height}this.virtualHeightCacheList.splice(t,1)},updateVirtualListRender(){this.finalUseVirtualList&&(this.updateVirtualListFromDataChange=!0,this.$nextTick((()=>{this.getCellHeightRetryCount.fixed=0,this.realTotalData.length?this.cellHeightMode===i.Enum.CellHeightMode.Fixed&&this.isFirstPage&&this._updateFixedCellHeight():this._resetDynamicListState(!this.isUserPullDown),this._updateVirtualScroll(this.oldScrollTop)})))},_virtualListInit(){this.$nextTick((()=>{t.u.delay((()=>{this._getNodeClientRect(".zp-scroll-view").then((t=>{t&&(this.pagingOrgTop=t[0].top,this.virtualPageHeight=t[0].height)}))}))}))},_updateFixedCellHeight(){this.finalFixedCellHeight?this.virtualCellHeight=this.finalFixedCellHeight:this.$nextTick((()=>{t.u.delay((()=>{this._getNodeClientRect("#zp-id-0",this.finalUseInnerList).then((t=>{if(t)this.virtualCellHeight=t[0].height,this._updateVirtualScroll(this.oldScrollTop);else{if(this.getCellHeightRetryCount.fixed>10)return;this.getCellHeightRetryCount.fixed++,this._updateFixedCellHeight()}}))}),e.c.delayTime,"updateFixedCellHeightDelay")}))},_updateDynamicCellHeight(i,l="bottom"){const a="top"===l,h=this.virtualHeightCacheList,s=a?[]:h;let r=0;this.$nextTick((()=>{t.u.delay((async()=>{for(let t=0;t<i.length;t++){const e=await this._getNodeClientRect(`#zp-id-${i[t][this.virtualCellIndexKey]}`,this.finalUseInnerList),n=e?e[0].height:0;if(!e)return void(this.getCellHeightRetryCount.dynamic<=10&&(h.splice(h.length-t,t),this.getCellHeightRetryCount.dynamic++,this._updateDynamicCellHeight(i,l)));const o=s.length?s.slice(-1)[0]:null,u=o?o.totalHeight:0;s.push({height:n,lastTotalHeight:u,totalHeight:u+n}),a&&(r+=n)}if(a&&i.length){for(let t=0;t<h.length;t++){const e=h[t];e.lastTotalHeight+=r,e.totalHeight+=r}this.virtualHeightCacheList=s.concat(h)}this._updateVirtualScroll(this.oldScrollTop)}),e.c.delayTime,"updateDynamicCellHeightDelay")}))},_setCellIndex(l,a="bottom"){let h=0;const s=this.virtualCellIndexKey;if([i.Enum.QueryFrom.Refresh,i.Enum.QueryFrom.Reload].indexOf(this.queryFrom)>=0&&this._resetDynamicListState(),this.totalData.length){if("bottom"===a){h=this.realTotalData.length;const t=this.realTotalData.length?this.realTotalData.slice(-1)[0]:null;t&&void 0!==t[s]&&(h=t[s]+1)}else if("top"===a){const t=this.realTotalData.length?this.realTotalData[0]:null;t&&void 0!==t[s]&&(h=t[s]-l.length)}}else this._resetDynamicListState();for(let i=0;i<l.length;i++){let a=l[i];a&&"[object Object]"===Object.prototype.toString.call(a)||(a={item:a}),a[e.c.listCellIndexUniqueKey]&&(a=t.u.deepCopy(a)),a[s]=h+i,a[e.c.listCellIndexUniqueKey]=`${this.virtualListKey}-${a[s]}`,l[i]=a}this.getCellHeightRetryCount.dynamic=0,this.cellHeightMode===i.Enum.CellHeightMode.Dynamic&&this._updateDynamicCellHeight(l,a)},_updateVirtualScroll(e,l=0){const a=t.u.getTime();if(0===e&&this._resetTopRange(),0!==e&&this.virtualScrollTimeStamp&&a-this.virtualScrollTimeStamp<=this.virtualScrollDisTimeStamp)return;this.virtualScrollTimeStamp=a;let h=0;const s=this.cellHeightMode;if(s===i.Enum.CellHeightMode.Fixed)h=parseInt(e/this.virtualCellHeight)||0,this._updateFixedTopRangeIndex(h),this._updateFixedBottomRangeIndex(h);else if(s===i.Enum.CellHeightMode.Dynamic){const t=l>0?"top":"bottom",i=this.virtualRangePageHeight,a=e-i,h=e+this.finalVirtualPageHeight+i;let s=0,r=0,n=!1;const o=this.virtualHeightCacheList,u=o?o.slice(-1)[0]:null;let g=this.virtualTopRangeIndex;if("bottom"===t)for(let e=g;e<o.length;e++){const t=o[e];if(t&&t.totalHeight>a){this.virtualTopRangeIndex=e,this.virtualPlaceholderTopHeight=t.lastTotalHeight;break}}else{let t=!1;for(let e=g;e>=0;e--){const i=o[e];if(i&&i.totalHeight<a){this.virtualTopRangeIndex=e,this.virtualPlaceholderTopHeight=i.lastTotalHeight,t=!0;break}}!t&&this._resetTopRange()}for(let e=this.virtualTopRangeIndex;e<o.length;e++){const t=o[e];if(t&&t.totalHeight>h){s=e,r=u.totalHeight-t.totalHeight,n=!0;break}}n&&0!==this.virtualBottomRangeIndex?(this.virtualBottomRangeIndex=s,this.virtualPlaceholderBottomHeight=r):(this.virtualBottomRangeIndex=this.realTotalData.length?this.realTotalData.length-1:this.pageSize,this.virtualPlaceholderBottomHeight=0),this._updateVirtualList()}},_updateFixedTopRangeIndex(t){let e=0===this.virtualCellHeight?0:t-(parseInt(this.finalVirtualPageHeight/this.virtualCellHeight)||1)*this.preloadPage;e*=this.virtualListCol,e=Math.max(0,e),this.virtualTopRangeIndex=e,this.virtualPlaceholderTopHeight=e/this.virtualListCol*this.virtualCellHeight},_updateFixedBottomRangeIndex(t){let e=0===this.virtualCellHeight?this.pageSize:t+(parseInt(this.finalVirtualPageHeight/this.virtualCellHeight)||1)*(this.preloadPage+1);e*=this.virtualListCol,e=Math.min(this.realTotalData.length,e),this.virtualBottomRangeIndex=e,this.virtualPlaceholderBottomHeight=(this.realTotalData.length-e)*this.virtualCellHeight/this.virtualListCol,this._updateVirtualList()},_updateVirtualList(){(this.updateVirtualListFromDataChange||this.lastVirtualTopRangeIndex!==this.virtualTopRangeIndex||this.lastVirtualBottomRangeIndex!==this.virtualBottomRangeIndex)&&(this.updateVirtualListFromDataChange=!1,this.lastVirtualTopRangeIndex=this.virtualTopRangeIndex,this.lastVirtualBottomRangeIndex=this.virtualBottomRangeIndex,this.virtualList=this.realTotalData.slice(this.virtualTopRangeIndex,this.virtualBottomRangeIndex+1))},_resetDynamicListState(t=!1){this.virtualHeightCacheList=[],t&&(this.virtualList=[]),this.virtualTopRangeIndex=0,this.virtualPlaceholderTopHeight=0},_resetTopRange(){this.virtualTopRangeIndex=0,this.virtualPlaceholderTopHeight=0,this._updateVirtualList()},_checkVirtualListScroll(){this.finalUseVirtualList&&this.$nextTick((()=>{this._getNodeClientRect(".zp-paging-touch-view").then((t=>{const e=t?t[0].top:0;(!t||e===this.pagingOrgTop&&0!==this.virtualPlaceholderTopHeight)&&this._updateVirtualScroll(0)}))}))},_innerCellClick(t,e){this.$emit("innerCellClick",t,e)}}};exports.virtualListModule=l;
