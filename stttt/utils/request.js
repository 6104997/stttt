const baseUrl = 'https://api.app.gys9.com'
export function request(config = {}) {
	let {
		url,
		method = "GET",
		data
	} = config
	url = baseUrl + url
	return new Promise((resolve, reject) => {
		wx.request({
			url,
			method,
			data,
			header: {
				'content-type': 'application/json' // 默认值
			},
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