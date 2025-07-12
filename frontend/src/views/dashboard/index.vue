<template>
  <div class="dashboard-container">
    <!-- 头部欢迎区域 -->
    <div class="welcome-section">
      <div class="greeting">
        <h1>欢迎回来</h1>
        <p>{{ currentTime }}</p>
      </div>
      <div class="weather-info">
        <div class="weather-icon">☀️</div>
        <div class="weather-text">
          <span class="temperature">{{ weather.temperature }}°</span>
          <span class="condition">{{ weather.condition }}</span>
        </div>
      </div>
    </div>

    <!-- 快捷功能卡片 -->
    <div class="quick-actions">
      <div class="action-card" @click="quickCall">
        <div class="card-icon">🚗</div>
        <div class="card-title">快速叫车</div>
        <div class="card-subtitle">预估3分钟</div>
      </div>
      <div class="action-card" @click="viewOrders">
        <div class="card-icon">📋</div>
        <div class="card-title">我的订单</div>
        <div class="card-subtitle">{{ orderCount }}个订单</div>
      </div>
      <div class="action-card" @click="viewWallet">
        <div class="card-icon">💰</div>
        <div class="card-title">我的钱包</div>
        <div class="card-subtitle">余额 ¥{{ balance }}</div>
      </div>
    </div>

    <!-- 数据统计 -->
    <div class="stats-section">
      <h2>本月数据</h2>
      <div class="stats-grid">
        <div class="stat-item">
          <div class="stat-value">{{ stats.trips }}</div>
          <div class="stat-label">总行程</div>
        </div>
        <div class="stat-item">
          <div class="stat-value">{{ stats.distance }}</div>
          <div class="stat-label">总里程(km)</div>
        </div>
        <div class="stat-item">
          <div class="stat-value">{{ stats.savings }}</div>
          <div class="stat-label">节省费用(¥)</div>
        </div>
        <div class="stat-item">
          <div class="stat-value">{{ stats.carbon }}</div>
          <div class="stat-label">减排(kg)</div>
        </div>
      </div>
    </div>

    <!-- 最近订单 -->
    <div class="recent-orders">
      <div class="section-header">
        <h2>最近订单</h2>
        <el-button text type="primary" @click="$router.push('/order')">查看全部</el-button>
      </div>
      <div class="order-list">
        <div 
          v-for="order in recentOrders" 
          :key="order.id"
          class="order-item"
          @click="viewOrderDetail(order.id)"
        >
          <div class="order-icon">
            <el-icon><Van /></el-icon>
          </div>
          <div class="order-info">
            <div class="order-route">{{ order.from }} → {{ order.to }}</div>
            <div class="order-time">{{ order.time }}</div>
          </div>
          <div class="order-status" :class="order.status">
            {{ getStatusText(order.status) }}
          </div>
        </div>
      </div>
    </div>

    <!-- 个人成就 -->
    <div class="achievements">
      <h2>个人成就</h2>
      <div class="achievement-list">
        <div 
          v-for="achievement in achievements" 
          :key="achievement.id"
          class="achievement-item"
          :class="{ unlocked: achievement.unlocked }"
        >
          <div class="achievement-icon">{{ achievement.icon }}</div>
          <div class="achievement-info">
            <div class="achievement-title">{{ achievement.title }}</div>
            <div class="achievement-desc">{{ achievement.description }}</div>
          </div>
          <div class="achievement-status">
            {{ achievement.unlocked ? '已获得' : achievement.progress }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElIcon } from 'element-plus'
import { Van } from '@element-plus/icons-vue'

const router = useRouter()

// 当前时间
const currentTime = ref('')

// 天气信息
const weather = ref({
  temperature: 22,
  condition: '晴朗'
})

// 用户统计数据
const stats = ref({
  trips: 15,
  distance: 128.5,
  savings: 89,
  carbon: 12.3
})

// 钱包余额
const balance = ref(156.8)

// 订单数量
const orderCount = ref(3)

// 最近订单
const recentOrders = ref([
  {
    id: 1,
    from: '北京南站',
    to: '中关村软件园',
    time: '今天 14:30',
    status: 'completed'
  },
  {
    id: 2,
    from: '望京SOHO',
    to: '三里屯太古里',
    time: '昨天 19:45',
    status: 'completed'
  },
  {
    id: 3,
    from: '国贸CBD',
    to: '朝阳公园',
    time: '昨天 12:15',
    status: 'cancelled'
  }
])

// 个人成就
const achievements = ref([
  {
    id: 1,
    icon: '🚗',
    title: '出行达人',
    description: '完成100次行程',
    unlocked: false,
    progress: '15/100'
  },
  {
    id: 2,
    icon: '🌱',
    title: '环保先锋',
    description: '减排100kg碳排放',
    unlocked: false,
    progress: '12.3/100'
  },
  {
    id: 3,
    icon: '⭐',
    title: '五星乘客',
    description: '获得100个五星评价',
    unlocked: true,
    progress: '已达成'
  },
  {
    id: 4,
    icon: '💎',
    title: 'VIP会员',
    description: '成为VIP会员',
    unlocked: true,
    progress: '已达成'
  }
])

// 更新当前时间
const updateTime = () => {
  const now = new Date()
  const hour = now.getHours()
  let greeting = ''
  
  if (hour < 6) greeting = '凌晨好'
  else if (hour < 9) greeting = '早上好'
  else if (hour < 12) greeting = '上午好'
  else if (hour < 14) greeting = '中午好'
  else if (hour < 18) greeting = '下午好'
  else if (hour < 22) greeting = '晚上好'
  else greeting = '夜深了'
  
  currentTime.value = `${greeting}，${now.toLocaleDateString('zh-CN', {
    month: 'long',
    day: 'numeric',
    weekday: 'long'
  })}`
}

// 获取订单状态文本
const getStatusText = (status: string) => {
  const statusMap: Record<string, string> = {
    completed: '已完成',
    cancelled: '已取消',
    pending: '进行中'
  }
  return statusMap[status] || '未知'
}

// 快速叫车
const quickCall = () => {
  ElMessage.success('正在为您安排车辆...')
  router.push('/home')
}

// 查看订单
const viewOrders = () => {
  router.push('/order')
}

// 查看钱包
const viewWallet = () => {
  ElMessage.info('钱包功能正在开发中...')
}

// 查看订单详情
const viewOrderDetail = (orderId: number) => {
  ElMessage.info(`查看订单 #${orderId} 详情`)
}

// 组件挂载时初始化
onMounted(() => {
  updateTime()
  // 每分钟更新一次时间
  setInterval(updateTime, 60000)
})
</script>

<style lang="scss" scoped>
.dashboard-container {
  min-height: 100vh;
  background: #f5f5f5;
  padding: 20px;
  padding-bottom: 100px; // 为底部导航留空间
}

// 欢迎区域
.welcome-section {
  background: linear-gradient(135deg, #ff6600, #ff8c42);
  border-radius: 16px;
  padding: 24px;
  color: white;
  margin-bottom: 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  
  .greeting {
    h1 {
      font-size: 24px;
      font-weight: bold;
      margin: 0 0 8px 0;
    }
    
    p {
      font-size: 14px;
      opacity: 0.9;
      margin: 0;
    }
  }
  
  .weather-info {
    display: flex;
    align-items: center;
    gap: 12px;
    
    .weather-icon {
      font-size: 32px;
    }
    
    .weather-text {
      display: flex;
      flex-direction: column;
      
      .temperature {
        font-size: 20px;
        font-weight: bold;
      }
      
      .condition {
        font-size: 12px;
        opacity: 0.8;
      }
    }
  }
}

// 快捷操作卡片
.quick-actions {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 12px;
  margin-bottom: 24px;
  
  .action-card {
    background: white;
    border-radius: 12px;
    padding: 20px;
    text-align: center;
    cursor: pointer;
    transition: all 0.3s ease;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
    
    &:hover {
      transform: translateY(-4px);
      box-shadow: 0 4px 16px rgba(0, 0, 0, 0.15);
    }
    
    .card-icon {
      font-size: 32px;
      margin-bottom: 12px;
    }
    
    .card-title {
      font-size: 16px;
      font-weight: 600;
      color: #333;
      margin-bottom: 4px;
    }
    
    .card-subtitle {
      font-size: 12px;
      color: #666;
    }
  }
}

// 数据统计
.stats-section {
  background: white;
  border-radius: 16px;
  padding: 24px;
  margin-bottom: 24px;
  
  h2 {
    font-size: 18px;
    color: #333;
    margin: 0 0 20px 0;
    font-weight: 600;
  }
  
  .stats-grid {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    gap: 16px;
    
    .stat-item {
      text-align: center;
      padding: 16px;
      background: #f8f9fa;
      border-radius: 12px;
      
      .stat-value {
        font-size: 24px;
        font-weight: bold;
        color: #ff6600;
        margin-bottom: 4px;
      }
      
      .stat-label {
        font-size: 12px;
        color: #666;
      }
    }
  }
}

// 最近订单
.recent-orders {
  background: white;
  border-radius: 16px;
  padding: 24px;
  margin-bottom: 24px;
  
  .section-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
    
    h2 {
      font-size: 18px;
      color: #333;
      margin: 0;
      font-weight: 600;
    }
  }
  
  .order-list {
    .order-item {
      display: flex;
      align-items: center;
      padding: 16px 0;
      border-bottom: 1px solid #f0f0f0;
      cursor: pointer;
      transition: background-color 0.2s;
      
      &:last-child {
        border-bottom: none;
      }
      
      &:hover {
        background-color: #f8f9fa;
        border-radius: 8px;
        padding: 16px 12px;
        margin: 0 -12px;
      }
      
      .order-icon {
        width: 40px;
        height: 40px;
        background: #ff6600;
        border-radius: 50%;
        display: flex;
        align-items: center;
        justify-content: center;
        color: white;
        margin-right: 16px;
        font-size: 18px;
      }
      
      .order-info {
        flex: 1;
        
        .order-route {
          font-size: 14px;
          color: #333;
          font-weight: 500;
          margin-bottom: 4px;
        }
        
        .order-time {
          font-size: 12px;
          color: #666;
        }
      }
      
      .order-status {
        font-size: 12px;
        padding: 4px 8px;
        border-radius: 12px;
        font-weight: 500;
        
        &.completed {
          background: #e8f5e8;
          color: #52c41a;
        }
        
        &.cancelled {
          background: #ffebe6;
          color: #ff4d4f;
        }
        
        &.pending {
          background: #e6f7ff;
          color: #1890ff;
        }
      }
    }
  }
}

// 个人成就
.achievements {
  background: white;
  border-radius: 16px;
  padding: 24px;
  
  h2 {
    font-size: 18px;
    color: #333;
    margin: 0 0 20px 0;
    font-weight: 600;
  }
  
  .achievement-list {
    .achievement-item {
      display: flex;
      align-items: center;
      padding: 16px 0;
      border-bottom: 1px solid #f0f0f0;
      
      &:last-child {
        border-bottom: none;
      }
      
      &.unlocked {
        .achievement-icon {
          background: #ff6600;
        }
        
        .achievement-title {
          color: #ff6600;
        }
      }
      
      .achievement-icon {
        width: 48px;
        height: 48px;
        background: #f0f0f0;
        border-radius: 50%;
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 20px;
        margin-right: 16px;
      }
      
      .achievement-info {
        flex: 1;
        
        .achievement-title {
          font-size: 14px;
          color: #333;
          font-weight: 500;
          margin-bottom: 4px;
        }
        
        .achievement-desc {
          font-size: 12px;
          color: #666;
        }
      }
      
      .achievement-status {
        font-size: 12px;
        color: #999;
      }
    }
  }
}

// 移动端适配
@media (max-width: 768px) {
  .dashboard-container {
    padding: 16px;
  }
  
  .quick-actions {
    grid-template-columns: repeat(3, 1fr);
    gap: 8px;
    
    .action-card {
      padding: 16px 8px;
      
      .card-icon {
        font-size: 24px;
        margin-bottom: 8px;
      }
      
      .card-title {
        font-size: 14px;
      }
      
      .card-subtitle {
        font-size: 11px;
      }
    }
  }
  
  .stats-grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 12px;
    
    .stat-item {
      padding: 12px;
      
      .stat-value {
        font-size: 20px;
      }
    }
  }
  
  .welcome-section {
    padding: 20px;
    
    .greeting h1 {
      font-size: 20px;
    }
    
    .weather-text {
      .temperature {
        font-size: 18px;
      }
    }
  }
}
</style> 