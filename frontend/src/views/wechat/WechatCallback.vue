<template>
  <div class="wechat-callback">
    <div v-if="loading" class="loading">
      <el-icon class="is-loading" size="40"><Loading /></el-icon>
      <p>正在处理登录请求...</p>
    </div>
    
    <div v-else-if="error" class="error">
      <el-icon size="40" color="#f56c6c"><WarningFilled /></el-icon>
      <p>{{ error }}</p>
      <el-button type="primary" @click="goToLogin">返回登录</el-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { Loading, WarningFilled } from '@element-plus/icons-vue'
import { wechatApi } from '@/api'
import { useUserStore } from '@/store'

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()

const loading = ref(true)
const error = ref('')

// 处理微信回调
const handleCallback = async () => {
  const { code, state } = route.query
  
  if (!code || !state || Array.isArray(code) || Array.isArray(state)) {
    error.value = '登录参数错误'
    loading.value = false
    return
  }
  
  try {
    const response = await wechatApi.handleCallback({ code, state })
    
    if (response.code === 0 && response.data) {
      // 更新用户状态
      userStore.user = response.data.userInfo
      userStore.token = response.data.token
      userStore.loginTime = Date.now()
      
      ElMessage.success('登录成功！')
      await router.replace('/')
    } else {
      error.value = response.msg || '登录失败'
    }
  } catch (err) {
    console.error('处理微信回调失败:', err)
    error.value = '登录失败，请重试'
  } finally {
    loading.value = false
  }
}

// 返回登录页
const goToLogin = () => {
  router.replace('/login')
}

onMounted(() => {
  handleCallback()
})
</script>

<style lang="scss" scoped>
.wechat-callback {
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: #f5f7fa;
  
  .loading,
  .error {
    text-align: center;
    
    p {
      margin: 20px 0;
      color: #666;
      font-size: 16px;
    }
  }
}
</style> 