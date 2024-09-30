<script>
	import {
		GetLogin
	} from '@/api/request.js'

	export default {
		onLaunch: function() {
			const Openpid = wx.getStorageSync('Openpid')
			//console.log('App onLaunch',Openpid)
			if (Openpid === "") {
				wx.login({
					success: async (res) => {
						await GetLogin(res.code)
					}

				})

			}

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