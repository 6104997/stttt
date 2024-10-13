"use strict";
const common_vendor = require("../../common/vendor.js");
const api_request = require("../../api/request.js");
if (!Array) {
  const _easycom_custom_bav_bar2 = common_vendor.resolveComponent("custom-bav-bar");
  const _easycom_uni_icons2 = common_vendor.resolveComponent("uni-icons");
  const _easycom_common_title2 = common_vendor.resolveComponent("common-title");
  const _easycom_common_listOfRooms2 = common_vendor.resolveComponent("common-listOfRooms");
  const _easycom_component_articleList2 = common_vendor.resolveComponent("component-articleList");
  const _easycom_z_paging2 = common_vendor.resolveComponent("z-paging");
  (_easycom_custom_bav_bar2 + _easycom_uni_icons2 + _easycom_common_title2 + _easycom_common_listOfRooms2 + _easycom_component_articleList2 + _easycom_z_paging2)();
}
const _easycom_custom_bav_bar = () => "../../components/custom-bav-bar/custom-bav-bar.js";
const _easycom_uni_icons = () => "../../uni_modules/uni-icons/components/uni-icons/uni-icons.js";
const _easycom_common_title = () => "../../components/common-title/common-title.js";
const _easycom_common_listOfRooms = () => "../../components/common-listOfRooms/common-listOfRooms.js";
const _easycom_component_articleList = () => "../../components/component-articleList/component-articleList.js";
const _easycom_z_paging = () => "../../uni_modules/z-paging/components/z-paging/z-paging.js";
if (!Math) {
  (_easycom_custom_bav_bar + _easycom_uni_icons + _easycom_common_title + _easycom_common_listOfRooms + _easycom_component_articleList + _easycom_z_paging)();
}
const _sfc_main = {
  __name: "index",
  setup(__props) {
    const bannerList = common_vendor.ref([]);
    const ListOfAnnouncements = common_vendor.ref([]);
    const articleList = common_vendor.ref([]);
    const paging = common_vendor.ref(null);
    const roomdata = common_vendor.ref(null);
    const roommdata = common_vendor.ref([]);
    const getBannerList = async () => {
      let res = await api_request.GetImage();
      if (res.code === 0) {
        bannerList.value = res.data;
      }
    };
    const queryList = async (pageNo, pageSize) => {
      await api_request.getAListOfArticles(pageNo, pageSize, "1").then((res) => {
        paging.value.complete(res.data.list);
      }).catch((res) => {
        paging.value.complete(false);
      });
    };
    const GetListOfAnnouncements = async () => {
      let res = await api_request.getAListOfArticles("1", "5", "0");
      if (res.code === 0) {
        ListOfAnnouncements.value = res.data.list;
      }
    };
    const AListOfRooms = async () => {
      let resa = await api_request.getAListOfRooms();
      roomdata.value = resa.data;
      roommdata.value = resa.data;
      console.log(resa);
      console.log(roomdata.value);
    };
    setInterval(function() {
      AListOfRooms();
    }, 5e3);
    AListOfRooms();
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
        d: common_vendor.f(roommdata.value, (item, k0, i0) => {
          return {
            a: "1cf27b2a-4-" + i0 + ",1cf27b2a-0",
            b: common_vendor.p({
              roomNumber: item.roomNumber,
              roomPassword: item.roomPassword,
              currentStatus: item.currentStatus,
              CreatedAt: common_vendor.unref(api_request.convertTimeFormat)(item.CreatedAt)
            }),
            c: item.ID
          };
        }),
        e: common_vendor.f(ListOfAnnouncements.value, (item, k0, i0) => {
          return {
            a: "1cf27b2a-6-" + i0 + ",1cf27b2a-0",
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
        f: common_vendor.f(articleList.value, (item, k0, i0) => {
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
        g: common_vendor.sr(paging, "1cf27b2a-0", {
          "k": "paging"
        }),
        h: common_vendor.o(queryList),
        i: common_vendor.o(($event) => articleList.value = $event),
        j: common_vendor.p({
          ["default-page-size"]: 5,
          modelValue: articleList.value
        })
      };
    };
  }
};
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__scopeId", "data-v-1cf27b2a"]]);
_sfc_main.__runtimeHooks = 2;
wx.createPage(MiniProgramPage);
