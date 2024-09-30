<template>
	<!-- 文章展示插件 -->
	<view class="image-card">
		<view class="image-section">
			<image class="image" :src="images" mode="aspectFill"></image>
		</view>
		<view class="content-section">
			<text class="title">{{ title }}</text>
			<view class="info">
				<text class="author">{{ author }}</text>
				<text class="date">{{CreatedAt}}</text>
			</view>
		</view>
	</view>
</template>

<script setup>
defineProps({
  title: {
    type: String,
    default: '文章标题'
  },
  author:{
	  type: String,
	  default: '作者'
  },
  CreatedAt:{
	  type: String,
	  default: '发布时间'
  },
  images:{
	  type: String,
	  default: '/static/c.png'
  }
  
})
</script>

<style lang="scss" scoped>
	.image-card {
		width: 690rpx;
		border-radius: 8px;
		box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
		overflow: hidden;
		background-color: #fff;
		margin: 20rpx auto 20rpx;
		
	}

	.image-section {
		
		
		height: 120px;
	}

	.image {
		
		border-radius: 8px;
		width: 100%;
		height: 100%;
	}

	.content-section {
		padding: 10px;
	}

	.title {
		font-size: 14px;
		font-weight: bold;
		margin-bottom: 6px;
	}

	.info {
		display: flex;
		justify-content: space-between;
		font-size: 12px;
		color: #888;
	}

	.author,
	.date {
		flex: 1;
	}
</style>