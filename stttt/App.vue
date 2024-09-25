<script>
	export default {
		onLaunch: function() {
			// wx.login({
			// 	success: async (res) => {
			// 		//const UserData = await RequestApi("/user_api_openpid?openpid=" + res.code)
			// 		console.log(res.code)
			// 	}
			
			// })
			console.log('App Launch')
		},
		onShow: function() {
			console.log('App Show')
		},
		onHide: function() {
			console.log('App Hide')
		}
	}
</script>

<style lang="scss">
	/*每个页面公共css */
	@import "common/style/common-style.scss"
</style>
