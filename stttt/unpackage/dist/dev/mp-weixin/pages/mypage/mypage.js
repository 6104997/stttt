"use strict";
const common_vendor = require("../../common/vendor.js");
const api_request = require("../../api/request.js");
if (!Array) {
  const _easycom_uni_icons2 = common_vendor.resolveComponent("uni-icons");
  const _easycom_common_userlist2 = common_vendor.resolveComponent("common-userlist");
  (_easycom_uni_icons2 + _easycom_common_userlist2)();
}
const _easycom_uni_icons = () => "../../uni_modules/uni-icons/components/uni-icons/uni-icons.js";
const _easycom_common_userlist = () => "../../components/common-userlist/common-userlist.js";
if (!Math) {
  (_easycom_uni_icons + _easycom_common_userlist)();
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
    Login();
    common_vendor.onMounted(async () => {
      if (common_vendor.cookies.get("x-token")) {
        let resa = await api_request.GetUserinfo();
        console.log(resa);
      }
    });
    return (_ctx, _cache) => {
      return common_vendor.e({
        a: User.value
      }, User.value ? {
        b: User.value.headPortrait,
        c: common_vendor.t(User.value.nickname),
        d: common_vendor.t(User.value.ID)
      } : {}, {
        e: common_vendor.p({
          type: "contact-filled",
          size: "30",
          color: "#23b389"
        }),
        f: common_vendor.p({
          type: "right",
          size: "15",
          color: "#aaa"
        }),
        g: common_vendor.p({
          type: "weixin",
          size: "30",
          color: "#23b389"
        }),
        h: common_vendor.p({
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
