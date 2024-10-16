"use strict";
const common_vendor = require("../common/vendor.js");
const utils_request = require("../utils/request.js");
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
const GetImage = () => {
  return utils_request.request({
    url: "/arouselImage/GetImage",
    method: "GET"
  });
};
const getArticleManagementByUUID = (uuid) => {
  return utils_request.request({
    url: "/articleManagement/getArticleManagementByUUID?uuid=" + uuid,
    method: "GET"
  });
};
const getAListOfArticles = (page, pageSize, articleClassificationId) => {
  return utils_request.request({
    url: "/articleManagement/getAListOfArticles?page=" + page + "&pageSize=" + pageSize + "&articleClassificationId=" + articleClassificationId,
    method: "GET"
  });
};
const getAListOfRooms = () => {
  return utils_request.request({
    url: "/goldenRoomForm/getAListOfGoldenHouses",
    method: "GET"
  });
};
const getNavigationist = () => {
  return utils_request.request({
    url: "/categoricalNavigationManagement/getAListOfNavigationCategories",
    method: "GET"
  });
};
const getcustomGetArticles = (url, page, pageSize) => {
  return utils_request.request({
    url: "/" + url + "&page=" + page + "&pageSize=" + pageSize,
    method: "GET"
  });
};
const convertTimeFormat = (tim) => {
  let date = new Date(tim);
  return `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(2, "0")}-${String(date.getDate()).padStart(2, "0")} ${String(date.getHours()).padStart(2, "0")}:${String(date.getMinutes()).padStart(2, "0")}:${String(date.getSeconds()).padStart(2, "0")}`;
};
function uploadFile(filePath, uuid) {
  return new Promise((resolve, reject) => {
    common_vendor.wx$1.uploadFile({
      url: "https://api.app.gys9.com/fileUploadAndDownload/upload",
      // 仅为示例，非真实的接口地址
      filePath,
      name: "file",
      formData: {
        "name": uuid
      },
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
exports.GetImage = GetImage;
exports.GetLogin = GetLogin;
exports.GetUpdateTheImage = GetUpdateTheImage;
exports.GetUpnikname = GetUpnikname;
exports.GetUserinfo = GetUserinfo;
exports.convertTimeFormat = convertTimeFormat;
exports.getAListOfArticles = getAListOfArticles;
exports.getAListOfRooms = getAListOfRooms;
exports.getArticleManagementByUUID = getArticleManagementByUUID;
exports.getNavigationist = getNavigationist;
exports.getcustomGetArticles = getcustomGetArticles;
exports.uploadFile = uploadFile;
