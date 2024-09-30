<template>
	<view class="noticeLayout">
		<view class="title">
			<view class="tag" v-if="detali.select">
				<uni-tag inverted text="置顶" type="error" />
			</view>
			<view class="font">{{detali.title}}</view>			
		</view>
		
		<view class="info">
			<view class="item">{{detali.author}}</view>					
			<view class="item">
				{{detali.CreatedAt}}
			</view>	
		</view>
		
		<mp-html :content="detali.content" />
		
		
		<view class="count">
			热度 {{detali.view_count}}	
		</view>
	</view>
	<ad unit-id="adunit-a2ae79342f0f4a70" ad-intervals="30" ad-type="video" ad-theme="white" bindload="adLoad"
	binderror="adError" bindclose="adClose"></ad>
</template>

<script setup>
import { getArticleManagementByUUID,convertTimeFormat } from '@/api/request.js';
import { ref } from 'vue';
import {onLoad} from "@dcloudio/uni-app"
const detali = ref({})



const getArticleManagementByUUIDs = async () => {
	const res = await getArticleManagementByUUID(noticeuuid);
	detali.value = res.data
	detali.value.CreatedAt = convertTimeFormat(res.data.CreatedAt)
	
}
let noticeuuid
onLoad((e)=>{
	noticeuuid = e.uuid
	getArticleManagementByUUIDs()
})

const onShareAppMessage = () => {
		title: '守塔金房助手-首页'
		imageUrl: 'https://stamdin.gys9.com/20240403102814.png'
	}




</script>

<style lang="scss" scoped>
.noticeLayout{
	padding:30rpx;
		.title{
			font-size: 40rpx;
			color:#111;
			line-height: 1.6em;
			padding-bottom:30rpx;
			display: flex;
			.tag{
				transform: scale(0.8);
				transform-origin: left center;
				flex-shrink: 0;	
			}
			.font{
				padding-left:6rpx;
			}
		}
		.info{
			display: flex;
			align-items: center;
			color:#999;
			font-size: 28rpx;
			.item{
				padding-right: 20rpx;
			}
		}
		.content{
			padding:50rpx 0;
		}
		.count{
			color:#999;
			font-size: 28rpx;
		}
}
</style>