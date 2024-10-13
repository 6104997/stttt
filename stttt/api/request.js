const baseUrl = 'https://api.app.gys9.com/'
import {
	request
} from '@/utils/request.js'
import cookies from 'weapp-cookie'

//获取用户信息表单
export const GetUserinfo = () => {
	return request({
		url: '/appUser/getUserinfo',
		method: "GET"
	})
}

//登录
export const GetLogin = (openpidcode) => {
	return request({
		url: '/appUser/Login?openpid=' + openpidcode,
		method: "GET"
	})

}


//修改头像
export const GetUpdateTheImage = (url) => {
	return request({
		url: '/appUser/updateTheImage?image=' + url,
		method: "GET"
	})

}

//修改昵称
export const GetUpnikname = (nickname) => {
	return request({
		url: '/appUser/upnickname?nickname=' + nickname,
		method: "GET"
	})
}

//获取公告内容列表
export const getAListOfAnnouncements = () => {
	return request({
		url: '/announcementManagement/getAListOfAnnouncements',
		method: "GET"
	})
}

//获取轮播图
export const GetImage = () => {
	return request({
		url: '/arouselImage/GetImage',
		method: "GET"
	})
}
//根据UUID查询文章
export const getArticleManagementByUUID = (uuid) => {
	return request({
		url: '/articleManagement/getArticleManagementByUUID?uuid=' + uuid,
		method: "GET"
	})
}
//获取文章列表
export const getAListOfArticles = (page, pageSize, articleClassificationId) => {
	return request({
		url: '/articleManagement/getAListOfArticles?page=' + page + '&pageSize=' + pageSize +
			'&articleClassificationId=' + articleClassificationId,
		method: "GET"
	})
}

//获取房间列表
export const getAListOfRooms = () => {
	return request({
		url: '/goldenRoomForm/getAListOfGoldenHouses',
		method: "GET"
	})
}


//获取导航列表数据
export const getNavigationist = () => {
	return request({
		url: '/categoricalNavigationManagement/getAListOfNavigationCategories',
		method: "GET"
	})
}

//自定义获取文章

export const getcustomGetArticles = (url, page, pageSize) => {
	return request({
		url: '/' + url + '&page=' + page + '&pageSize=' + pageSize,
		method: "GET"
	})
}

//转换时间格式
export const convertTimeFormat = (tim) => {
	let date = new Date(tim)
	return `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(2, '0')}-${String(date.getDate()).padStart(2, '0')} ${String(date.getHours()).padStart(2, '0')}:${String(date.getMinutes()).padStart(2, '0')}:${String(date.getSeconds()).padStart(2, '0')}`;
}


export function uploadFile(filePath, uuid) {
	return new Promise((resolve, reject) => {
		wx.uploadFile({
			url: 'https://api.app.gys9.com/fileUploadAndDownload/upload', // 仅为示例，非真实的接口地址
			filePath: filePath,
			name: 'file',
			formData: {
				'name': uuid
			},
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