"use strict";
Object.defineProperty(exports, Symbol.toStringTag, { value: "Module" });
const common_vendor = require("./common/vendor.js");
const api_request = require("./api/request.js");
if (!Math) {
  "./pages/index/index.js";
  "./pages/mypage/mypage.js";
  "./pages/personalCenter/personalCenter.js";
  "./pages/classification/classification.js";
  "./pages/detail/detail.js";
  "./pages/user-profile/user-profile.js";
}
const _sfc_main = {
  onLaunch: function() {
    const Openpid = common_vendor.wx$1.getStorageSync("Openpid");
    if (Openpid === "") {
      common_vendor.wx$1.login({
        success: async (res) => {
          await api_request.GetLogin(res.code);
        }
      });
    }
  },
  onShow: function() {
    console.log("App Show");
  },
  onHide: function() {
    console.log("App Hide");
  }
};
function createApp() {
  const app = common_vendor.createSSRApp(_sfc_main);
  return {
    app
  };
}
createApp().app.mount("#app");
exports.createApp = createApp;
