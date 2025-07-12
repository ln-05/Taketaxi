<template>
  <div class="taxi-page">
    <div id="taxi-map" class="taxi-map"></div>
    <div class="taxi-info">
      <div><span class="label">起点：</span>{{ startAddr }}</div>
      <div><span class="label">终点：</span>{{ dest }}</div>
      <div v-if="distance"><span class="label">距离：</span>{{ distance }}</div>
      <div v-if="duration"><span class="label">预计用时：</span>{{ duration }}</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
const route = useRoute();
const dest = ref(route.query.dest as string || '');
const startAddr = ref('定位中...');
const distance = ref('');
const duration = ref('');

onMounted(() => {
  if (!window.BMap) {
    const script = document.createElement('script');
    script.src = 'https://api.map.baidu.com/api?v=3.0&ak=N3RoKOXTLUJgX6bchYUFqntOtcwTwD7u&callback=onBMapCallback';
    document.body.appendChild(script);
    window.onBMapCallback = () => {
      showRoute();
    };
  } else {
    showRoute();
  }

  function showRoute() {
    if (!dest.value) return;
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        (pos) => {
          const lng = pos.coords.longitude;
          const lat = pos.coords.latitude;
          // WGS84转BD09
          // @ts-ignore
          const convertor = new window.BMap.Convertor();
          const pointArr = [new window.BMap.Point(lng, lat)];
          convertor.translate(pointArr, 1, 5, (data) => {
            if (data.status === 0) {
              const bdPoint = data.points[0];
              // @ts-ignore
              const map = new window.BMap.Map('taxi-map');
              map.centerAndZoom(bdPoint, 15);
              map.enableScrollWheelZoom(true);
              // 起点逆地理编码
              // @ts-ignore
              const geoc = new window.BMap.Geocoder();
              geoc.getLocation(bdPoint, (rs) => {
                if (rs && rs.addressComponents) {
                  const comp = rs.addressComponents;
                  startAddr.value = comp.street + comp.streetNumber;
                }
              });
              // 目的地转坐标
              // @ts-ignore
              const local = new window.BMap.LocalSearch(map, {
                onSearchComplete: function(results) {
                  if (local.getStatus() === window.BMAP_STATUS_SUCCESS) {
                    const destPoi = results.getPoi(0);
                    if (destPoi) {
                      // 路线规划
                      // @ts-ignore
                      const driving = new window.BMap.DrivingRoute(map, {
                        renderOptions: { map, autoViewport: true },
                        onSearchComplete: function(results) {
                          if (driving.getStatus() === window.BMAP_STATUS_SUCCESS) {
                            const plan = results.getPlan(0);
                            if (plan) {
                              distance.value = plan.getDistance(true);
                              duration.value = plan.getDuration(true);
                            }
                          }
                        }
                      });
                      driving.search(bdPoint, destPoi.point);
                    }
                  }
                }
              });
              local.search(dest.value);
            }
          });
        },
        () => {
          startAddr.value = '定位失败';
        },
        { enableHighAccuracy: true, timeout: 5000 }
      );
    } else {
      startAddr.value = '不支持定位';
    }
  }
});
</script>

<style scoped>
.taxi-page {
  max-width: 480px;
  margin: 0 auto;
  background: #f7f8fa;
  min-height: 100vh;
}
.taxi-map {
  width: 100%;
  height: 60vh;
  min-height: 340px;
  border-radius: 16px;
  margin-bottom: 18px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.08);
}
.taxi-info {
  background: #fff;
  border-radius: 12px;
  box-shadow: 0 1px 4px rgba(0,0,0,0.04);
  padding: 18px 18px 12px 18px;
  font-size: 1.08rem;
  color: #222;
  margin: 0 8px;
}
.label {
  color: #888;
  font-size: 0.98rem;
  margin-right: 4px;
}
</style> 