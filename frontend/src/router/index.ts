import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '@/store'

const routes = [
  {
    path: '/',
    redirect: '/login'
  },
  {
    path: '/home',
    name: 'Home',
    component: () => import('@/views/home/index.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/auth/Login.vue'),
    meta: { requiresAuth: false }
  },
  {
    path: '/verify-code',
    name: 'VerifyCode',
    component: () => import('@/views/auth/VerifyCode.vue'),
    meta: { requiresAuth: false }
  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: () => import('@/views/dashboard/index.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/user',
    name: 'User',
    component: () => import('@/views/user/index.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/order',
    name: 'Order',
    component: () => import('@/views/order/index.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/owner',
    name: 'Owner',
    component: () => import('@/views/owner/index.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/tickets',
    name: 'Tickets',
    component: () => import('@/views/tickets/index.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/welfare',
    name: 'Welfare',
    component: () => import('@/views/welfare/index.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/settings',
    name: 'Settings',
    component: () => import('@/views/settings/index.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/wechat/callback',
    name: 'WechatCallback',
    component: () => import('@/views/wechat/WechatCallback.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// 路由守卫
router.beforeEach(async (to, from, next) => {
  console.log(`=== 路由守卫开始 ===`)
  console.log(`路由守卫: ${from.path} -> ${to.path}`)
  
  // 避免循环重定向
  if (to.path === from.path) {
    console.log('路由守卫: 避免循环重定向，直接通过')
    next()
    return
  }

  const userStore = useUserStore()
  console.log('路由守卫: 用户状态检查', {
    isLoggedIn: userStore.isLoggedIn,
    hasUser: !!userStore.user,
    hasToken: !!userStore.token,
    targetRequiresAuth: to.meta.requiresAuth
  })
  
  // 检查路由是否需要认证
  if (to.meta.requiresAuth) {
    // 检查用户是否已登录
    if (userStore.isLoggedIn) {
      console.log('路由守卫: ✅ 用户已登录，允许访问:', to.path)
      next()
    } else {
      console.log('路由守卫: ❌ 用户未登录，重定向到登录页')
      // 避免重复重定向到登录页
      if (to.path !== '/login') {
        next({ path: '/login', replace: true })
      } else {
        next()
      }
    }
  } else {
    // 不需要认证的页面
    if (to.path === '/login' && userStore.isLoggedIn) {
      console.log('路由守卫: 🔄 已登录用户访问登录页，重定向到首页')
      // 已登录用户访问登录页，重定向到首页
      next({ path: '/home', replace: true })
    } else {
      console.log('路由守卫: ✅ 访问公开页面:', to.path)
      next()
    }
  }
  console.log(`=== 路由守卫结束 ===`)
})

export default router 