<template>
  <div class="search-page">
    <!-- 搜索头部 -->
    <div class="search-header">
      <el-button 
        circle 
        size="small" 
        class="back-btn"
        @click="$emit('close')"
      >
        <el-icon><ArrowLeft /></el-icon>
      </el-button>
      <div class="search-input-wrapper">
        <el-input
          v-model="searchKeyword"
          placeholder="请输入起点"
          size="large"
          @input="handleSearch"
          @focus="showHistory = true"
          autofocus
          clearable
        >
          <template #prefix>
            <el-icon><Search /></el-icon>
          </template>
        </el-input>
      </div>
    </div>

    <!-- 快捷标签 -->
    <div class="quick-tabs">
      <div 
        v-for="tab in quickTabs" 
        :key="tab.id"
        class="tab-item"
        :class="{ active: activeTab === tab.id }"
        @click="switchTab(tab.id)"
      >
        <el-icon>
          <component :is="tab.icon" />
        </el-icon>
        <span>{{ tab.name }}</span>
      </div>
    </div>

    <!-- 内容区域 -->
    <div class="search-content">
      <!-- 搜索加载状态 -->
      <div v-if="isSearching" class="loading-state">
        <el-icon class="loading-icon"><Loading /></el-icon>
        <p>正在搜索...</p>
      </div>
      
      <!-- 搜索结果 -->
      <div v-else-if="searchKeyword && searchResults.length" class="search-results">
        <div class="result-title">搜索结果</div>
        <div 
          v-for="(result, index) in searchResults" 
          :key="index"
          class="result-item"
          @click="selectResult(result)"
        >
          <div class="result-icon">
            <el-icon><LocationFilled /></el-icon>
          </div>
          <div class="result-content">
            <div class="result-name">{{ result.name }}</div>
            <div class="result-address">{{ result.address }}</div>
            <div class="result-distance" v-if="result.distance">{{ result.distance }}</div>
          </div>
          <el-button size="small" text @click.stop="addToFavorites(result)">
            <el-icon><StarFilled /></el-icon>
          </el-button>
        </div>
      </div>

      <!-- 历史记录 -->
      <div v-else-if="showHistory" class="history-section">
        <!-- 搜索历史 -->
        <div v-if="searchHistory.length" class="history-group">
          <div class="history-header">
            <span class="history-title">历史记录</span>
            <el-button size="small" text @click="clearHistory">清空</el-button>
          </div>
          <div 
            v-for="item in searchHistory" 
            :key="item.id"
            class="history-item"
            @click="selectHistory(item)"
          >
            <div class="history-icon">
              <el-icon><Clock /></el-icon>
            </div>
                          <div class="history-content">
                <div class="history-name">{{ item.keyword }}</div>
                <div class="history-address">{{ item.endLocation || item.end_location }}</div>
                <div class="history-meta">
                  <span class="search-count">搜索次数: {{ item.searchCount || item.search_count }}</span>
                  <span class="created-at">{{ item.createdAt || item.created_at }}</span>
                </div>
              </div>
            <el-button size="small" text @click.stop="removeHistory(item)">
              <el-icon><Close /></el-icon>
            </el-button>
          </div>
          <el-pagination
            v-if="total > pageSize"
            v-model:current-page="currentPage"
            v-model:page-size="pageSize"
            :total="total"
            layout="prev, pager, next"
            @current-change="fetchSearchHistory"
          />
        </div>

        <!-- 常用地址 -->
        <div v-if="commonAddresses.length" class="common-group">
          <div class="common-header">
            <span class="common-title">常用地址</span>
          </div>
          <div 
            v-for="(address, index) in commonAddresses" 
            :key="index"
            class="common-item"
            @click="selectCommon(address)"
          >
            <div class="common-icon" :style="{ backgroundColor: address.color }">
              <el-icon>
                <component :is="address.icon" />
              </el-icon>
            </div>
            <div class="common-content">
              <div class="common-label">{{ address.label }}</div>
              <div class="common-address">{{ address.address || '设置常用地址' }}</div>
            </div>
            <el-icon class="arrow-icon"><ArrowRight /></el-icon>
          </div>
        </div>

        <!-- 附近地点 -->
        <div v-if="nearbyPlaces.length" class="nearby-group">
          <div class="nearby-header">
            <span class="nearby-title">附近地点</span>
          </div>
          <div 
            v-for="(place, index) in nearbyPlaces" 
            :key="index"
            class="nearby-item"
            @click="selectNearby(place)"
          >
            <div class="nearby-icon">
              <el-icon><LocationFilled /></el-icon>
            </div>
            <div class="nearby-content">
              <div class="nearby-name">{{ place.name }}</div>
              <div class="nearby-address">{{ place.address }}</div>
              <div class="nearby-distance">{{ place.distance }}</div>
            </div>
          </div>
        </div>
      </div>

      <!-- 无搜索结果状态 -->
      <div v-else-if="searchKeyword && !isSearching && searchResults.length === 0" class="no-results-state">
        <el-icon class="empty-icon"><Search /></el-icon>
        <p>未找到相关地点</p>
        <p class="sub-text">请尝试其他关键词</p>
      </div>

      <!-- 空状态 -->
      <div v-else-if="!searchKeyword && !showHistory" class="empty-state">
        <el-icon class="empty-icon"><Search /></el-icon>
        <p>请输入搜索内容</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { useUserStore } from '@/store'
import { searchHistoryApi } from '@/api'
import { ElMessage } from 'element-plus'
import { 
  ArrowLeft, 
  Search, 
  LocationFilled, 
  StarFilled,
  Clock,
  Close,
  ArrowRight,
  HomeFilled,
  OfficeBuilding,
  Star,
  Loading
} from '@element-plus/icons-vue'

interface SearchResult {
  name: string
  address: string
  distance?: string
  lat?: number
  lng?: number
}

interface HistoryItem {
  id: number
  keyword: string
  startLocation?: string
  start_location?: string
  endLocation?: string
  end_location?: string
  searchCount?: number
  search_count?: number
  createdAt?: string
  created_at?: string
}

interface CommonAddress {
  label: string
  address?: string
  icon: any
  color: string
}

interface NearbyPlace {
  name: string
  address: string
  distance: string
}

interface QuickTab {
  id: string
  name: string
  icon: any
}

const props = defineProps<{
  searchHistory: HistoryItem[]
  commonAddresses: CommonAddress[]
  nearbyPlaces: NearbyPlace[]
  currentCity?: string
  currentLocation?: { lng: number, lat: number }
}>()

const emit = defineEmits<{
  'close': []
  'select-result': [result: SearchResult]
  'clear-history': []
  'remove-history': [index: number]
  'add-to-favorites': [result: SearchResult]
}>()

const userStore = useUserStore()
const searchKeyword = ref('')
const showHistory = ref(true)
const activeTab = ref('home')
const searchResults = ref<SearchResult[]>([])
const searchHistory = ref<HistoryItem[]>([])
const currentPage = ref(1)
const pageSize = ref(10)
const total = ref(0)
const isSearching = ref(false)
const searchTimer = ref<ReturnType<typeof setTimeout> | null>(null)

const quickTabs: QuickTab[] = [
  { id: 'home', name: '家', icon: HomeFilled },
  { id: 'company', name: '公司', icon: OfficeBuilding },
  { id: 'favorites', name: '收藏夹', icon: Star }
]

// 获取搜索历史
const fetchSearchHistory = async () => {
  if (!userStore.user?.id) {
    ElMessage.warning('请先登录')
    return
  }
  
  try {
    const { data } = await searchHistoryApi.getSearchHistory({
      user_id: Number(userStore.user.id),
      page: currentPage.value,
      page_size: pageSize.value
    })
    searchHistory.value = data.items
    total.value = data.total
  } catch (error) {
    console.error('获取搜索历史失败:', error)
    ElMessage.error('获取搜索历史失败')
  }
}

// 添加搜索历史
const addSearchHistory = async (result: SearchResult) => {
  if (!userStore.user?.id) {
    ElMessage.warning('请先登录')
    return
  }

  try {
    await searchHistoryApi.addSearchHistory({
      user_id: Number(userStore.user.id),
      keyword: result.name,
      start_location: '',  // 根据实际情况设置
      end_location: result.address
    })
    await fetchSearchHistory()
  } catch (error) {
    console.error('添加搜索历史失败:', error)
    ElMessage.error('添加搜索历史失败')
  }
}

// 删除单条搜索历史
const removeHistory = async (item: HistoryItem) => {
  if (!userStore.user?.id) {
    ElMessage.warning('请先登录')
    return
  }

  try {
    await searchHistoryApi.deleteSearchHistory(item.id, {
      user_id: Number(userStore.user.id)
    })
    await fetchSearchHistory()
    ElMessage.success('删除成功')
  } catch (error) {
    console.error('删除搜索历史失败:', error)
    ElMessage.error('删除搜索历史失败')
  }
}

// 清空搜索历史
const clearHistory = async () => {
  if (!userStore.user?.id) {
    ElMessage.warning('请先登录')
    return
  }

  try {
    await searchHistoryApi.clearSearchHistory({
      user_id: Number(userStore.user.id)
    })
    searchHistory.value = []
    total.value = 0
    ElMessage.success('清空成功')
  } catch (error) {
    console.error('清空搜索历史失败:', error)
    ElMessage.error('清空搜索历史失败')
  }
}

// 处理搜索
const handleSearch = async (keyword: string) => {
  // 清除之前的定时器
  if (searchTimer.value) {
    clearTimeout(searchTimer.value)
  }

  if (!keyword.trim()) {
    searchResults.value = []
    showHistory.value = true
    isSearching.value = false
    return
  }

  showHistory.value = false
  
  // 防抖处理：延迟300ms执行搜索
  searchTimer.value = setTimeout(async () => {
    try {
      isSearching.value = true
      
      // 导入地图服务
      const { baiduMapService } = await import('@/utils/baiduMapService')
      
      // 调用真实搜索API
      const results = await baiduMapService.searchPlaces(
        keyword, 
        props.currentCity || '北京市',
        props.currentLocation
      )
      
      searchResults.value = results.map(result => ({
        name: result.name,
        address: result.address,
        distance: result.distance,
        lng: result.lng,
        lat: result.lat
      }))
      
      if (results.length === 0) {
        ElMessage.info('未找到相关地点，请尝试其他关键词')
      }
    } catch (error) {
      console.error('搜索失败:', error)
      ElMessage.error('搜索失败，请稍后重试')
      searchResults.value = []
    } finally {
      isSearching.value = false
    }
  }, 300)
}

// 切换标签
const switchTab = (tabId: string) => {
  activeTab.value = tabId
  // 可以根据标签显示不同内容
}

// 选择搜索结果
const selectResult = async (result: SearchResult) => {
  await addSearchHistory(result)
  emit('select-result', result)
}

// 选择历史记录
const selectHistory = async (item: HistoryItem) => {
  // 更新搜索关键词
  searchKeyword.value = item.keyword
  
  // 尝试重新搜索获取详细信息
  try {
    const { baiduMapService } = await import('@/utils/baiduMapService')
    const results = await baiduMapService.searchPlaces(
      item.keyword,
      props.currentCity || '北京市',
      props.currentLocation
    )
    
    if (results.length > 0) {
      // 找到匹配的结果，选择第一个
      emit('select-result', {
        name: results[0].name,
        address: results[0].address,
        lng: results[0].lng,
        lat: results[0].lat,
        distance: results[0].distance
      })
    } else {
      // 没找到结果，使用历史记录中的信息
      emit('select-result', {
        name: item.keyword,
        address: item.endLocation || item.end_location || '地址未知'
      })
    }
  } catch (error) {
    console.error('搜索历史记录失败:', error)
    // 出错时使用历史记录中的信息
    emit('select-result', {
      name: item.keyword,
      address: item.endLocation || item.end_location || '地址未知'
    })
  }
}

// 选择常用地址
const selectCommon = (address: CommonAddress) => {
  if (address.address) {
    emit('select-result', {
      name: address.label,
      address: address.address
    })
  }
}

// 选择附近地点
const selectNearby = (place: NearbyPlace) => {
  emit('select-result', {
    name: place.name,
    address: place.address
  })
}

onMounted(() => {
  fetchSearchHistory()
})

onUnmounted(() => {
  // 清理定时器
  if (searchTimer.value) {
    clearTimeout(searchTimer.value)
  }
})

// 添加到收藏
const addToFavorites = (result: SearchResult) => {
  emit('add-to-favorites', result)
}
</script>

<style lang="scss" scoped>
.search-page {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: white;
  z-index: 2000;
  display: flex;
  flex-direction: column;
}

.search-header {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 16px;
  background: white;
  border-bottom: 1px solid #f0f0f0;
  
  .back-btn {
    flex-shrink: 0;
    width: 32px;
    height: 32px;
    
    .el-icon {
      font-size: 16px;
    }
  }
  
  .search-input-wrapper {
    flex: 1;
  }
}

.quick-tabs {
  display: flex;
  padding: 12px 16px;
  gap: 16px;
  border-bottom: 1px solid #f0f0f0;
  
  .tab-item {
    display: flex;
    align-items: center;
    gap: 6px;
    padding: 8px 12px;
    border-radius: 16px;
    cursor: pointer;
    transition: all 0.2s;
    
    &:hover {
      background: #f5f5f5;
    }
    
    &.active {
      background: #e6f7ff;
      color: #1890ff;
    }
    
    .el-icon {
      font-size: 14px;
    }
    
    span {
      font-size: 13px;
    }
  }
}

.search-content {
  flex: 1;
  overflow-y: auto;
  padding: 0 16px;
}

.search-results {
  .result-title {
    font-size: 14px;
    color: #666;
    margin: 16px 0 8px;
  }
  
  .result-item {
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 12px 0;
    border-bottom: 1px solid #f0f0f0;
    cursor: pointer;
    
    &:hover {
      background: #fafafa;
    }
    
    .result-icon {
      width: 32px;
      height: 32px;
      border-radius: 50%;
      background: #1890ff;
      display: flex;
      align-items: center;
      justify-content: center;
      
      .el-icon {
        color: white;
        font-size: 16px;
      }
    }
    
    .result-content {
      flex: 1;
      
      .result-name {
        font-size: 15px;
        color: #333;
        font-weight: 500;
        margin-bottom: 4px;
      }
      
      .result-address {
        font-size: 13px;
        color: #666;
        margin-bottom: 2px;
      }
      
      .result-distance {
        font-size: 12px;
        color: #999;
      }
    }
  }
}

.history-section {
  .history-group,
  .common-group,
  .nearby-group {
    margin: 16px 0;
    
    .history-header,
    .common-header,
    .nearby-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 8px;
      
      .history-title,
      .common-title,
      .nearby-title {
        font-size: 14px;
        color: #666;
        font-weight: 500;
      }
    }
  }
  
  .history-item,
  .nearby-item {
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 12px 0;
    border-bottom: 1px solid #f8f8f8;
    cursor: pointer;
    
    &:hover {
      background: #fafafa;
    }
    
    .history-icon,
    .nearby-icon {
      width: 28px;
      height: 28px;
      border-radius: 50%;
      background: #f0f0f0;
      display: flex;
      align-items: center;
      justify-content: center;
      
      .el-icon {
        color: #666;
        font-size: 14px;
      }
    }
    
    .history-content,
    .nearby-content {
      flex: 1;
      
      .history-name,
      .nearby-name {
        font-size: 14px;
        color: #333;
        margin-bottom: 2px;
      }
      
      .history-address,
      .nearby-address {
        font-size: 12px;
        color: #666;
      }
      
      .nearby-distance {
        font-size: 11px;
        color: #999;
        margin-top: 2px;
      }
    }
  }
  
  .common-item {
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 12px 0;
    border-bottom: 1px solid #f8f8f8;
    cursor: pointer;
    
    &:hover {
      background: #fafafa;
    }
    
    .common-icon {
      width: 32px;
      height: 32px;
      border-radius: 50%;
      display: flex;
      align-items: center;
      justify-content: center;
      
      .el-icon {
        color: white;
        font-size: 16px;
      }
    }
    
    .common-content {
      flex: 1;
      
      .common-label {
        font-size: 14px;
        color: #333;
        font-weight: 500;
        margin-bottom: 2px;
      }
      
      .common-address {
        font-size: 12px;
        color: #666;
      }
    }
    
    .arrow-icon {
      color: #ccc;
      font-size: 14px;
    }
  }
}

.history-content {
  .history-meta {
    display: flex;
    gap: 12px;
    font-size: 12px;
    color: #999;
    margin-top: 4px;
    
    .search-count {
      color: #1890ff;
    }
  }
}

.loading-state,
.empty-state,
.no-results-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 200px;
  
  .empty-icon {
    font-size: 48px;
    color: #ddd;
    margin-bottom: 16px;
  }
  
  .loading-icon {
    font-size: 32px;
    color: #1890ff;
    margin-bottom: 16px;
    animation: rotate 1s linear infinite;
  }
  
  p {
    font-size: 14px;
    color: #999;
    margin: 4px 0;
  }
  
  .sub-text {
    font-size: 12px;
    color: #ccc;
  }
}

@keyframes rotate {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}
</style> 