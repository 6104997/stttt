"use strict";
const common_vendor = require("../../common/vendor.js");
const api_request = require("../../api/request.js");
if (!Array) {
  const _easycom_custom_bav_bar2 = common_vendor.resolveComponent("custom-bav-bar");
  const _easycom_uni_icons2 = common_vendor.resolveComponent("uni-icons");
  const _easycom_common_title2 = common_vendor.resolveComponent("common-title");
  const _easycom_component_articleList2 = common_vendor.resolveComponent("component-articleList");
  const _easycom_z_paging2 = common_vendor.resolveComponent("z-paging");
  (_easycom_custom_bav_bar2 + _easycom_uni_icons2 + _easycom_common_title2 + _easycom_component_articleList2 + _easycom_z_paging2)();
}
const _easycom_custom_bav_bar = () => "../../components/custom-bav-bar/custom-bav-bar.js";
const _easycom_uni_icons = () => "../../uni_modules/uni-icons/components/uni-icons/uni-icons.js";
const _easycom_common_title = () => "../../components/common-title/common-title.js";
const _easycom_component_articleList = () => "../../components/component-articleList/component-articleList.js";
const _easycom_z_paging = () => "../../uni_modules/z-paging/components/z-paging/z-paging.js";
if (!Math) {
  (_easycom_custom_bav_bar + _easycom_uni_icons + _easycom_common_title + _easycom_component_articleList + _easycom_z_paging)();
}
const _sfc_main = {
  __name: "index",
  setup(__props) {
    const bannerList = common_vendor.ref([]);
    const ListOfAnnouncements = common_vendor.ref([]);
    const articleList = common_vendor.ref({});
    const paging = common_vendor.ref(null);
    const getBannerList = async () => {
      let res = await api_request.GetImage();
      if (res.code === 0) {
        bannerList.value = res.data;
      }
    };
    const queryList = (pageNo, pageSize) => {
      api_request.getAListOfArticles(pageNo, pageSize, "1").then((res) => {
        paging.value.complete(res.data.list);
      }).catch((res) => {
        paging.value.complete(false);
      });
    };
    const GetListOfAnnouncements = async () => {
      let res = await api_request.getAListOfAnnouncements();
      if (res.code === 0) {
        ListOfAnnouncements.value = res.data;
      }
    };
    const GetAListOfArticless = async () => {
      let res = await api_request.getAListOfArticles();
      articleList.value = res.data.list;
      console.log(articleList.value);
    };
    GetAListOfArticless();
    GetListOfAnnouncements();
    getBannerList();
    return (_ctx, _cache) => {
      return {
        a: common_vendor.p({
          title: "守塔不能停"
        }),
        b: common_vendor.f(bannerList.value, (item, index, i0) => {
          return {
            a: item.image,
            b: index
          };
        }),
        c: common_vendor.p({
          type: "calendar",
          size: "18"
        }),
        d: common_vendor.f(2, (item, k0, i0) => {
          return {
            a: "1cf27b2a-4-" + i0 + ",1cf27b2a-0",
            b: "1cf27b2a-5-" + i0 + ",1cf27b2a-0",
            c: "1cf27b2a-6-" + i0 + ",1cf27b2a-0"
          };
        }),
        e: common_vendor.p({
          type: "calendar",
          size: "20"
        }),
        f: common_vendor.p({
          type: "loop",
          size: "20"
        }),
        g: common_vendor.p({
          type: "eye",
          size: "20"
        }),
        h: common_vendor.f(articleList.value, (item, k0, i0) => {
          return {
            a: "1cf27b2a-8-" + i0 + ",1cf27b2a-0",
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
        i: common_vendor.sr(paging, "1cf27b2a-0", {
          "k": "paging"
        }),
        j: common_vendor.o(queryList),
        k: common_vendor.o(($event) => articleList.value = $event),
        l: common_vendor.p({
          ["default-page-size"]: 1,
          modelValue: articleList.value
        })
      };
    };
  }
};
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__scopeId", "data-v-1cf27b2a"]]);
wx.createPage(MiniProgramPage);
