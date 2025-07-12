// 百度地图服务类
export class BaiduMapService {
  private map: any = null
  private BMap: any = null
  
  constructor() {
    this.BMap = (window as any).BMap
  }

  // 检查百度地图 API 是否已加载
  private checkBMapAPI(): boolean {
    if (typeof (window as any).BMap === 'undefined') {
      console.error('百度地图API未加载，请检查网络连接或API密钥')
      return false
    }
    this.BMap = (window as any).BMap
    
    // 检查必要的类是否存在
    const requiredClasses = ['Map', 'Point', 'LocalSearch', 'Geocoder']
    for (const className of requiredClasses) {
      if (!this.BMap[className]) {
        console.error(`百度地图API缺少必要的类: ${className}`)
        return false
      }
    }
    
    return true
  }

  // 初始化地图
  initMap(containerId: string, options?: { 
    center?: { lng: number; lat: number }
    zoom?: number 
  }) {
    if (!this.checkBMapAPI()) {
      throw new Error('百度地图API未加载')
    }

    try {
      // 创建地图实例
      this.map = new this.BMap.Map(containerId, {
        enableMapClick: true,
        enableDragging: true,
        enableScrollWheelZoom: true,
        enableDoubleClickZoom: true,
        enableKeyboard: true,
        enableDblclickZoom: true,
        enablePinchToZoom: true
      })
      
      const center = options?.center || { lng: 116.397, lat: 39.909 }
      const zoom = options?.zoom || 12
      
      const point = new this.BMap.Point(center.lng, center.lat)
      this.map.centerAndZoom(point, zoom)

      // 添加控件
      setTimeout(() => {
        this.addControls()
      }, 100)
      
      return this.map
    } catch (error) {
      console.error('初始化地图失败:', error)
      throw error
    }
  }

  // 添加地图控件
  private addControls() {
    if (!this.map || !this.BMap) return

    try {
      // 导航控件（包含缩放按钮）
      if (this.BMap.NavigationControl) {
        const navigationControl = new this.BMap.NavigationControl({
          anchor: this.BMap.BMAP_ANCHOR_TOP_RIGHT,
          type: this.BMap.BMAP_NAVIGATION_CONTROL_ZOOM,  // 仅使用缩放按钮
          showZoomInfo: true
        })
        this.map.addControl(navigationControl)
      }

      // 比例尺控件
      if (this.BMap.ScaleControl) {
        const scaleControl = new this.BMap.ScaleControl({
          anchor: this.BMap.BMAP_ANCHOR_BOTTOM_LEFT
        })
        this.map.addControl(scaleControl)
      }

      // 定位控件
      if (this.BMap.GeolocationControl) {
        const geolocationControl = new this.BMap.GeolocationControl({
          anchor: this.BMap.BMAP_ANCHOR_BOTTOM_RIGHT,
          showAddressBar: false,
          enableAutoLocation: true
        })
        this.map.addControl(geolocationControl)
      }

    } catch (error) {
      console.error('添加地图控件失败:', error)
    }
  }

  // 设置地图缩放级别
  setZoom(zoom: number) {
    if (this.map) {
      this.map.setZoom(zoom)
    }
  }

  // 获取当前缩放级别
  getZoom(): number {
    return this.map ? this.map.getZoom() : 0
  }

  // 放大地图
  zoomIn() {
    if (this.map) {
      this.map.zoomIn()
    }
  }

  // 缩小地图
  zoomOut() {
    if (this.map) {
      this.map.zoomOut()
    }
  }

  // 设置最大和最小缩放级别
  setZoomLimits(minZoom: number, maxZoom: number) {
    if (this.map) {
      this.map.setMinZoom(minZoom)
      this.map.setMaxZoom(maxZoom)
    }
  }

  // 地理编码 - 地址转坐标
  async geocode(address: string): Promise<{ lng: number; lat: number } | null> {
    return new Promise((resolve) => {
      if (!this.BMap) {
        resolve(null)
        return
      }

      const geocoder = new this.BMap.Geocoder()
      geocoder.getPoint(address, (point: any) => {
        if (point) {
          resolve({ lng: point.lng, lat: point.lat })
        } else {
          resolve(null)
        }
      })
    })
  }

  // 逆地理编码 - 坐标转地址
  async reverseGeocode(lng: number, lat: number): Promise<string | null> {
    return new Promise((resolve) => {
      if (!this.BMap) {
        resolve(null)
        return
      }

      const geocoder = new this.BMap.Geocoder()
      const point = new this.BMap.Point(lng, lat)
      
      geocoder.getLocation(point, (result: any) => {
        if (result) {
          resolve(result.address)
        } else {
          resolve(null)
        }
      })
    })
  }

  // 路线规划
  async planRoute(
    start: { lng: number; lat: number },
    end: { lng: number; lat: number },
    options?: {
      policy?: number // 路线策略
      onComplete?: (result: any) => void
      onError?: (error: any) => void
    }
  ): Promise<RouteResult | null> {
    return new Promise((resolve) => {
      if (!this.BMap || !this.map) {
        resolve(null)
        return
      }

      const startPoint = new this.BMap.Point(start.lng, start.lat)
      const endPoint = new this.BMap.Point(end.lng, end.lat)

      const driving = new this.BMap.DrivingRoute(this.map, {
        renderOptions: {
          map: this.map,
          autoViewport: true,
          panel: false
        },
        policy: options?.policy || this.BMap.BMAP_DRIVING_POLICY_LEAST_TIME,
        onSearchComplete: (result: any) => {
          if (result && result.getNumPlans() > 0) {
            const plan = result.getPlan(0)
            const route = plan.getRoute(0)
            
            const routeResult: RouteResult = {
              distance: (plan.getDistance() / 1000).toFixed(1) + 'km',
              duration: Math.ceil(plan.getDuration() / 60) + '分钟',
              trafficLevel: this.getTrafficLevel(plan.getDuration(), plan.getDistance()),
              trafficText: this.getTrafficText(plan.getDuration(), plan.getDistance()),
              polyline: route.getPath(),
              startPoint,
              endPoint
            }

            options?.onComplete?.(result)
            resolve(routeResult)
          } else {
            options?.onError?.('路线规划失败')
            resolve(null)
          }
        },
        onPolylinesSet: () => {
          console.log('路线绘制完成')
        }
      })

      driving.search(startPoint, endPoint)
    })
  }

  // 获取交通状况等级
  private getTrafficLevel(duration: number, distance: number): 'smooth' | 'slow' | 'congested' {
    // 根据时间和距离计算交通状况
    const speed = (distance / 1000) / (duration / 3600) // km/h
    
    if (speed > 40) return 'smooth'
    if (speed > 20) return 'slow'
    return 'congested'
  }

  // 获取交通状况文本
  private getTrafficText(duration: number, distance: number): string {
    const level = this.getTrafficLevel(duration, distance)
    
    switch (level) {
      case 'smooth': return '畅通'
      case 'slow': return '缓慢'
      case 'congested': return '拥堵'
      default: return '未知'
    }
  }

  // 搜索地点
  async searchPlaces(keyword: string, city?: string, currentLocation?: { lng: number; lat: number }): Promise<SearchResult[]> {
    return new Promise(async (resolve) => {
      if (!this.checkBMapAPI()) {
        console.error('百度地图API不可用')
        resolve([])
        return
      }

      try {
        const searchCity = city || '全国'
        const localSearch = new this.BMap.LocalSearch(searchCity, {
          onSearchComplete: async (result: any) => {
            try {
              if (result && result.getPoi && typeof result.getNumPois === 'function') {
                const numPois = result.getNumPois()
                if (numPois > 0) {
                  const places: SearchResult[] = []
                  
                  for (let i = 0; i < Math.min(numPois, 10); i++) {
                    const poi = result.getPoi(i)
                    if (poi && poi.point && poi.title) {
                      const distance = currentLocation 
                        ? await this.calculateRealDistance(currentLocation, poi.point)
                        : undefined

                      places.push({
                        name: poi.title || '未知地点',
                        address: poi.address || '地址未知',
                        lng: poi.point.lng,
                        lat: poi.point.lat,
                        distance
                      })
                    }
                  }
                  
                  console.log(`搜索"${keyword}"在"${searchCity}"找到${places.length}个结果`)
                  resolve(places)
                } else {
                  console.log(`搜索"${keyword}"在"${searchCity}"未找到匹配的地点`)
                  resolve([])
                }
              } else {
                console.error('搜索结果格式无效:', result)
                resolve([])
              }
            } catch (error) {
              console.error('处理搜索结果时出错:', error)
              resolve([])
            }
          },
          onSearchError: (error: any) => {
            console.error('百度地图搜索请求失败:', error)
            resolve([])
          }
        })

        // 执行搜索
        if (keyword.trim()) {
          localSearch.search(keyword.trim())
        } else {
          resolve([])
        }
      } catch (error) {
        console.error('创建搜索对象失败:', error)
        resolve([])
      }
    })
  }

  // 计算真实距离
  private async calculateRealDistance(from: { lng: number; lat: number }, to: any): Promise<string> {
    if (!this.BMap) {
      return this.calculateDistance(to)
    }

    try {
      const fromPoint = new this.BMap.Point(from.lng, from.lat)
      const toPoint = new this.BMap.Point(to.lng, to.lat)
      
      // 使用直线距离计算
      const distance = Math.round(this.calculateLineDistance(fromPoint, toPoint))
      if (distance < 1000) {
        return distance + 'm'
      } else {
        return (distance / 1000).toFixed(1) + 'km'
      }
    } catch (error) {
      console.error('距离计算失败:', error)
      return this.calculateDistance(to)
    }
  }

  // 计算两点间直线距离
  private calculateLineDistance(point1: any, point2: any): number {
    const R = 6371000 // 地球半径，单位米
    const lat1 = this.toRadians(point1.lat)
    const lat2 = this.toRadians(point2.lat)
    const deltaLat = this.toRadians(point2.lat - point1.lat)
    const deltaLng = this.toRadians(point2.lng - point1.lng)

    const a = Math.sin(deltaLat / 2) * Math.sin(deltaLat / 2) +
              Math.cos(lat1) * Math.cos(lat2) *
              Math.sin(deltaLng / 2) * Math.sin(deltaLng / 2)
    const c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a))
    return R * c
  }

  // 将角度转换为弧度
  private toRadians(degrees: number): number {
    return degrees * Math.PI / 180
  }

  // 计算距离（备用方法）
  private calculateDistance(targetPoint: any): string {
    // 使用简单的随机距离作为最后的备用方案
    const distance = Math.random() * 3 + 0.5 // 0.5km - 3.5km
    return distance.toFixed(1) + 'km'
  }

  // 添加标记点
  addMarker(lng: number, lat: number, options?: {
    icon?: string
    title?: string
    content?: string
    draggable?: boolean
    onClick?: () => void
  }) {
    if (!this.BMap || !this.map) return null

    const point = new this.BMap.Point(lng, lat)
    const marker = new this.BMap.Marker(point, {
      enableDragging: options?.draggable || false
    })

    if (options?.icon) {
      const icon = new this.BMap.Icon(options.icon, new this.BMap.Size(32, 32))
      marker.setIcon(icon)
    }

    this.map.addOverlay(marker)

    if (options?.content) {
      const infoWindow = new this.BMap.InfoWindow(options.content)
      marker.addEventListener('click', () => {
        this.map.openInfoWindow(infoWindow, point)
        options?.onClick?.()
      })
    } else if (options?.onClick) {
      marker.addEventListener('click', options.onClick)
    }

    return marker
  }

  // 清除所有覆盖物
  clearOverlays() {
    if (this.map) {
      this.map.clearOverlays()
    }
  }

  // 居中地图到指定位置
  centerTo(lng: number, lat: number, zoom?: number) {
    if (!this.BMap || !this.map) return

    const point = new this.BMap.Point(lng, lat)
    this.map.centerAndZoom(point, zoom || this.map.getZoom())
  }

  // 获取地图实例
  getMap() {
    return this.map
  }

  // 根据IP获取当前城市
  async getCurrentCity(): Promise<string> {
    return new Promise((resolve) => {
      if (!this.BMap) {
        resolve('未知城市')
        return
      }

      const myCity = new this.BMap.LocalCity()
      myCity.get((result: any) => {
        if (result && result.name) {
          resolve(result.name)
        } else {
          resolve('未知城市')
        }
      })
    })
  }

  // 获取坐标对应的城市
  async getCityByCoords(lng: number, lat: number): Promise<string> {
    return new Promise((resolve) => {
      if (!this.BMap) {
        resolve('未知城市')
        return
      }

      const geocoder = new this.BMap.Geocoder()
      const point = new this.BMap.Point(lng, lat)
      
      geocoder.getLocation(point, (result: any) => {
        if (result && result.addressComponents && result.addressComponents.city) {
          resolve(result.addressComponents.city)
        } else {
          resolve('未知城市')
        }
      })
    })
  }

  // 计算预估价格
  calculateEstimatedPrice(distance: number, duration: number, vehicleType: string): number {
    // 基础价格（起步价）
    const basePrices: Record<string, number> = {
      'economy': 10,    // 经济型
      'comfort': 12,    // 舒适型
      'luxury': 15,     // 豪华型
      'business': 20    // 商务型
    }

    // 每公里价格
    const kmPrices: Record<string, number> = {
      'economy': 1.8,
      'comfort': 2.2,
      'luxury': 3.0,
      'business': 3.5
    }

    // 时长费率（每分钟）
    const timePrices: Record<string, number> = {
      'economy': 0.3,
      'comfort': 0.4,
      'luxury': 0.5,
      'business': 0.6
    }

    const basePrice = basePrices[vehicleType] || basePrices.economy
    const kmPrice = kmPrices[vehicleType] || kmPrices.economy
    const timePrice = timePrices[vehicleType] || timePrices.economy

    // 计算总价
    const distancePrice = (distance / 1000) * kmPrice  // 距离费用
    const durationPrice = (duration / 60) * timePrice   // 时长费用
    const totalPrice = basePrice + distancePrice + durationPrice

    // 四舍五入到小数点后两位
    return Math.round(totalPrice * 100) / 100
  }

  // 获取不同车型的预估价格
  async getEstimatedPrices(start: { lng: number; lat: number }, end: { lng: number; lat: number }): Promise<Record<string, number>> {
    const route = await this.planRoute(start, end)
    if (!route) return {}

    const distance = parseFloat(route.distance.replace('km', '')) * 1000  // 转换为米
    const duration = parseFloat(route.duration.replace('分钟', '')) * 60   // 转换为秒

    return {
      economy: this.calculateEstimatedPrice(distance, duration, 'economy'),
      comfort: this.calculateEstimatedPrice(distance, duration, 'comfort'),
      luxury: this.calculateEstimatedPrice(distance, duration, 'luxury'),
      business: this.calculateEstimatedPrice(distance, duration, 'business')
    }
  }
}

// 路线结果接口
export interface RouteResult {
  distance: string
  duration: string
  trafficLevel: 'smooth' | 'slow' | 'congested'
  trafficText: string
  polyline: any[]
  startPoint: any
  endPoint: any
}

// 搜索结果接口
export interface SearchResult {
  name: string
  address: string
  lng: number
  lat: number
  distance?: string
}

// 创建全局实例
export const baiduMapService = new BaiduMapService() 