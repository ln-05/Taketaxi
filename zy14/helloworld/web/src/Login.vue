<template>
  <div class="container">
    <div class="header">
      <h2>嗨，欢迎回来</h2>
      <p class="desc">登录后更精彩，美好出行即将开始</p>
    </div>
    <div class="form">
      <div class="input-group">
        <span class="country-code">+86</span>
        <input type="tel" maxlength="11" placeholder="198 5152 6101" v-model="mobile" class="mobile-input" />
      </div>
      <div class="protocol">
        <input type="checkbox" id="protocol" v-model="checked" />
        <label for="protocol">
          阅读并同意 <a href="#" class="protocol-link">服务协议及滴滴出行基本功能个人信息处理规则</a>
        </label>
      </div>
      <button class="next-btn" :disabled="!checked" @click="goCaptcha">下一步</button>
      <div class="login-problem">登录遇到问题</div>
    </div>
    <div class="login-bottom">
      <img src="./assets/wechat.svg" alt="微信登录" class="wechat-icon" @click="showQr = true" />
    </div>
    <div v-if="showQr" class="qr-modal" @click.self="showQr = false">
      <div class="qr-content">
        <img src="./assets/qr-placeholder.png" alt="微信扫码登录" class="qr-img" />
        <div class="qr-tip">请使用微信扫码登录</div>
        <button class="close-btn" @click="showQr = false">关闭</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
const mobile = ref('');
const checked = ref(false);
const showQr = ref(false);
const router = useRouter();
const goCaptcha = () => {
  if (checked.value) {
    router.push({ path: '/captcha', query: { phone: mobile.value || '198 5152 6101' } });
  }
};
</script>

<style scoped>
/* 样式同原App.vue，可直接复制 */
.container {
  max-width: 400px;
  margin: 0 auto;
  padding: 32px 16px 0 16px;
  font-family: 'PingFang SC', 'Microsoft YaHei', Arial, sans-serif;
}
.header {
  text-align: left;
  margin-bottom: 32px;
}
.header h2 {
  font-size: 2rem;
  margin: 0 0 8px 0;
  font-weight: 700;
}
.desc {
  color: #888;
  font-size: 1rem;
}
.form {
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.04);
  padding: 24px 16px 16px 16px;
  margin-bottom: 32px;
}
.input-group {
  display: flex;
  align-items: center;
  margin-bottom: 16px;
  border: 1px solid #eee;
  border-radius: 8px;
  padding: 8px 12px;
  background: #fafbfc;
}
.country-code {
  color: #222;
  font-size: 1rem;
  margin-right: 8px;
}
.mobile-input {
  border: none;
  outline: none;
  font-size: 1rem;
  flex: 1;
  background: transparent;
}
.protocol {
  font-size: 0.95rem;
  margin-bottom: 16px;
  color: #888;
}
.protocol-link {
  color: #ff8800;
  text-decoration: underline;
}
.next-btn {
  width: 100%;
  background: #13294b;
  color: #fff;
  border: none;
  border-radius: 8px;
  padding: 12px 0;
  font-size: 1.1rem;
  font-weight: 600;
  margin-bottom: 12px;
  cursor: pointer;
  transition: background 0.2s;
}
.next-btn:disabled {
  background: #ccc;
  cursor: not-allowed;
}
.login-problem {
  text-align: right;
  color: #888;
  font-size: 0.95rem;
  margin-top: 4px;
}
.login-bottom {
  display: flex;
  justify-content: center;
  margin-top: 32px;
}
.wechat-icon {
  width: 48px;
  height: 48px;
  cursor: pointer;
  border-radius: 50%;
  box-shadow: 0 2px 8px rgba(0,0,0,0.08);
  background: #fff;
  padding: 8px;
  transition: box-shadow 0.2s;
}
.wechat-icon:hover {
  box-shadow: 0 4px 16px rgba(0,0,0,0.12);
}
.qr-modal {
  position: fixed;
  left: 0; top: 0; right: 0; bottom: 0;
  background: rgba(0,0,0,0.35);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}
.qr-content {
  background: #fff;
  border-radius: 12px;
  padding: 32px 24px 24px 24px;
  text-align: center;
  box-shadow: 0 4px 24px rgba(0,0,0,0.12);
}
.qr-img {
  width: 180px;
  height: 180px;
  margin-bottom: 16px;
}
.qr-tip {
  font-size: 1.1rem;
  color: #222;
  margin-bottom: 16px;
}
.close-btn {
  background: #13294b;
  color: #fff;
  border: none;
  border-radius: 8px;
  padding: 8px 24px;
  font-size: 1rem;
  cursor: pointer;
}
</style> 