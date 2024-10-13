"use strict";
const common_vendor = require("../../common/vendor.js");
const api_request = require("../../api/request.js");
const utils_system = require("../../utils/system.js");
const _sfc_main = {
  __name: "personalCenter",
  setup(__props) {
    let temporaryAddress = "";
    const Userinfos = common_vendor.ref({
      image: null,
      uuid: null,
      nickname: null
    });
    async function onChooseAvatar(e) {
      temporaryAddress = e.detail.avatarUrl;
      Userinfos.value.image = temporaryAddress;
    }
    function Upnickname(event) {
      Userinfos.value.nickname = event.target.value;
    }
    async function Update() {
      if (temporaryAddress) {
        const res = await api_request.uploadFile(temporaryAddress, Userinfos.value.uuid);
        console.log(res);
        const data = JSON.parse(res.data);
        if (data.code == 0) {
          const urls = data.data.file.url;
          console.log(urls);
          console.log(await api_request.GetUpdateTheImage(urls));
        }
      }
      console.log(await api_request.GetUpnikname(Userinfos.value.nickname));
      common_vendor.index.navigateBack({});
    }
    common_vendor.onMounted(async () => {
      if (common_vendor.cookies.get("x-token")) {
        let resa = await api_request.GetUserinfo();
        if (resa.code == 0) {
          Userinfos.value.image = resa.data.headPortrait;
          Userinfos.value.uuid = resa.data.uuid;
          Userinfos.value.nickname = resa.data.nickname;
        }
        console.log(resa);
      } else {
        common_vendor.wx$1.login({
          success: async (res) => {
            if (res.code != null) {
              let resa = await api_request.GetLogin(res.code);
              console.log(resa);
              Userinfos.value.image = resa.data.user.headPortrait;
              Userinfos.value.uuid = resa.data.user.uuid;
              Userinfos.value.nickname = resa.data.user.nickname;
            }
          }
        });
      }
    });
    return (_ctx, _cache) => {
      return {
        a: common_vendor.unref(utils_system.getNavBarheight)() + "px",
        b: Userinfos.value.image,
        c: common_vendor.o(onChooseAvatar),
        d: common_vendor.o(Upnickname),
        e: Userinfos.value.nickname,
        f: common_vendor.o(($event) => Userinfos.value.nickname = $event.detail.value),
        g: common_vendor.t(Userinfos.value.uuid),
        h: common_vendor.o(Update)
      };
    };
  }
};
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["__scopeId", "data-v-7b035f71"]]);
_sfc_main.__runtimeHooks = 2;
wx.createPage(MiniProgramPage);
