<template>
  <div class="search-page">
    <div class="search-header">
      <button class="back-btn" @click="router.back()">←</button>
      <input
        v-model="searchText"
        class="search-input"
        placeholder="请输入终点"
        @keyup.enter="onSearch"
      />
      <button class="search-btn" @click="onSearch">搜索</button>
    </div>
    <div class="quick-entry">
      <div class="entry" @click="setHome">
        <span class="icon home">🏠</span>
        <span>{{ homeAddr || '设置家地址' }}</span>
      </div>
      <div class="entry" @click="setCompany">
        <span class="icon company">🏢</span>
        <span>{{ companyAddr || '设置公司地址' }}</span>
      </div>
      <div class="entry" @click="chooseOnMap">
        <span class="icon map">📍</span>
        <span>地图选点</span>
      </div>
    </div>
    <div class="search-tabs">
      <button :class="{active: tab==='recommend'}" @click="tab='recommend'">为你推荐</button>
      <button :class="{active: tab==='history'}" @click="tab='history'">历史记录</button>
      <button class="clear-btn" v-if="tab==='history' && history.length" @click="clearHistory">清空</button>
    </div>
    <div v-if="tab==='recommend'" class="recommend-section">
      <div class="recommend-title"><span class="recommend-icon">🧡</span>为你推荐</div>
      <div v-if="recommend.length === 0" class="empty">暂无推荐</div>
      <div v-for="item in recommend" :key="item" class="recommend-item" @click="fillSearch(item)">
        <div class="rec-main">{{ item }}</div>
        <!-- 可扩展副标题、距离等 -->
      </div>
    </div>
    <div v-if="tab==='history'" class="history-section">
      <div class="history-title">历史记录 <span class="clear-btn" v-if="history.length" @click="clearHistory">清空</span></div>
      <div v-if="history.length === 0" class="empty">暂无历史记录</div>
      <div v-for="item in history" :key="item" class="history-item" @click="fillSearch(item)">
        <div class="his-main">{{ item }}</div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
const searchText = ref('');
const history = ref<string[]>([]);
const recommend = ref<string[]>([]);
const tab = ref<'recommend'|'history'>('recommend');
const router = useRouter();
const homeAddr = ref('');
const companyAddr = ref('');

function loadHistory() {
  const arr = JSON.parse(localStorage.getItem('search_history') || '[]');
  history.value = arr;
  recommend.value = arr.slice(-2).reverse(); // 最近2条
}
function saveHistory(val: string) {
  let arr = JSON.parse(localStorage.getItem('search_history') || '[]');
  arr = arr.filter((v: string) => v !== val);
  arr.push(val);
  localStorage.setItem('search_history', JSON.stringify(arr));
  loadHistory();
}
function onSearch() {
  const val = searchText.value.trim();
  if (!val) return;
  saveHistory(val);
  searchText.value = '';
  tab.value = 'recommend';
  router.push({ path: '/taxi', query: { dest: val } });
}
function clearHistory() {
  localStorage.removeItem('search_history');
  loadHistory();
}
function fillSearch(val: string) {
  searchText.value = val;
}
function setHome() {
  const addr = prompt('请输入家地址', homeAddr.value || '');
  if (addr && addr.trim()) {
    homeAddr.value = addr.trim();
    localStorage.setItem('home_addr', homeAddr.value);
  }
}
function setCompany() {
  const addr = prompt('请输入公司地址', companyAddr.value || '');
  if (addr && addr.trim()) {
    companyAddr.value = addr.trim();
    localStorage.setItem('company_addr', companyAddr.value);
  }
}
function chooseOnMap() {
  alert('地图选点功能可后续接入地图组件实现');
}
onMounted(() => {
  loadHistory();
  homeAddr.value = localStorage.getItem('home_addr') || '';
  companyAddr.value = localStorage.getItem('company_addr') || '';
});
</script>

<style scoped>
.search-page {
  max-width: 480px;
  margin: 0 auto;
  padding: 16px 8px;
  background: #f7f8fa;
  min-height: 100vh;
}
.search-header {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 16px;
}
.back-btn {
  background: #fff;
  border: none;
  border-radius: 8px;
  font-size: 1.3rem;
  padding: 0 10px;
  cursor: pointer;
  color: #222;
  box-shadow: 0 1px 4px rgba(0,0,0,0.04);
}
.search-input {
  flex: 1;
  padding: 10px 14px;
  border: none;
  border-radius: 12px;
  background: #fff;
  font-size: 1rem;
  outline: none;
  box-shadow: 0 1px 4px rgba(0,0,0,0.04);
}
.search-btn {
  padding: 8px 16px;
  border: none;
  border-radius: 8px;
  background: #ff7a00;
  color: #fff;
  font-size: 1rem;
  cursor: pointer;
}
.quick-entry {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 18px;
  gap: 8px;
}
.entry {
  flex: 1;
  background: #fff;
  border-radius: 12px;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 10px 0 6px 0;
  font-size: 0.98rem;
  cursor: pointer;
  box-shadow: 0 1px 4px rgba(0,0,0,0.04);
  transition: box-shadow 0.2s;
}
.entry:active {
  box-shadow: 0 2px 8px rgba(0,0,0,0.10);
}
.icon {
  font-size: 1.3rem;
  margin-bottom: 2px;
}
.search-tabs {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 12px;
}
.search-tabs button {
  background: #fff;
  border: none;
  border-radius: 8px;
  padding: 6px 16px;
  font-size: 1rem;
  cursor: pointer;
}
.search-tabs .active {
  background: #ff7a00;
  color: #fff;
}
.clear-btn {
  margin-left: auto;
  background: #f7f8fa !important;
  color: #ff7a00 !important;
}
.card-list {
  background: #fff;
  border-radius: 12px;
  padding: 12px 8px;
  min-height: 80px;
  margin-bottom: 18px;
  box-shadow: 0 1px 4px rgba(0,0,0,0.04);
}
.card-title {
  font-weight: bold;
  font-size: 1.08rem;
  margin-bottom: 8px;
  color: #222;
}
.item {
  padding: 8px 0;
  border-bottom: 1px solid #f0f0f0;
  font-size: 1rem;
  cursor: pointer;
}
.item:last-child {
  border-bottom: none;
}
.empty {
  color: #aaa;
  text-align: center;
  padding: 24px 0;
}
.recommend-section {
  background: #fff7f0;
  border-radius: 12px;
  margin-bottom: 12px;
  padding: 0 0 8px 0;
}
.recommend-title {
  font-weight: bold;
  font-size: 1.05rem;
  color: #ff7a00;
  padding: 12px 12px 4px 12px;
  display: flex;
  align-items: center;
}
.recommend-icon {
  font-size: 1.1rem;
  margin-right: 6px;
}
.recommend-item {
  padding: 12px 12px 8px 36px;
  font-size: 1rem;
  border-bottom: 1px solid #ffe3c2;
  cursor: pointer;
  background: transparent;
}
.recommend-item:last-child {
  border-bottom: none;
}
.rec-main {
  color: #222;
}
.history-section {
  background: #fff;
  border-radius: 12px;
  margin-bottom: 12px;
  padding: 0 0 8px 0;
}
.history-title {
  font-weight: bold;
  font-size: 1.05rem;
  color: #222;
  padding: 12px 12px 4px 12px;
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.history-item {
  padding: 12px 12px 8px 36px;
  font-size: 1rem;
  border-bottom: 1px solid #f0f0f0;
  cursor: pointer;
  background: transparent;
}
.history-item:last-child {
  border-bottom: none;
}
.his-main {
  color: #222;
}
</style> 