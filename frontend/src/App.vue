<template>
  <div id="app">
    <!-- 登录页面不需要布局 -->
    <template v-if="isLoginPage">
      <router-view />
    </template>
    <!-- 其他页面使用布局组件 -->
    <template v-else>
      <Layout />
    </template>
  </div>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router'
import { onMounted, computed } from 'vue'
import Layout from '@/components/Layout/index.vue'
import { useUserStore } from '@/store'

const route = useRoute()
const userStore = useUserStore()

// 计算是否为登录相关页面（不需要Layout的页面）
const isLoginPage = computed(() => {
  return route.path === '/login' || route.path === '/verify-code'
})

// 应用启动时初始化用户信息（只初始化一次）
onMounted(() => {
  console.log('App mounted: 初始化用户状态')
  userStore.initUser()
})
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

#app {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen',
    'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue',
    sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  height: 100vh;
  overflow: hidden;
}

body {
  margin: 0;
  padding: 0;
  background-color: #f5f5f5;
}

/* 移动端适配 */
@media (max-width: 768px) {
  html {
    font-size: 14px;
  }
}

/* 滚动条样式 */
::-webkit-scrollbar {
  width: 4px;
}

::-webkit-scrollbar-track {
  background: #f1f1f1;
}

::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 2px;
}

::-webkit-scrollbar-thumb:hover {
  background: #a8a8a8;
}
</style> 