"use strict";
const common_vendor = require("../../common/vendor.js");
const api_request = require("../../api/request.js");
if (!Array) {
  const _easycom_uni_icons2 = common_vendor.resolveComponent("uni-icons");
  const _easycom_common_title2 = common_vendor.resolveComponent("common-title");
  (_easycom_uni_icons2 + _easycom_common_title2)();
}
const _easycom_uni_icons = () => "../../uni_modules/uni-icons/components/uni-icons/uni-icons.js";
const _easycom_common_title = () => "../../components/common-title/common-title.js";
if (!Math) {
  (_easycom_uni_icons + _easycom_common_title)();
}
const _sfc_main = {
  __name: "index",
  setup(__props) {
    const bannerList = common_vendor.ref([]);
    const ListOfAnnouncements = common_vendor.ref([]);
    const getBannerList = async () => {
      let res = await api_request.RequestApi("/arouselImage/GetImage");
      if (res.code === 0) {
        bannerList.value = res.data;
      }
    };
    const GetListOfAnnouncements = async () => {
      let res = await api_request.RequestApi("/announcementManagement/getAListOfAnnouncements");
      if (res.code === 0) {
        ListOfAnnouncements.value = res.data;
      }
    };
    GetListOfAnnouncements();
    getBannerList();
    return (_ctx, _cache) => {
      return {
        a: common_vendor.f(bannerList.value, (item, index, i0) => {
          return {
            a: item.image,
            b: index
          };
        }),
        b: common_vendor.p({
          type: "calendar",
          size: "18"
        }),
        c: common_vendor.f(2, (item, k0, i0) => {
          return {
            a: "1cf27b2a-2-" + i0,
            b: "1cf27b2a-3-" + i0,
            c: "1cf27b2a-4-" + i0
          };
        }),
        d: common_vendor.p({
          type: "calendar",
          size: "20"
        }),
        e: common_vendor.p({
          type: "loop",
          size: "20"
        }),
        f: common_vendor.p({
          type: "eye",
          size: "20"
        })
      };
    };
  }
};
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__scopeId", "data-v-1cf27b2a"]]);
wx.createPage(MiniProgramPage);
