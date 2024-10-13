"use strict";
const common_vendor = require("../../common/vendor.js");
require("../../utils/system.js");
const api_request = require("../../api/request.js");
if (!Array) {
  const _easycom_custom_bav_bar2 = common_vendor.resolveComponent("custom-bav-bar");
  const _easycom_common_title2 = common_vendor.resolveComponent("common-title");
  const _easycom_component_articleList2 = common_vendor.resolveComponent("component-articleList");
  const _easycom_z_paging2 = common_vendor.resolveComponent("z-paging");
  (_easycom_custom_bav_bar2 + _easycom_common_title2 + _easycom_component_articleList2 + _easycom_z_paging2)();
}
const _easycom_custom_bav_bar = () => "../../components/custom-bav-bar/custom-bav-bar.js";
const _easycom_common_title = () => "../../components/common-title/common-title.js";
const _easycom_component_articleList = () => "../../components/component-articleList/component-articleList.js";
const _easycom_z_paging = () => "../../uni_modules/z-paging/components/z-paging/z-paging.js";
if (!Math) {
  (_easycom_custom_bav_bar + _easycom_common_title + _easycom_component_articleList + _easycom_z_paging)();
}
const _sfc_main = {
  __name: "classification",
  setup(__props) {
    common_vendor.ref([]);
    common_vendor.ref(0);
    const ListOfAnnouncements = common_vendor.ref([]);
    const paging = common_vendor.ref(null);
    common_vendor.ref([]);
    const listOfArticles = common_vendor.ref([]);
    const listOfArticlesIndex = common_vendor.ref(0);
    common_vendor.ref("tap0");
    const customGetArticless = common_vendor.ref([]);
    const urll = common_vendor.ref();
    const GetListOfAnnouncements = async () => {
      let res = await api_request.getAListOfArticles("1", "1", "0");
      if (res.code === 0) {
        ListOfAnnouncements.value = res.data.list;
      }
    };
    const GetNavigationists = async () => {
      listOfArticlesIndex.value = 0;
      let res = await api_request.getNavigationist();
      listOfArticles.value = res.data;
      res = await api_request.getcustomGetArticles(
        "articleManagement/getAListOfArticles?theLatestArticles=1&articleClassificationId=1"
      );
      customGetArticless.value = res.data.list;
      urll.value = "articleManagement/getAListOfArticles?theLatestArticles=1&articleClassificationId=1";
      queryList(1, 5);
    };
    const queryList = async (pageNo, pageSize) => {
      await api_request.getcustomGetArticles(urll.value, pageNo, pageSize).then((res) => {
        console.log(urll.value, res);
        paging.value.complete(res.data.list);
      }).catch((res) => {
        paging.value.complete(false);
      });
    };
    const changeTab = async (index, navigationUrl) => {
      if (listOfArticlesIndex.value === index) {
        return;
      }
      listOfArticlesIndex.value = index;
      paging.value = null;
      customGetArticless.value = [];
      urll.value = navigationUrl;
      queryList(1, 5);
    };
    common_vendor.onShow(() => {
      GetNavigationists();
    });
    GetListOfAnnouncements();
    GetNavigationists();
    return (_ctx, _cache) => {
      return {
        a: common_vendor.p({
          title: "分类"
        }),
        b: common_vendor.f(ListOfAnnouncements.value, (item, k0, i0) => {
          return {
            a: "5fbab323-3-" + i0 + ",5fbab323-0",
            b: common_vendor.p({
              images: item.preview_the_image,
              title: item.title,
              author: item.author,
              CreatedAt: common_vendor.unref(api_request.convertTimeFormat)(item.created_at)
            }),
            c: "/pages/detail/detail?uuid=" + item.uuid,
            d: item.uuid
          };
        }),
        c: common_vendor.f(listOfArticles.value, (item, index, i0) => {
          return {
            a: common_vendor.t(item.navigationName),
            b: common_vendor.n(listOfArticlesIndex.value === index ? "f-active-color" : "f-color"),
            c: index,
            d: common_vendor.o(($event) => changeTab(index, item.navigationUrl), index)
          };
        }),
        d: common_vendor.f(customGetArticless.value, (item, k0, i0) => {
          return {
            a: "5fbab323-4-" + i0 + ",5fbab323-0",
            b: common_vendor.p({
              images: item.preview_the_image,
              title: item.title,
              author: item.author,
              CreatedAt: common_vendor.unref(api_request.convertTimeFormat)(item.created_at)
            }),
            c: "/pages/detail/detail?uuid=" + item.uuid,
            d: item.uuid
          };
        }),
        e: common_vendor.sr(paging, "5fbab323-0", {
          "k": "paging"
        }),
        f: common_vendor.o(queryList),
        g: common_vendor.o(($event) => customGetArticless.value = $event),
        h: common_vendor.p({
          ["default-page-size"]: 5,
          modelValue: customGetArticless.value
        })
      };
    };
  }
};
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__scopeId", "data-v-5fbab323"]]);
_sfc_main.__runtimeHooks = 2;
wx.createPage(MiniProgramPage);
