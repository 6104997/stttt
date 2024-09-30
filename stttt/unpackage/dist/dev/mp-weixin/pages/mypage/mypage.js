"use strict";
const common_vendor = require("../../common/vendor.js");
const utils_system = require("../../utils/system.js");
const api_request = require("../../api/request.js");
if (!Array) {
  const _easycom_uni_icons2 = common_vendor.resolveComponent("uni-icons");
  const _easycom_common_userlist2 = common_vendor.resolveComponent("common-userlist");
  const _easycom_common_userlist_button2 = common_vendor.resolveComponent("common-userlist-button");
  (_easycom_uni_icons2 + _easycom_common_userlist2 + _easycom_common_userlist_button2)();
}
const _easycom_uni_icons = () => "../../uni_modules/uni-icons/components/uni-icons/uni-icons.js";
const _easycom_common_userlist = () => "../../components/common-userlist/common-userlist.js";
const _easycom_common_userlist_button = () => "../../components/common-userlist-button/common-userlist-button.js";
if (!Math) {
  (_easycom_uni_icons + _easycom_common_userlist + _easycom_common_userlist_button)();
}
const _sfc_main = {
  __name: "mypage",
  setup(__props) {
    const User = common_vendor.ref();
    const Login = async () => {
      if (common_vendor.cookies.get("x-token")) {
        let resa = await api_request.GetUserinfo();
        User.value = resa.data;
      } else {
        common_vendor.wx$1.login({
          success: async (res) => {
            if (res.code != null) {
              let resa = await api_request.GetLogin(res.code);
              User.value = resa.data.user;
            }
          }
        });
      }
    };
    const clickJump = () => {
      common_vendor.index.navigateTo({
        url: "/pages/personalCenter/personalCenter"
      });
    };
    Login();
    common_vendor.onShow(async () => {
      if (common_vendor.cookies.get("x-token")) {
        let resa = await api_request.GetUserinfo();
        User.value = resa.data;
        console.log(resa.data);
      }
    });
    common_vendor.onMounted(async () => {
      if (common_vendor.cookies.get("x-token")) {
        let resa = await api_request.GetUserinfo();
        console.log(resa);
      }
    });
    return (_ctx, _cache) => {
      return common_vendor.e({
        a: common_vendor.unref(utils_system.getNavBarheight)() + "px",
        b: User.value
      }, User.value ? {
        c: User.value.headPortrait,
        d: common_vendor.o(clickJump),
        e: common_vendor.t(User.value.nickname),
        f: common_vendor.o(clickJump),
        g: common_vendor.t(User.value.ID)
      } : {}, {
        h: common_vendor.p({
          type: "contact-filled",
          size: "30",
          color: "#23b389"
        }),
        i: common_vendor.p({
          type: "right",
          size: "15",
          color: "#aaa"
        }),
        j: common_vendor.p({
          type: "weixin",
          size: "30",
          color: "#23b389"
        }),
        k: common_vendor.p({
          type: "right",
          size: "15",
          color: "#aaa"
        })
      });
    };
  }
};
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__scopeId", "data-v-449a23d9"]]);
wx.createPage(MiniProgramPage);
