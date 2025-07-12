<template>
  <div class="service-grid">
    <!-- 主要服务 -->
    <div class="main-services">
      <div 
        v-for="service in mainServices" 
        :key="service.id"
        class="service-item main"
        @click="$emit('select-service', service.id)"
      >
        <div class="service-icon" :style="{ backgroundColor: service.color }">
          <span class="icon-emoji">{{ service.icon }}</span>
        </div>
        <span class="service-name">{{ service.name }}</span>
      </div>
    </div>

    <!-- 次要服务 -->
    <div class="sub-services">
      <div 
        v-for="service in subServices" 
        :key="service.id"
        class="service-item"
        @click="$emit('select-service', service.id)"
      >
        <div class="service-icon" :style="{ backgroundColor: service.color }">
          <span class="icon-emoji">{{ service.icon }}</span>
        </div>
        <span class="service-name">{{ service.name }}</span>
      </div>
    </div>

    <!-- 服务协议 -->
    <div class="service-footer">
      <span>《服务协议与规则》</span>
    </div>
  </div>
</template>

<script setup lang="ts">
interface Service {
  id: string
  name: string
  icon: string
  color: string
}

const mainServices: Service[] = [
  { id: 'taxi', name: '打车', icon: '🚗', color: '#ff6b35' },
  { id: 'carpool', name: '顺风车', icon: '🚙', color: '#4ecdc4' },
  { id: 'driver', name: '代驾', icon: '👨‍💼', color: '#45b7d1' },
  { id: 'intercity', name: '城际拼车', icon: '🚐', color: '#96ceb4' },
  { id: 'bike', name: '青桔骑行', icon: '🚲', color: '#feca57' }
]

const subServices: Service[] = [
  { id: 'discount', name: '特价拼车', icon: '🎯', color: '#ff9ff3' },
  { id: 'delivery', name: '快送跑腿', icon: '📦', color: '#54a0ff' },
  { id: 'loan', name: '借钱', icon: '💰', color: '#5f27cd' },
  { id: 'fuel', name: '加油/充电', icon: '⛽', color: '#00d2d3' },
  { id: 'transit', name: '公交地铁', icon: '🚇', color: '#ff6348' },
  { id: 'business', name: '企业用车', icon: '🏢', color: '#6c5ce7' },
  { id: 'cargo', name: '送货', icon: '📦', color: '#fd79a8' },
  { id: 'rental', name: '滴滴租车', icon: '🚗', color: '#ff9f43' },
  { id: 'ticket', name: '火车票机票', icon: '🎫', color: '#ee5a6f' },
  { id: 'more', name: '更多服务', icon: '⋯', color: '#c7ecee' }
]

defineEmits<{
  'select-service': [serviceId: string]
}>()
</script>

<style lang="scss" scoped>
.service-grid {
  background: white;
  padding: 20px 16px 0;
}

.main-services {
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
  padding-bottom: 20px;
  border-bottom: 1px solid #f0f0f0;
}

.sub-services {
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  gap: 16px;
  margin-bottom: 20px;
}

.service-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  transition: transform 0.2s;
  
  &:active {
    transform: scale(0.95);
  }
  
  &.main {
    .service-icon {
      width: 56px;
      height: 56px;
    }
    
    .service-name {
      font-size: 14px;
    }
  }
  
  .service-icon {
    width: 48px;
    height: 48px;
    border-radius: 12px;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.2s;
    
    .icon-emoji {
      font-size: 24px;
    }
  }
  
  .service-name {
    font-size: 12px;
    color: #333;
    text-align: center;
    line-height: 1.2;
  }
}

.service-footer {
  text-align: center;
  padding: 16px 0;
  border-top: 1px solid #f0f0f0;
  
  span {
    font-size: 12px;
    color: #999;
  }
}
</style> 