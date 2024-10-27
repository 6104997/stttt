"use strict";const e=require("./js/z-paging-main.js"),o=require("../../../../common/vendor.js"),r=e=>{e.wxsCallMethods||(e.wxsCallMethods=[]),e.wxsCallMethods.push("_handleListTouchstart","_handleRefresherTouchstart","_handleTouchDirectionChange","_handleScrollViewBounce","_handleWxsPullingDown","_handleRefresherTouchmove","_handleRefresherTouchend","_handlePropUpdate","_handleWxsPullingDownStatusChange")},a={};if(!Array){(o.resolveComponent("z-paging-refresh")+o.resolveComponent("z-paging-load-more")+o.resolveComponent("z-paging-empty-view"))()}Math,r(e._sfc_main),"function"==typeof a&&a(e._sfc_main);const t=o._export_sfc(e._sfc_main,[["render",function(e,r,a,t,l,s){return o.e({a:-1===e.cssSafeAreaInsetBottom},(e.cssSafeAreaInsetBottom,{}),{b:e.showF2&&e.showRefresherF2},e.showF2&&e.showRefresherF2?{c:o.o((()=>{})),d:o.s({transform:e.f2Transform,transition:"transform .2s linear",height:e.superContentHeight+"px","z-index":e.f2ZIndex})}:{},{e:!e.usePageScroll&&e.zSlots.top},!e.usePageScroll&&e.zSlots.top?{}:e.usePageScroll&&e.zSlots.top?{g:o.o((()=>{})),h:o.s({top:`${e.windowTop}px`,"z-index":e.topZIndex})}:{},{f:e.usePageScroll&&e.zSlots.top,i:e.zSlots.left},e.zSlots.left?{j:e.finalIsOldWebView?1:""}:{},{k:e.finalRefresherFixedBacHeight>0},e.finalRefresherFixedBacHeight>0?{l:o.s({background:e.refresherFixedBackground,height:`${e.finalRefresherFixedBacHeight}px`})}:{},{m:e.showRefresher},e.showRefresher?o.e({n:e.useRefresherStatusBarPlaceholder},e.useRefresherStatusBarPlaceholder?{o:o.s({height:`${e.statusBarHeight}px`})}:{},{p:!(e.zSlots.refresherComplete&&e.refresherStatus===e.R.Complete||e.zSlots.refresherF2&&e.refresherStatus===e.R.GoF2)},e.zSlots.refresherComplete&&e.refresherStatus===e.R.Complete||e.zSlots.refresherF2&&e.refresherStatus===e.R.GoF2?{}:{q:o.r("refresher",{refresherStatus:e.refresherStatus})},{r:e.zSlots.refresherComplete&&e.refresherStatus===e.R.Complete},e.zSlots.refresherComplete&&e.refresherStatus===e.R.Complete||e.zSlots.refresherF2&&e.refresherStatus===e.R.GoF2||e.showCustomRefresher?{}:{v:o.sr("refresh","d64cc1e0-0"),w:o.s({height:e.finalRefresherThreshold-e.finalRefresherThresholdPlaceholder+"px"}),x:o.p({status:e.refresherStatus,defaultThemeStyle:e.finalRefresherThemeStyle,defaultText:e.finalRefresherDefaultText,pullingText:e.finalRefresherPullingText,refreshingText:e.finalRefresherRefreshingText,completeText:e.finalRefresherCompleteText,goF2Text:e.finalRefresherGoF2Text,defaultImg:e.refresherDefaultImg,pullingImg:e.refresherPullingImg,refreshingImg:e.refresherRefreshingImg,completeImg:e.refresherCompleteImg,refreshingAnimated:e.refresherRefreshingAnimated,showUpdateTime:e.showRefresherUpdateTime,updateTimeKey:e.refresherUpdateTimeKey,updateTimeTextMap:e.finalRefresherUpdateTimeTextMap,imgStyle:e.refresherImgStyle,titleStyle:e.refresherTitleStyle,updateTimeStyle:e.refresherUpdateTimeStyle,unit:e.unit})},{s:e.zSlots.refresherF2&&e.refresherStatus===e.R.GoF2,t:!e.showCustomRefresher,y:o.s({height:`${e.finalRefresherThreshold}px`,background:e.refresherBackground}),z:o.s({"margin-top":`-${e.finalRefresherThreshold+e.refresherThresholdUpdateTag}px`,background:e.refresherBackground,opacity:e.isTouchmoving?1:0})}):{},{A:e.showLoading&&e.zSlots.loading&&!e.loadingFullFixed},(e.showLoading&&e.zSlots.loading&&e.loadingFullFixed,{}),{B:e.finalUseInnerList},e.finalUseInnerList?o.e({C:e.finalUseVirtualList},e.finalUseVirtualList?{D:o.f(e.virtualList,((r,a,t)=>o.e(e.useCompatibilityMode?{}:{a:"cell-"+t,b:o.r("cell",{item:r,index:e.virtualTopRangeIndex+a},t)},{c:`zp-id-${r[e.virtualCellIndexKey]}`,d:r.zp_unique_index,e:o.o((o=>e._innerCellClick(r,e.virtualTopRangeIndex+a)),r.zp_unique_index)}))),E:e.useCompatibilityMode,F:o.s(e.innerCellStyle)}:{G:o.f(e.realTotalData,((r,a,t)=>({a:"cell-"+t,b:o.r("cell",{item:r,index:a},t),c:a,d:o.o((o=>e._innerCellClick(r,a)),a)})))},{H:o.s(e.innerListStyle)}):{},{I:e.useChatRecordMode&&e.realTotalData.length>=e.defaultPageSize&&(e.loadingStatus!==e.M.NoMore||e.zSlots.chatNoMore)&&(e.realTotalData.length||e.showChatLoadingWhenReload&&e.showLoading)&&!e.isFirstPageAndNoMore},e.useChatRecordMode&&e.realTotalData.length>=e.defaultPageSize&&(e.loadingStatus!==e.M.NoMore||e.zSlots.chatNoMore)&&(e.realTotalData.length||e.showChatLoadingWhenReload&&e.showLoading)&&!e.isFirstPageAndNoMore?o.e({J:e.loadingStatus===e.M.NoMore&&e.zSlots.chatNoMore},e.loadingStatus===e.M.NoMore&&e.zSlots.chatNoMore?{}:o.e({K:e.zSlots.chatLoading},e.zSlots.chatLoading?{L:o.r("chatLoading",{loadingMoreStatus:e.loadingStatus})}:{M:o.o((o=>e._onLoadingMore("click"))),N:o.p({zConfig:e.zLoadMoreConfig})}),{O:o.s(e.chatRecordRotateStyle)}):{},{P:e.useVirtualList},e.useVirtualList?{Q:o.s({height:e.virtualPlaceholderBottomHeight+"px"})}:{},{R:e.showLoadingMoreDefault},e.showLoadingMoreDefault||e.showLoadingMoreLoading||e.showLoadingMoreNoMore||e.showLoadingMoreFail?{}:e.showLoadingMoreCustom?{W:o.o((o=>e._onLoadingMore("click"))),X:o.p({zConfig:e.zLoadMoreConfig})}:{},{S:e.showLoadingMoreLoading,T:e.showLoadingMoreNoMore,U:e.showLoadingMoreFail,V:e.showLoadingMoreCustom,Y:e.safeAreaInsetBottom&&e.useSafeAreaPlaceholder&&!e.useChatRecordMode},e.safeAreaInsetBottom&&e.useSafeAreaPlaceholder&&!e.useChatRecordMode?{Z:o.s({height:e.safeAreaBottom+"px"})}:{},{aa:o.s({transform:e.virtualPlaceholderTopHeight>0?`translateY(${e.virtualPlaceholderTopHeight}px)`:"none"}),ab:o.s(e.finalPagingContentStyle),ac:e.showEmpty},e.showEmpty?o.e({ad:e.zSlots.empty},e.zSlots.empty?{ae:o.r("empty",{isLoadFailed:e.isLoadFailed})}:{af:o.o(e._emptyViewReload),ag:o.o(e._emptyViewClick),ah:o.p({emptyViewImg:e.finalEmptyViewImg,emptyViewText:e.finalEmptyViewText,showEmptyViewReload:e.finalShowEmptyViewReload,emptyViewReloadText:e.finalEmptyViewReloadText,isLoadFailed:e.isLoadFailed,emptyViewStyle:e.emptyViewStyle,emptyViewTitleStyle:e.emptyViewTitleStyle,emptyViewImgStyle:e.emptyViewImgStyle,emptyViewReloadStyle:e.emptyViewReloadStyle,emptyViewZIndex:e.emptyViewZIndex,emptyViewFixed:e.emptyViewFixed,unit:e.unit})},{ai:e.emptyViewCenter?1:"",aj:o.s(e.emptyViewSuperStyle),ak:o.s(e.chatRecordRotateStyle)}):{},{al:o.s({justifyContent:e.useChatRecordMode?"flex-end":"flex-start"}),am:o.s(e.scrollViewInStyle),an:o.s({transform:e.finalRefresherTransform,transition:e.refresherTransition}),ao:e.wxsPropType,ap:e.finalRefresherThreshold,aq:e.refresherF2Enabled,ar:e.finalRefresherF2Threshold,as:e.isIos,at:e.loading||e.isRefresherInComplete,av:e.useChatRecordMode,aw:e.refresherEnabled,ax:e.useCustomRefresher,ay:e.wxsPageScrollTop,az:e.wxsScrollTop,aA:e.refresherMaxAngle,aB:e.refresherNoTransform,aC:e.refresherAngleEnableChangeContinued,aD:e.usePageScroll,aE:e.watchTouchDirectionChange,aF:e.isTouchmoving,aG:e.finalRefresherOutRate,aH:e.finalRefresherPullRate,aI:e.hasTouchmove,aJ:e.usePageScroll?"":1,aK:e.showScrollbar?"":1,aL:o.s(e.chatRecordRotateStyle),aM:e.scrollTop,aN:e.scrollX,aO:e.finalScrollable,aP:e.finalEnableBackToTop,aQ:e.showScrollbar,aR:e.finalScrollWithAnimation,aS:e.scrollIntoView,aT:e.finalLowerThreshold,aU:e.finalRefresherEnabled&&!e.useCustomRefresher,aV:e.finalRefresherThreshold,aW:e.finalRefresherDefaultStyle,aX:e.refresherBackground,aY:e.finalRefresherTriggered,aZ:o.o(((...o)=>e._scroll&&e._scroll(...o))),ba:o.o(((...o)=>e._onScrollToLower&&e._onScrollToLower(...o))),bb:o.o(((...o)=>e._onScrollToUpper&&e._onScrollToUpper(...o))),bc:o.o(((...o)=>e._onRestore&&e._onRestore(...o))),bd:o.o((o=>e._onRefresh(!0))),be:e.finalIsOldWebView?1:"",bf:o.s(e.scrollViewContainerStyle),bg:e.zSlots.right},e.zSlots.right?{bh:e.finalIsOldWebView?1:""}:{},{bi:e.usePageScroll?"":1,bj:o.s(e.finalScrollViewStyle),bk:!e.usePageScroll&&e.zSlots.bottom},!e.usePageScroll&&e.zSlots.bottom?{}:e.usePageScroll&&e.zSlots.bottom?{bm:o.o((()=>{})),bn:o.s({bottom:`${e.windowBottom}px`})}:{},{bl:e.usePageScroll&&e.zSlots.bottom,bo:e.useChatRecordMode&&e.autoAdjustPositionWhenChat},e.useChatRecordMode&&e.autoAdjustPositionWhenChat?{bp:o.s({height:e.chatRecordModeSafeAreaBottom+"px"}),bq:o.s({height:e.keyboardHeight+"px"})}:{},{br:e.bottomBgColor,bs:e.showBackToTopClass},e.showBackToTopClass?o.e({bt:e.zSlots.backToTop},e.zSlots.backToTop?{}:{bv:e.useChatRecordMode&&!e.backToTopImg.length?1:"",bw:e.backToTopImg.length?e.backToTopImg:e.base64BackToTop},{bx:o.n(e.finalBackToTopClass),by:o.s(e.finalBackToTopStyle),bz:o.o(((...o)=>e._backToTopClick&&e._backToTopClick(...o)))}):{},{bA:e.showLoading&&e.zSlots.loading&&e.loadingFullFixed},(e.showLoading&&e.zSlots.loading&&e.loadingFullFixed,{}),{bB:e.usePageScroll?"":1,bC:!e.usePageScroll&&e.fixed?1:"",bD:e.usePageScroll?1:"",bE:e.renderPropScrollTop<1?1:"",bF:e.useChatRecordMode?1:"",bG:o.s(e.finalPagingStyle)})}],["__scopeId","data-v-d64cc1e0"]]);wx.createComponent(t);