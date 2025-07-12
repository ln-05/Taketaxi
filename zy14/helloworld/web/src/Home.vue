
<template>
  <div class="home-container">
    <div id="baidu-map" class="map"></div>
    <div class="map-overlay">
      <div class="location">{{ city }}</div>
      <div class="scan-btns">
        <div class="scan-item">
          <img :src="bike" alt="扫一扫" class="scan-icon" />
          <div class="scan-label">扫一扫</div>
        </div>
      </div>
    </div>
    <!-- 主卡片：上车点+目的地输入卡片 -->
    <div class="main-card">
      <div class="pickup-row">
        <span class="dot-green"></span>
        <span class="pickup-label">从 <span class="pickup-location">{{ pickupLocation }}</span> 上车</span>
      </div>
      <div class="dest-row">
        <span class="dot-orange"></span>
        <span class="dest-label">您想去哪儿？</span>
      </div>
      <input class="dest-search-input" type="text" placeholder="请输入目的地" readonly @click="router.push('/search')" style="cursor:pointer;" />
      <div class="card-actions">
        <button class="card-action-btn">预约</button>
        <button class="card-action-btn">帮叫车</button>
        <button class="card-action-btn">接送机</button>
      </div>
    </div>
    <div class="feature-icons">
      <div class="feature-item" v-for="item in features" :key="item.label">
        <img :src="item.icon" :alt="item.label" class="feature-icon" />
        <div class="feature-label">{{ item.label }}</div>
      </div>
    </div>
    <div class="bottom-nav">
      <div class="nav-item" v-for="item in navs" :key="item.label">
        <img :src="item.icon" :alt="item.label" class="nav-icon" />
        <div class="nav-label">{{ item.label }}</div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import taxi from './assets/taxi.png';
import shunfeng from './assets/shunfeng.png';
import daijia from './assets/daijia.png';
import bike from './assets/bike.png';
import subway from './assets/subway.png';
import delivery from './assets/delivery.png';
import home from './assets/home.png';
import order from './assets/order.png';
import profile from './assets/profile.png';
import { useRouter } from 'vue-router';

const city = ref('定位中...');
const pickupLocation = ref('定位中...');
const features = [
  { label: '打车', icon: taxi },
  { label: '顺风车', icon: shunfeng },
  { label: '代驾', icon: daijia },
  { label: '青桔骑行', icon: bike },
  { label: '地铁', icon: subway },
  { label: '送货', icon: delivery },
];
const navs = [
  { label: '首页', icon: home },
  { label: '订单', icon: order },
  { label: '我的', icon: profile },
];

const router = useRouter();

onMounted(() => {
  if (!window.BMap) {
    const script = document.createElement('script');
    script.src = 'https://api.map.baidu.com/api?v=3.0&ak=N3RoKOXTLUJgX6bchYUFqntOtcwTwD7u&callback=onBMapCallback';
    document.body.appendChild(script);
    window.onBMapCallback = () => {
      locateAndShow();
    };
  } else {
    locateAndShow();
  }

  function locateAndShow() {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        (pos) => {
          const lng = pos.coords.longitude;
          const lat = pos.coords.latitude;
          // 坐标转换 WGS84 -> BD09
          // @ts-ignore
          const convertor = new window.BMap.Convertor();
          const pointArr = [new window.BMap.Point(lng, lat)];
          convertor.translate(pointArr, 1, 5, (data) => {
            if (data.status === 0) {
              const bdPoint = data.points[0];
              // @ts-ignore
              const map = new window.BMap.Map('baidu-map');
              map.centerAndZoom(bdPoint, 16);
              map.enableScrollWheelZoom(true);
              // @ts-ignore
              const geoc = new window.BMap.Geocoder();
              geoc.getLocation(bdPoint, (rs) => {
                if (rs && rs.addressComponents) {
                  const comp = rs.addressComponents;
                  city.value = comp.city || '未知城市';
                  pickupLocation.value = comp.street + comp.streetNumber;
                } else {
                  city.value = '未知城市';
                  pickupLocation.value = '未知位置';
                }
              });
            }
          });
        },
        (err) => {
          city.value = '定位失败';
          pickupLocation.value = '定位失败';
        },
        { enableHighAccuracy: true, timeout: 5000 }
      );
    } else {
      city.value = '不支持定位';
      pickupLocation.value = '不支持定位';
    }
  }
});
</script>

<style scoped>
.home-container {
  position: relative;
  width: 100vw;
  min-height: 80vh;
  background: #f7f8fa;
}
.map {
  width: 100vw;
  height: 260px;
}
.map-overlay {
  position: absolute;
  left: 0; top: 0; width: 100vw;
  pointer-events: none;
}
.location {
  position: absolute;
  left: 24px;
  top: 16px;
  right: auto;
  font-size: 1.2rem;
  color: #222;
  font-weight: 600;
  background: rgba(255,255,255,0.8);
  border-radius: 16px;
  padding: 4px 16px;
  pointer-events: auto;
}
.scan-btns {
  position: absolute;
  right: 24px; top: 64px;
  display: flex;
  gap: 16px;
}
.scan-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  background: #fff;
  border-radius: 24px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.08);
  padding: 8px 12px;
  pointer-events: auto;
}
.scan-icon {
  width: 32px;
  height: 32px;
}
.scan-label {
  font-size: 0.95rem;
  color: #222;
  margin-top: 2px;
}
.feature-icons {
  margin: 24px auto 0 auto;
  width: 92vw;
  max-width: 480px;
  display: flex;
  flex-wrap: wrap;
  gap: 0;
  justify-content: flex-start;
}
.feature-item {
  width: 33.33%;
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 12px;
}
.feature-icon {
  width: 36px;
  height: 36px;
  margin-bottom: 4px;
}
.feature-label {
  font-size: 0.95rem;
  color: #222;
}
.bottom-nav {
  position: fixed;
  left: 0; right: 0; bottom: 0;
  height: 60px;
  background: #fff;
  box-shadow: 0 -2px 8px rgba(0,0,0,0.04);
  display: flex;
  justify-content: space-around;
  align-items: center;
  z-index: 10;
  margin-bottom: 24px;
}
.nav-item {
  display: flex;
  flex-direction: column;
  align-items: center;
}
.nav-icon {
  width: 28px;
  height: 28px;
}
.nav-label {
  font-size: 0.95rem;
  color: #222;
}
.main-card {
  position: relative;
  margin: -24px auto 0 auto;
  width: 92vw;
  max-width: 480px;
  background: #fff;
  border-radius: 18px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.10);
  padding: 18px 16px 12px 16px;
  z-index: 3;
}
.pickup-row, .dest-row {
  display: flex;
  align-items: center;
  margin-bottom: 8px;
}
.dot-green {
  width: 8px; height: 8px; border-radius: 50%;
  background: #1aad19; margin-right: 8px;
}
.dot-orange {
  width: 8px; height: 8px; border-radius: 50%;
  background: #ff7a00; margin-right: 8px;
}
.pickup-label {
  color: #222;
  font-size: 1rem;
}
.pickup-location {
  color: #1aad19;
  font-weight: 600;
}
.dest-label {
  color: #222;
  font-size: 1.1rem;
  font-weight: 600;
}
.card-actions {
  display: flex;
  gap: 16px;
  margin-top: 8px;
}
.card-action-btn {
  background: #f7f8fa;
  border: none;
  border-radius: 12px;
  padding: 6px 18px;
  font-size: 1rem;
  color: #222;
  box-shadow: 0 1px 4px rgba(0,0,0,0.04);
  cursor: pointer;
}
.dest-search-input {
  width: 100%;
  margin: 0 0 10px 0;
  padding: 10px 14px;
  border: none;
  border-radius: 12px;
  background: #f7f8fa;
  box-shadow: 0 1px 4px rgba(0,0,0,0.04);
  font-size: 1rem;
  outline: none;
  color: #222;
  transition: box-shadow 0.2s;
}
.dest-search-input:focus {
  box-shadow: 0 2px 8px rgba(0,0,0,0.10);
  background: #fff;
}
</style> 