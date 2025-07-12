<template>
  <div class="user-center">
    <!-- 用户信息头部 -->
    <div class="user-header">
      <div class="user-info">
        <el-avatar :size="60" :src="userInfo.avatar" class="user-avatar">
          <el-icon><User /></el-icon>
        </el-avatar>
        <div class="user-details">
          <h3>{{ maskedPhone }}</h3>
          <div class="user-tags">
            <span class="tag">里程 ></span>
            <span class="tag-divider">|</span>
            <span class="tag">周报 ></span>
          </div>
        </div>
      </div>
      <div class="user-actions">
        <el-button text @click="handleMessage">消息</el-button>
        <el-button text @click="handleSettings">设置</el-button>
      </div>
    </div>

    <!-- 会员卡片 -->
    <div class="member-card">
      <div class="member-info">
        <span class="member-level">V1 会员</span>
        <span class="member-benefit">行程意外险 · 5次</span>
      </div>
      <el-button type="primary" plain size="small">查看</el-button>
    </div>

    <!-- 订单相关功能 -->
    <div class="function-grid">
      <div class="grid-item" @click="handleAllOrders">
        <el-icon><Document /></el-icon>
        <span>全部订单</span>
      </div>
      <div class="grid-item" @click="handlePending">
        <el-icon><Clock /></el-icon>
        <span>待出发</span>
      </div>
      <div class="grid-item" @click="handlePayment">
        <el-icon><Money /></el-icon>
        <span>待支付</span>
      </div>
      <div class="grid-item" @click="handleInvoice">
        <el-icon><Files /></el-icon>
        <span>开发票</span>
      </div>
      <div class="grid-item" @click="handleCustomerService">
        <el-icon><Headset /></el-icon>
        <span>客服</span>
      </div>
    </div>

    <!-- 钱包功能 -->
    <div class="wallet-section">
      <div class="wallet-item" @click="handleWallet">
        <el-icon><Wallet /></el-icon>
        <div class="wallet-info">
          <span class="wallet-title">我的钱包</span>
        </div>
      </div>
      
      <div class="wallet-stats">
        <div class="stat-item" @click="handleCoupons">
          <div class="stat-number">5<span class="unit">张</span></div>
          <div class="stat-label">优惠卡券</div>
        </div>
        <div class="stat-item" @click="handleCashback">
          <div class="stat-number">18.8<span class="unit">元</span></div>
          <div class="stat-label">单单返现</div>
        </div>
        <div class="stat-item" @click="handleFinance">
          <div class="stat-number">1<span class="unit">个</span></div>
          <div class="stat-label">金融福利</div>
        </div>
        <div class="stat-item recommended" @click="handleCredit">
          <div class="stat-number">20<span class="unit">万</span></div>
          <div class="stat-label">预估可借</div>
          <span class="recommend-tag">推荐</span>
        </div>
      </div>
    </div>

    <!-- 其他功能 -->
    <div class="other-functions">
      <div class="function-item" @click="handleAchievement">
        <el-icon><Trophy /></el-icon>
        <span>出行成就</span>
      </div>
      <div class="function-item" @click="handleFeedback">
        <el-icon><Microphone /></el-icon>
        <span>公众评议</span>
      </div>
      <div class="function-item" @click="handleReview">
        <el-icon><EditPen /></el-icon>
        <span>评价有奖</span>
      </div>
      <div class="function-item" @click="handlePayment">
        <el-icon><CreditCard /></el-icon>
        <span>滴滴支付</span>
      </div>
    </div>

    <!-- 活动推广 -->
    <div class="promotion-section">
      <div class="promotion-item" @click="handleSaveMoney">
        <div class="promotion-icon save-money">🎯</div>
        <span>省钱套餐</span>
      </div>
      <div class="promotion-item" @click="handleDailyReward">
        <div class="promotion-icon daily-reward">🎁</div>
        <span>天天神券</span>
      </div>
      <div class="promotion-item" @click="handleWelfare">
        <div class="promotion-icon welfare">🐶</div>
        <span>福利群</span>
      </div>
      <div class="promotion-item" @click="handleGold">
        <div class="promotion-icon gold">💰</div>
        <span>福利金</span>
      </div>
      <div class="promotion-item" @click="handleMall">
        <div class="promotion-icon mall">🛍️</div>
        <span>积分商城</span>
      </div>
    </div>

    <!-- 车主服务 -->
    <div class="driver-service">
      <h4>车主服务</h4>
      <p class="service-desc">有车的人就用TA</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import {
  User, Document, Clock, Money, Files, Headset,
  Wallet, Trophy, Microphone, EditPen, CreditCard
} from '@element-plus/icons-vue'
import { useUserStore } from '@/store'

const router = useRouter()
const userStore = useUserStore()

// 用户信息
const userInfo = computed(() => userStore.user || {
  phone: '18138005937',
  avatar: '',
  name: '用户'
})

// 手机号掩码处理
const maskedPhone = computed(() => {
  const phone = userInfo.value.phone || '18138005937'
  return phone.replace(/(\d{3})\d{4}(\d{4})/, '$1****$2')
})

// 功能处理函数
const handleMessage = () => {
  ElMessage.info('消息功能开发中')
}

const handleSettings = () => {
  router.push('/settings')
}

const handleAllOrders = () => {
  router.push('/order')
}

const handlePending = () => {
  ElMessage.info('待出发订单')
}

const handlePayment = () => {
  ElMessage.info('待支付订单')
}

const handleInvoice = () => {
  ElMessage.info('开发票功能')
}

const handleCustomerService = () => {
  ElMessage.info('客服功能')
}

const handleWallet = () => {
  ElMessage.info('我的钱包')
}

const handleCoupons = () => {
  ElMessage.info('优惠卡券')
}

const handleCashback = () => {
  ElMessage.info('单单返现')
}

const handleFinance = () => {
  ElMessage.info('金融福利')
}

const handleCredit = () => {
  ElMessage.info('预估可借')
}

const handleAchievement = () => {
  ElMessage.info('出行成就')
}

const handleFeedback = () => {
  ElMessage.info('公众评议')
}

const handleReview = () => {
  ElMessage.info('评价有奖')
}

const handleSaveMoney = () => {
  ElMessage.info('省钱套餐')
}

const handleDailyReward = () => {
  ElMessage.info('天天神券')
}

const handleWelfare = () => {
  ElMessage.info('福利群')
}

const handleGold = () => {
  ElMessage.info('福利金')
}

const handleMall = () => {
  ElMessage.info('积分商城')
}
</script>

<style lang="scss" scoped>
.user-center {
  background-color: #f5f5f5;
  min-height: 100vh;
  padding-bottom: 20px;
}

.user-header {
  background: white;
  padding: 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  
  .user-info {
    display: flex;
    align-items: center;
    gap: 15px;
    
    .user-avatar {
      background-color: #ff6b35;
    }
    
    .user-details {
      h3 {
        margin: 0 0 8px 0;
        font-size: 18px;
        color: #333;
      }
      
      .user-tags {
        display: flex;
        align-items: center;
        gap: 8px;
        color: #666;
        font-size: 14px;
        
        .tag-divider {
          color: #ddd;
        }
      }
    }
  }
  
  .user-actions {
    display: flex;
    gap: 15px;
  }
}

.member-card {
  background: linear-gradient(135deg, #a8e6cf 0%, #88d8c0 100%);
  margin: 10px 20px;
  padding: 15px 20px;
  border-radius: 12px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  
  .member-info {
    display: flex;
    align-items: center;
    gap: 10px;
    
    .member-level {
      font-weight: bold;
      color: #333;
    }
    
    .member-benefit {
      color: #666;
      font-size: 14px;
    }
  }
}

.function-grid {
  background: white;
  margin: 10px 20px;
  padding: 20px;
  border-radius: 12px;
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  gap: 20px;
  
  .grid-item {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 8px;
    cursor: pointer;
    
    .el-icon {
      font-size: 24px;
      color: #333;
    }
    
    span {
      font-size: 14px;
      color: #666;
    }
    
    &:hover {
      opacity: 0.7;
    }
  }
}

.wallet-section {
  background: white;
  margin: 10px 20px;
  border-radius: 12px;
  overflow: hidden;
  
  .wallet-item {
    padding: 20px;
    display: flex;
    align-items: center;
    gap: 15px;
    cursor: pointer;
    border-bottom: 1px solid #f0f0f0;
    
    .el-icon {
      font-size: 24px;
      color: #ff6b35;
    }
    
    .wallet-title {
      font-size: 16px;
      color: #333;
    }
  }
  
  .wallet-stats {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    
    .stat-item {
      padding: 20px 15px;
      text-align: center;
      cursor: pointer;
      position: relative;
      
      &:not(:last-child) {
        border-right: 1px solid #f0f0f0;
      }
      
      .stat-number {
        font-size: 18px;
        font-weight: bold;
        color: #333;
        margin-bottom: 5px;
        
        .unit {
          font-size: 14px;
          font-weight: normal;
        }
      }
      
      .stat-label {
        font-size: 12px;
        color: #666;
      }
      
      &.recommended {
        .recommend-tag {
          position: absolute;
          top: 8px;
          right: 8px;
          background: #ff6b35;
          color: white;
          font-size: 10px;
          padding: 2px 6px;
          border-radius: 8px;
        }
      }
      
      &:hover {
        background-color: #f8f8f8;
      }
    }
  }
}

.other-functions {
  background: white;
  margin: 10px 20px;
  padding: 20px;
  border-radius: 12px;
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 20px;
  
  .function-item {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 8px;
    cursor: pointer;
    
    .el-icon {
      font-size: 24px;
      color: #333;
    }
    
    span {
      font-size: 14px;
      color: #666;
    }
    
    &:hover {
      opacity: 0.7;
    }
  }
}

.promotion-section {
  background: white;
  margin: 10px 20px;
  padding: 20px;
  border-radius: 12px;
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  gap: 20px;
  
  .promotion-item {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 8px;
    cursor: pointer;
    
    .promotion-icon {
      width: 40px;
      height: 40px;
      border-radius: 20px;
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 20px;
    }
    
    span {
      font-size: 14px;
      color: #666;
    }
    
    &:hover {
      opacity: 0.7;
    }
  }
}

.driver-service {
  background: white;
  margin: 10px 20px;
  padding: 30px 20px;
  border-radius: 12px;
  text-align: center;
  
  h4 {
    margin: 0 0 10px 0;
    font-size: 18px;
    color: #333;
  }
  
  .service-desc {
    margin: 0;
    color: #666;
    font-size: 14px;
  }
}
</style> 