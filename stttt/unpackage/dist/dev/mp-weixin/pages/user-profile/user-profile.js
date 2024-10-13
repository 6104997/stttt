"use strict";
const common_vendor = require("../../common/vendor.js");
const common_assets = require("../../common/assets.js");
const _sfc_main = {
  data() {
    return {
      birthday: ""
    };
  },
  methods: {
    onDateChange(e) {
      this.birthday = e.detail.value;
    },
    saveProfile() {
      common_vendor.index.showToast({ title: "保存成功" });
    },
    logout() {
      common_vendor.index.showModal({
        title: "提示",
        content: "确定要退出登录吗？",
        success: (res) => {
          if (res.confirm) {
            common_vendor.index.showToast({ title: "已退出登录" });
          }
        }
      });
    }
  }
};
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return {
    a: common_assets._imports_0,
    b: common_vendor.t($data.birthday || "请选择 >"),
    c: common_vendor.o((...args) => $options.onDateChange && $options.onDateChange(...args)),
    d: common_vendor.o((...args) => $options.saveProfile && $options.saveProfile(...args)),
    e: common_vendor.o((...args) => $options.logout && $options.logout(...args))
  };
}
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render], ["__scopeId", "data-v-ef78d882"]]);
wx.createPage(MiniProgramPage);
