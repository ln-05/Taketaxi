<template>
  <div class="verify-code-container">
    <!-- 头部返回按钮 -->
    <div class="header">
      <div class="back-btn" @click="goBack">
        <el-icon><ArrowLeft /></el-icon>
      </div>
    </div>
    
    <!-- 标题区域 -->
    <div class="title-section">
      <h1>请输入验证码</h1>
      <p class="phone-info">验证码已发送至 +86 {{ formatPhone(phone) }}</p>
    </div>
    
    <!-- 验证码输入区域 -->
    <div class="code-section">
      <div class="code-inputs">
        <input
          v-for="(_, index) in codeDigits"
          :key="index"
          :ref="el => setInputRef(el, index)"
          v-model="codeDigits[index]"
          type="tel"
          maxlength="1"
          class="code-input"
          :class="{ 
            focused: focusedIndex === index,
            filled: codeDigits[index] !== ''
          }"
          @input="onInput(index, $event)"
          @keydown="onKeydown(index, $event)"
          @focus="onFocus(index)"
          @blur="onBlur"
        />
      </div>
    </div>
    
    <!-- 重新发送倒计时 -->
    <div class="resend-section">
      <div class="resend-content">
        <span v-if="countdown > 0" class="countdown-text">
          {{ countdown }}s后重新发送
        </span>
        <button v-else class="resend-btn" @click="resendCode" :disabled="resending">
          {{ resending ? '发送中...' : '重新发送' }}
        </button>
      </div>
    </div>
    
    <!-- 登录按钮 -->
    <div class="login-section">
      <button 
        class="login-btn"
        :class="{ disabled: !isCodeComplete || logging }"
        :disabled="!isCodeComplete || logging"
        @click="submitLogin"
      >
        {{ logging ? '登录中...' : '登录' }}
      </button>
    </div>
    
    <!-- 短信发送成功提示 -->
    <div v-if="showSmsSuccess" class="sms-success-toast">
      <div class="success-content">
        <el-icon class="success-icon"><Check /></el-icon>
        <span>短信已发送</span>
      </div>
    </div>
    
    <!-- 调试信息 -->
    <div class="debug-section" v-if="isDev">
      <el-button type="warning" size="small" @click="fillTestCode">
        填入测试验证码
      </el-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted, nextTick } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { ElMessage, ElIcon } from 'element-plus'
import { ArrowLeft, Check } from '@element-plus/icons-vue'
import { useUserStore } from '@/store'
import { userApi } from '@/api'

const router = useRouter()
const route = useRoute()
const userStore = useUserStore()

// 获取路由参数
const phone = ref(route.query.phone as string || '')

// 验证码相关状态
const codeDigits = reactive(['', '', '', '', '', ''])
const inputRefs = ref<(HTMLInputElement | null)[]>([])
const focusedIndex = ref(-1)
const logging = ref(false)
const resending = ref(false)
const showSmsSuccess = ref(false)

// 倒计时
const countdown = ref(60)
const countdownTimer = ref<NodeJS.Timeout | null>(null)

// 开发模式
const isDev = ref(process.env.NODE_ENV === 'development')

// 计算属性
const isCodeComplete = computed(() => {
  return codeDigits.every(digit => digit.length === 1)
})

const verificationCode = computed(() => {
  return codeDigits.join('')
})

// 设置输入框引用
const setInputRef = (el: any, index: number) => {
  if (el) {
    inputRefs.value[index] = el as HTMLInputElement
  }
}

// 格式化手机号显示
const formatPhone = (phoneNumber: string) => {
  if (!phoneNumber) return ''
  // 将手机号格式化为 181 6863 5937 的形式
  return phoneNumber.replace(/(\d{3})(\d{4})(\d{4})/, '$1 $2 $3')
}

// 返回上一页
const goBack = () => {
  router.go(-1)
}

// 处理输入
const onInput = (index: number, event: Event) => {
  const target = event.target as HTMLInputElement
  const value = target.value.replace(/\D/g, '') // 只允许数字
  
  if (value.length > 1) {
    // 如果粘贴了多位数字，分配到各个输入框
    const digits = value.split('').slice(0, 6)
    digits.forEach((digit, i) => {
      if (index + i < 6) {
        codeDigits[index + i] = digit
      }
    })
    
    // 移动焦点到最后一个填充的输入框或下一个空的输入框
    const nextIndex = Math.min(index + digits.length, 5)
    nextTick(() => {
      inputRefs.value[nextIndex]?.focus()
    })
  } else {
    codeDigits[index] = value
    
    // 自动跳转到下一个输入框
    if (value && index < 5) {
      nextTick(() => {
        inputRefs.value[index + 1]?.focus()
      })
    }
  }
  
  // 如果验证码填完了，自动提交
  if (isCodeComplete.value && verificationCode.value.length === 6) {
    setTimeout(() => {
      submitLogin()
    }, 500)
  }
}

// 处理键盘事件
const onKeydown = (index: number, event: KeyboardEvent) => {
  if (event.key === 'Backspace') {
    if (!codeDigits[index] && index > 0) {
      // 如果当前输入框为空且按退格键，跳转到前一个输入框
      codeDigits[index - 1] = ''
      nextTick(() => {
        inputRefs.value[index - 1]?.focus()
      })
    } else {
      codeDigits[index] = ''
    }
  } else if (event.key === 'ArrowLeft' && index > 0) {
    inputRefs.value[index - 1]?.focus()
  } else if (event.key === 'ArrowRight' && index < 5) {
    inputRefs.value[index + 1]?.focus()
  }
}

// 处理焦点
const onFocus = (index: number) => {
  focusedIndex.value = index
}

const onBlur = () => {
  focusedIndex.value = -1
}

// 重新发送验证码
const resendCode = async () => {
  if (!phone.value) {
    ElMessage.error('手机号码不能为空')
    return
  }
  
  try {
    resending.value = true
    
    console.log('重新发送验证码到:', phone.value)
    
    const response = await userApi.sendSms({
      phone: phone.value,
      source: 'Login'
    })
    
    console.log('重发短信响应:', response)
    
    if (response.code === 200) {
      ElMessage.success('验证码已重新发送')
      showSmsSuccessToast()
      startCountdown()
      
      // 清空当前验证码
      codeDigits.forEach((_, index) => {
        codeDigits[index] = ''
      })
      
      // 焦点回到第一个输入框
      nextTick(() => {
        inputRefs.value[0]?.focus()
      })
    } else {
      ElMessage.error(response.msg || '发送失败，请重试')
    }
  } catch (error) {
    console.error('重发验证码失败:', error)
    ElMessage.error('发送失败，请检查网络连接')
  } finally {
    resending.value = false
  }
}

// 显示短信发送成功提示
const showSmsSuccessToast = () => {
  showSmsSuccess.value = true
  setTimeout(() => {
    showSmsSuccess.value = false
  }, 3000)
}

// 开始倒计时
const startCountdown = () => {
  if (countdownTimer.value) {
    clearInterval(countdownTimer.value)
  }
  
  countdown.value = 60
  countdownTimer.value = setInterval(() => {
    countdown.value--
    if (countdown.value <= 0) {
      clearInterval(countdownTimer.value!)
      countdownTimer.value = null
    }
  }, 1000)
}

// 提交登录
const submitLogin = async () => {
  if (!isCodeComplete.value) {
    ElMessage.warning('请输入完整的验证码')
    return
  }
  
  if (!phone.value) {
    ElMessage.error('手机号码不能为空')
    return
  }
  
  try {
    logging.value = true
    
    console.log('开始验证码登录:', {
      phone: phone.value,
      code: verificationCode.value
    })
    
    const result = await userStore.login({
      phone: phone.value,
      sendSms: verificationCode.value
    })
    
    console.log('登录结果:', result)
    
    if (result.success) {
      ElMessage.success('登录成功！正在跳转...')
      
      // 等待状态更新
      await nextTick()
      
      // 跳转到首页
      console.log('跳转到滴滴首页')
      await router.replace('/home')
    } else {
      ElMessage.error(result.message || '验证码错误，请重试')
      
      // 清空验证码，重新输入
      codeDigits.forEach((_, index) => {
        codeDigits[index] = ''
      })
      
      nextTick(() => {
        inputRefs.value[0]?.focus()
      })
    }
  } catch (error) {
    console.error('登录失败:', error)
    ElMessage.error('登录失败，请重试')
    
    // 清空验证码
    codeDigits.forEach((_, index) => {
      codeDigits[index] = ''
    })
    
    nextTick(() => {
      inputRefs.value[0]?.focus()
    })
  } finally {
    logging.value = false
  }
}

// 填入测试验证码（开发用）
const fillTestCode = () => {
  const testCode = '123456'
  testCode.split('').forEach((digit, index) => {
    codeDigits[index] = digit
  })
  
  nextTick(() => {
    inputRefs.value[5]?.focus()
  })
}

// 组件挂载时的初始化
onMounted(() => {
  // 检查是否有手机号
  if (!phone.value) {
    ElMessage.error('缺少手机号参数')
    router.replace('/login')
    return
  }
  
  // 自动聚焦第一个输入框
  nextTick(() => {
    inputRefs.value[0]?.focus()
  })
  
  // 开始倒计时
  startCountdown()
  
  // 显示短信发送成功提示
  showSmsSuccessToast()
  
  console.log('验证码页面初始化，手机号:', phone.value)
})

// 组件销毁时清理定时器（备用函数）
// const cleanup = () => {
//   if (countdownTimer.value) {
//     clearInterval(countdownTimer.value)
//     countdownTimer.value = null
//   }
// }
</script>

<style lang="scss" scoped>
.verify-code-container {
  height: 100vh;
  background: #ffffff;
  display: flex;
  flex-direction: column;
  padding: 0;
  overflow: hidden;
  position: relative;
}

/* 头部区域 */
.header {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 60px;
  display: flex;
  align-items: center;
  padding: 0 20px;
  z-index: 10;
  
  .back-btn {
    width: 44px;
    height: 44px;
    display: flex;
    align-items: center;
    justify-content: center;
    color: #333;
    cursor: pointer;
    border-radius: 50%;
    transition: background-color 0.2s;
    
    &:hover {
      background-color: rgba(0, 0, 0, 0.05);
    }
    
    .el-icon {
      font-size: 20px;
    }
  }
}

/* 标题区域 */
.title-section {
  margin-top: 100px;
  padding: 30px 30px 80px;
  
  h1 {
    font-size: 24px;
    color: #333;
    font-weight: 600;
    margin: 0 0 20px 0;
  }
  
  .phone-info {
    font-size: 16px;
    color: #666;
    margin: 0;
    line-height: 1.4;
  }
}

/* 验证码输入区域 */
.code-section {
  padding: 0 30px;
  margin-bottom: 50px;
  
  .code-inputs {
    display: flex;
    justify-content: center;
    gap: 15px;
    
    .code-input {
      width: 50px;
      height: 60px;
      border: 1px solid #ddd;
      border-radius: 6px;
      text-align: center;
      font-size: 28px;
      font-weight: normal;
      color: #333;
      background: #fff;
      outline: none;
      transition: all 0.3s ease;
      
      &:focus, &.focused {
        border-color: #ff6600;
        border-width: 2px;
        box-shadow: none;
      }
      
      &.filled {
        border-color: #ff6600;
        color: #ff6600;
      }
      
      &::placeholder {
        color: transparent;
      }
    }
  }
}

/* 重新发送区域 */
.resend-section {
  padding: 0 30px;
  margin-bottom: 40px;
  
  .resend-content {
    display: flex;
    justify-content: flex-end;
    
    .countdown-text {
      color: #999;
      font-size: 16px;
    }
    
    .resend-btn {
      background: none;
      border: none;
      color: #ff6600;
      font-size: 16px;
      cursor: pointer;
      text-decoration: underline;
      padding: 0;
      
      &:hover:not(:disabled) {
        color: #e55a00;
      }
      
      &:disabled {
        color: #ccc;
        cursor: not-allowed;
        text-decoration: none;
      }
    }
  }
}

/* 登录按钮 */
.login-section {
  padding: 0 20px;
  margin-bottom: 40px;
  
  .login-btn {
    width: 100%;
    height: 50px;
    background: #ff6600;
    color: white;
    border: none;
    border-radius: 25px;
    font-size: 16px;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.2s;
    
    &:hover:not(.disabled) {
      background: #e55a00;
    }
    
    &.disabled {
      background: #ddd;
      color: #999;
      cursor: not-allowed;
    }
  }
}

/* 短信发送成功提示 */
.sms-success-toast {
  position: fixed;
  bottom: 200px;
  left: 50%;
  transform: translateX(-50%);
  background: #6b7280;
  color: white;
  padding: 20px 40px;
  border-radius: 12px;
  z-index: 1000;
  animation: slideUp 0.3s ease;
  min-width: 200px;
  
  .success-content {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 12px;
    font-size: 16px;
    
    .success-icon {
      width: 24px;
      height: 24px;
      background: white;
      border-radius: 50%;
      display: flex;
      align-items: center;
      justify-content: center;
      color: #4caf50;
      font-size: 14px;
    }
  }
}

@keyframes slideUp {
  from {
    transform: translateX(-50%) translateY(20px);
    opacity: 0;
  }
  to {
    transform: translateX(-50%) translateY(0);
    opacity: 1;
  }
}

/* 调试区域 */
.debug-section {
  position: absolute;
  bottom: 20px;
  right: 20px;
}

/* 移动端适配 */
@media (max-width: 768px) {
  .title-section {
    padding: 20px 20px 60px;
    margin-top: 80px;
    
    h1 {
      font-size: 22px;
    }
    
    .phone-info {
      font-size: 15px;
    }
  }
  
  .code-section {
    padding: 0 20px;
    margin-bottom: 40px;
    
    .code-inputs {
      gap: 12px;
      
      .code-input {
        width: 45px;
        height: 55px;
        font-size: 24px;
      }
    }
  }
  
  .resend-section {
    padding: 0 20px;
    margin-bottom: 30px;
  }
}

/* 安全区域适配 */
@supports (padding: max(0px)) {
  .verify-code-container {
    padding-bottom: max(20px, env(safe-area-inset-bottom));
  }
}
</style> 