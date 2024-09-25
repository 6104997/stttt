<template>
	<view class="classuser PageBg">
		<view class="userLayout" v-if="User">
			<view class="avatar">
				<image :src='User.headPortrait' mode="aspectFill"></image>
			</view>
			<view class="userName weui-input">{{User.nickname}}</view>



			<view class="userId">ID:{{User.ID}}</view>
			<view class="userInfo">
			</view>
		</view>


		<common-userlist>
			<template #listicons>
				<view class="deta">
					<uni-icons type="contact-filled" size="30" color="#23b389"></uni-icons>
				</view>
			</template>
			<template #listname>
				个人中心
			</template>
			<template #listrighticons>
				<uni-icons type="right" size="15" color="#aaa"></uni-icons>
			</template>
		</common-userlist>
		<common-userlist>
			<template #listicons>
				<view class="deta">
					<uni-icons type="weixin" size="30" color="#23b389"></uni-icons>
				</view>
			</template>
			<template #listname>
				联系客服
			</template>
			<template #listrighticons>
				<uni-icons type="right" size="15" color="#aaa"></uni-icons>
			</template>
		</common-userlist>










	</view>
</template>


<script setup>
	import cookies from 'weapp-cookie'

	import {
		onMounted,
		ref
	} from 'vue';
	import {
		RequestApi,
		GetLogin,
		GetUserinfo
	} from '@/api/request.js'

	const User = ref()
	const Login = async() => {

		if (cookies.get("x-token"))  {
			let resa = await GetUserinfo()
			User.value = resa.data

		}else{
			wx.login({
				success: async (res) => {
          if (res.code != null) {
            let resa = await GetLogin(res.code)
            User.value = resa.data.user
          }

				}

			})
		}



	}

	//personalCenter



	Login()

	onMounted(async () => {
		if (cookies.get("x-token")) {
			let resa = await GetUserinfo()
			console.log(resa);
		}

	})
</script>

<style lang="scss" scoped>
	.classuser {
		.userLayout {
			display: flex;
			align-items: center;
			justify-content: center;
			flex-direction: column;
			padding: 50rpx 0;

			.avatar {
				width: 160rpx;
				height: 160rpx;
				border-radius: 50%;
				overflow: hidden;

				image {
					width: 100%;
					height: 100%;
				}
			}

			.userName {
				font-size: 40rpx;
				color: #333333;
				padding: 20rpx 0 10rpx;

			}

			.userId {
				font-size: 28rpx;
				color: #aaaaaa;
			}
		}


	}
</style>
