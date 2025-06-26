<template>
  <div class="map-container">
    <div v-if="isLoading" class="loading-mask">
      <el-loading-spinner></el-loading-spinner>
      <p>地图加载中...</p>
    </div>
    <div id="container"></div>
  </div>
</template>

<script setup>
import { onMounted, onUnmounted, ref, defineExpose } from "vue";
import AMapLoader from "@amap/amap-jsapi-loader";

const map = ref(null);
const geocoder = ref(null);
const currentMarker = ref(null);
const AMapInstance = ref(null);
const isLoading = ref(true);
const loadError = ref(null);
const MAX_RETRIES = 3; // 最大重试次数

onMounted(() => {
  window._AMapSecurityConfig = {
    securityJsCode: "1205b29ddf312e5ed5270bb9ed3c31fe",
  };
  AMapLoader.load({
    key: "37662ba7b24b5005fdffac9f8fb1501d",
    version: "2.0",
    plugins: ["AMap.Geocoder", "AMap.InfoWindow"] // 添加InfoWindow插件
  })
      .then((AMap) => {
        AMapInstance.value = AMap;
        map.value = new AMap.Map("container", {
          viewMode: "3D",
          zoom: 5,
          center: [110, 35]
        });

        geocoder.value = new AMap.Geocoder({
          city: "全国"
        });

        // 监听地图加载完成事件
        map.value.on("complete", () => {
          isLoading.value = false;
          console.log("地图加载完成");
        });
      })
      .catch((error) => {
        loadError.value = error;
        isLoading.value = false;
        console.error("地图加载失败:", error);
      });
});

onUnmounted(() => {
  // 清理资源
  if (currentMarker.value) {
    map.value.remove(currentMarker.value);
    currentMarker.value = null;
  }

  if (map.value) {
    map.value.destroy();
    map.value = null;
  }

  geocoder.value = null;
  AMapInstance.value = null;
});

const locateAddress = (address, retries = 0) => {
  console.log('定位请求:', address);

  return new Promise((resolve, reject) => {
    // 参数验证
    if (!address || typeof address !== "string") {
      reject({ code: "INVALID_PARAM", message: "无效的地址参数" });
      return;
    }

    // 确保地图和地理编码器都已初始化
    if (!map.value || !geocoder.value || !AMapInstance.value) {
      const missingDeps = [];
      if (!map.value) missingDeps.push('map');
      if (!geocoder.value) missingDeps.push('geocoder');
      if (!AMapInstance.value) missingDeps.push('AMapInstance');

      reject({
        code: "DEPENDENCY_MISSING",
        message: `缺少依赖: ${missingDeps.join(', ')}`
      });
      return;
    }

    // 清除旧标记
    if (currentMarker.value) {
      map.value.remove(currentMarker.value);
      currentMarker.value = null;
    }

    console.log('开始地理编码请求:', address);

    // 添加超时处理，延长到20秒
    const timeoutId = setTimeout(() => {
      if (retries < MAX_RETRIES) {
        console.log(`请求超时，正在进行第 ${retries + 1} 次重试...`);
        clearTimeout(timeoutId);
        locateAddress(address, retries + 1).then(resolve).catch(reject);
      } else {
        reject({
          code: "REQUEST_TIMEOUT",
          message: "地理编码请求超时，请检查网络连接"
        });
      }
    }, 2000);

    try {
      // 地理编码
      geocoder.value.getLocation(address, (status, result) => {
        // 清除超时计时器
        clearTimeout(timeoutId);

        console.log('地理编码回调执行:', status, result);

        if (status === "complete" && result.info === "OK") {
          if (!result.geocodes || result.geocodes.length === 0) {
            reject({
              code: "NO_RESULT",
              message: "未找到匹配的地理编码结果"
            });
            return;
          }

          const location = result.geocodes[0].location;

          // 创建标记
          currentMarker.value = new AMapInstance.value.Marker({
            position: [location.lng, location.lat],
            title: address,
            draggable: false,
            animation: "AMAP_ANIMATION_DROP",
            offset: new AMapInstance.value.Pixel(-10, -30)
          });

          // 添加信息窗口
          const infoWindow = new AMapInstance.value.InfoWindow({
            content: `<div class="info-window">${address}</div>`,
            offset: new AMapInstance.value.Pixel(0, -30),
            autoMove: true
          });

          // 点击标记显示信息窗口
          currentMarker.value.on("click", () => {
            infoWindow.open(map.value, currentMarker.value.getPosition());
          });

          // 添加到地图
          currentMarker.value.setMap(map.value);

          // 平滑调整地图视图
          map.value.setZoomAndCenter(15, [location.lng, location.lat], {
            animate: true
          });

          resolve({ location, address });
        } else {
          // 返回详细错误信息
          reject({
            code: "GEOCODE_FAILED",
            message: `地址解析失败: ${address}`,
            originalResult: result
          });
        }
      });
    } catch (error) {
      // 清除超时计时器
      clearTimeout(timeoutId);

      reject({
        code: "INTERNAL_ERROR",
        message: `地理编码过程中发生错误: ${error.message}`,
        error: error
      });
    }
  });
};

defineExpose({locateAddress});

</script>

<style scoped>
#container {
  width: 100%;
  height: 100%;
}

.loading-mask {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  background-color: rgba(255, 255, 255, 0.8);
  z-index: 1000;
}

.info-window {
  padding: 8px 12px;
  font-size: 14px;
  color: #333;
}
</style>