<template>
  <div class="echarts-base" :style="{ width: width, height: height }"></div>
</template>

<script>
import * as echarts from 'echarts'

export default {
  name: 'EChartsBase',
  props: {
    width: {
      type: String,
      default: '100%'
    },
    height: {
      type: String,
      default: '400px'
    }
  },
  data() {
    return {
      chartInstance: null
    }
  },
  mounted() {
    this.initChart()
  },
  beforeDestroy() {
    if (this.chartInstance) {
      this.chartInstance.dispose()
    }
  },
  methods: {
    initChart() {
      this.chartInstance = echarts.init(this.$el)
      // 可在这里统一配置默认主题、响应式等
      window.addEventListener('resize', this.handleResize)
    },
    handleResize() {
      if (this.chartInstance) {
        this.chartInstance.resize()
      }
    },
    // 暴露给父组件设置option的方法
    setOption(option) {
      if (this.chartInstance) {
        this.chartInstance.setOption(option)
      }
    }
  }
}
</script>

<style scoped>
.echarts-base {
  overflow: hidden;
}
</style>