<template>
	<view class="homoLayout PageBg">
		<z-paging ref="paging" v-model="articleList" @query="queryList" :default-page-size="1">
			<custom-bav-bar title="守塔不能停"></custom-bav-bar>
			<view class="banner">
				<!--  获取轮播图并渲染    -->
				<swiper circular indicator-dots indicator-color="rgba(255,255,255,0.5)" indicator-active-color="#fff"
					autoplay :interval="3000">
					<swiper-item v-for="(item,index) in bannerList" :key="index">
						<image :src="item.image" mode="aspectFill"></image>
					</swiper-item>
				</swiper>
			</view>

			<common-title>
				<template #name>金房信息</template>
				<template #custom>
					<view class="date">
						<uni-icons type="calendar" size="18"></uni-icons>
					</view>
				</template>
			</common-title>

			<view class="common-box-list" v-for="item in 2">
				<view class="common-box-list-fj">
					<view class="left">
						<uni-icons type="calendar" size="20"></uni-icons>
						<view class="text">房间号:
							<text>8815156153</text>

						</view>
					</view>
					<view class="right">
						<view class="text">房间密码:8888</view>
					</view>
				</view>


				<view class="common-box-list-zt">
					<uni-icons type="loop" size="20"></uni-icons>
					<view class="text">状态:<text>未开始</text></view>
				</view>
				<view class="common-box-list-sj">

					<view class="left">
						<uni-icons type="eye" size="20"></uni-icons>
						<view class="text">创建时间:<text>未开始</text></view>
					</view>
					<view class="right">
						<button type="primary" size="mini">复制房间号</button>
					</view>
				</view>

			</view>
			<!-- <ad unit-id="adunit-a2ae79342f0f4a70" ad-intervals="30" ad-type="video" ad-theme="white" bindload="adLoad"
			binderror="adError" bindclose="adClose"></ad> -->
			<common-title>
				<template #name>公告</template>
				<template #custom>
					<view class="date">

					</view>
				</template>
			</common-title>
			<view v-for="item in articleList" :key="item.uuid">
				<navigator :url="'/pages/detail/detail?uuid='+item.uuid">
					<component-articleList :images="item.preview_the_image" :title="item.title" :author="item.author"
						:CreatedAt="convertTimeFormat(item.created_at)"></component-articleList>
				</navigator>
			</view>
			<view :style="{ height:'20px'}"></view>
		</z-paging>
	</view>

</template>

<script setup>
	import {
		ref
	} from 'vue'
	import {
		getAListOfAnnouncements,
		GetImage,
		getAListOfArticles,
		convertTimeFormat
	} from '@/api/request.js'
	//通过 http://127.0.0.1:8891/arouselImage/GetImage 获取图片列表
	const bannerList = ref([])
	const ListOfAnnouncements = ref([])
	const articleList = ref({})
	const paging = ref(null)
	const getBannerList = async () => {
		let res = await GetImage()
		if (res.code === 0) {
			bannerList.value = res.data
		}
	}

const queryList = (pageNo, pageSize) => {
		// 此处请求仅为演示，请替换为自己项目中的请求
        getAListOfArticles(pageNo,pageSize,'1').then(res => {
			// 将请求结果通过complete传给z-paging处理，同时也代表请求结束，这一行必须调用
            paging.value.complete(res.data.list);
        }).catch(res => {
			// 如果请求失败写paging.value.complete(false);
			// 注意，每次都需要在catch中写这句话很麻烦，z-paging提供了方案可以全局统一处理
			// 在底层的网络请求抛出异常时，写uni.$emit('z-paging-error-emit');即可
			paging.value.complete(false);
		})
    }




	const GetListOfAnnouncements = async () => {
		//GetListOfAnnouncements
		//获取公告内容列表
		let res = await getAListOfAnnouncements()
		if (res.code === 0) {

			ListOfAnnouncements.value = res.data

		}
	}
	const GetAListOfArticless = async () => {
		let res = await getAListOfArticles()

		articleList.value = res.data.list

		console.log(articleList.value);
	}



	//获取文章
	GetAListOfArticless()
	//获取公告内容列表
	GetListOfAnnouncements()
	//获取轮播图
	getBannerList()

	//公告列表
</script>

<style lang="scss" scoped>
	.homoLayout {
		.banner {
			width: 750rpx;
			padding: 30rpx 0;

			swiper {
				width: 750rpx;
				height: 340rpx;

				&-item {
					width: 100%;
					height: 100%;
					padding: 0 30rpx;

					image {
						width: 100%;
						height: 100%;
						border-radius: 10rpx;
					}
				}
			}
		}

		.common-box-list {
			width: 690rpx;
			height: 200rpx;
			margin: 20rpx auto 20rpx;
			border: 1px solid #eee;
			border-radius: 10rpx;
			box-shadow: 0 0 30rpx rgba(0, 0, 0, 0.05);
			overflow: hidden;
			white-space: nowrap;
			text-overflow: ellipsis;

			.common-box-list-fj {
				display: flex;
				justify-content: space-between;
				align-items: center;
				padding: 0 20rpx;
				height: 50rpx;
				//border-bottom: 1px solid #eee;

				.left {
					display: flex;

					:deep() {
						.uni-icons {
							color: $barand-theme-color !important;
						}
					}

					.text {
						color: $text-font-color-2;
					}
				}

				.right {
					display: flex;

					.text {
						color: $text-font-color-2;
					}
				}
			}

			.common-box-list-zt {
				display: flex;
				//justify-content: space-between;
				align-items: center;
				padding: 0 20rpx;
				height: 50rpx;

				:deep() {
					.uni-icons {
						color: $barand-theme-color !important;
					}

					.text {
						display: flex;
						color: $text-font-color-2;
					}
				}
			}

			.common-box-list-sj {
				display: flex;
				//justify-content: space-between;
				align-items: center;
				padding: 0 20rpx 0;
				height: 50rpx;

				.text {
					display: flex;

				}

				.left {
					display: flex;
					flex-grow: 1;

					.text {
						color: $text-font-color-2;
					}

					:deep() {
						.uni-icons {
							color: $barand-theme-color !important;
						}
					}
				}

				.right {
					display: flex;
					justify-content: flex-end;
				}
			}
		}
	}
</style>