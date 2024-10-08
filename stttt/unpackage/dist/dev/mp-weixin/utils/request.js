"use strict";
const common_vendor = require("../common/vendor.js");
const baseUrl = "https://api.app.gys9.com";
function request(config = {}) {
  let {
    url,
    method = "GET",
    data
  } = config;
  url = baseUrl + url;
  return new Promise((resolve, reject) => {
    common_vendor.wx$1.request({
      url,
      method,
      data,
      header: {
        "content-type": "application/json"
        // 默认值
      },
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
exports.request = request;
