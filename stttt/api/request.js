const baseUrl = 'http://127.0.0.1:8891'
import {
	request
} from '@/utils/request.js'
import cookies from 'weapp-cookie'
function RequestApi(url, data = {}, method = 'GET') {
	return new Promise((resolve, reject) => {
		wx.request({
			url: baseUrl + url,
			header: {
				'content-type': 'application/json' // 默认值
			},
			method,
			data,
			success: (res) => {

				resolve(res.data)
			},
			fail: (err) => {
				console.log(err);
				reject(err)
			},
		})

	})
}

function GetRequestApi(url, data = {}, method = 'GET') {
	return new Promise((resolve, reject) => {
		wx.request({
			url: baseUrl + url,
			header: {
				'content-type': 'application/json' // 默认值
			},
			method,
			data,
			success: (res) => {

				resolve(res.data)
			},
			fail: (err) => {
				console.log(err);
				reject(err)
			},
		})

	})
}

function PostRequestApi(url, data = {}, method = 'POST') {
	return new Promise((resolve, reject) => {
		wx.request({
			url: baseUrl + url,
			header: {
				'content-type': 'application/json' // 默认值
			},
			method,
			data,
			success: (res) => {

				resolve(res.data)
			},
			fail: (err) => {
				console.log(err);
				reject(err)
			},
		})

	})
}


export const GetUserinfo = () => {
	return request({
		url: '/appUser/getUserinfo',
		method: "GET"
	})
}

export const GetLogin = (openpidcode) => {
	return request({
		url: '/appUser/Login?openpid=' + openpidcode,
		method: "GET"
	})

}


export const GetUpdateTheImage = (url) => {
	return request({
		url: '/appUser/updateTheImage?image=' + url,
		method: "GET"
	})

}
export const GetUpnikname = (nickname) => {
	return request({
		url: '/appUser/upnickname?nickname=' + nickname,
		method: "GET"
	})
}
export function uploadFile(filePath) {
    return new Promise((resolve, reject) => {
        wx.uploadFile({
            url: 'http://localhost:8891/fileUploadAndDownload/upload', // 仅为示例，非真实的接口地址
            filePath: filePath,
            name: 'file',
            formData: {},
            header: {
                'x-token': cookies.get("x-token")
            },
            success: resolve, // 成功时调用 resolve
            fail: reject // 失败时调用 reject
        });
    });
}



export {
	RequestApi,
	GetRequestApi,
	PostRequestApi
}
