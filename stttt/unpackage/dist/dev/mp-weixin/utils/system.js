"use strict";
const common_vendor = require("../common/vendor.js");
let SYSTEM_INFO = common_vendor.index.getSystemInfoSync();
const getStatusBarHeight = () => SYSTEM_INFO.statusBarHeight || 0;
const getTitbarHeight = () => {
  let { top, height } = common_vendor.index.getMenuButtonBoundingClientRect();
  return height + (top - SYSTEM_INFO.statusBarHeight) * 2;
};
const getNavBarheight = () => getStatusBarHeight() + getTitbarHeight();
exports.getNavBarheight = getNavBarheight;
exports.getStatusBarHeight = getStatusBarHeight;
exports.getTitbarHeight = getTitbarHeight;
