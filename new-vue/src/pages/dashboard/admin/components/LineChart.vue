<template>
  <div class="chart-container">
    <!-- 添加 v-if 确保数据存在 -->
    <div ref="chartRef" style="width: 500px; height: 600px"></div>
  </div>
</template>

<script setup lang="ts">
import * as echarts from 'echarts';
import { reactive, ref, onMounted } from 'vue';

const chartRef = ref(null);

onMounted(() => {
  // ⚠️ Make sure the element exists
  if (!chartRef.value) {
    console.error('Chart element not found');
    return;
  }
  const myChart = echarts.init(chartRef.value);
  const chartOptions = reactive({
    xAxis: {
      type: 'category' as const,
      data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
    },
    yAxis: {
      type: 'value' as const,
      min: 0, // 从 0 开始
      max: 300, // 顶到 300
      interval: 50, // 刻度间隔
    },
    series: [
      {
        data: [150, 230, 224, 218, 135, 147, 260],
        type: 'line' as const,
        smooth: true,
      },
    ],
  });

  myChart.setOption(chartOptions, true);
});
</script>
