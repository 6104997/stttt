<template>
	<!-- 暂时没名字的插件 -->
	<view class="common-title">
		<view class="name">
			<slot name="name"></slot>
		</view>
		<view class="custom">
			<slot name="custom"></slot>
		</view>
	</view>
</template>

<script setup>

</script>

<style lang="scss" scoped>
.common-title{
	display: flex;
	justify-content: space-between;
	align-items: center;
	padding:0 30rpx;
	// color:#23b389 ;
	.name{
		font-size: 40rpx
	}
}
</style>