<template>
  <div class="trip-planner-card">
    <!-- 起点 -->
    <div class="location-row start" @click="$emit('edit-start')">
      <div class="location-icon">
        <el-icon><LocationFilled /></el-icon>
      </div>
      <div class="location-info">
        <span v-if="startPoint" class="location-text">{{ startPoint.name }}</span>
        <span v-else class="placeholder">您在哪儿上车</span>
      </div>
    </div>

    <!-- 终点 -->
    <div class="location-row end" @click="$emit('edit-end')">
      <div class="location-icon">
        <el-icon><Position /></el-icon>
      </div>
      <div class="location-info">
        <span v-if="endPoint" class="location-text">{{ endPoint.name }}</span>
        <span v-else class="placeholder">您想去哪儿</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { LocationFilled, Position } from '@element-plus/icons-vue'
import type { LocationPoint } from '@/types'

defineProps<{
  startPoint: LocationPoint | null
  endPoint: LocationPoint | null
}>()

defineEmits<{
  'edit-start': []
  'edit-end': []
}>()
</script>

<style lang="scss" scoped>
.trip-planner-card {
  background: white;
  margin: 16px;
  padding: 8px 16px;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0,0,0,0.1);
}
.location-row {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 0;
  cursor: pointer;
  &.start {
    border-bottom: 1px solid #f0f0f0;
  }
  .location-icon {
    width: 24px;
    height: 24px;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    .el-icon { font-size: 14px; color: white; }
  }
  &.start .location-icon { background: #52c41a; }
  &.end .location-icon { background: #ff4d4f; }
  .location-info {
    flex: 1;
    .location-text { font-size: 15px; color: #333; font-weight: 500; }
    .placeholder { font-size: 15px; color: #999; }
  }
}
</style> 