<template>
  <div class="settings-page">
    <!-- 头部导航 -->
    <div class="header">
      <div class="back-btn" @click="goBack">
        <el-icon><ArrowLeft /></el-icon>
      </div>
      <h1 class="title">设置</h1>
    </div>

    <!-- 设置列表 -->
    <div class="settings-content">
      <!-- 账户设置 -->
      <div class="setting-section">
        <div class="section-title">账户设置</div>
        
        <div class="setting-item" @click="handleProfile">
          <div class="item-left">
            <el-icon class="item-icon"><User /></el-icon>
            <span class="item-text">个人信息</span>
          </div>
          <div class="item-right">
            <el-icon class="arrow-icon"><ArrowRight /></el-icon>
          </div>
        </div>

        <div class="setting-item" @click="handleSecurity">
          <div class="item-left">
            <el-icon class="item-icon"><Lock /></el-icon>
            <span class="item-text">账户与安全</span>
          </div>
          <div class="item-right">
            <el-icon class="arrow-icon"><ArrowRight /></el-icon>
          </div>
        </div>
      </div>

      <!-- 通用设置 -->
      <div class="setting-section">
        <div class="section-title">通用设置</div>
        
        <div class="setting-item" @click="handleNotification">
          <div class="item-left">
            <el-icon class="item-icon"><Bell /></el-icon>
            <span class="item-text">消息通知</span>
          </div>
          <div class="item-right">
            <el-icon class="arrow-icon"><ArrowRight /></el-icon>
          </div>
        </div>

        <div class="setting-item" @click="handlePayment">
          <div class="item-left">
            <el-icon class="item-icon"><CreditCard /></el-icon>
            <span class="item-text">支付设置</span>
          </div>
          <div class="item-right">
            <el-icon class="arrow-icon"><ArrowRight /></el-icon>
          </div>
        </div>
      </div>

      <!-- 退出登录 -->
      <div class="logout-section">
        <el-button 
          class="logout-btn"
          @click="handleLogout"
          :loading="isLoggingOut"
        >
          {{ isLoggingOut ? '退出中...' : '退出登录' }}
        </el-button>
      </div>
    </div>

    <!-- 退出确认对话框 -->
    <el-dialog
      v-model="showLogoutDialog"
      title="退出登录"
      width="300px"
      center
    >
      <div class="logout-dialog-content">
        <p>确定要退出登录吗？</p>
        <p class="logout-hint">退出后需要重新输入手机号登录</p>
      </div>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="showLogoutDialog = false">取消</el-button>
          <el-button 
            type="danger" 
            @click="confirmLogout"
            :loading="isLoggingOut"
          >
            确定退出
          </el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import {
  ArrowLeft,
  ArrowRight,
  User,
  Lock,
  Bell,
  CreditCard
} from '@element-plus/icons-vue'
import { useUserStore } from '@/store'

const router = useRouter()
const userStore = useUserStore()

// 状态管理
const isLoggingOut = ref(false)
const showLogoutDialog = ref(false)

// 返回上一页
const goBack = () => {
  router.go(-1)
}

// 设置功能处理函数
const handleProfile = () => {
  ElMessage.info('个人信息功能开发中...')
}

const handleSecurity = () => {
  ElMessage.info('账户与安全功能开发中...')
}

const handleNotification = () => {
  ElMessage.info('消息通知功能开发中...')
}

const handlePayment = () => {
  ElMessage.info('支付设置功能开发中...')
}

// 退出登录处理
const handleLogout = () => {
  showLogoutDialog.value = true
}

// 确认退出登录
const confirmLogout = async () => {
  try {
    isLoggingOut.value = true
    
    console.log('用户确认退出登录')
    
    // 调用Store的退出登录方法
    userStore.logout()
    
    // 延迟一下，给用户反馈
    await new Promise(resolve => setTimeout(resolve, 500))
    
    // 关闭对话框
    showLogoutDialog.value = false
    
    // 显示退出成功消息
    ElMessage.success('已退出登录')
    
    // 跳转到登录页面
    console.log('退出登录成功，跳转到登录页面')
    await router.replace('/login')
    
  } catch (error) {
    console.error('退出登录失败:', error)
    ElMessage.error('退出登录失败，请重试')
  } finally {
    isLoggingOut.value = false
  }
}
</script>

<style lang="scss" scoped>
.settings-page {
  min-height: 100vh;
  background: #f5f5f5;
  padding-bottom: 100px;
}

.header {
  background: white;
  height: 60px;
  display: flex;
  align-items: center;
  padding: 0 20px;
  position: sticky;
  top: 0;
  z-index: 100;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  
  .back-btn {
    width: 40px;
    height: 40px;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    border-radius: 50%;
    transition: background-color 0.2s;
    
    &:hover {
      background-color: rgba(0, 0, 0, 0.05);
    }
    
    .el-icon {
      font-size: 20px;
      color: #333;
    }
  }
  
  .title {
    flex: 1;
    text-align: center;
    font-size: 18px;
    font-weight: 600;
    color: #333;
    margin: 0;
    margin-left: -40px;
  }
}

.settings-content {
  padding: 20px;
}

.setting-section {
  background: white;
  border-radius: 12px;
  margin-bottom: 20px;
  overflow: hidden;
  
  .section-title {
    padding: 16px 20px 12px;
    font-size: 14px;
    color: #666;
    font-weight: 500;
    background: #f8f9fa;
    border-bottom: 1px solid #f0f0f0;
  }
  
  .setting-item {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 16px 20px;
    cursor: pointer;
    transition: background-color 0.2s;
    border-bottom: 1px solid #f0f0f0;
    
    &:last-child {
      border-bottom: none;
    }
    
    &:hover {
      background-color: #f8f9fa;
    }
    
    .item-left {
      display: flex;
      align-items: center;
      gap: 16px;
      
      .item-icon {
        font-size: 20px;
        color: #666;
      }
      
      .item-text {
        font-size: 16px;
        color: #333;
      }
    }
    
    .item-right {
      display: flex;
      align-items: center;
      gap: 8px;
      
      .arrow-icon {
        font-size: 16px;
        color: #ccc;
      }
    }
  }
}

.logout-section {
  margin-top: 40px;
  
  .logout-btn {
    width: 100%;
    height: 50px;
    background: #ff4757;
    color: white;
    border: none;
    border-radius: 25px;
    font-size: 16px;
    font-weight: 500;
    transition: all 0.2s;
    
    &:hover:not(.is-loading) {
      background: #ff3838;
    }
    
    &.is-loading {
      background: #ff6b7a;
    }
  }
}

.logout-dialog-content {
  text-align: center;
  padding: 20px 0;
  
  p {
    margin: 0 0 12px 0;
    font-size: 16px;
    color: #333;
  }
  
  .logout-hint {
    font-size: 14px;
    color: #666;
  }
}

.dialog-footer {
  display: flex;
  gap: 12px;
  
  .el-button {
    flex: 1;
    height: 44px;
    border-radius: 22px;
  }
}
</style> 