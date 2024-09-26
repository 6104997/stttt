"use strict";
const common_vendor = require("../../common/vendor.js");
const _sfc_main = {
  __name: "component-articleList",
  props: {
    title: {
      type: String,
      default: "文章标题"
    },
    author: {
      type: String,
      default: "作者"
    },
    CreatedAt: {
      type: String,
      default: "发布时间"
    },
    images: {
      type: String,
      default: "/static/c.png"
    }
  },
  setup(__props) {
    return (_ctx, _cache) => {
      return {
        a: __props.images,
        b: common_vendor.t(__props.title),
        c: common_vendor.t(__props.author),
        d: common_vendor.t(__props.CreatedAt)
      };
    };
  }
};
const Component = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__scopeId", "data-v-06782a64"]]);
wx.createComponent(Component);
