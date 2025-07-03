<template>
  <view class="container">
    <view class="nav-bar">
      <text class="back" @click="goBack">&#8592;</text>
      <text class="nav-title">验证码</text>
    </view>
    <view class="main-content">
      <text class="title">请输入验证码</text>
      <text class="subtitle">验证码已发送至 {{ phone }}</text>
      <view class="code-box">
        <input
          class="code-input-real"
          type="number"
          inputmode="numeric"
          maxlength="6"
          v-model="code"
          @input="onInput"
          :focus="focus"
        />
        <view class="code-item" v-for="(item, idx) in 6" :key="idx">
          <text v-if="code[idx]">{{ code[idx] }}</text>
        </view>
      </view>
      <view class="countdown">
        <text v-if="countdown > 0">{{ countdown }}s 后重新发送</text>
        <text v-else class="resend" @click="resendCode">重新发送</text>
      </view>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {
      phone: '', // 从页面参数获取
      code: '',
      countdown: 60,
      timer: null,
      focus: false
    };
  },
  onLoad(options) {
    this.phone = options.phone || '';
    this.sendCode(); // 页面加载时自动发送验证码
    this.startCountdown();
    setTimeout(() => {
      this.focus = true;
    }, 200);
  },
  beforeDestroy() {
    clearInterval(this.timer);
  },
  methods: {
    goBack() {
      uni.navigateBack();
    },
    onInput(e) {
      this.code = e.target.value.slice(0, 6);
      if (this.code.length === 6) {
        this.login();
      }
    },
    login() {
      uni.request({
        url: 'http://127.0.0.1:8080/user/login',
        method: 'POST',
        data: {
          mobile: this.phone.replace(/\s/g, ''),
          SendSmsCode: this.code
        },
        success: (res) => {
          console.log('login返回', res);
          console.log('code类型', typeof res.data.code, '值', res.data.code);
          if (res.data.code == 200 || res.data.code == "200") {
            uni.setStorageSync('token', res.data.token);
            console.log('准备跳转首页', res.data);
            uni.reLaunch({
              url: '/pages/index/didi',
              fail: (err) => {
                console.log('跳转失败', err);
                uni.showToast({ title: '跳转失败', icon: 'none' });
              }
            });
          } else {
            console.log('登录失败分支', res.data);
            uni.showToast({
              title: res.data.message || '验证码错误',
              icon: 'none'
            });
            this.code = '';
          }
        },
        fail: (err) => {
          console.log('请求失败', err);
          uni.showToast({
            title: '网络错误',
            icon: 'none'
          });
          this.code = '';
        }
      });
    },
    startCountdown() {
      this.countdown = 60;
      this.timer && clearInterval(this.timer);
      this.timer = setInterval(() => {
        if (this.countdown > 0) {
          this.countdown--;
        } else {
          clearInterval(this.timer);
        }
      }, 1000);
    },
    sendCode() {
      if (!this.phone) {
        uni.showToast({ title: '手机号缺失', icon: 'none' });
        return;
      }
      uni.request({
        url: 'http://127.0.0.1:8080/user/sendSms',
        method: 'POST',
        data: {
          mobile: this.phone.replace(/\s/g, '')
        },
        success: (res) => {
          if (res.data.code === 200) {
            uni.showToast({ title: '验证码已发送', icon: 'success' });
          } else {
            uni.showToast({ title: res.data.message || '发送失败', icon: 'none' });
          }
        },
        fail: () => {
          uni.showToast({ title: '网络错误', icon: 'none' });
        }
      });
    },
    resendCode() {
      this.sendCode();
      this.startCountdown();
    }
  }
};
</script>

<style scoped>
.container {
  background: #fff;
  min-height: 100vh;
  box-sizing: border-box;
}
.nav-bar {
  display: flex;
  align-items: center;
  height: 88rpx;
  border-bottom: 1px solid #f0f0f0;
  position: relative;
}
.back {
  font-size: 40rpx;
  color: #333;
  margin-left: 24rpx;
  margin-right: 24rpx;
  cursor: pointer;
}
.nav-title {
  flex: 1;
  text-align: center;
  font-size: 36rpx;
  font-weight: bold;
  color: #222;
  margin-right: 64rpx;
}
.main-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 60rpx;
  padding: 0 5vw;
}
.title {
  font-size: 44rpx;
  font-weight: bold;
  margin-bottom: 16rpx;
  color: #222;
}
.subtitle {
  font-size: 28rpx;
  color: #888;
  margin-bottom: 40rpx;
}
.code-box {
  display: flex;
  justify-content: center;
  align-items: center;
  position: relative;
  width: 90vw;
  max-width: 420rpx;
  height: 80rpx;
  margin-bottom: 24rpx;
  background: #fff;
  border-radius: 24rpx;
  box-shadow: 0 4rpx 24rpx rgba(0,0,0,0.04);
}
.code-input-real {
  position: absolute;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  opacity: 0;
  z-index: 2;
  caret-color: #108ee9;
}
.code-item {
  width: 48rpx;
  height: 56rpx;
  margin: 0 10rpx;
  background: #fff;
  border-radius: 12rpx;
  border: 1.5rpx solid #eee;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 36rpx;
  color: #222;
  box-shadow: 0 2rpx 8rpx rgba(0,0,0,0.03);
  z-index: 1;
}
.countdown {
  width: 90vw;
  max-width: 420rpx;
  text-align: right;
  color: #888;
  font-size: 26rpx;
  margin-bottom: 24rpx;
}
.resend {
  color: #108ee9;
  cursor: pointer;
}
.didi-container {
  background: #f7f8fa;
  min-height: 100vh;
  padding: 0 16px 100px 16px;
  box-sizing: border-box;
  max-width: 430px; /* iPhone 14 Pro Max 竖屏宽度 */
  margin: 0 auto;
}
.didi-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 56px;
  margin-top: 12px;
  margin-bottom: 8px;
}
.didi-title {
  font-size: 22px;
  font-weight: bold;
  flex: 1;
  text-align: center;
}
.didi-header-btns {
  display: flex;
  align-items: center;
  gap: 8px;
}
.didi-header-btn {
  background: #fff;
  border: none;
  border-radius: 22px;
  padding: 0 14px;
  font-size: 15px;
  height: 36px;
  display: flex;
  align-items: center;
  box-shadow: 0 2px 8px rgba(0,0,0,0.04);
}
.didi-header-btn.icon-only {
  width: 36px;
  padding: 0;
  justify-content: center;
}
.didi-header-icon {
  width: 20px;
  height: 20px;
}
#amap-container {
  width: 100%;
  height: 180px;
  border-radius: 16px;
  overflow: hidden;
  margin: 12px 0 16px 0;
  box-shadow: 0 2px 12px rgba(0,0,0,0.06);
}
</style> 