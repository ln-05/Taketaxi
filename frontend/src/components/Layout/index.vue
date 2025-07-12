<template>
  <div class="mobile-layout">
    <!-- 主内容区域 -->
    <div class="main-content">
      <router-view />
    </div>
    
    <!-- 底部导航栏 -->
    <div class="bottom-navigation">
      <div class="nav-items">
        <div 
          v-for="item in navItems" 
          :key="item.path"
          :class="['nav-item', { 
            active: currentPath === item.path,
            navigating: navigating
          }]"
          :aria-label="`导航到${item.label}页面`"
          :aria-pressed="currentPath === item.path"
          role="button"
          tabindex="0"
          @click="navigateTo(item.path)"
          @keydown.enter="(event: KeyboardEvent) => handleKeydown(event, item.path)"
          @keydown.space="(event: KeyboardEvent) => handleKeydown(event, item.path)"
        >
          <div class="nav-icon">
            <component :is="item.icon" />
          </div>
          <span class="nav-text">{{ item.label }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import {
  HomeFilled,
  Van,
  Tickets,
  ShoppingBag,
  User
} from '@element-plus/icons-vue'
import type { Component } from 'vue'

interface NavItem {
  path: string
  label: string
  icon: Component
}

const router = useRouter()
const route = useRoute()

// 导航项配置
const navItems: NavItem[] = [
  {
    path: '/home',
    label: '首页',
    icon: HomeFilled
  },
  {
    path: '/owner',
    label: '车主',
    icon: Van
  },
  {
    path: '/tickets',
    label: '订票',
    icon: Tickets
  },
  {
    path: '/welfare',
    label: '领打车费',
    icon: ShoppingBag
  },
  {
    path: '/user',
    label: '我的',
    icon: User
  }
]

// 当前活跃路径的computed属性（用于响应式地判断当前路由）
const currentPath = computed(() => route.path)

// 导航跳转（防重复点击）
const navigating = ref(false)

// 键盘事件处理
const handleKeydown = (event: KeyboardEvent, path: string) => {
  event.preventDefault()
  navigateTo(path)
}

const navigateTo = async (path: string) => {
  // 防止重复点击和重复跳转
  if (path === currentPath.value || navigating.value) {
    return
  }
  
  try {
    navigating.value = true
    await router.push(path)
  } catch (error) {
    console.error('底部导航跳转失败:', error)
  } finally {
    // 使用 setTimeout 确保动画完成后才允许下次跳转
    setTimeout(() => {
      navigating.value = false
    }, 300)
  }
}
</script>

<style lang="scss" scoped>
.mobile-layout {
  height: 100vh;
  display: flex;
  flex-direction: column;
  background: #f5f5f5;
}

.main-content {
  flex: 1;
  overflow-y: auto;
  padding-bottom: 70px; /* 为底部导航留出空间 */
  background: #f8f9fa;
}

.bottom-navigation {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  background: #fff;
  border-top: 1px solid #e4e7ed;
  z-index: 1000;
  padding-bottom: env(safe-area-inset-bottom); /* iOS 安全区域适配 */
}

.nav-items {
  display: flex;
  padding: 8px 0;
}

.nav-item {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 6px 4px;
  cursor: pointer;
  transition: all 0.3s ease;
  
  .nav-icon {
    font-size: 22px;
    margin-bottom: 4px;
    transition: all 0.3s ease;
    color: #999;
    
    .el-icon {
      font-size: 22px;
    }
  }
  
  .nav-text {
    font-size: 10px;
    color: #999;
    transition: all 0.3s ease;
    line-height: 1.2;
  }
  
  &.active {
    .nav-icon {
      color: #ff6600;
      transform: scale(1.1);
    }
    
    .nav-text {
      color: #ff6600;
      font-weight: 500;
    }
  }
  
  &.navigating {
    pointer-events: none; // 防止重复点击
    opacity: 0.6;
    
    .nav-icon {
      transform: scale(0.9);
    }
  }
  
  &:active {
    transform: scale(0.95);
  }
}

/* 响应式设计 */
@media (max-width: 768px) {
  .nav-item {
    padding: 4px 2px;
    
    .nav-icon {
      font-size: 20px;
      
      .el-icon {
        font-size: 20px;
      }
    }
    
    .nav-text {
      font-size: 9px;
    }
  }
}

/* 适配不同屏幕密度 */
@media (-webkit-min-device-pixel-ratio: 2) {
  .bottom-navigation {
    border-top-width: 0.5px;
  }
}

/* 暗色主题适配 */
@media (prefers-color-scheme: dark) {
  .bottom-navigation {
    background: #1f1f1f;
    border-top-color: #333;
  }
  
  .nav-item {
    .nav-icon {
      color: #666;
    }
    
    .nav-text {
      color: #666;
    }
    
    &.active {
      .nav-icon {
        color: #ff6600;
      }
      
      .nav-text {
        color: #ff6600;
      }
    }
  }
}
</style> 