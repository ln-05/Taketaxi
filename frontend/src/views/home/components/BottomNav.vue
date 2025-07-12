<template>
  <div class="bottom-nav">
    <div 
      v-for="item in navItems" 
      :key="item.id"
      class="nav-item"
      :class="{ active: activeTab === item.id }"
      @click="$emit('change-tab', item.id)"
    >
      <el-icon class="nav-icon">
        <component :is="item.icon" />
      </el-icon>
      <span class="nav-text">{{ item.name }}</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { 
  HomeFilled, 
  Van, 
  List, 
  CreditCard, 
  User 
} from '@element-plus/icons-vue'

interface NavItem {
  id: string
  name: string
  icon: any
}

interface Props {
  activeTab?: string
}

withDefaults(defineProps<Props>(), {
  activeTab: 'home'
})

const navItems: NavItem[] = [
  { id: 'home', name: '首页', icon: HomeFilled },
  { id: 'car', name: '车主', icon: Van },
  { id: 'order', name: '订单', icon: List },
  { id: 'payment', name: '修订车费', icon: CreditCard },
  { id: 'profile', name: '我的', icon: User }
]

defineEmits<{
  'change-tab': [tabId: string]
}>()
</script>

<style lang="scss" scoped>
.bottom-nav {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  height: 56px;
  background: white;
  border-top: 1px solid #f0f0f0;
  display: flex;
  justify-content: space-around;
  align-items: center;
  z-index: 1000;
  padding-bottom: env(safe-area-inset-bottom);
}

.nav-item {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 4px;
  cursor: pointer;
  transition: all 0.2s;
  
  &:active {
    transform: scale(0.95);
  }
  
  .nav-icon {
    font-size: 22px;
    color: #999;
    transition: color 0.2s;
  }
  
  .nav-text {
    font-size: 11px;
    color: #999;
    transition: color 0.2s;
  }
  
  &.active {
    .nav-icon,
    .nav-text {
      color: #ff6b35;
    }
  }
}
</style> 