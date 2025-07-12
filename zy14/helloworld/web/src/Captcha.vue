<template>
  <div class="captcha-container">
    <div class="captcha-header">
      <h2>请输入验证码</h2>
      <p class="desc">验证码已发送至 <span class="phone">+86 {{ phone }}</span></p>
    </div>
    <div class="captcha-inputs">
      <input v-for="(v, i) in 6" :key="i" maxlength="1" type="tel" class="captcha-box" v-model="codes[i]" @input="onInput(i, $event)" />
    </div>
    <div class="captcha-actions">
      <span class="resend" v-if="countdown === 0" @click="resend">重新发送</span>
      <span class="countdown" v-else>{{ countdown }}s 后重新发送</span>
      <span class="password-login">使用密码登录</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const route = useRoute();
const router = useRouter();
const phone = route.query.phone || '198 5152 6101';
const codes = ref<string[]>(['', '', '', '', '', '']);
const countdown = ref(60);
let timer: any = null;

const startCountdown = () => {
  countdown.value = 60;
  timer = setInterval(() => {
    if (countdown.value > 0) {
      countdown.value--;
    } else {
      clearInterval(timer);
    }
  }, 1000);
};

const resend = () => {
  startCountdown();
};

const onInput = (i: number, e: Event) => {
  const target = e.target as HTMLInputElement;
  if (target.value && i < 5) {
    const next = document.querySelectorAll<HTMLInputElement>('.captcha-box')[i + 1];
    next && next.focus();
  }
  // 检查是否全部输入完毕
  if (codes.value.every(c => c.length === 1)) {
    setTimeout(() => {
      router.push('/home');
    }, 200);
  }
};

onMounted(() => {
  startCountdown();
});
</script>

<style scoped>
.captcha-container {
  max-width: 400px;
  margin: 0 auto;
  padding: 32px 16px 0 16px;
  font-family: 'PingFang SC', 'Microsoft YaHei', Arial, sans-serif;
}
.captcha-header {
  margin-bottom: 32px;
}
.captcha-header h2 {
  font-size: 2rem;
  font-weight: 700;
  margin-bottom: 8px;
}
.captcha-header .desc {
  color: #888;
  font-size: 1rem;
}
.captcha-header .phone {
  color: #222;
  font-weight: 600;
}
.captcha-inputs {
  display: flex;
  justify-content: space-between;
  margin: 32px 0 24px 0;
}
.captcha-box {
  width: 48px;
  height: 48px;
  border: 1.5px solid #ff8800;
  border-radius: 8px;
  text-align: center;
  font-size: 2rem;
  outline: none;
  margin: 0 4px;
  background: #fff;
  box-shadow: 0 2px 8px rgba(0,0,0,0.04);
}
.captcha-actions {
  display: flex;
  justify-content: space-between;
  align-items: center;
  color: #888;
  font-size: 1rem;
}
.captcha-actions .resend {
  color: #ff8800;
  cursor: pointer;
}
.captcha-actions .countdown {
  color: #888;
}
.captcha-actions .password-login {
  color: #888;
  cursor: pointer;
}
</style> 