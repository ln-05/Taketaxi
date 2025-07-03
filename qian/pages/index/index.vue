<template>
	<view class="container">
		<view class="header">
			<text class="back" @click="goBack">&#8592;</text>
			<view class="logo">&#128663;</view>
			<text class="title">嗨，欢迎回来</text>
			<text class="subtitle">登录后更精彩，美好出行即将开始</text>
		</view>
		<view class="form">
			<view class="input-group">
				<picker mode="selector" :range="countryCodes" @change="onCountryChange">
					<view class="country-code">+{{ countryCode }}</view>
				</picker>
				<input class="phone-input" type="text" maxlength="11" placeholder="请输入手机号" v-model="phone" @input="onPhoneInput" />
			</view>
			<view class="agreement">
				<checkbox-group @change="onCheckChange">
					<checkbox value="agree" :checked="checkedArr.includes('agree')" />
				</checkbox-group>
				<text>阅读并同意</text>
				<text class="link">服务协议</text>
				<text>及滴滴出行基本功能个人信息处理规则</text>
			</view>
			<view class="next-btn" :class="{disabled: !canNext}" @click="nextStep">下一步</view>
		</view>
		<!-- 社交登录图标，放在页面底部 -->
		<view class="login-icons">
			<view class="login-icon-item" v-for="icon in icons" :key="icon.name">
				<view class="login-icon-img-wrap" @click="icon.name === '微信' ? showWechatQrcode() : null">
					<image :src="icon.src" class="login-icon-img" mode="aspectFill" />
				</view>
				<text class="login-icon-text">{{ icon.name }}</text>
			</view>
		</view>
		<!-- 微信二维码弹窗 -->
		<view v-if="showQrcode" class="qrcode-modal">
			<view class="qrcode-modal-mask" @click="closeQrcode"></view>
			<view class="qrcode-modal-content">
				<image class="qrcode-img" src="http://localhost:8080/wechat/login_qrcode" mode="aspectFit" />
				<view class="qrcode-close-btn" @click="closeQrcode">关闭</view>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				countryCodes: ['86', '1', '44'],
				countryCode: '86',
				phone: '',
				checkedArr: [],
				icons: [
					{ src: 'https://img.icons8.com/color/48/000000/alipay.png', name: '支付宝' },
					{ src: 'https://img.icons8.com/color/48/000000/weixing.png', name: '微信' },
					{ src: 'https://img.icons8.com/color/48/000000/approval.png', name: '身份认证' },
					{ src: 'https://img.icons8.com/ios-filled/50/000000/mac-os.png', name: 'Apple' },
					{ src: 'https://img.icons8.com/color/48/000000/phone.png', name: '手机号' }
				],
				showQrcode: false
			};
		},
		computed: {
			canNext() {
				return String(this.phone).length === 11 && this.checkedArr.includes('agree');
			}
		},
		methods: {
			goBack() {
				uni.navigateBack();
			},
			onCountryChange(e) {
				this.countryCode = this.countryCodes[e.detail.value];
			},
			onPhoneInput(e) {
				this.phone = String(e.detail.value || '');
			},
			nextStep() {
				if (!this.canNext) return;
				uni.navigateTo({
					url: `/pages/index/verify?phone=${this.phone}`
				});
			},
			onCheckChange(e) {
				this.checkedArr = e.detail.value;
			},
			showWechatQrcode() {
				this.showQrcode = true;
			},
			closeQrcode() {
				this.showQrcode = false;
			},
			onWechatLoginSuccess() {
				uni.redirectTo({ url: '/pages/index/didi' });
			}
		}
	};
</script>

<style scoped>
	.container {
		background: #fff;
		min-height: 100vh;
		padding: 0 5vw;
		box-sizing: border-box;
	}

	.header {
		margin-top: 60rpx;
		text-align: left;
	}

	.back {
		font-size: 40rpx;
		color: #333;
		margin-bottom: 20rpx;
		display: inline-block;
	}

	.logo {
		font-size: 80rpx;
		color: #ff8200;
		margin: 20rpx 0 10rpx 0;
	}

	.title {
		font-size: 44rpx;
		font-weight: bold;
		margin-bottom: 10rpx;
		display: block;
	}

	.subtitle {
		font-size: 28rpx;
		color: #888;
		margin-bottom: 40rpx;
		display: block;
	}

	.form {
		margin-top: 40rpx;
	}

	.input-group {
		display: flex;
		align-items: center;
		border: 1px solid #eee;
		border-radius: 16rpx;
		padding: 0 20rpx;
		background: #fafafa;
		margin-bottom: 24rpx;
	}

	.country-code {
		font-size: 32rpx;
		color: #333;
		margin-right: 16rpx;
	}

	.phone-input {
		flex: 1;
		font-size: 32rpx;
		border: none;
		background: transparent;
		outline: none;
		padding: 24rpx 0;
	}

	.agreement {
		display: flex;
		align-items: center;
		font-size: 24rpx;
		color: #888;
		margin-bottom: 24rpx;
	}

	.link {
		color: #ff8200;
		margin: 0 4rpx;
	}

	.next-btn {
		width: 100%;
		background: #ff8200;
		color: #fff;
		font-size: 32rpx;
		border-radius: 16rpx;
		padding: 24rpx 0;
		margin-bottom: 16rpx;
		text-align: center;
		transition: background 0.2s;
	}

	.next-btn.disabled {
		background: #ddd;
		color: #aaa;
		pointer-events: none;
	}

	.login-icons {
		display: flex;
		justify-content: center;
		align-items: center;
		gap: 56rpx;
		margin: 120rpx 0 0 0;
	}

	.login-icon-item {
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	.login-icon-img-wrap {
		width: 72rpx;
		height: 72rpx;
		border-radius: 50%;
		background: #f7f7f7;
		display: flex;
		align-items: center;
		justify-content: center;
		box-shadow: 0 4rpx 16rpx rgba(0,0,0,0.10);
		margin-bottom: 10rpx;
	}

	.login-icon-img {
		width: 48rpx;
		height: 48rpx;
		border-radius: 50%;
		background: transparent;
	}

	.login-icon-text {
		font-size: 22rpx;
		color: #333;
	}

	.qrcode-modal {
		position: fixed;
		left: 0;
		top: 0;
		width: 100vw;
		height: 100vh;
		z-index: 9999;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.qrcode-modal-mask {
		position: absolute;
		left: 0;
		top: 0;
		width: 100vw;
		height: 100vh;
		background: rgba(0,0,0,0.35);
	}

	.qrcode-modal-content {
		position: relative;
		background: #fff;
		border-radius: 18rpx;
		padding: 48rpx 32rpx 32rpx 32rpx;
		z-index: 10;
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	.qrcode-img {
		width: 320rpx;
		height: 320rpx;
		margin-bottom: 32rpx;
		background: #f7f7f7;
		border-radius: 12rpx;
	}

	.qrcode-close-btn {
		color: #ff8200;
		font-size: 30rpx;
		margin-top: 8rpx;
		padding: 12rpx 36rpx;
		border-radius: 8rpx;
		background: #fff0e1;
		box-shadow: 0 2rpx 8rpx rgba(0,0,0,0.06);
		cursor: pointer;
	}
</style>
