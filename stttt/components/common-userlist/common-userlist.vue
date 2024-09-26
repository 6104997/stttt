<template>
	<view class="common-userlist">
		<view class="section">
			<view class="list">
				<view class="left">
					<slot name="listicons"></slot>
					<view class="text">
						<slot name="listname"></slot>
					</view>
				</view>
				<view class="right">
					<view class="text">
						<slot name="listda"></slot>
					</view>
					<slot name="listrighticons"></slot>

				</view>
				
			</view>
		</view>
	</view>
</template>

<script>
</script>

<style lang="scss" scoped>
	.common-userlist {
		.section {
			width: 690rpx;
			margin: 0 auto;
			border: 1px solid #eee;
			border-radius: 10rpx;
			box-shadow: 0 0 30rpx rgba(0, 0, 0, 0.05);
			overflow: hidden;
			white-space: nowrap;
			text-overflow: ellipsis;

			.list {
				display: flex;
				justify-content: space-between;
				align-items: center;
				padding: 0 20rpx;
				height: 100rpx;
				border-bottom: 1px solid #eee;

				&:last-child {
					border-bottom: 0;
				}

				.left {
					display: flex;
					align-items: center;
					color: #666;

					.text {
						padding-left: 20rpx;
					}
				}

				.right {
					font-size: 40rpx;
					display: flex;
					align-items: center;

					.text {
						font-size: 28rpx;
						color: #aaa;
					}
				}
			}
		}
	}
</style>