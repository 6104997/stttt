"use strict";
const common_vendor = require("../../common/vendor.js");
const api_request = require("../../api/request.js");
if (!Array) {
  const _easycom_uni_tag2 = common_vendor.resolveComponent("uni-tag");
  const _easycom_mp_html2 = common_vendor.resolveComponent("mp-html");
  (_easycom_uni_tag2 + _easycom_mp_html2)();
}
const _easycom_uni_tag = () => "../../uni_modules/uni-tag/components/uni-tag/uni-tag.js";
const _easycom_mp_html = () => "../../uni_modules/mp-html/components/mp-html/mp-html.js";
if (!Math) {
  (_easycom_uni_tag + _easycom_mp_html)();
}
const _sfc_main = {
  __name: "detail",
  setup(__props) {
    const detali = common_vendor.ref({});
    const getArticleManagementByUUIDs = async () => {
      const res = await api_request.getArticleManagementByUUID(noticeuuid);
      detali.value = res.data;
      detali.value.CreatedAt = api_request.convertTimeFormat(res.data.CreatedAt);
    };
    let noticeuuid;
    common_vendor.onLoad((e) => {
      noticeuuid = e.uuid;
      getArticleManagementByUUIDs();
    });
    return (_ctx, _cache) => {
      return common_vendor.e({
        a: detali.value.select
      }, detali.value.select ? {
        b: common_vendor.p({
          inverted: true,
          text: "置顶",
          type: "error"
        })
      } : {}, {
        c: common_vendor.t(detali.value.title),
        d: common_vendor.t(detali.value.author),
        e: common_vendor.t(detali.value.CreatedAt),
        f: common_vendor.p({
          content: detali.value.content
        }),
        g: common_vendor.t(detali.value.view_count)
      });
    };
  }
};
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__scopeId", "data-v-eca06f3c"]]);
_sfc_main.__runtimeHooks = 2;
wx.createPage(MiniProgramPage);
