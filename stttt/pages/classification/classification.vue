<template>
	<view class="classclassification PageBg">
		<z-paging ref="paging" v-model="customGetArticless" @query="queryList" :default-page-size="5">
			<custom-bav-bar title="分类"></custom-bav-bar>
			
			<common-title>
				<template #name>置顶公告</template>
				<template #custom>
					<view class="date">
					</view>
				</template>
			</common-title>
			<view v-for="item in ListOfAnnouncements" :key="item.uuid">
				<navigator :url="'/pages/detail/detail?uuid='+item.uuid">
					<component-articleList :images="item.preview_the_image" :title="item.title" :author="item.author"
						:CreatedAt="convertTimeFormat(item.created_at)"></component-articleList>
				</navigator>
			</view>
			<ad unit-id="adunit-a2ae79342f0f4a70" ad-intervals="30" ad-type="video" ad-theme="white" bindload="adLoad"
				binderror="adError" bindclose="adClose"></ad>
			<scroll-view class="scroll-content" scroll-x="true">
				<view class="scroll-item" v-for="(item,index) in listOfArticles" :key="index"
					@tap='changeTab(index,item.navigationUrl)'>
					<text
						:class='listOfArticlesIndex===index? "f-active-color":"f-color"'>{{item.navigationName}}</text>
				</view>
			</scroll-view>
			

			<view v-for="item in customGetArticless" :key="item.uuid">
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
	} from 'vue';
	import {
		getNavBarheight
	} from '@/utils/system.js'

	import {
		convertTimeFormat,
		getAListOfArticles,
		getNavigationist,
		getcustomGetArticles

	} from '@/api/request.js'
	
	import { onShow } from '@dcloudio/uni-app'
	const bannerList = ref([])

	const current = ref(0);
	const ListOfAnnouncements = ref([])

	
	const paging = ref(null)
	const items = ref([])
	const listOfArticles = ref([]);
	const listOfArticlesIndex = ref(0)
	const scrollIntoIndex = ref('tap0')
	const customGetArticless = ref([])
	const urll = ref()
	const GetListOfAnnouncements = async () => {
		//GetListOfAnnouncements
		//获取公告内容列表
		let res = await getAListOfArticles("1", "1", '0')
		if (res.code === 0) {

			ListOfAnnouncements.value = res.data.list

		}
	}

	//获取导航列表
	const GetNavigationists = async () => {
		listOfArticlesIndex.value = 0
		let res = await getNavigationist()
		listOfArticles.value = res.data
		res = await getcustomGetArticles(
			"articleManagement/getAListOfArticles?theLatestArticles=1&articleClassificationId=1")
		customGetArticless.value = res.data.list
		urll.value = "articleManagement/getAListOfArticles?theLatestArticles=1&articleClassificationId=1"
		queryList(1,5)
	}
	
	const queryList = async (pageNo, pageSize) => {
		// 此处请求仅为演示，请替换为自己项目中的请求
		await getcustomGetArticles(urll.value,pageNo, pageSize).then(res => {
			// 将请求结果通过complete传给z-paging处理，同时也代表请求结束，这一行必须调用
			console.log(urll.value,res);
			paging.value.complete(res.data.list);
			
		}).catch(res => {
			// 如果请求失败写paging.value.complete(false);
			// 注意，每次都需要在catch中写这句话很麻烦，z-paging提供了方案可以全局统一处理
			// 在底层的网络请求抛出异常时，写uni.$emit('z-paging-error-emit');即可
			paging.value.complete(false);
		})
	}
	
	
	
	//导航点击
	const changeTab = async (index, navigationUrl) => {
		if (listOfArticlesIndex.value === index) {
			return
		}
		listOfArticlesIndex.value = index
		paging.value = null
		customGetArticless.value =[]
		//let ress = await getcustomGetArticles(navigationUrl)
		//console.log(ress,navigationUrl);
		urll.value = navigationUrl
		queryList(1,5)
		// customGetArticless.value = ress.data.list
		//paging.value.complete(ress.data.list);
	}
	const onShareAppMessage = () => {
			title: '守塔金房助手-首页'
			imageUrl: 'https://stamdin.gys9.com/20240403102814.png'
		}
	onShow( () =>{
		GetNavigationists()
	})
	//获取公告内容列表
	GetListOfAnnouncements()
	//获取导航数据
	GetNavigationists()
</script>

<style lang="scss" scoped>
	.classclassification {
		.scroll-content {
			width: 690rpx;
			height: 80rpx;
			white-space: nowrap;

			.scroll-item {
				display: inline-block;
				padding: 10rpx 20rpx;
				font-size: 32rpx;

				.f-active-color {
					color: #42b7fb;
					padding: 10rpx 0;
					border-bottom: 6rpx solid #42b7fb;
				}

				.f-color {
					color: #636263;
				}
			}

		}

		swiper {
			width: 750rpx;
			height: 340rpx;

			&-item {
				width: 100%;
				height: 100%;
				padding: 0 30rpx;


			}
		}


	}
</style>