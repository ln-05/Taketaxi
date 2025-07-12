<template>
  <div class="home-page">
    <!-- 地图区域 -->
    <MapSection
      :current-city="currentCity"
      :location-loading="locationLoading"
      :location-error="locationError"
      :map-loading="mapLoading"
      :destination="destination"
      :current-location-text="currentLocation"
      :route-data="routeData"
      @show-city-selector="showCitySelector"
      @refresh-location="refreshLocation"
      @center-map="centerMap"
      @show-destination-search="showDestinationSearch"
      @adjust-pickup="adjustPickup"
      @select-destination="selectQuickDestination"
      @show-route-options="showRouteOptions"
      @book-ride="bookRide"
    />
    
    <!-- 搜索页面 -->
    <SearchPage
      v-if="showSearchPage"
      :search-history="searchHistory"
      :common-addresses="commonAddresses"
      :nearby-places="nearbyPlaces"
      :current-city="currentCity"
      :current-location="currentCoords || undefined"
      @close="closeSearchPage"
      @select-result="selectSearchResult"
      @clear-history="clearSearchHistory"
      @remove-history="removeSearchHistory"
      @add-to-favorites="addToFavorites"
    />
    
    <!-- 服务网格 -->
    <ServiceGrid @select-service="goToService" />
    
    <!-- 底部导航 -->
    <BottomNav :active-tab="activeTab" @change-tab="handleTabChange" />

    <!-- 目的地搜索弹窗 -->
    <el-dialog v-model="destinationDialogVisible" title="选择目的地" width="90%">
      <div class="destination-search">
        <el-input
          v-model="searchKeyword"
          placeholder="搜索地点"
          size="large"
          @input="searchLocation"
        >
          <template #prefix>
            <el-icon><Search /></el-icon>
          </template>
        </el-input>
        
        <div class="search-results">
          <div 
            v-for="(item, index) in searchResults" 
            :key="index"
            class="search-result-item"
            @click="selectDestination(item)"
          >
            <el-icon><LocationFilled /></el-icon>
            <div class="result-info">
              <div class="result-name">{{ item.name }}</div>
              <div class="result-address">{{ item.address }}</div>
            </div>
          </div>
        </div>
      </div>
    </el-dialog>

    <!-- 城市选择弹窗 -->
    <el-dialog 
      v-model="citySelectorVisible" 
      title="选择城市" 
      width="90%"
      :show-close="false"
    >
      <div class="city-selector-content">
        <div class="current-city">
          <div class="section-title">当前城市</div>
          <div class="city-item current" @click="selectCity(currentCity)">
            <el-icon><Location /></el-icon>
            <span>{{ currentCity }}</span>
            <el-icon class="check"><Check /></el-icon>
          </div>
        </div>
        
        <div class="hot-cities">
          <div class="section-title">热门城市</div>
          <div class="city-grid">
            <div 
              v-for="city in hotCities" 
              :key="city"
              class="city-item"
              @click="selectCity(city)"
            >
              {{ city }}
            </div>
          </div>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { 
  Location, Check, Search, LocationFilled
} from '@element-plus/icons-vue'

// 导入子组件
import MapSection from './components/MapSection.vue'
import ServiceGrid from './components/ServiceGrid.vue'
import BottomNav from './components/BottomNav.vue'
import SearchPage from './components/SearchPage.vue'

// 导入地图服务
import { baiduMapService, type RouteResult } from '@/utils/baiduMapService'
import { HomeFilled, OfficeBuilding, Connection } from '@element-plus/icons-vue'

const router = useRouter()

// 响应式数据
const activeTab = ref('home')
const currentCity = ref('选择城市')
const currentLocation = ref('正在定位中...')
const locationLoading = ref(true)
const destination = ref('')
const destinationDialogVisible = ref(false)
const citySelectorVisible = ref(false)
const searchKeyword = ref('')
const currentCoords = ref<{ lat: number; lng: number } | null>(null)
const locationError = ref('')
const mapLoading = ref(true)
const showSearchPage = ref(false)
const routeData = ref<RouteResult | undefined>(undefined)

// 搜索历史和常用地址
const searchHistory = ref([
  { name: '苏州市吴中区国瑞.熙墅', address: '江苏省苏州市中心区', timestamp: Date.now() - 86400000 },
  { name: '上海市浦东新区东禄素质教育产业园区', address: '上海市浦东新区', timestamp: Date.now() - 172800000 }
])

const commonAddresses = ref([
  { label: '家', address: '设置家庭地址', icon: HomeFilled, color: '#52c41a' },
  { label: '公司', address: '设置公司地址', icon: OfficeBuilding, color: '#1890ff' },
  { label: '收藏夹', address: '', icon: Connection, color: '#722ed1' }
])

const nearbyPlaces = ref([
  { name: '虹山湖公园管理处', address: '贵州省安顺市西秀区虹山湖30号', distance: '32m' },
  { name: '虹泽·山水云著营销中心', address: '贵州省安顺市西秀区', distance: '61m' },
  { name: '小坡上', address: '贵州省安顺市西秀区东关大街道小坡上', distance: '38m' }
])

const searchResults = ref([
  { name: '北京首都国际机场', address: '北京市朝阳区机场路' },
  { name: '北京南站', address: '北京市丰台区永外大街车站路' },
  { name: '天安门广场', address: '北京市东城区长安街' },
  { name: '故宫博物院', address: '北京市东城区景山前街4号' },
  { name: '王府井大街', address: '北京市东城区王府井大街' }
])

// 热门城市数据
const hotCities = ref([
  '北京', '上海', '广州', '深圳', '杭州', '南京',
  '武汉', '成都', '重庆', '西安', '天津', '苏州'
])

// 百度地图实例
let map: any = null

// 显示城市选择器
const showCitySelector = () => {
  citySelectorVisible.value = true
}

// 选择城市
const selectCity = async (city: string) => {
  currentCity.value = city
  citySelectorVisible.value = false
  ElMessage.success(`已切换到${city}`)
  
  // 当切换城市时，尝试获取该城市的中心坐标
  try {
    const cityCoords = await baiduMapService.geocode(city)
    if (cityCoords) {
      // 更新当前坐标为城市中心
      currentCoords.value = { lat: cityCoords.lat, lng: cityCoords.lng }
      
      // 获取城市中心的地址描述
      const address = await baiduMapService.reverseGeocode(cityCoords.lng, cityCoords.lat)
      if (address) {
        currentLocation.value = address
      }
      
      // 重新初始化地图
      setTimeout(() => {
        initBaiduMap()
      }, 500)
    }
  } catch (error) {
    console.error('切换城市失败:', error)
    // 如果获取城市坐标失败，仍然可以进行定位
    getCurrentLocation()
  }
}

// 处理底部导航切换
const handleTabChange = (tabId: string) => {
  activeTab.value = tabId
  
  const routeMap: Record<string, string> = {
    'home': '/',
    'car': '/owner',
    'order': '/order',
    'payment': '/payment',
    'profile': '/user'
  }
  
  if (routeMap[tabId]) {
    router.push(routeMap[tabId])
  }
}

// 获取当前位置
const getCurrentLocation = async () => {
  try {
    locationLoading.value = true
    locationError.value = ''
    
    // 检查是否支持地理定位API
    if (!navigator.geolocation) {
      currentLocation.value = '不支持定位'
      locationLoading.value = false
      ElMessage.warning('您的浏览器不支持地理定位功能')
      return
    }

    // 定位配置选项
    const options = {
      enableHighAccuracy: true, // 启用高精度定位
      timeout: 10000,          // 10秒超时
      maximumAge: 300000       // 5分钟缓存
    }

    navigator.geolocation.getCurrentPosition(
      async (position) => {
        console.log('定位成功:', position)
        
        const { latitude, longitude } = position.coords
        currentCoords.value = { lat: latitude, lng: longitude }
        
        // 显示坐标信息
        const coordsText = `${latitude.toFixed(4)}, ${longitude.toFixed(4)}`
        currentLocation.value = coordsText
        
        // 使用百度地图服务获取地址和城市信息
        try {
          const address = await baiduMapService.reverseGeocode(longitude, latitude)
          if (address) {
            currentLocation.value = address
          }
          
          // 获取当前城市
          const city = await baiduMapService.getCityByCoords(longitude, latitude)
          if (city && city !== '未知城市') {
            currentCity.value = city
          }
        } catch (error) {
          console.warn('逆地理编码失败:', error)
          // 保持坐标显示
        }
        
        locationLoading.value = false
        ElMessage.success('定位成功')
        
        // 定位成功后初始化地图
        setTimeout(() => {
          initBaiduMap()
        }, 500)
      },
      (error) => {
        console.error('定位失败:', error)
        locationLoading.value = false
        
        let errorMessage = '定位失败'
        switch (error.code) {
          case error.PERMISSION_DENIED:
            errorMessage = '定位权限被拒绝'
            locationError.value = '请在浏览器设置中允许定位权限'
            ElMessage({
              message: '需要您的位置权限才能提供精准服务',
              type: 'warning',
              duration: 5000,
              showClose: true
            })
            // 使用默认城市位置
            setTimeout(() => {
              if (!currentCoords.value) {
                initDefaultMap()
              }
            }, 1000)
            break
          case error.POSITION_UNAVAILABLE:
            errorMessage = '位置信息不可用'
            locationError.value = '无法获取位置信息'
            ElMessage.error('位置信息不可用')
            break
          case error.TIMEOUT:
            errorMessage = '定位超时'
            locationError.value = '定位请求超时，请重试'
            ElMessage.warning('定位超时，使用默认位置')
            // 使用默认位置
            setTimeout(() => {
              if (!currentCoords.value) {
                initDefaultMap()
              }
            }, 1000)
            break
          default:
            errorMessage = '未知定位错误'
            locationError.value = '发生未知错误'
            ElMessage.error('定位失败')
            break
        }
        
        currentLocation.value = errorMessage
      },
      options
    )
    
  } catch (error) {
    console.error('获取位置失败:', error)
    currentLocation.value = '定位异常'
    locationLoading.value = false
    ElMessage.error('定位服务异常')
  }

}

// 刷新位置
const refreshLocation = () => {
  ElMessage.info('正在重新定位...')
  getCurrentLocation()
}

// 居中地图
const centerMap = () => {
  if (currentCoords.value && map) {
    const { lat, lng } = currentCoords.value
    const BMap = (window as any).BMap
    const point = new BMap.Point(lng, lat)
    map.centerAndZoom(point, 15)
    ElMessage.success(`地图已居中到: ${lat.toFixed(4)}, ${lng.toFixed(4)}`)
  } else if (!currentCoords.value) {
    ElMessage.info('暂无定位信息，无法居中地图')
  } else if (!map) {
    ElMessage.warning('地图尚未初始化')
  }
}

// 显示目的地搜索
const showDestinationSearch = () => {
  showSearchPage.value = true
}

// 搜索位置
const searchLocation = async (keyword: string) => {
  if (!keyword.trim()) {
    // 恢复默认搜索结果
    searchResults.value = [
      { name: '北京首都国际机场', address: '北京市朝阳区机场路' },
      { name: '北京南站', address: '北京市丰台区永外大街车站路' },
      { name: '天安门广场', address: '北京市东城区长安街' },
      { name: '故宫博物院', address: '北京市东城区景山前街4号' },
      { name: '王府井大街', address: '北京市东城区王府井大街' }
    ]
    return
  }

  try {
    // 模拟搜索API调用
    console.log('搜索关键词:', keyword)
    
    // 模拟搜索结果
    const mockResults = [
      { name: `${keyword} - 地点1`, address: `${keyword}相关地址1` },
      { name: `${keyword} - 地点2`, address: `${keyword}相关地址2` },
      { name: `${keyword} - 地点3`, address: `${keyword}相关地址3` }
    ]
    
    // 如果有当前位置，添加"附近的xxx"
    if (currentCoords.value) {
      mockResults.unshift({
        name: `附近的${keyword}`,
        address: `当前位置附近的${keyword}`
      })
    }
    
    searchResults.value = mockResults
  } catch (error) {
    console.error('搜索失败:', error)
    ElMessage.error('搜索失败，请重试')
  }
}

// 选择目的地
const selectDestination = (location: any) => {
  destination.value = location.name
  destinationDialogVisible.value = false
  
  // 清空搜索关键词
  searchKeyword.value = ''
  
  ElMessage.success(`已选择目的地: ${location.name}`)
  
  // 如果有当前位置和目的地，初始化地图
  if (currentCoords.value) {
    setTimeout(() => {
      initBaiduMap()
    }, 1000)
  }
}

// 前往服务
const goToService = (serviceId: string) => {
  const serviceMap: Record<string, string> = {
    'taxi': '打车服务',
    'carpool': '顺风车',
    'driver': '代驾服务',
    'intercity': '城际拼车',
    'bike': '青桔骑行',
    'discount': '特价拼车',
    'delivery': '快送跑腿',
    'loan': '借钱服务',
    'fuel': '加油充电',
    'transit': '公交地铁',
    'business': '企业用车',
    'cargo': '送货',
    'rental': '滴滴租车',
    'ticket': '火车票机票',
    'more': '更多服务'
  }
  
  ElMessage.info(`即将跳转到: ${serviceMap[serviceId]}`)
  
  // 根据服务类型跳转到不同页面
  if (serviceId === 'taxi') {
    if (!currentCoords.value) {
      ElMessage.warning('请先允许定位权限获取当前位置')
      return
    }
    if (!destination.value) {
      ElMessage.warning('请先选择目的地')
      showDestinationSearch()
      return
    }
    // 有起点和终点，跳转到订单页面
    ElMessage.success('准备为您叫车...')
    router.push('/order')
  } else if (serviceId === 'more') {
    ElMessage.info('更多服务正在开发中...')
  }
}

// 初始化百度地图（更新定位后的地图）
const initBaiduMap = async () => {
  if (!currentCoords.value) {
    console.error('当前位置不可用')
    return
  }
  
  mapLoading.value = true
  
  try {
    // 使用地图服务初始化地图
    map = baiduMapService.initMap('baidu-map', {
      center: { lng: currentCoords.value.lng, lat: currentCoords.value.lat },
      zoom: destination.value ? 13 : 15
    })
    
    // 清除之前的覆盖物
    baiduMapService.clearOverlays()
    
    // 添加起点标记
    baiduMapService.addMarker(
      currentCoords.value.lng, 
      currentCoords.value.lat, 
      {
        content: '📍 当前位置',
        title: '当前位置'
      }
    )
    
    // 如果有目的地和路线数据，显示路线
    if (destination.value && routeData.value) {
      try {
        // 从路线数据中获取终点坐标
        const endCoords = routeData.value.endPoint
        
        if (endCoords) {
          // 添加终点标记
          baiduMapService.addMarker(
            endCoords.lng,
            endCoords.lat,
            {
              content: `🎯 ${destination.value}`,
              title: destination.value
            }
          )
          
          console.log('地图标记已更新，显示路线')
        }
        
      } catch (error) {
        console.error('目的地处理失败:', error)
      }
    }
    
    mapLoading.value = false
    console.log('百度地图初始化完成')
    
  } catch (error) {
    console.error('百度地图初始化失败:', error)
    mapLoading.value = false
    ElMessage.error('地图初始化失败，请刷新页面重试')
  }
}

// 组件挂载时获取位置和初始化地图
onMounted(async () => {
  // 首先尝试根据IP获取城市
  try {
    const ipCity = await baiduMapService.getCurrentCity()
    if (ipCity && ipCity !== '未知城市') {
      currentCity.value = ipCity
    }
  } catch (error) {
    console.warn('IP定位城市失败:', error)
  }
  
  getCurrentLocation()
  
  // 延迟一点初始化地图，确保DOM已渲染
  setTimeout(() => {
    initDefaultMap()
  }, 1000)
})

// 初始化默认地图（显示默认位置）
const initDefaultMap = () => {
  const BMap = (window as any).BMap
  if (!BMap) {
    console.warn('百度地图API未加载，稍后重试...')
    // 延迟重试
    setTimeout(() => {
      if ((window as any).BMap) {
        initDefaultMap()
      } else {
        mapLoading.value = false
        ElMessage.error('地图加载失败，请刷新页面重试')
      }
    }, 2000)
    return
  }
  
  const mapContainer = document.getElementById('baidu-map')
  if (!mapContainer) {
    console.warn('地图容器不存在，稍后重试...')
    setTimeout(() => {
      initDefaultMap()
    }, 500)
    return
  }
  
  try {
    // 使用默认位置初始化地图（如果还没有定位信息）
    if (!currentCoords.value && !map) {
      // 使用北京天安门作为默认中心点
      const defaultPoint = new BMap.Point(116.397, 39.909)
      
      map = new BMap.Map('baidu-map')
      map.centerAndZoom(defaultPoint, 12)
      map.enableScrollWheelZoom(true)
      
      // 添加地图控件
      const navigationControl = new BMap.NavigationControl({
        anchor: (window as any).BMAP_ANCHOR_TOP_RIGHT,
        type: (window as any).BMAP_NAVIGATION_CONTROL_SMALL
      })
      map.addControl(navigationControl)
      
      const scaleControl = new BMap.ScaleControl({
        anchor: (window as any).BMAP_ANCHOR_BOTTOM_LEFT
      })
      map.addControl(scaleControl)
      
      mapLoading.value = false
      console.log('默认地图初始化完成')
    } else {
      mapLoading.value = false
    }
  } catch (error) {
    console.error('默认地图初始化失败:', error)
    mapLoading.value = false
    ElMessage.error('地图初始化失败')
  }
}

// 新增的方法处理
const adjustPickup = () => {
  ElMessage.info('调整上车点功能')
}

const selectQuickDestination = (type: string) => {
  ElMessage.info(`选择了快捷选项: ${type}`)
}

const showRouteOptions = () => {
  ElMessage.info('显示路线选项')
}

const bookRide = () => {
  ElMessage.info('预约出行')
  router.push('/order')
}

const closeSearchPage = () => {
  showSearchPage.value = false
}

const selectSearchResult = async (result: any) => {
  destination.value = result.name
  showSearchPage.value = false
  
  // 添加到搜索历史
  searchHistory.value.unshift({
    name: result.name,
    address: result.address,
    timestamp: Date.now()
  })
  
  // 限制历史记录数量
  if (searchHistory.value.length > 10) {
    searchHistory.value = searchHistory.value.slice(0, 10)
  }
  
  ElMessage.success(`已选择目的地: ${result.name}`)
  
  // 如果有起点和终点，计算路线
  if (currentCoords.value && result.lng && result.lat) {
    try {
      const route = await baiduMapService.planRoute(
        { lng: currentCoords.value.lng, lat: currentCoords.value.lat },
        { lng: result.lng, lat: result.lat },
        {
          onComplete: (routeResult: any) => {
            console.log('路线规划完成:', routeResult)
          },
          onError: (error: any) => {
            console.error('路线规划出错:', error)
            ElMessage.error('路线规划失败')
          }
        }
      )
      
      if (route) {
        routeData.value = route
        ElMessage.success(`路线规划成功: ${route.distance}, ${route.duration}`)
        
        // 重新初始化地图以显示路线
        setTimeout(() => {
          initBaiduMap()
        }, 500)
      } else {
        ElMessage.warning('无法规划路线，请检查起点和终点')
      }
    } catch (error) {
      console.error('路线规划失败:', error)
      ElMessage.error('路线规划失败')
    }
  } else if (!currentCoords.value) {
    ElMessage.warning('请先允许定位权限获取当前位置')
  } else {
    ElMessage.warning('目的地坐标信息不完整')
  }
}

const clearSearchHistory = () => {
  searchHistory.value = []
  ElMessage.success('历史记录已清空')
}

const removeSearchHistory = (index: number) => {
  searchHistory.value.splice(index, 1)
  ElMessage.success('已删除')
}

const addToFavorites = (result: any) => {
  ElMessage.success(`已添加到收藏: ${result.name}`)
}
</script>

<style lang="scss" scoped>
.home-page {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  background: #f5f7fa;
  padding-bottom: 56px; // 为底部导航留出空间
}

// 城市选择弹窗样式
.city-selector-content {
  .section-title {
    font-size: 12px;
    color: #666;
    margin: 16px 0 8px 0;
    padding: 0 4px;
  }
  
  .current-city {
    .city-item.current {
      display: flex;
      align-items: center;
      gap: 8px;
      padding: 12px;
      background: #e3f2fd;
      border-radius: 8px;
      cursor: pointer;
      
      .el-icon {
        color: #1890ff;
      }
      
      span {
        flex: 1;
        color: #1890ff;
        font-weight: 500;
      }
      
      .check {
        color: #1890ff;
      }
    }
  }
  
  .hot-cities {
    .city-grid {
      display: grid;
      grid-template-columns: repeat(3, 1fr);
      gap: 8px;
      
      .city-item {
        padding: 12px;
        text-align: center;
        background: #f8f9fa;
        border-radius: 8px;
        cursor: pointer;
        transition: all 0.2s ease;
        
        &:hover {
          background: #e9ecef;
        }
      }
    }
  }
}

// 目的地搜索弹窗样式
.destination-search-content {
  .el-input {
    margin-bottom: 20px;
  }
  
  .search-results {
    max-height: 400px;
    overflow-y: auto;
  }
  
  .search-result-item {
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 12px;
    border-bottom: 1px solid #f0f0f0;
    cursor: pointer;
    transition: background-color 0.2s;
    
    &:hover {
      background-color: #f8f9fa;
    }
    
    .el-icon {
      color: #1890ff;
      font-size: 18px;
    }
    
    .result-info {
      flex: 1;
      
      .result-name {
        font-weight: 500;
        color: #333;
        margin-bottom: 4px;
      }
      
      .result-address {
        font-size: 12px;
        color: #666;
      }
    }
  }
}
</style>