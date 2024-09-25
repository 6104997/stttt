"use strict";
const common_vendor = require("../common/vendor.js");
const baseUrl = "http://127.0.0.1:8891";
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
