"use strict";
const common_vendor = require("../../common/vendor.js");
if (!Array) {
  const _easycom_uni_icons2 = common_vendor.resolveComponent("uni-icons");
  _easycom_uni_icons2();
}
const _easycom_uni_icons = () => "../../uni_modules/uni-icons/components/uni-icons/uni-icons.js";
if (!Math) {
  _easycom_uni_icons();
}
const _sfc_main = {
  __name: "common-listOfRooms",
  props: {
    roomNumber: {
      type: String,
      default: "房间号"
    },
    roomPassword: {
      type: String,
      default: "房间密码"
    },
    currentStatus: {
      type: String,
      default: "状态"
    },
    CreatedAt: {
      type: String,
      default: "创建时间"
    }
  },
  setup(__props) {
    const copy = (text) => {
      common_vendor.index.setClipboardData({
        data: text,
        success: function() {
          common_vendor.index.showToast({
            title: "复制成功",
            icon: "none"
          });
        },
        fail: function(err) {
          common_vendor.index.showToast({
            title: "复制失败",
            icon: "none"
          });
        }
      });
    };
    return (_ctx, _cache) => {
      return {
        a: common_vendor.p({
          type: "calendar",
          size: "20"
        }),
        b: common_vendor.t(__props.roomNumber),
        c: common_vendor.t(__props.roomPassword),
        d: common_vendor.p({
          type: "loop",
          size: "20"
        }),
        e: common_vendor.t(__props.CreatedAt),
        f: common_vendor.p({
          type: "eye",
          size: "20"
        }),
        g: common_vendor.t(__props.currentStatus),
        h: common_vendor.o(($event) => copy(__props.roomNumber))
      };
    };
  }
};
const Component = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__scopeId", "data-v-13556ce8"]]);
wx.createComponent(Component);
