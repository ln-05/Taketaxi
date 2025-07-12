<template>
  <div class="didi-login-container">
    <!-- 头部返回按钮 -->
    <div class="header">
      <div class="back-btn" @click="goBack">
        <el-icon><ArrowLeft /></el-icon>
      </div>
    </div>
    
    <!-- 滴滴Logo -->
    <div class="logo-section">
      <div class="didi-logo">
        <div class="logo-icon">滴</div>
      </div>
    </div>
    
    <!-- 标题区域 -->
    <div class="title-section">
      <h1>登录后更精彩</h1>
      <p>美好出行即将开始</p>
    </div>
    
    <!-- 手机号输入区域 -->
    <div class="phone-section">
      <div class="phone-input-container" :class="{ 'error': !phoneValidation.isValid && loginForm.phone }">
        <div class="country-code">
          <span>+86</span>
          <el-icon><ArrowDown /></el-icon>
        </div>
        <div class="phone-input">
          <input 
            v-model="loginForm.phone"
            type="tel"
            placeholder="请输入手机号"
            maxlength="11"
          />
        </div>
      </div>
      <div class="validation-message" v-if="loginForm.phone && !phoneValidation.isValid">
        {{ phoneValidation.message }}
      </div>
    </div>
    
    <!-- 服务协议 -->
    <div class="agreement-section">
      <label class="agreement-checkbox">
        <input type="checkbox" v-model="agreedToTerms" />
        <span class="checkmark"></span>
        <span class="agreement-text">
          阅读并同意 
          <span class="link">服务协议及滴滴出行基本功能个人信息处理规则</span>
        </span>
      </label>
    </div>
    
    <!-- 下一步按钮 -->
    <div class="next-section">
      <button 
        class="next-btn"
        :class="{ disabled: !canProceed }"
        :disabled="!canProceed"
        @click="nextStep"
      >
        下一步
      </button>
    </div>
    
    <!-- 登录问题 -->
    <div class="help-section">
      <span class="help-link" @click="showHelp">登录遇到问题</span>
    </div>
    
    <!-- 其他登录方式 -->
    <div class="other-login-section">
      <div class="login-methods">
        <div class="login-method" @click="loginWithAlipay">
          <div class="method-icon alipay">支</div>
        </div>
        <div class="login-method" @click="loginWithWechat">
          <div class="method-icon wechat">微信</div>
        </div>
        <div class="login-method" @click="loginWithId">
          <div class="method-icon id-auth">身份认证</div>
        </div>
        <div class="login-method" @click="quickLogin">
          <div class="method-icon phone">📱</div>
        </div>
      </div>
    </div>
    
    <!-- 调试按钮 (开发时使用) -->
    <div class="debug-section" v-if="isDev">
      <el-button 
        type="warning" 
        size="small"
        @click="debugInfo"
      >
        🔍 调试信息
      </el-button>
    </div>
  </div>

  <!-- 微信登录弹窗 -->
  <el-dialog
    v-model="wechatDialogVisible"
    title="微信扫码登录"
    width="400px"
    :close-on-click-modal="false"
    :close-on-press-escape="false"
    @close="stopPolling"
  >
    <div class="qr-container">
      <div v-if="loading" class="loading-container">
        <el-icon class="is-loading" size="40"><Loading /></el-icon>
        <p>正在生成登录二维码...</p>
      </div>
      
      <div v-else-if="qrcodeUrl" class="qr-code">
        <img :src="qrcodeUrl" alt="微信登录二维码" />
        <div class="qr-tips">
          <p>请使用微信扫描二维码登录</p>
          <p class="sub-tip">扫码后在手机上确认登录</p>
        </div>
        
        <div v-if="loginStatus === 'scanned'" class="scan-success">
          <el-icon size="40" color="#67c23a"><SuccessFilled /></el-icon>
          <p>扫码成功！请在手机上确认</p>
        </div>
        
        <div v-if="loginStatus === 'expired'" class="qr-expired">
          <p>二维码已过期</p>
          <el-button type="primary" @click="refreshQRCode">刷新二维码</el-button>
        </div>
      </div>
      
      <div v-else class="error-container">
        <el-icon size="40" color="#f56c6c"><WarningFilled /></el-icon>
        <p>二维码生成失败</p>
        <el-button type="primary" @click="generateQRCode">重新生成</el-button>
      </div>
    </div>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, reactive, nextTick, computed, watch, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ArrowLeft, ArrowDown } from '@element-plus/icons-vue'
import { useUserStore } from '@/store'
import { userApi, wechatApi } from '@/api'
import { Loading, SuccessFilled, WarningFilled } from '@element-plus/icons-vue'

const router = useRouter()
const userStore = useUserStore()

// 表单数据
const loginForm = reactive({
  phone: '',
  code: ''
})

// 表单验证规则
const loginRules = {
  phone: [
    { required: true, message: '请输入手机号', trigger: 'blur' },
    { pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号', trigger: 'blur' }
  ],
  code: [
    { required: true, message: '请输入验证码', trigger: 'blur' },
    { pattern: /^\d{4}$/, message: '验证码为4位数字', trigger: 'blur' }
  ]
}

// 手机号验证状态
const phoneValidation = reactive({
  isValid: false,
  message: ''
})

// 验证手机号
const validatePhone = (phone: string) => {
  if (!phone) {
    phoneValidation.isValid = false
    phoneValidation.message = '请输入手机号'
    return
  }
  
  if (!/^1[3-9]\d{9}$/.test(phone)) {
    phoneValidation.isValid = false
    phoneValidation.message = '请输入正确的手机号'
    return
  }
  
  phoneValidation.isValid = true
  phoneValidation.message = ''
}

// 监听手机号变化
watch(() => loginForm.phone, (newPhone) => {
  validatePhone(newPhone)
})

// 状态
const loading = ref(false)
const smsDisabled = ref(false)
const smsCountdown = ref(0)
const smsButtonText = ref('获取验证码')
const loginFormRef = ref()
const isDev = ref(process.env.NODE_ENV === 'development')
const agreedToTerms = ref(false)

// 计算属性：是否可以进行下一步
const canProceed = computed(() => {
  return phoneValidation.isValid && agreedToTerms.value
})

// 发送短信验证码
const sendSms = async () => {
  if (!loginForm.phone || !/^1[3-9]\d{9}$/.test(loginForm.phone)) {
    ElMessage.error('请输入正确的手机号')
    return
  }
  
  try {
    loading.value = true
    const response = await userApi.sendSms({
      Phone: loginForm.phone,
      Source: 'Login'
    })
    
    ElMessage.success('验证码发送成功')
    startSmsCountdown()
  } catch (error) {
    console.error('发送验证码失败:', error)
    ElMessage.error('发送验证码失败，请稍后重试')
  } finally {
    loading.value = false
  }
}

// 开始短信倒计时
const startSmsCountdown = () => {
  smsDisabled.value = true
  smsCountdown.value = 60
  
  const timer = setInterval(() => {
    smsCountdown.value--
    smsButtonText.value = `${smsCountdown.value}s后重试`
    
    if (smsCountdown.value <= 0) {
      clearInterval(timer)
      smsDisabled.value = false
      smsButtonText.value = '获取验证码'
    }
  }, 1000)
}

// 统一的登录后跳转处理
const handleLoginSuccess = async () => {
  console.log('登录成功，准备跳转到仿滴滴首页')
  ElMessage.success('登录成功！正在跳转到首页...')
  
  // 等待下一个tick确保Store状态已更新
  await nextTick()
  
  try {
    console.log('当前登录状态:', userStore.isLoggedIn)
    console.log('即将跳转到仿滴滴首页 (/)')
    
    // 确保跳转到首页
    await router.replace('/')
    console.log('路由跳转成功，已到达仿滴滴首页')
  } catch (error) {
    console.error('路由跳转失败:', error)
    // 备用方案：强制跳转
    window.location.href = '/'
  }
}

// 正常登录处理
const handleLogin = async () => {
  if (!loginFormRef.value) return
  
  try {
    const valid = await loginFormRef.value.validate()
    if (!valid) return
    
    loading.value = true
    
    const response = await userApi.login({
      Phone: loginForm.phone,
      SendSms: loginForm.code
    })
    
    if (response && response.data) {
      userStore.user = {
        id: response.data.ID || '',
        phone: response.data.Phone || loginForm.phone,
        name: response.data.Name || '',
        avatar: response.data.Avatar || ''
      }
      userStore.token = response.data.Token || ''
      userStore.loginTime = Date.now()
      await handleLoginSuccess()
    } else {
      ElMessage.error('登录失败，返回数据异常')
    }
  } catch (error) {
    console.error('登录过程出错:', error)
    ElMessage.error('登录失败，请重试')
  } finally {
    loading.value = false
  }
}

// 新增：安全的调试信息
const debugInfo = () => {
  if (!isDev.value) return
  
  const safeInfo = {
    'Store状态': {
      isLoggedIn: userStore.isLoggedIn,
      loginTime: userStore.loginTime
    },
    '当前路由': router.currentRoute.value.path,
    'App组件状态': {
      isLoginPage: router.currentRoute.value.path === '/login',
      shouldShowLayout: router.currentRoute.value.path !== '/login'
    }
  }
  
  console.log('=== 调试信息 ===', safeInfo)
  ElMessage.info('调试信息已输出到控制台')
  
  // 额外检查Layout组件是否存在
  const layoutElements = document.querySelectorAll('.mobile-layout, .bottom-navigation')
  console.log('页面中的Layout元素:', {
    mobileLayout: layoutElements.length > 0 ? '找到' : '未找到',
    bottomNavigation: document.querySelector('.bottom-navigation') ? '找到' : '未找到',
    navItems: document.querySelectorAll('.nav-item').length + '个导航项'
  })
}

// 显示帮助信息
const showHelp = () => {
  ElMessage.info('登录遇到问题？请联系客服或使用其他登录方式')
}

// 其他登录方式（暂未实现）
const loginWithAlipay = () => {
  ElMessage.info('支付宝登录功能正在开发中...')
}

// const layoutElements = document.querySelectorAll('.mobile-layout, .bottom-navigation')
// console.log('页面中的Layout元素:', {
//   mobileLayout: layoutElements.length > 0 ? '找到' : '未找到',
//   bottomNavigation: document.querySelector('.bottom-navigation') ? '找到' : '未找到',
//   navItems: document.querySelectorAll('.nav-item').length + '个导航项'
// })

const loginWithWechat = () => {
  showWechatQR()
}

const loginWithId = () => {
  ElMessage.info('身份认证登录功能正在开发中...')
}

// 快速登录（仅开发环境可用）
const quickLogin = async () => {
  if (!isDev.value) {
    ElMessage.warning('快速登录仅在开发环境可用')
    return
  }
  
  try {
    loading.value = true
    
    // 使用环境变量或配置文件中的测试账号
    const response = await userApi.login({
      Phone: import.meta.env.VITE_TEST_PHONE || '18888888888',
      SendSms: import.meta.env.VITE_TEST_SMS || '1234'
    })
    
    // 更新用户状态
    if (response && response.data) {
      userStore.user = {
        id: response.data.ID || '',
        phone: response.data.Phone || '18888888888',
        name: response.data.Name || '',
        avatar: response.data.Avatar || ''
      }
      userStore.token = response.data.Token || ''
      userStore.loginTime = Date.now()
      ElMessage.success('快速登录成功！正在跳转...')
      await router.replace('/')
    }
  } catch (error) {
    console.error('快速登录失败:', error)
    ElMessage.error('快速登录失败，请重试')
  } finally {
    loading.value = false
  }
}

// 新的滴滴风格登录方法
// 返回按钮
const goBack = () => {
  // 如果有历史记录，则返回上一页
  if (window.history.length > 1) {
    router.go(-1)
  } else {
    // 否则跳转到首页
    router.push('/')
  }
}

// 下一步（发送短信验证码并跳转到验证码页面）
const nextStep = async () => {
  if (!canProceed.value) {
    ElMessage.warning(
      !agreedToTerms.value 
        ? '请同意服务协议' 
        : phoneValidation.message || '请输入正确的手机号'
    )
    return
  }
  
  try {
    loading.value = true
    
    // 调用后端发送短信接口
    await userApi.sendSms({
      Phone: loginForm.phone,
      Source: 'Login'
    })
    
    ElMessage.success('验证码已发送')
    
    // 跳转到验证码输入页面
    await router.push({
      path: '/verify-code',
      query: {
        phone: loginForm.phone,
        timestamp: Date.now() // 添加时间戳防止重放
      }
    })
  } catch (error) {
    console.error('发送验证码失败:', error)
    ElMessage.error('发送验证码失败，请检查网络连接')
  } finally {
    loading.value = false
  }
}

// 微信登录相关状态
const wechatDialogVisible = ref(false)
const qrcodeUrl = ref('')
const sessionId = ref('')
const loginStatus = ref<'waiting' | 'scanned' | 'success' | 'expired'>('waiting')
let pollTimer: number | null = null
let expireTimer: number | null = null

// 显示微信登录二维码
const showWechatQR = async () => {
  wechatDialogVisible.value = true
  loginStatus.value = 'waiting'
  await generateQRCode()
}

// 生成二维码
const generateQRCode = async () => {
  loading.value = true
  
  try {
    // 获取二维码
    const response = await wechatApi.getLoginQRCode(window.location.host)
    
    if (response instanceof Blob) {
      qrcodeUrl.value = URL.createObjectURL(response)
      sessionId.value = `temp_session_${Date.now()}`
      
      // 开始轮询检查登录状态
      startPolling()
      
      // 设置二维码过期时间（5分钟）
      expireTimer = window.setTimeout(() => {
        loginStatus.value = 'expired'
        stopPolling()
      }, 5 * 60 * 1000)
    } else {
      throw new Error('二维码生成失败')
    }
  } catch (error) {
    console.error('生成二维码失败:', error)
    ElMessage.error('二维码生成失败，请重试')
  } finally {
    loading.value = false
  }
}

// 开始轮询检查登录状态
const startPolling = () => {
  if (!sessionId.value) return
  
  pollTimer = window.setInterval(async () => {
    try {
      const response = await wechatApi.checkSignature({
        signature: sessionId.value,
        timestamp: Date.now().toString(),
        nonce: Math.random().toString(36).substring(7),
        echostr: ''
      })
      
      if (response) {
        // 处理微信登录响应
        stopPolling()
        wechatDialogVisible.value = false
        await handleLoginSuccess()
      }
    } catch (error) {
      console.error('检查登录状态失败:', error)
    }
  }, 2000)
}

// 停止轮询
const stopPolling = () => {
  if (pollTimer) {
    clearInterval(pollTimer)
    pollTimer = null
  }
  
  if (expireTimer) {
    clearTimeout(expireTimer)
    expireTimer = null
  }
}

// 刷新二维码
const refreshQRCode = async () => {
  stopPolling()
  await generateQRCode()
}

// 组件卸载时清理定时器
onUnmounted(() => {
  stopPolling()
})
</script>

<style lang="scss" scoped>
.didi-login-container {
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

/* Logo区域 */
.logo-section {
  margin-top: 80px;
  display: flex;
  justify-content: center;
  padding: 20px;
  
  .didi-logo {
    .logo-icon {
      width: 60px;
      height: 60px;
      background: #ff6600;
      border-radius: 50%;
      display: flex;
      align-items: center;
      justify-content: center;
      color: white;
      font-size: 28px;
      font-weight: bold;
      box-shadow: 0 4px 12px rgba(255, 102, 0, 0.3);
    }
  }
}

/* 标题区域 */
.title-section {
  text-align: center;
  padding: 30px 20px 50px;
  
  h1 {
    font-size: 28px;
    color: #333;
    font-weight: bold;
    margin: 0 0 8px 0;
  }
  
  p {
    font-size: 16px;
    color: #666;
    margin: 0;
  }
}

/* 手机号输入区域 */
.phone-section {
  padding: 0 20px;
  margin-bottom: 40px;
  
  .phone-input-container {
    display: flex;
    border-bottom: 1px solid #eee;
    padding-bottom: 10px;
    
    .country-code {
      display: flex;
      align-items: center;
      gap: 8px;
      padding-right: 20px;
      color: #333;
      font-size: 16px;
      border-right: 1px solid #eee;
      margin-right: 20px;
      cursor: pointer;
      
      .el-icon {
        font-size: 14px;
        color: #999;
      }
    }
    
    .phone-input {
      flex: 1;
      
      input {
        border: none;
        outline: none;
        width: 100%;
        font-size: 16px;
        color: #333;
        background: transparent;
        
        &::placeholder {
          color: #999;
        }
      }
    }
  }

  /* 新增：手机号验证样式 */
  .phone-input-container {
    &.error {
      border-bottom-color: #ff4d4f;
    }
  }
  
  .validation-message {
    margin-top: 8px;
    color: #ff4d4f;
    font-size: 12px;
  }
}

/* 服务协议区域 */
.agreement-section {
  padding: 0 20px;
  margin-bottom: 40px;
  
  .agreement-checkbox {
    display: flex;
    align-items: flex-start;
    gap: 12px;
    cursor: pointer;
    
    input[type="checkbox"] {
      display: none;
    }
    
    .checkmark {
      width: 20px;
      height: 20px;
      border: 2px solid #ddd;
      border-radius: 4px;
      position: relative;
      flex-shrink: 0;
      margin-top: 2px;
      transition: all 0.2s;
      
      &::after {
        content: '';
        position: absolute;
        display: none;
        left: 6px;
        top: 2px;
        width: 5px;
        height: 10px;
        border: solid white;
        border-width: 0 2px 2px 0;
        transform: rotate(45deg);
      }
    }
    
    input:checked + .checkmark {
      background-color: #ff6600;
      border-color: #ff6600;
      
      &::after {
        display: block;
      }
    }
    
    .agreement-text {
      font-size: 14px;
      color: #666;
      line-height: 1.5;
      
      .link {
        color: #ff6600;
        text-decoration: underline;
        cursor: pointer;
      }
    }
  }
}

/* 下一步按钮 */
.next-section {
  padding: 0 20px;
  margin-bottom: 30px;
  
  .next-btn {
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

/* 帮助区域 */
.help-section {
  text-align: center;
  margin-bottom: 40px;
  
  .help-link {
    color: #999;
    font-size: 14px;
    cursor: pointer;
    text-decoration: underline;
    
    &:hover {
      color: #ff6600;
    }
  }
}

/* 其他登录方式 */
.other-login-section {
  position: absolute;
  bottom: 80px;
  left: 0;
  right: 0;
  padding: 0 20px;
  
  .login-methods {
    display: flex;
    justify-content: center;
    gap: 30px;
    
    .login-method {
      cursor: pointer;
      transition: transform 0.2s;
      
      &:hover {
        transform: scale(1.1);
      }
      
      .method-icon {
        width: 50px;
        height: 50px;
        border-radius: 50%;
        display: flex;
        align-items: center;
        justify-content: center;
        color: white;
        font-size: 14px;
        font-weight: bold;
        
        &.alipay {
          background: #1677ff;
          font-size: 20px;
        }
        
        &.wechat {
          background: #07c160;
          font-size: 12px;
        }
        
        &.id-auth {
          background: #dc3545;
          font-size: 10px;
        }
        
        &.phone {
          background: #666;
          font-size: 20px;
        }
      }
    }
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
    padding: 20px 20px 40px;
    
    h1 {
      font-size: 24px;
    }
    
    p {
      font-size: 14px;
    }
  }
  
  .phone-section {
    margin-bottom: 30px;
  }
  
  .agreement-section {
    margin-bottom: 30px;
  }
  
  .other-login-section {
    bottom: 60px;
    
    .login-methods {
      gap: 20px;
      
      .login-method .method-icon {
        width: 45px;
        height: 45px;
      }
    }
  }
}

/* 安全区域适配 */
@supports (padding: max(0px)) {
  .didi-login-container {
    padding-bottom: max(20px, env(safe-area-inset-bottom));
  }
}

.qr-container {
  text-align: center;
  padding: 20px;
  
  .loading-container,
  .error-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 16px;
    
    p {
      color: #666;
      margin: 0;
    }
  }
  
  .qr-code {
    position: relative;
    
    img {
      width: 200px;
      height: 200px;
      border: 1px solid #e4e7ed;
      border-radius: 8px;
    }
    
    .qr-tips {
      margin: 20px 0;
      
      p {
        margin: 4px 0;
        
        &.sub-tip {
          font-size: 14px;
          color: #999;
        }
      }
    }
    
    .scan-success,
    .qr-expired {
      position: absolute;
      top: 0;
      left: 0;
      right: 0;
      bottom: 0;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      background: rgba(255, 255, 255, 0.9);
      border-radius: 8px;
      
      p {
        margin: 12px 0;
        font-size: 16px;
        color: #333;
      }
    }
  }
}
</style> 