<template>
  <div class="map-container">
    <!-- 地图区域 -->
    <div class="map-section" :style="{ height: mapHeight + 'px' }">
      <!-- 顶部工具栏 -->
      <div class="map-toolbar">
        <div class="location-info" @click="emit('show-city-selector')">
          <el-icon><Location /></el-icon>
          <span>{{ currentCity }}</span>
          <el-icon><ArrowDown /></el-icon>
        </div>
        
        <div class="location-status">
          <template v-if="locationLoading">
            <el-icon class="is-loading"><Loading /></el-icon>
            <span>正在定位...</span>
          </template>
          <template v-else-if="locationError">
            <el-icon color="#ff4d4f"><Warning /></el-icon>
            <span class="error-text">{{ locationError }}</span>
            <el-button size="small" text @click="emit('refresh-location')">
              重试
            </el-button>
          </template>
          <template v-else>
            <span>{{ currentLocationText }}</span>
            <el-button size="small" text @click="emit('center-map')">
              <el-icon><Aim /></el-icon>
            </el-button>
          </template>
        </div>
      </div>

      <!-- 地图容器 -->
      <div id="baidu-map" class="baidu-map"></div>
    </div>

    <!-- 拖动分隔条 -->
    <div class="resize-handle" 
         @mousedown="startResize"
         @touchstart="startResize">
      <div class="handle-line"></div>
    </div>

    <!-- 路线选择和搜索区域 -->
    <div class="bottom-section">
      <!-- 路线选择面板 -->
      <div v-if="routeData" class="route-selection">
        <div class="route-header">
          <div class="title">推荐路线</div>
          <div class="toggle" @click="toggleRouteDisplay">
            {{ isRouteVisible ? '隐藏路线' : '显示路线' }}
          </div>
        </div>
        
        <div class="route-info">
          <div class="time">
            <el-icon><Timer /></el-icon>
            <span>{{ formatDuration(routeData.duration) }}</span>
          </div>
          <div class="distance">
            <el-icon><Location /></el-icon>
            <span>{{ formatDistance(routeData.distance) }}</span>
          </div>
          <div class="traffic" :class="routeData.trafficLevel">
            <el-icon><InfoFilled /></el-icon>
            <span>{{ routeData.trafficText }}</span>
          </div>
        </div>

        <div class="route-actions">
          <el-button type="primary" class="book-btn" @click="emit('book-ride')">
            立即叫车
          </el-button>
          <el-button class="options-btn" @click="showRouteOptions">
            更多选项
          </el-button>
        </div>
      </div>

      <!-- 底部搜索栏 -->
      <div class="bottom-toolbar" v-if="!mapLoading">
        <div class="destination-input" @click="emit('show-destination-search')">
          <el-input
            :model-value="destination"
            placeholder="您要去哪儿？"
            readonly
            :suffix-icon="Search"
          />
        </div>
      </div>
    </div>

    <!-- 加载状态 -->
    <div class="map-loading" v-if="mapLoading">
      <el-icon class="is-loading"><Loading /></el-icon>
      <span>地图加载中...</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import { baiduMapService } from '@/utils/baiduMapService'
import {
  Location,
  ArrowDown,
  Loading,
  Warning,
  Aim,
  Search,
  Timer,
  InfoFilled
} from '@element-plus/icons-vue'

const props = defineProps<{
  currentCity: string
  locationLoading: boolean
  locationError: string
  mapLoading: boolean
  destination: string
  currentLocationText: string
  routeData?: {
    distance: string
    duration: string
    trafficLevel: 'smooth' | 'slow' | 'congested'
    trafficText: string
    polyline: any[]
    startPoint: any
    endPoint: any
  }
}>()

const emit = defineEmits<{
  (e: 'show-city-selector'): void
  (e: 'refresh-location'): void
  (e: 'center-map'): void
  (e: 'show-destination-search'): void
  (e: 'adjust-pickup'): void
  (e: 'select-destination', type: string): void
  (e: 'show-route-options'): void
  (e: 'book-ride'): void
}>()

// 路线显示控制
const isRouteVisible = ref(true)

// 切换路线显示状态
const toggleRouteDisplay = () => {
  isRouteVisible.value = !isRouteVisible.value
  if (isRouteVisible.value) {
    showRoute()
  } else {
    hideRoute()
  }
}

// 显示路线
const showRoute = () => {
  if (!props.routeData) return
  
  baiduMapService.clearOverlays()
  
  // 添加起点标记
  if (props.routeData.startPoint) {
    baiduMapService.addMarker(
      props.routeData.startPoint.lng,
      props.routeData.startPoint.lat,
      { title: '起点', content: '📍 当前位置' }
    )
  }

  // 添加终点标记
  if (props.routeData.endPoint) {
    baiduMapService.addMarker(
      props.routeData.endPoint.lng,
      props.routeData.endPoint.lat,
      { title: '终点', content: `🎯 ${props.destination}` }
    )
  }

  // 绘制路线
  if (props.routeData.polyline && props.routeData.polyline.length > 0) {
    const BMap = (window as any).BMap
    const points = props.routeData.polyline.map(p => new BMap.Point(p.lng, p.lat))
    const polyline = new BMap.Polyline(points, {
      strokeColor: "#1890ff",
      strokeWeight: 6,
      strokeOpacity: 0.8
    })
    baiduMapService.getMap().addOverlay(polyline)
    baiduMapService.getMap().setViewport(points)
  }
}

// 隐藏路线
const hideRoute = () => {
  baiduMapService.clearOverlays()
  
  // 只保留起点标记
  if (props.routeData?.startPoint) {
    baiduMapService.addMarker(
      props.routeData.startPoint.lng,
      props.routeData.startPoint.lat,
      { title: '当前位置', content: '📍 当前位置' }
    )
  }
}

// 监听路线数据变化
watch(() => props.routeData, (newRouteData) => {
  if (newRouteData && isRouteVisible.value) {
    showRoute()
  }
}, { deep: true })

// 格式化时间
const formatDuration = (duration: string) => {
  const minutes = parseInt(duration)
  if (isNaN(minutes)) return '计算中...'
  if (minutes < 60) return `${minutes}分钟`
  const hours = Math.floor(minutes / 60)
  const remainingMinutes = minutes % 60
  return remainingMinutes > 0 ? `${hours}小时${remainingMinutes}分钟` : `${hours}小时`
}

// 格式化距离
const formatDistance = (distance: string) => {
  const km = parseFloat(distance)
  if (isNaN(km)) return '计算中...'
  if (km < 1) return `${Math.round(km * 1000)}米`
  return `${km.toFixed(1)}公里`
}

// 显示路线选项
const showRouteOptions = () => {
  emit('show-route-options')
}

// 地图高度控制
const mapHeight = ref(window.innerHeight * 0.7) // 初始高度为70%
const minMapHeight = 200 // 最小高度
const maxMapHeight = window.innerHeight - 150 // 最大高度，留出底部空间

// 拖动调整大小相关
const isResizing = ref(false)
const startY = ref(0)
const startHeight = ref(0)

// 开始拖动
const startResize = (e: MouseEvent | TouchEvent) => {
  isResizing.value = true
  if (e instanceof MouseEvent) {
    startY.value = e.clientY
  } else {
    startY.value = e.touches[0].clientY
  }
  startHeight.value = mapHeight.value

  // 添加事件监听
  document.addEventListener('mousemove', handleResize)
  document.addEventListener('mouseup', stopResize)
  document.addEventListener('touchmove', handleResize)
  document.addEventListener('touchend', stopResize)
}

// 处理拖动
const handleResize = (e: MouseEvent | TouchEvent) => {
  if (!isResizing.value) return

  let currentY
  if (e instanceof MouseEvent) {
    currentY = e.clientY
  } else {
    currentY = e.touches[0].clientY
  }

  const deltaY = currentY - startY.value
  let newHeight = startHeight.value - deltaY

  // 限制高度范围
  newHeight = Math.max(minMapHeight, Math.min(newHeight, maxMapHeight))
  mapHeight.value = newHeight

  // 更新地图大小
  const map = baiduMapService.getMap()
  if (map) {
    map.invalidateSize()
  }
}

// 停止拖动
const stopResize = () => {
  isResizing.value = false
  document.removeEventListener('mousemove', handleResize)
  document.removeEventListener('mouseup', stopResize)
  document.removeEventListener('touchmove', handleResize)
  document.removeEventListener('touchend', stopResize)
}

onMounted(() => {
  try {
    // 初始化地图
    baiduMapService.initMap('baidu-map', {
      center: { lng: 121.4737, lat: 31.2304 },  // 上海中心位置
      zoom: 12
    })
    
    // 地图初始化成功后，自动触发 addControls 方法添加控件
  } catch (error) {
    console.error('地图初始化失败:', error)
  }
})
</script>

<style lang="scss" scoped>
.map-container {
  position: relative;
  width: 100%;
  height: 100vh;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.map-section {
  position: relative;
  width: 100%;
  min-height: 200px;
  transition: height 0.05s ease;
}

.baidu-map {
  width: 100%;
  height: 100%;
}

.resize-handle {
  width: 100%;
  height: 24px;
  background: #fff;
  cursor: row-resize;
  user-select: none;
  touch-action: none;
  display: flex;
  justify-content: center;
  align-items: center;
  box-shadow: 0 -2px 8px rgba(0,0,0,0.05);
  z-index: 100;

  .handle-line {
    width: 40px;
    height: 4px;
    background: #e0e0e0;
    border-radius: 2px;
    
    &:hover {
      background: #d0d0d0;
    }
  }
}

.bottom-section {
  flex: 1;
  min-height: 150px;
  background: #f5f7fa;
  position: relative;
  overflow-y: auto;
}

.map-toolbar {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  z-index: 100; // 提高层级
  padding: 12px 16px; // 增加左右内边距
  background: linear-gradient(to bottom, rgba(255,255,255,0.95), rgba(255,255,255,0.85));
  backdrop-filter: blur(10px);
  border-bottom: 1px solid rgba(0,0,0,0.05);

  .location-info {
    display: inline-flex; // 改为inline-flex使宽度自适应
    align-items: center;
    gap: 4px;
    padding: 8px 12px;
    background: #fff;
    border-radius: 20px;
    box-shadow: 0 2px 8px rgba(0,0,0,0.08);
    cursor: pointer;
    
    .el-icon {
      color: #1890ff;
    }
  }

  .location-status {
    margin-top: 8px;
    display: flex;
    align-items: center;
    gap: 8px;
    font-size: 14px;
    color: #666;
    background: rgba(255,255,255,0.8);
    padding: 4px 8px;
    border-radius: 4px;

    .error-text {
      color: #ff4d4f;
    }
  }
}

.route-selection {
  position: relative; // 改为相对定位
  margin: 16px;
  background: #fff;
  border-radius: 16px; // 增加圆角
  padding: 16px;
  box-shadow: 0 4px 16px rgba(0,0,0,0.12);
  z-index: 99; // 确保在地图控件之上

  .route-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 12px;

    .title {
      font-size: 16px;
      font-weight: 500;
      color: #333;
    }

    .toggle {
      color: #1890ff;
      font-size: 14px;
      cursor: pointer;
      padding: 4px 8px;
      border-radius: 4px;
      
      &:hover {
        background: rgba(24,144,255,0.1);
      }
    }
  }

  .route-info {
    display: flex;
    align-items: center;
    gap: 16px;
    margin-bottom: 16px;
    padding: 12px;
    background: #f5f7fa;
    border-radius: 12px;

    .time, .distance, .traffic {
      display: flex;
      align-items: center;
      gap: 6px;
      color: #666;
      font-size: 14px;
      padding: 4px 8px;
      background: rgba(255,255,255,0.8);
      border-radius: 6px;

      .el-icon {
        font-size: 16px;
        color: #1890ff;
      }
    }

    .traffic {
      &.smooth { 
        color: #52c41a;
        background: rgba(82,196,26,0.1);
      }
      &.slow { 
        color: #faad14;
        background: rgba(250,173,20,0.1);
      }
      &.congested { 
        color: #ff4d4f;
        background: rgba(255,77,79,0.1);
      }
    }
  }

  .route-actions {
    display: flex;
    gap: 12px;

    .book-btn, .options-btn {
      flex: 1;
      height: 44px; // 增加按钮高度
      border-radius: 22px;
      font-size: 16px;
    }

    .book-btn {
      background: #1890ff;
      border-color: #1890ff;
      
      &:hover {
        background: #40a9ff;
        border-color: #40a9ff;
      }
    }

    .options-btn {
      border-color: #d9d9d9;
      color: #666;
      
      &:hover {
        color: #1890ff;
        border-color: #1890ff;
        background: rgba(24,144,255,0.05);
      }
    }
  }
}

.bottom-toolbar {
  position: relative; // 改为相对定位
  padding: 16px;
  z-index: 98; // 确保在地图控件之上

  .destination-input {
    :deep(.el-input__inner) {
      border-radius: 24px;
      padding-left: 20px;
      height: 48px;
      font-size: 16px;
      box-shadow: 0 4px 16px rgba(0,0,0,0.12);
      background: rgba(255,255,255,0.95);
      backdrop-filter: blur(10px);
      
      &::placeholder {
        color: #999;
      }
    }
  }
}

.map-loading {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 12px;
  color: #666;
  background: rgba(255,255,255,0.9);
  padding: 20px 32px;
  border-radius: 12px;
  box-shadow: 0 4px 16px rgba(0,0,0,0.12);
  backdrop-filter: blur(10px);
  z-index: 101; // 最高层级

  .el-icon {
    font-size: 32px;
    color: #1890ff;
  }
}
</style> 