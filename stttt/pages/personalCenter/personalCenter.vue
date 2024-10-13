<template>
	<view class="container PageBg">
		<view :style="{height:getNavBarheight()+'px'}"></view>
		<!-- <image class="avatar" src="/static/logo.png"></image> -->

		<button class="imageButton" open-type="chooseAvatar" @chooseavatar="onChooseAvatar">
			<image class="image" :src='Userinfos.image'></image>
		</button>

		<view class="info-section">
			<view class="info-item">
				<text>昵称</text>
				<input v-model="Userinfos.nickname" type="nickname" class="weui-input" placeholder="请输入昵称"
					@blur="Upnickname" />
			</view>
			<view class="info-item">
				<text>UUid</text>
				<view class="phone">
					<text class="info-value">{{Userinfos.uuid}}</text>

				</view>
			</view>
			<view class="info-item">
				<text>当前版本</text>
				<view class="phone">
					<text class="info-value">1.0.0</text>

				</view>
			</view>

			<view class="info-item">
				<text>温馨提示</text>
				<view class="phone">
					<text class="info-value">更换头像点击头像即可，更换昵称点击昵称即可</text>

				</view>
			</view>


		</view>

		<button class="save" type="primary" @click="Update">保存</button>
		<!-- <button class="logout">退出登录</button> -->
		<ad unit-id="adunit-a2ae79342f0f4a70" ad-intervals="30" ad-type="video" ad-theme="white" bindload="adLoad"
			binderror="adError" bindclose="adClose"></ad>
	</view>
</template>
<script setup>
	import cookies from 'weapp-cookie'
	import {
		onMounted,
		ref
	} from 'vue';
	import {
		GetUserinfo,
		GetUpdateTheImage,
		uploadFile,
		GetLogin,
		GetUpnikname
	} from '@/api/request.js'
	import {
		getNavBarheight
	} from '@/utils/system.js'
	let temporaryAddress = ''
	const Userinfos = ref({
		image: null,
		uuid: null,
		nickname: null
	})
	async function onChooseAvatar(e) {
		temporaryAddress = e.detail.avatarUrl
		Userinfos.value.image = temporaryAddress

	}

	function Upnickname(event) {
		Userinfos.value.nickname = event.target.value
	}

	async function Update() {
		if (temporaryAddress) {
			const res = await uploadFile(temporaryAddress, Userinfos.value.uuid);
			console.log(res);
			const data = JSON.parse(res.data);
			

			if (data.code == 0) {
				const urls = data.data.file.url;
				console.log(urls);
				console.log(await GetUpdateTheImage(urls))



			}

		}
		console.log(await GetUpnikname(Userinfos.value.nickname));
		uni.navigateBack({})

	}

	onMounted(async () => {
		if (cookies.get("x-token")) {
			let resa = await GetUserinfo()
			if (resa.code == 0) {
				Userinfos.value.image = resa.data.headPortrait
				Userinfos.value.uuid = resa.data.uuid
				Userinfos.value.nickname = resa.data.nickname
			}
			console.log(resa);
		} else {
			wx.login({
				success: async (res) => {
					if (res.code != null) {
						let resa = await GetLogin(res.code)
						console.log(resa);
						Userinfos.value.image = resa.data.user.headPortrait
						Userinfos.value.uuid = resa.data.user.uuid
						Userinfos.value.nickname = resa.data.user.nickname
					}

				}

			})
		}
		const onShareAppMessage = () => {
			title: '守塔金房助手-首页'
			imageUrl: 'https://stamdin.gys9.com/20240403102814.png'
		}
	})
</script>
<style lang="scss" scoped>
	.container {
		padding: 20px;
		background-color: #f5f5f5;

		.imageButton {
			width: 160rpx;
			height: 160rpx;
			padding: 0;
			border-radius: 50%;

			.image {
				margin: 0;
				width: 100%;
				height: 100%;
			}
		}
	}

	.profile-picture {
		display: flex;
		justify-content: center;
	}



	.info-section {
		background-color: white;
		margin: 20px 0;
		padding: 15px;
		border-radius: 10px;
	}

	.info-item {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 15px;
		border-bottom: 1rpx solid $text-font-color-3;
	}

	.info-value {
		color: $text-font-color-3;
		font-size: 24rpx;
	}

	.phone {
		display: flex;
		align-items: center;
	}

	.change-phone {
		color: #1aad19;
		margin-left: 10px;
	}

	.gender {
		display: flex;
		align-items: center;
	}



	.save {
		background-color: #1aad19;
		color: white;
		padding: 10px;
		border-radius: 5px;
		margin-bottom: 20px;
	}

	.logout {
		background-color: #ffffff;
		color: #333;
		padding: 10px;
		border-radius: 5px;
	}
</style>