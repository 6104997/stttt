"use strict";const e=require("../common/vendor.js");exports.request=function(t={}){let{url:o,method:s="GET",data:r}=t;return o="https://api.app.gys9.com"+o,new Promise(((t,a)=>{e.wx$1.request({url:o,method:s,data:r,header:{"content-type":"application/json"},success:e=>{t(e.data)},fail:e=>{console.log(e),a(e)}})}))};