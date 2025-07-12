<template>
  <div class="wechat-integration">
    <div class="page-header">
      <h1 class="page-title">微信登录</h1>
      <p class="page-description">扫描二维码，使用微信快速登录</p>
    </div>

    <el-row :gutter="24">
      <el-col :span="12">
        <el-card title="微信扫码登录">
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
                <p class="session-tip">会话ID: {{ sessionId }}</p>
              </div>
              <div class="qr-actions">
                <el-button @click="refreshQRCode" :loading="refreshing">
                  刷新二维码
                </el-button>
              </div>
            </div>
            
            <div v-else class="error-container">
              <el-icon size="40" color="#f56c6c"><WarningFilled /></el-icon>
              <p>二维码生成失败</p>
              <el-button type="primary" @click="generateQRCode">
                重新生成
              </el-button>
            </div>
          </div>
        </el-card>
      </el-col>

      <el-col :span="12">
        <el-card title="登录状态">
          <div class="login-status">
            <div v-if="loginStatus === 'waiting'" class="status-waiting">
              <el-icon size="40" color="#409eff"><Clock /></el-icon>
              <p>等待扫码...</p>
              <p class="status-desc">请使用微信扫描左侧二维码</p>
            </div>
            
            <div v-else-if="loginStatus === 'scanned'" class="status-scanned">
              <el-icon size="40" color="#67c23a"><SuccessFilled /></el-icon>
              <p>扫码成功！</p>
              <p class="status-desc">请在手机上确认登录</p>
            </div>
            
            <div v-else-if="loginStatus === 'success'" class="status-success">
              <el-icon size="40" color="#67c23a"><CircleCheckFilled /></el-icon>
              <p>登录成功！</p>
              <p class="status-desc">正在跳转到系统...</p>
            </div>
            
            <div v-else-if="loginStatus === 'expired'" class="status-expired">
              <el-icon size="40" color="#e6a23c"><WarningFilled /></el-icon>
              <p>二维码已过期</p>
              <p class="status-desc">请刷新二维码重新扫描</p>
            </div>
          </div>
        </el-card>

        <!-- 功能说明 -->
        <el-card title="功能说明" class="feature-info">
          <ul class="feature-list">
            <li>
              <el-icon><Check /></el-icon>
              快速登录，无需记住密码
            </li>
            <li>
              <el-icon><Check /></el-icon>
              安全可靠，微信授权保护
            </li>
            <li>
              <el-icon><Check /></el-icon>
              自动注册，首次使用即可登录
            </li>
            <li>
              <el-icon><Check /></el-icon>
              同步头像昵称等基本信息
            </li>
          </ul>
        </el-card>
      </el-col>
    </el-row>

    <!-- 用户信息展示 -->
    <el-card v-if="userInfo" title="用户信息" class="user-info-card">
      <div class="user-info">
        <div class="user-avatar">
          <el-avatar :size="80" :src="userInfo.avatar" />
        </div>
        <div class="user-details">
          <h3>{{ userInfo.nickname }}</h3>
          <p>OpenID: {{ userInfo.openid }}</p>
          <p>用户ID: {{ userInfo.userId }}</p>
          <p v-if="accessToken">访问令牌: {{ accessToken.substring(0, 20) }}...</p>
          <div class="user-actions">
            <el-button type="primary" @click="goToDashboard">
              进入系统
            </el-button>
            <el-button @click="logout">
              退出登录
            </el-button>
          </div>
        </div>
      </div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { 
  Loading, 
  WarningFilled, 
  Clock, 
  SuccessFilled, 
  CircleCheckFilled, 
  Check 
} from '@element-plus/icons-vue'
import { wechatApi } from '@/api/wechat'
import { useUserStore } from '@/store'

// 定义用户信息类型
interface UserInfo {
  openid: string
  nickname: string
  avatar: string
  userId: number
}

// 定义登录状态类型
type LoginStatus = 'waiting' | 'scanned' | 'success' | 'expired'

const router = useRouter()
const userStore = useUserStore()

// 状态管理
const loading = ref(false)
const refreshing = ref(false)
const qrcodeUrl = ref('')
const sessionId = ref('')
const loginStatus = ref<LoginStatus>('waiting')
const userInfo = ref<UserInfo | null>(null)
const accessToken = ref('')

// 轮询定时器
let pollTimer: number | null = null
let expireTimer: number | null = null

// 生成二维码
const generateQRCode = async () => {
  loading.value = true
  loginStatus.value = 'waiting'
  
  try {
    // 获取二维码
    const response = await wechatApi.generateLoginQR(window.location.host)
    
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
      
      ElMessage.success('二维码生成成功')
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

// 刷新二维码
const refreshQRCode = async () => {
  refreshing.value = true
  stopPolling()
  
  if (expireTimer) {
    clearTimeout(expireTimer)
  }
  
  try {
    await generateQRCode()
  } finally {
    refreshing.value = false
  }
}

// 开始轮询检查登录状态
const startPolling = () => {
  if (!sessionId.value) return
  
  pollTimer = window.setInterval(async () => {
    try {
      const response = await wechatApi.checkLoginStatus(sessionId.value)
      
      if (response && response.data) {
        const { status, user_info, access_token } = response.data
        
        if (status !== loginStatus.value) {
          loginStatus.value = status as LoginStatus
          
          if (status === 'success' && user_info) {
            // 登录成功
            userInfo.value = user_info
            accessToken.value = access_token || ''
            
            // 保存用户信息到store
            await userStore.login({
              phone: '',
              sendSms: '',
              name: user_info.nickname,
              avatar: user_info.avatar
            })
            
            ElMessage.success('登录成功！')
            stopPolling()
            
            // 3秒后跳转到仪表盘
            setTimeout(() => {
              router.push('/dashboard')
            }, 3000)
          } else if (status === 'expired') {
            stopPolling()
            ElMessage.warning('二维码已过期，请刷新重试')
          }
        }
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
}

// 进入系统
const goToDashboard = () => {
  router.push('/dashboard')
}

// 退出登录
const logout = () => {
  userStore.logout()
  userInfo.value = null
  accessToken.value = ''
  loginStatus.value = 'waiting'
  ElMessage.success('已退出登录')
}

onMounted(() => {
  generateQRCode()
})

onUnmounted(() => {
  stopPolling()
  if (expireTimer) {
    clearTimeout(expireTimer)
  }
})
</script>

<style lang="scss" scoped>
.wechat-integration {
  padding: 24px;
  
  .page-header {
    margin-bottom: 24px;
    text-align: center;
    
    .page-title {
      font-size: 28px;
      font-weight: 600;
      color: #333;
      margin-bottom: 8px;
    }
    
    .page-description {
      color: #666;
      font-size: 16px;
    }
  }
  
  .qr-container {
    text-align: center;
    padding: 40px 20px;
    
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
          
          &.session-tip {
            font-size: 12px;
            color: #ccc;
            font-family: monospace;
          }
        }
      }
      
      .qr-actions {
        margin-top: 20px;
      }
    }
  }
  
  .login-status {
    text-align: center;
    padding: 40px 20px;
    
    .status-waiting,
    .status-scanned,
    .status-success,
    .status-expired {
      display: flex;
      flex-direction: column;
      align-items: center;
      gap: 16px;
      
      p {
        margin: 0;
        
        &:first-of-type {
          font-size: 18px;
          font-weight: 600;
          color: #333;
        }
        
        &.status-desc {
          color: #666;
          font-size: 14px;
        }
      }
    }
  }
  
  .feature-info {
    margin-top: 24px;
    
    .feature-list {
      list-style: none;
      padding: 0;
      margin: 0;
      
      li {
        display: flex;
        align-items: center;
        gap: 8px;
        padding: 8px 0;
        color: #666;
        
        .el-icon {
          color: #67c23a;
        }
      }
    }
  }
  
  .user-info-card {
    margin-top: 24px;
    
    .user-info {
      display: flex;
      align-items: center;
      gap: 20px;
      
      .user-details {
        flex: 1;
        
        h3 {
          margin: 0 0 12px 0;
          font-size: 20px;
          color: #333;
        }
        
        p {
          margin: 4px 0;
          color: #666;
          font-size: 14px;
        }
        
        .user-actions {
          margin-top: 16px;
          
          .el-button {
            margin-right: 12px;
          }
        }
      }
    }
  }
}
</style> 