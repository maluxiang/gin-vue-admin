<template>
  <div class="map-container">
    <!-- 地图容器 -->
    <div ref="mapChart" class="map-chart"></div>

    <!-- 右侧设备管理面板 -->
    <div class="right-panel">
      <!-- 面板头部 -->
      <div class="panel-header">
        <div class="info-panel-toggle">
          <span>信息面板：</span>
          <el-switch v-model="infoPanelOpen" active-text="开" inactive-text="关" />
        </div>
        <el-button type="primary" @click="handleRegionManage">区域管理</el-button>
      </div>

      <!-- 筛选区域 -->
      <div class="filter-section" v-if="infoPanelOpen">
        <el-form :inline="true" :model="filterForm" class="filter-form">
          <el-form-item label="查看区域：">
            <el-select v-model="filterForm.region" placeholder="全部" @change="handleFilter">
              <el-option label="全部" value="all"></el-option>
              <el-option label="区域1" value="region1"></el-option>
              <el-option label="区域2" value="region2"></el-option>
              <!-- 可根据实际需求扩展区域选项 -->
            </el-select>
          </el-form-item>
          <el-form-item label="设备信息：">
            <el-select v-model="filterForm.deviceInfoType" placeholder="设备ID" @change="handleFilter">
              <el-option label="设备ID" value="deviceId"></el-option>
              <el-option label="设备名称" value="deviceName"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-input v-model="filterForm.keyword" placeholder="请输入" @keyup.enter="handleFilter"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleFilter">查询</el-button>
            <el-button @click="handleReset">重置</el-button>
          </el-form-item>
        </el-form>
      </div>

      <!-- 标签页区域 -->
      <el-tabs v-model="activeTab" class="tab-section" v-if="infoPanelOpen">
        <el-tab-pane label="设备清单" name="deviceList">
          <el-table :data="deviceTableData" border style="width: 100%">
            <el-table-column prop="deviceId" label="设备ID"></el-table-column>
            <el-table-column prop="deviceName" label="设备名称"></el-table-column>
            <el-table-column label="操作">
              <template #default="scope">
                <el-button type="text" @click="handleDetail(scope.row)">详情</el-button>
                <el-button type="text" @click="handleDelete(scope.row)">删除</el-button>
                <el-button type="text" @click="handleLocate(scope.row)">定位</el-button>
              </template>
            </el-table-column>
          </el-table>
          <el-pagination
              @size-change="handleSizeChange"
              @current-change="handleCurrentChange"
              :current-page="currentPage"
              :page-sizes="[5, 8, 10, 15]"
              :page-size="pageSize"
              layout="total, sizes, prev, pager, next, jumper"
              :total="total"
              style="margin-top: 10px; text-align: right;"
          />
        </el-tab-pane>
        <el-tab-pane label="设备告警" name="deviceAlarm">
          <el-table :data="alarmTableData" border style="width: 100%">
            <el-table-column prop="alarmId" label="告警ID"></el-table-column>
            <el-table-column prop="deviceId" label="设备ID"></el-table-column>
            <el-table-column prop="alarmType" label="告警类型"></el-table-column>
            <el-table-column prop="alarmTime" label="告警时间"></el-table-column>
            <el-table-column label="操作">
              <template #default="scope">
                <el-button type="text" @click="handleAlarmDetail(scope.row)">处理</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, watch } from 'vue';
import * as echarts from 'echarts';
import axios from 'axios';
import { useRouter } from 'vue-router';
import {
  ElSwitch,
  ElButton,
  ElSelect,
  ElOption,
  ElForm,
  ElFormItem,
  ElInput,
  ElTabs,
  ElTabPane,
  ElTable,
  ElTableColumn,
  ElPagination
} from 'element-plus';

const router = useRouter();

// 地图实例相关
const mapChart = ref(null);
let myChart = null;

// 设备数据相关
const deviceList = ref([
  {
    deviceId: 'Smart 3200',
    deviceName: '智能烟感',
    status: '在线',
    lng: 116.407526,
    lat: 39.90403,
    area: '北京市' // 新增字段：设备所属地区
  },
  {
    deviceId: 'Smart 3200',
    deviceName: '智能烟感',
    status: '离线',
    lng: 121.473701,
    lat: 31.230416,
    area: '上海市'
  },
  {
    deviceId: 'Smart 3600',
    deviceName: '火灾主机',
    status: '故障',
    lng: 113.264435,
    lat: 23.12911,
    area: '广东省'
  },
  {
    deviceId: 'Smart 6000',
    deviceName: '智能门禁',
    status: '在线',
    lng: 114.057868,
    lat: 22.543099,
    area: '广东省'
  },
  {
    deviceId: 'Smart 6100',
    deviceName: '智能门禁',
    status: '在线',
    lng: 104.065735,
    lat: 30.572269,
    area: '四川省'
  },
  {
    deviceId: 'Smart 8000',
    deviceName: '智能空开',
    status: '离线',
    lng: 120.15358,
    lat: 30.287458,
    area: '浙江省'
  },
  {
    deviceId: 'Smart 2600',
    deviceName: '燃气探测器',
    status: '故障',
    lng: 118.767413,
    lat: 32.041544,
    area: '江苏省'
  },
  {
    deviceId: 'Smart 2400',
    deviceName: '视频主机',
    status: '在线',
    lng: 114.298572,
    lat: 30.584355,
    area: '湖北省'
  },
]);

const alarmTableData = ref([
  { alarmId: 'A001', deviceId: 'Smart 3200', alarmType: '烟雾报警', alarmTime: '2025-06-24 10:00' },
  { alarmId: 'A002', deviceId: 'Smart 3600', alarmType: '火灾报警', alarmTime: '2025-06-24 10:10' },
]);

// 右侧面板相关状态
const infoPanelOpen = ref(true);
const filterForm = ref({
  region: 'all',
  deviceInfoType: 'deviceId',
  keyword: ''
});
const activeTab = ref('deviceList');
const deviceTableData = ref([]);
const currentPage = ref(1);
const pageSize = ref(8);
const total = ref(15);

// 地图初始化
onMounted(async () => {
  try {
    const response = await axios.get('/china.json');
    echarts.registerMap('china', response.data);
    myChart = echarts.init(mapChart.value);
    renderMap();

    initDeviceTable();

    window.addEventListener('resize', () => {
      myChart && myChart.resize();
    });
  } catch (error) {
    console.error('地图初始化失败:', error);
  }
});

onUnmounted(() => {
  if (myChart) {
    myChart.dispose();
  }
  window.removeEventListener('resize', () => {
    myChart && myChart.resize();
  });
});

// 渲染地图
const renderMap = () => {
  const lineData = deviceList.value.map(device => ({
    coords: [
      [device.lng, device.lat],
      [device.lng, device.lat + 3]
    ],
    status: device.status,
    deviceInfo: device
  }));
  const scatterData = deviceList.value.map(device => ({
    value: [device.lng, device.lat + 0],
    status: device.status,
    deviceInfo: device
  }));

  const option = {
    backgroundColor: 'rgba(15, 23, 42, 0.95)',
    title: {
      text: '全国设备地理位置分布图',
      left: 'center',
      textStyle: {
        color: '#e2e8f0',
        fontSize: 20,
        fontWeight: 500
      }
    },
    tooltip: {
      trigger: 'item',
      backgroundColor: 'rgba(15, 23, 42, 0.85)',
      borderColor: 'rgba(54, 207, 251, 0.5)',
      borderWidth: 1,
      textStyle: {
        color: '#E2E8F0'
      },
      formatter: (params) => {
        const device = params.data.deviceInfo;
        return `
          <div style="font-weight:bold; color:#36CFFB">${device.deviceName}</div>
          <div style="color:#94A3B8">状态: <span style="color:${getStatusColor(device.status)}">${device.status}</span></div>
          <div style="color:#94A3B8">位置: ${device.lng.toFixed(2)}, ${device.lat.toFixed(2)}</div>
          <div style="color:#94A3B8">地区: ${device.area}</div>
        `;
      }
    },
    geo: {
      map: 'china',
      roam: true,
      zoom: 1.2,
      label: {
        show: true,
        color: '#94a3b8'
      },
      itemStyle: {
        areaColor: 'rgba(30, 41, 59, 0.8)',
        borderColor: 'rgba(54, 207, 251, 0.3)',
        borderWidth: 0.8
      },
      emphasis: {
        itemStyle: {
          areaColor: 'rgba(54, 207, 251, 0.1)'
        }
      }
    },
    series: [
      {
        name: '设备光柱',
        type: 'lines',
        coordinateSystem: 'geo',
        data: lineData,
        effect: {
          show: true,
          period: 4,
          trailLength: 0.4,
          color: '#36CFFB',
          symbolSize: 6
        },
        lineStyle: {
          color: (params) => getStatusColor(params.data.status),
          width: 2,
          opacity: 0.8,
          curveness: 0
        }
      },
      {
        name: '光柱顶端',
        type: 'scatter',
        coordinateSystem: 'geo',
        data: scatterData,
        symbolSize: 6,
        itemStyle: {
          color: (params) => getStatusColor(params.data.status)
        },
        label: {
          show: false
        }
      }
    ]
  };

  myChart.setOption(option);

  // 地图区域点击事件：筛选对应地区设备
  myChart.on('click', (params) => {
    if (params.componentType === 'geo') {
      const areaName = params.name;
      filterDeviceByArea(areaName);
    }
  });
};

// 根据设备状态获取颜色
const getStatusColor = (status) => {
  switch (status) {
    case '在线': return '#22c55e';
    case '离线': return '#64748b';
    case '故障': return '#ef4444';
    default: return '#3b82f6';
  }
};

// 初始化设备表格数据
const initDeviceTable = () => {
  deviceTableData.value = deviceList.value.map(device => ({
    deviceId: device.deviceId,
    deviceName: device.deviceName,
    area: device.area // 把地区也放进去，方便后续用
  }));
  total.value = deviceTableData.value.length;
};

// 处理区域管理点击（可扩展实际逻辑，如打开弹窗等）
const handleRegionManage = () => {
  console.log('点击区域管理');
};

// 处理筛选
const handleFilter = () => {
  const { region, deviceInfoType, keyword } = filterForm.value;
  let filteredData = [...deviceList.value];

  if (region!== 'all') {
    filteredData = filteredData.filter(device => device.lng > 110);
  }

  if (keyword) {
    filteredData = filteredData.filter(device =>
        device[deviceInfoType].toLowerCase().includes(keyword.toLowerCase())
    );
  }

  deviceTableData.value = filteredData.map(device => ({
    deviceId: device.deviceId,
    deviceName: device.deviceName,
    area: device.area
  }));
  total.value = deviceTableData.value.length;
  currentPage.value = 1;
};

// 处理重置筛选
const handleReset = () => {
  filterForm.value = {
    region: 'all',
    deviceInfoType: 'deviceId',
    keyword: ''
  };
  initDeviceTable();
};

// 处理表格分页大小变更
const handleSizeChange = (newSize) => {
  pageSize.value = newSize;
};

// 处理表格页码变更
const handleCurrentChange = (newPage) => {
  currentPage.value = newPage;
};

// 处理设备详情，跳转到设备列表页
const handleDetail = (row) => {
  console.log('查看设备详情:', row);
  // 跳转到设备列表页，可根据需要传递参数，比如设备ID等
  router.push({ name: 'EquManage', params: { deviceId: row.deviceId } });
};

const handleLocate = (row) => {
  if (!myChart || !row.lng || !row.lat) return;

  // 通过 ECharts 动作，设置地图视图中心
  myChart.dispatchAction({
    type: 'geoRoam',
    center: [row.lng, row.lat], // 设备的经纬度
    zoom: 10 // 可根据需求调整缩放级别，值越大地图越大
  });

  // 可选：高亮定位的设备点，增强交互反馈
  highlightMapPoint(row);
};

// 高亮定位的设备点（可选功能，与上面的 handleLocate 搭配）
const highlightMapPoint = (row) => {
  // 先取消之前的高亮（如果需要）
  myChart.dispatchAction({
    type: 'downplay',
    seriesIndex: 1 // 对应 scatter 系列的索引（根据你的 series 配置调整）
  });

  // 找到对应设备在 scatter 数据中的索引
  const targetIndex = scatterData.value.findIndex(item =>
      item.deviceInfo.deviceId === row.deviceId
  );

  if (targetIndex > -1) {
    myChart.dispatchAction({
      type: 'highlight',
      seriesIndex: 1, // scatter 系列索引
      dataIndex: targetIndex
    });
  }
};

// 注意：scatterData 建议抽成响应式数据，方便查找，示例如下（可在 onMounted 里初始化）
const scatterData = ref([]);
onMounted(async () => {
  // ...原有地图初始化逻辑
  scatterData.value = deviceList.value.map(device => ({
    value: [device.lng, device.lat + 0],
    status: device.status,
    deviceInfo: device
  }));
  // ...
});

// 点击地图区域：筛选对应地区设备
const filterDeviceByArea = (areaName) => {
  const filteredData = deviceList.value.filter(device => device.area === areaName);
  deviceTableData.value = filteredData.map(device => ({
    deviceId: device.deviceId,
    deviceName: device.deviceName,
    area: device.area
  }));
  total.value = deviceTableData.value.length;
  currentPage.value = 1;
};

// 处理设备删除（可扩展实际逻辑，如调用接口删除并更新列表）
const handleDelete = (row) => {
  console.log('删除设备:', row);
  deviceList.value = deviceList.value.filter(device => device.deviceId!== row.deviceId);
  initDeviceTable();
  renderMap();
};

// 处理告警详情（可扩展实际逻辑，如打开弹窗处理告警）
const handleAlarmDetail = (row) => {
  console.log('处理告警:', row);
};

// 监听设备列表变化，更新表格和地图
watch(deviceList, () => {
  initDeviceTable();
  renderMap();
});
</script>

<style scoped>
.map-container {
  width: 100%;
  height: 750px;
  position: relative;
  display: flex;
  background: linear-gradient(135deg, #0F172A 0%, #1E293B 100%);
}

.map-chart {
  flex: 1;
  height: 100%;
}
.right-panel {
  width: 360px;
  background: rgba(15, 23, 42, 0.9);
  border-left: 1px solid rgba(54, 207, 251, 0.2);
  display: flex;
  flex-direction: column;
  color: #e2e8f0;
}

.panel-header {
  padding: 16px 20px;
  border-bottom: 1px solid rgba(54, 207, 251, 0.1);
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.info-panel-toggle {
  display: flex;
  align-items: center;
  gap: 8px;
}

.filter-section {
  padding: 16px 20px;
  border-bottom: 1px solid rgba(54, 207, 251, 0.1);
}

.filter-form {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.tab-section {
  flex: 1;
  padding: 10px;
  overflow: hidden
}
</style>