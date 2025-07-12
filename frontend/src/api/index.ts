import request from '@/utils/request'

// 用户相关接口
export const userApi = {
  // 发送短信验证码
  sendSms: (data: { Phone: string; Source: string }) => {
    return request.post('/v1/sms/send', data)
  },

  // 用户登录
  login: (data: { Phone: string; SendSms: string; Name?: string; Avatar?: string; Balance?: number }) => {
    return request.post('/v1/user/login', data)
  },

  // 实名认证
  realName: (data: { mobile: string; Name: string; Identification?: string }) => {
    return request.post('/v1/user/real', data)
  },

  // 下单
  placeOrder: (data: {
    user_id: number
    order_type: number
    vehicle_type: number
    start_location: string
    end_location: string
    scheduled_departure_time?: string
    coupon_id?: number
    OrderAmount?: number
  }) => {
    return request.post('/v1/order/place', data)
  }
}

// 微信相关接口
export const wechatApi = {
  // 获取登录二维码
  getLoginQRCode: (url: string) => {
    return request.get('/v1/wechat/login', {
      params: { url },
      responseType: 'blob'
    })
  },

  // 处理微信登录回调
  handleCallback: (params: { code: string; state: string }) => {
    return request.get('/v1/wechat/callback', { params })
  },

  // 微信签名验证
  checkSignature: (params: { signature: string; timestamp: string; nonce: string; echostr: string }) => {
    return request.get('/v1/wechat/check', { params })
  }
}

// 搜索历史记录相关接口
export const searchHistoryApi = {
  // 获取搜索历史
  getSearchHistory: (params: { user_id: number; page: number; page_size: number }) => {
    return request.get('/v1/search/history', { params })
  },

  // 添加搜索历史
  addSearchHistory: (data: {
    user_id: number
    keyword: string
    start_location: string
    end_location: string
  }) => {
    return request.post('/v1/search/history', data)
  },

  // 删除单条搜索历史
  deleteSearchHistory: (id: number, params: { user_id: number }) => {
    return request.delete(`/v1/search/history/${id}`, { params })
  },

  // 清空搜索历史
  clearSearchHistory: (params: { user_id: number }) => {
    return request.delete('/v1/search/history/all', { params })
  }
}

// 订单相关接口（预留）
export const orderApi = {
  // 可以在这里添加订单相关的API
}

// 地图相关接口（预留）
export const mapApi = {
  // 可以在这里添加地图相关的API
} 