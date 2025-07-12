<template>
  <div class="search-page">
    <div class="search-header">
      <el-button circle size="small" class="back-btn" @click="$emit('close')">
        <el-icon><ArrowLeft /></el-icon>
      </el-button>
      <div class="search-input-wrapper">
        <el-input
          v-model="searchKeyword"
          :placeholder="placeholder"
          size="large"
          @input="handleSearch"
          autofocus
          clearable
        >
          <template #prefix>
            <el-icon><Search /></el-icon>
          </template>
        </el-input>
      </div>
      <el-button v-if="mode === 'end'" text class="waypoint-btn" @click="$emit('add-waypoint')">
        + 途经点
      </el-button>
    </div>

    <div class="search-content">
      <div v-if="!searchKeyword">
        <div class="quick-actions">
          <div class="action-item" @click="selectQuickAction('home')">
            <el-icon class="icon-home"><HomeFilled /></el-icon>
            <div class="action-text">
              <span class="label">家</span>
              <span class="desc">设置家的地址</span>
            </div>
          </div>
          <div class="action-item" @click="selectQuickAction('company')">
            <el-icon class="icon-company"><OfficeBuilding /></el-icon>
            <div class="action-text">
              <span class="label">公司</span>
              <span class="desc">设置公司地址</span>
            </div>
          </div>
           <div class="action-item" @click="$emit('pick-on-map')">
            <el-icon class="icon-fav"><MapLocation /></el-icon>
            <div class="action-text">
              <span class="label">地图选点</span>
               <span class="desc">拖动图钉，精准定位</span>
            </div>
        </div>
      </div>

        <div v-if="mode === 'end' && recommendations.length" class="recommendations-section">
          <div class="section-header">
            <el-icon class="section-icon"><MagicStick /></el-icon>
            <span class="section-title">为您推荐</span>
          </div>
          <div class="item-list">
            <div v-for="(item, index) in recommendations" :key="index" class="list-item" @click="selectResult(item)">
              <div class="item-content">
                <span class="item-name">{{ item.name }}</span>
                <span class="item-address">{{ item.address }}</span>
            </div>
              <el-button text size="small" @click.stop="shareLocation(item)">分享</el-button>
            </div>
          </div>
        </div>

        <div v-if="searchHistory.length" class="history-section">
          <div class="section-header">
            <span class="section-title">历史记录</span>
            <el-button text size="small" @click="$emit('clear-history')">清空</el-button>
          </div>
          <div class="item-list">
            <div v-for="(item, index) in searchHistory" :key="index" class="list-item" @click="selectResult(item)">
              <div class="item-content">
                <span class="item-name">{{ item.name }}</span>
                <span class="item-address">{{ item.address }}</span>
            </div>
              <span class="item-distance" v-if="item.distance">{{ item.distance }}</span>
            </div>
          </div>
        </div>

        <div v-if="mode === 'start' && nearbyPlaces.length" class="nearby-section">
          <div class="section-header">
            <span class="section-title">附近地点</span>
          </div>
          <div class="item-list">
            <div v-for="(item, index) in nearbyPlaces" :key="index" class="list-item" @click="selectResult(item)">
              <div class="item-content">
                <span class="item-name">{{ item.name }}</span>
                <span class="item-address">{{ item.address }}</span>
            </div>
              <span class="item-distance">{{ item.distance }}</span>
            </div>
          </div>
        </div>
      </div>

      <div v-else class="search-results">
        <div v-for="(item, index) in searchResults" :key="index" class="list-item" @click="selectResult(item)">
          <div class="item-content">
            <span class="item-name">{{ item.name }}</span>
            <span class="item-address">{{ item.address }}</span>
          </div>
          <span class="item-distance" v-if="item.distance">{{ item.distance }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { baiduMapService, type SearchResult } from '@/utils/baiduMapService';
import { 
  ArrowLeft, Search, HomeFilled, OfficeBuilding, MagicStick, MapLocation
} from '@element-plus/icons-vue';

const props = defineProps<{
  mode: 'start' | 'end';
  placeholder: string;
  searchHistory: SearchResult[];
  recommendations: SearchResult[];
  nearbyPlaces: SearchResult[];
  currentCity?: string;
  currentLocation?: { lat: number; lng: number };
}>();

const emit = defineEmits<{
  (e: 'close'): void;
  (e: 'select-result', result: SearchResult): void;
  (e: 'clear-history'): void;
  (e: 'add-waypoint'): void;
  (e: 'pick-on-map'): void;
}>();

const searchKeyword = ref('');
const searchResults = ref<SearchResult[]>([]);

const handleSearch = async (keyword: string) => {
  if (!keyword.trim()) {
    searchResults.value = [];
    return;
  }
  searchResults.value = await baiduMapService.searchPlaces(
      keyword, 
      props.currentCity,
    props.currentLocation
  );
};

const selectResult = (result: SearchResult) => {
  emit('select-result', result);
};

const selectQuickAction = (action: 'home' | 'company' | 'fav') => {
  // This is a mock implementation. In a real app, you'd fetch these addresses.
  const MOCK_ADDRESS = {
    home: { name: '家', address: '模拟的家的地址', lat: 39.9, lng: 116.4 },
    company: { name: '公司', address: '模拟的公司地址', lat: 39.95, lng: 116.45 },
    fav: { name: '收藏点', address: '模拟的收藏点地址', lat: 39.85, lng: 116.35 },
  };
  emit('select-result', MOCK_ADDRESS[action]);
};

const shareLocation = (item: SearchResult) => {
  // Placeholder for sharing functionality
  console.log('Sharing:', item);
};
</script>

<style lang="scss" scoped>
.search-page {
  position: fixed;
  top: 0; left: 0; right: 0; bottom: 0;
  background: #f4f4f4;
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
    flex-shrink: 0;

  .waypoint-btn {
    color: #333;
    font-weight: 500;
  }
}

.search-content {
  flex-grow: 1;
  overflow-y: auto;
  background: white;
}

.quick-actions {
  display: flex;
  justify-content: space-around;
  padding: 16px;
  background: white;
  border-bottom: 10px solid #f4f4f4;

  .action-item {
    display: flex;
    align-items: center;
    gap: 8px;
    cursor: pointer;
    
    .el-icon { font-size: 20px; }
    .icon-home { color: #67c23a; }
    .icon-company { color: #409eff; }
    .icon-fav { color: #f56c6c; }

    .action-text {
      display: flex;
      flex-direction: column;
      .label { font-size: 15px; font-weight: 500; color: #333; }
      .desc { font-size: 12px; color: #999; }
    }
  }
}

.section-header {
    display: flex;
    align-items: center;
  justify-content: space-between;
  padding: 16px 16px 8px;
  .section-title { font-size: 14px; color: #666; }
  .section-icon { margin-right: 4px; color: #e6a23c; }
      }
      
.item-list {
  padding: 0 16px;
  .list-item {
    display: flex;
    align-items: center;
    padding: 12px 0;
    border-bottom: 1px solid #f4f4f4;
    cursor: pointer;
    
    .item-content {
      flex-grow: 1;
      .item-name { display: block; font-size: 15px; color: #333; }
      .item-address { display: block; font-size: 12px; color: #999; }
    }
    .item-distance { font-size: 12px; color: #999; margin-left: 12px; }
  }
}
.search-results {
  padding: 0 16px;
}
</style> 