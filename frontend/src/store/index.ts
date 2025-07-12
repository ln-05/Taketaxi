import { defineStore } from 'pinia'
import { ref } from 'vue'

export interface User {
  id: string
  phone: string
  name?: string
  avatar?: string
}

export interface LoginForm {
  phone: string
  sendSms: string
}

export interface LoginResult {
  success: boolean
  message: string
  user?: User | null
  token?: string
}

export const useUserStore = defineStore('user', () => {
  const user = ref<User | null>(null)
  const token = ref<string>('')
  const isLoggedIn = ref(false)
  const loginTime = ref<number>(0)

  // 登录方法
  const login = async (loginForm: LoginForm): Promise<LoginResult> => {
    console.log('Store: 开始登录流程')
    
    try {
      // 测试模式：直接设置登录成功状态
      const mockUser: User = {
        id: '1',
        phone: loginForm.phone,
        name: '测试用户',
        avatar: ''
      }
      
      const mockToken = 'test-token-' + Date.now()
      const currentTime = Date.now()
      
      // 先清除可能存在的旧状态
      user.value = null
      token.value = ''
      isLoggedIn.value = false
      loginTime.value = 0
      
      // 原子性设置新的登录状态
      user.value = mockUser
      token.value = mockToken
      loginTime.value = currentTime
      // 最后设置isLoggedIn，确保其他状态都已设置完成
      isLoggedIn.value = true
      
      // 持久化存储
      localStorage.setItem('token', mockToken)
      localStorage.setItem('user', JSON.stringify(mockUser))
      localStorage.setItem('loginTime', currentTime.toString())
      
      console.log('Store: 登录状态设置完成', {
        isLoggedIn: isLoggedIn.value,
        user: user.value?.phone,
        token: mockToken.substring(0, 20) + '...'
      })
      
      // 后台调用实际API（不影响登录结果）
      setTimeout(async () => {
        try {
          const response = await fetch('/api/login', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
              Phone: loginForm.phone,
              SendSms: loginForm.sendSms
            }),
          })
          const data = await response.json()
          console.log('后端API响应:', data)
        } catch (apiError) {
          console.log('后端API调用失败 (不影响登录):', apiError)
        }
      }, 0)
      
      return {
        success: true,
        message: '登录成功',
        user: mockUser,
        token: mockToken
      }
      
    } catch (error) {
      console.error('Store: 登录过程出错:', error)
      return {
        success: false,
        message: '登录失败，请重试'
      }
    }
  }

  // 登出方法
  const logout = () => {
    user.value = null
    token.value = ''
    isLoggedIn.value = false
    loginTime.value = 0
    localStorage.removeItem('token')
    localStorage.removeItem('user')
    localStorage.removeItem('loginTime')
  }

  // 初始化用户信息（从 localStorage 恢复）
  const initUser = () => {
    console.log('初始化用户状态')
    
    const storedToken = localStorage.getItem('token')
    const storedUser = localStorage.getItem('user')
    const storedLoginTime = localStorage.getItem('loginTime')
    
    if (storedToken && storedUser && storedLoginTime) {
      const loginTimestamp = parseInt(storedLoginTime)
      const currentTime = Date.now()
      const TWO_DAYS = 2 * 24 * 60 * 60 * 1000 // 2天的毫秒数
      
      // 检查是否在2天内
      if (currentTime - loginTimestamp < TWO_DAYS) {
        token.value = storedToken
        user.value = JSON.parse(storedUser)
        loginTime.value = loginTimestamp
        isLoggedIn.value = true
        
        console.log('用户登录状态已恢复')
      } else {
        console.log('登录已过期，清除存储')
        logout()
      }
    } else {
      console.log('无有效登录信息，保持未登录状态')
    }
  }

  return {
    user,
    token,
    isLoggedIn,
    loginTime,
    login,
    logout,
    initUser
  }
}) 