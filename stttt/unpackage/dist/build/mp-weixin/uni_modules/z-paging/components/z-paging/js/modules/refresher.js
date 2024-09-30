"use strict";const e=require("../../../../../../common/vendor.js"),r=require("../z-paging-utils.js"),t=require("../z-paging-enum.js"),s={props:{refresherThemeStyle:{type:String,default:r.u.gc("refresherThemeStyle","")},refresherImgStyle:{type:Object,default:r.u.gc("refresherImgStyle",{})},refresherTitleStyle:{type:Object,default:r.u.gc("refresherTitleStyle",{})},refresherUpdateTimeStyle:{type:Object,default:r.u.gc("refresherUpdateTimeStyle",{})},watchRefresherTouchmove:{type:Boolean,default:r.u.gc("watchRefresherTouchmove",!1)},loadingMoreThemeStyle:{type:String,default:r.u.gc("loadingMoreThemeStyle","")},refresherOnly:{type:Boolean,default:r.u.gc("refresherOnly",!1)},refresherDefaultDuration:{type:[Number,String],default:r.u.gc("refresherDefaultDuration",100)},refresherCompleteDelay:{type:[Number,String],default:r.u.gc("refresherCompleteDelay",0)},refresherCompleteDuration:{type:[Number,String],default:r.u.gc("refresherCompleteDuration",300)},refresherRefreshingScrollable:{type:Boolean,default:r.u.gc("refresherRefreshingScrollable",!0)},refresherCompleteScrollable:{type:Boolean,default:r.u.gc("refresherCompleteScrollable",!1)},useCustomRefresher:{type:Boolean,default:r.u.gc("useCustomRefresher",!0)},refresherFps:{type:[Number,String],default:r.u.gc("refresherFps",40)},refresherMaxAngle:{type:[Number,String],default:r.u.gc("refresherMaxAngle",40)},refresherAngleEnableChangeContinued:{type:Boolean,default:r.u.gc("refresherAngleEnableChangeContinued",!1)},refresherDefaultText:{type:[String,Object],default:r.u.gc("refresherDefaultText",null)},refresherPullingText:{type:[String,Object],default:r.u.gc("refresherPullingText",null)},refresherRefreshingText:{type:[String,Object],default:r.u.gc("refresherRefreshingText",null)},refresherCompleteText:{type:[String,Object],default:r.u.gc("refresherCompleteText",null)},refresherGoF2Text:{type:[String,Object],default:r.u.gc("refresherGoF2Text",null)},refresherDefaultImg:{type:String,default:r.u.gc("refresherDefaultImg",null)},refresherPullingImg:{type:String,default:r.u.gc("refresherPullingImg",null)},refresherRefreshingImg:{type:String,default:r.u.gc("refresherRefreshingImg",null)},refresherCompleteImg:{type:String,default:r.u.gc("refresherCompleteImg",null)},refresherRefreshingAnimated:{type:Boolean,default:r.u.gc("refresherRefreshingAnimated",!0)},refresherEndBounceEnabled:{type:Boolean,default:r.u.gc("refresherEndBounceEnabled",!0)},refresherEnabled:{type:Boolean,default:r.u.gc("refresherEnabled",!0)},refresherThreshold:{type:[Number,String],default:r.u.gc("refresherThreshold","80rpx")},refresherDefaultStyle:{type:String,default:r.u.gc("refresherDefaultStyle","black")},refresherBackground:{type:String,default:r.u.gc("refresherBackground","transparent")},refresherFixedBackground:{type:String,default:r.u.gc("refresherFixedBackground","transparent")},refresherFixedBacHeight:{type:[Number,String],default:r.u.gc("refresherFixedBacHeight",0)},refresherOutRate:{type:Number,default:r.u.gc("refresherOutRate",.65)},refresherF2Enabled:{type:Boolean,default:r.u.gc("refresherF2Enabled",!1)},refresherF2Threshold:{type:[Number,String],default:r.u.gc("refresherF2Threshold","200rpx")},refresherF2Duration:{type:[Number,String],default:r.u.gc("refresherF2Duration",200)},showRefresherF2:{type:Boolean,default:r.u.gc("showRefresherF2",!0)},refresherPullRate:{type:Number,default:r.u.gc("refresherPullRate",.75)},showRefresherUpdateTime:{type:Boolean,default:r.u.gc("showRefresherUpdateTime",!1)},refresherUpdateTimeKey:{type:String,default:r.u.gc("refresherUpdateTimeKey","default")},refresherVibrate:{type:Boolean,default:r.u.gc("refresherVibrate",!1)},refresherNoTransform:{type:Boolean,default:r.u.gc("refresherNoTransform",!1)},useRefresherStatusBarPlaceholder:{type:Boolean,default:r.u.gc("useRefresherStatusBarPlaceholder",!1)}},data:()=>({R:t.Enum.Refresher,refresherStatus:t.Enum.Refresher.Default,refresherTouchstartY:0,lastRefresherTouchmove:null,refresherReachMaxAngle:!0,refresherTransform:"translateY(0px)",refresherTransition:"",finalRefresherDefaultStyle:"black",refresherRevealStackCount:0,refresherCompleteTimeout:null,refresherCompleteSubTimeout:null,refresherEndTimeout:null,isTouchmovingTimeout:null,refresherTriggered:!1,isTouchmoving:!1,isTouchEnded:!1,isUserPullDown:!1,privateRefresherEnabled:-1,privateShowRefresherWhenReload:!1,customRefresherHeight:-1,showCustomRefresher:!1,doRefreshAnimateAfter:!1,isRefresherInComplete:!1,showF2:!1,f2Transform:"",pullDownTimeStamp:0,moveDis:0,oldMoveDis:0,currentDis:0,oldCurrentMoveDis:0,oldRefresherTouchmoveY:0,oldTouchDirection:"",oldEmitedTouchDirection:"",oldPullingDistance:-1,refresherThresholdUpdateTag:0}),watch:{refresherDefaultStyle:{handler(e){e.length&&(this.finalRefresherDefaultStyle=e)},immediate:!0},refresherStatus(e){e===t.Enum.Refresher.Loading&&this._cleanRefresherEndTimeout(),this.refresherVibrate&&(e===t.Enum.Refresher.ReleaseToRefresh||e===t.Enum.Refresher.GoF2)&&this._doVibrateShort(),this.$emit("refresherStatusChange",e),this.$emit("update:refresherStatus",e)},refresherEnabled(e){!e&&this.endRefresh()}},computed:{pullDownDisTimeStamp(){return 1e3/this.refresherFps},refresherThresholdUnitConverted(){return r.u.addUnit(this.refresherThreshold,this.unit)},finalRefresherEnabled(){return!this.useChatRecordMode&&(-1===this.privateRefresherEnabled?this.refresherEnabled:1===this.privateRefresherEnabled)},finalRefresherThreshold(){let e=this.refresherThresholdUnitConverted,t=!1;return e===r.u.addUnit(80,this.unit)&&(t=!0,this.showRefresherUpdateTime&&(e=r.u.addUnit(120,this.unit))),t&&this.customRefresherHeight>0?this.customRefresherHeight+this.finalRefresherThresholdPlaceholder:r.u.convertToPx(e)+this.finalRefresherThresholdPlaceholder},finalRefresherF2Threshold(){return r.u.convertToPx(r.u.addUnit(this.refresherF2Threshold,this.unit))},finalRefresherThresholdPlaceholder(){return this.useRefresherStatusBarPlaceholder?this.statusBarHeight:0},finalRefresherFixedBacHeight(){return r.u.convertToPx(this.refresherFixedBacHeight)},finalRefresherThemeStyle(){return this.refresherThemeStyle.length?this.refresherThemeStyle:this.defaultThemeStyle},finalRefresherOutRate(){let e=this.refresherOutRate;return e=Math.max(0,e),e=Math.min(1,e),e},finalRefresherPullRate(){let e=this.refresherPullRate;return e=Math.max(0,e),e},finalRefresherTransform(){return this.refresherNoTransform||"translateY(0px)"===this.refresherTransform?"none":this.refresherTransform},finalShowRefresherWhenReload(){return this.showRefresherWhenReload||this.privateShowRefresherWhenReload},finalRefresherTriggered(){return!(!this.finalRefresherEnabled||this.useCustomRefresher)&&this.refresherTriggered},showRefresher(){const e=this.finalRefresherEnabled||this.useCustomRefresher&&!this.useChatRecordMode;return this.active&&-1===this.customRefresherHeight&&e&&this.updateCustomRefresherHeight(),e},hasTouchmove(){return this.watchRefresherTouchmove}},methods:{endRefresh(){this.totalData=this.realTotalData,this._refresherEnd(),this._endSystemLoadingAndRefresh(),this._handleScrollViewBounce({bounce:!0}),this.$nextTick((()=>{this.refresherTriggered=!1}))},updateCustomRefresherHeight(){r.u.delay((()=>this.$nextTick(this._updateCustomRefresherHeight)))},closeF2(){this._handleCloseF2()},_onRefresh(e=!1,r=!0){(!e||this.finalRefresherEnabled&&!this.useCustomRefresher)&&(this.$emit("onRefresh"),this.$emit("Refresh"),this.loading||this.isRefresherInComplete||(this.loadingType=t.Enum.LoadingType.Refresher,this.nShowRefresherReveal||(this.isUserPullDown=r,this.isUserReload=!r,this._startLoading(!0),this.refresherTriggered=!0,this.reloadWhenRefresh&&r&&(this.useChatRecordMode?this._onLoadingMore("click"):this._reload(!1,!1,r)))))},_onRestore(){this.refresherTriggered="restore",this.$emit("onRestore"),this.$emit("Restore")},_handleRefresherTouchstart(e){!this.loading&&this.isTouchEnded&&(this.isTouchmoving=!1),this.loadingType=t.Enum.LoadingType.Refresher,this.isTouchmovingTimeout&&clearTimeout(this.isTouchmovingTimeout),this.isTouchEnded=!1,this.refresherTransition="",this.refresherTouchstartY=e.touchY,this.$emit("refresherTouchstart",this.refresherTouchstartY),this.lastRefresherTouchmove=e,this._cleanRefresherCompleteTimeout(),this._cleanRefresherEndTimeout()},_handleRefresherTouchmove(e,r){this.refresherReachMaxAngle=!0,this.isTouchmovingTimeout&&clearTimeout(this.isTouchmovingTimeout),this.isTouchmoving=!0,this.isTouchEnded=!1,e>=this.finalRefresherThreshold?this.refresherStatus=this.refresherF2Enabled&&e>=this.finalRefresherF2Threshold?t.Enum.Refresher.GoF2:t.Enum.Refresher.ReleaseToRefresh:this.refresherStatus=t.Enum.Refresher.Default,this.moveDis=e},_handleRefresherTouchend(e){this.isTouchmovingTimeout&&clearTimeout(this.isTouchmovingTimeout),this.refresherReachMaxAngle=!0,this.isTouchEnded=!0;const s=this.finalRefresherThreshold;e>=s&&(this.refresherStatus===t.Enum.Refresher.ReleaseToRefresh||this.refresherStatus===t.Enum.Refresher.GoF2)?this.refresherStatus===t.Enum.Refresher.GoF2?(this._handleGoF2(),this._refresherEnd()):(r.u.delay((()=>{this._emitTouchmove({pullingDistance:s,dy:this.moveDis-s})}),.1),this.moveDis=s,this.refresherStatus=t.Enum.Refresher.Loading,this._doRefresherLoad()):(this._refresherEnd(),this.isTouchmovingTimeout=r.u.delay((()=>{this.isTouchmoving=!1}),this.refresherDefaultDuration)),this.scrollEnable=!0,this.$emit("refresherTouchend",e)},_handleListTouchstart(){this.useChatRecordMode&&this.autoHideKeyboardWhenChat&&(e.index.hideKeyboard(),this.$emit("hidedKeyboard"))},_handleScrollViewBounce({bounce:e}){this.usePageScroll||this.scrollToTopBounceEnabled||(this.wxsScrollTop<=5?(this.refresherTransition="",this.scrollEnable=e):e&&(this.scrollEnable=e))},_handleWxsPullingDownStatusChange(e){this.wxsOnPullingDown=e,e&&!this.useChatRecordMode&&(this.renderPropScrollTop=0)},_handleWxsPullingDown({moveDis:e,diffDis:r}){this._emitTouchmove({pullingDistance:e,dy:r})},_handleTouchDirectionChange({direction:e}){this.$emit("touchDirectionChange",e)},_handlePropUpdate(){this.wxsPropType=r.u.getTime().toString()},_refresherEnd(e=!0,s=!1,h=!1,i=!0){if(this.loadingType===t.Enum.LoadingType.Refresher){const e=s&&(h||this.showRefresherWhenReload)?this.refresherCompleteDelay:0,i=e>0?t.Enum.Refresher.Complete:t.Enum.Refresher.Default;if(this.finalShowRefresherWhenReload){const e=this.refresherRevealStackCount;if(this.refresherRevealStackCount--,e>1)return}this._cleanRefresherEndTimeout(),this.refresherEndTimeout=r.u.delay((()=>{this.refresherStatus=i}),this.refresherStatus!==t.Enum.Refresher.Default&&i===t.Enum.Refresher.Default?this.refresherCompleteDuration:0),e>0&&(this.isRefresherInComplete=!0),this._cleanRefresherCompleteTimeout(),this.refresherCompleteTimeout=r.u.delay((()=>{let e=1;const h=this.refresherEndBounceEnabled&&s?"cubic-bezier(0.19,1.64,0.42,0.72)":"linear";s&&(e=this.refresherEndBounceEnabled?this.refresherCompleteDuration/1e3:this.refresherCompleteDuration/3e3),this.refresherTransition=`transform ${s?e:this.refresherDefaultDuration/1e3}s ${h}`,this.wxsPropType=this.refresherTransition+"end"+r.u.getTime(),this.moveDis=0,i===t.Enum.Refresher.Complete&&(this.refresherCompleteSubTimeout&&(clearTimeout(this.refresherCompleteSubTimeout),this.refresherCompleteSubTimeout=null),this.refresherCompleteSubTimeout=r.u.delay((()=>{this.$nextTick((()=>{this.refresherStatus=t.Enum.Refresher.Default,this.isRefresherInComplete=!1}))}),800*e)),this._emitTouchmove({pullingDistance:0,dy:this.moveDis})}),e)}i&&(r.u.delay((()=>this.loading=!1),e?10:0),h&&this._onRestore())},_handleGoF2(){!this.showF2&&this.refresherF2Enabled&&(this.$emit("refresherF2Change","go"),this.showRefresherF2&&(this.f2Transform=`translateY(${-this.superContentHeight}px)`,this.showF2=!0,r.u.delay((()=>{this.f2Transform="translateY(0px)"}),100,"f2ShowDelay")))},_handleCloseF2(){this.showF2&&this.refresherF2Enabled&&(this.$emit("refresherF2Change","close"),this.showRefresherF2&&(this.f2Transform=`translateY(${-this.superContentHeight}px)`,r.u.delay((()=>{this.showF2=!1,this.nF2Opacity=0}),this.refresherF2Duration,"f2CloseDelay")))},_doRefresherRefreshAnimate(){this._cleanRefresherCompleteTimeout();!this.doRefreshAnimateAfter&&this.finalShowRefresherWhenReload&&-1===this.customRefresherHeight&&this.refresherThreshold===r.u.addUnit(80,this.unit)?this.doRefreshAnimateAfter=!0:(this.refresherRevealStackCount++,this.wxsPropType="begin"+r.u.getTime(),this.moveDis=this.finalRefresherThreshold,this.refresherStatus=t.Enum.Refresher.Loading,this.isTouchmoving=!0,this.isTouchmovingTimeout&&clearTimeout(this.isTouchmovingTimeout),this._doRefresherLoad(!1))},_doRefresherLoad(e=!0){this._onRefresh(!1,e),this.loading=!0},_updateCustomRefresherHeight(){this._getNodeClientRect(".zp-custom-refresher-slot-view").then((e=>{this.customRefresherHeight=e?e[0].height:0,this.showCustomRefresher=this.customRefresherHeight>0,this.doRefreshAnimateAfter&&(this.doRefreshAnimateAfter=!1,this._doRefresherRefreshAnimate())}))},_emitTouchmove(e){e.viewHeight=this.finalRefresherThreshold,e.rate=e.viewHeight>0?e.pullingDistance/e.viewHeight:0,this.hasTouchmove&&this.oldPullingDistance!==e.pullingDistance&&this.$emit("refresherTouchmove",e),this.oldPullingDistance=e.pullingDistance},_cleanRefresherCompleteTimeout(){this.refresherCompleteTimeout=this._cleanTimeout(this.refresherCompleteTimeout)},_cleanRefresherEndTimeout(){this.refresherEndTimeout=this._cleanTimeout(this.refresherEndTimeout)}}};exports.refresherModule=s;
