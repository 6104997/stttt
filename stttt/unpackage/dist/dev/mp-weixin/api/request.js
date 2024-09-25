"use strict";
const common_vendor = require("../common/vendor.js");
const utils_request = require("../utils/request.js");
const baseUrl = "http://127.0.0.1:8891";
function RequestApi(url, data = {}, method = "GET") {
  return new Promise((resolve, reject) => {
    common_vendor.wx$1.request({
      url: baseUrl + url,
      header: {
        "content-type": "application/json"
        // 默认值
      },
      method,
      data,
      success: (res) => {
        resolve(res.data);
      },
      fail: (err) => {
        console.log(err);
        reject(err);
      }
    });
  });
}
const GetUserinfo = () => {
  return utils_request.request({
    url: "/appUser/getUserinfo",
    method: "GET"
  });
};
const GetLogin = (openpidcode) => {
  return utils_request.request({
    url: "/appUser/Login?openpid=" + openpidcode,
    method: "GET"
  });
};
const GetUpdateTheImage = (url) => {
  return utils_request.request({
    url: "/appUser/updateTheImage?image=" + url,
    method: "GET"
  });
};
const GetUpnikname = (nickname) => {
  return utils_request.request({
    url: "/appUser/upnickname?nickname=" + nickname,
    method: "GET"
  });
};
function uploadFile(filePath) {
  return new Promise((resolve, reject) => {
    common_vendor.wx$1.uploadFile({
      url: "http://localhost:8891/fileUploadAndDownload/upload",
      // 仅为示例，非真实的接口地址
      filePath,
      name: "file",
      formData: {},
      header: {
        "x-token": common_vendor.cookies.get("x-token")
      },
      success: resolve,
      // 成功时调用 resolve
      fail: reject
      // 失败时调用 reject
    });
  });
}
exports.GetLogin = GetLogin;
exports.GetUpdateTheImage = GetUpdateTheImage;
exports.GetUpnikname = GetUpnikname;
exports.GetUserinfo = GetUserinfo;
exports.RequestApi = RequestApi;
exports.uploadFile = uploadFile;
