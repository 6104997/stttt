<template>
	<!-- 房间展示插件 -->
	<view class="common-listOfRooms">
		<view class="common-box-list">
			<view class="common-box-list-fj">
				<view class="left">
					<uni-icons type="calendar" size="20"></uni-icons>
					<view class="text">房间号:
						<text>{{roomNumber}}</text>

					</view>
				</view>
				<view class="right">
					<view class="text">房间密码:{{roomPassword}}</view>
				</view>
			</view>
			<view class="common-box-list-zt">
				<uni-icons type="loop" size="20"></uni-icons>
				<view class="text">创建时间:<text>{{ CreatedAt}}</text></view>
			</view>
			<view class="common-box-list-sj">
				<view class="left">
					<uni-icons type="eye" size="20"></uni-icons>
					<view class="text">状态:<text>{{currentStatus}}</text></view>
				</view>
				<view class="right">
					<button type="primary" size="mini" @click="copy(roomNumber)">复制房间号</button>
				</view>
			</view>




		</view>
	</view>
</template>

<script setup>
	defineProps({
		roomNumber: {
			type: String,
			default: '房间号'
		},
		roomPassword: {
			type: String,
			default: '房间密码'
		},
		currentStatus: {
			type: String,
			default: '状态'
		},
		CreatedAt: {
			type: String,
			default: '创建时间'
		}
	})
	const copy = (text) => {


		uni.setClipboardData({
			data: text,
			
			success: function() {
				uni.showToast({
				        title: '复制成功',
				        icon: 'none'
				      });
			},
			fail: function(err) {
				uni.showToast({
				        title: '复制失败',
				        icon: 'none'
				      });

			}
		});
	};
</script>

<style lang="scss" scoped>
	.common-listOfRooms {
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
		}
	}
</style>