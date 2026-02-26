<template>
  <div class="dashboard-editor-container">
    <panel-group />
    <br />
    <line-chart />
  </div>
</template>

<script setup lang="ts">
// import GithubCorner from 'components/GithubCorner'
import PanelGroup from './components/PanelGroup.vue';
import LineChart from './components/LineChart.vue';
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
.dashboard-editor-container {
  padding: 32px;
  background-color: rgb(240, 242, 245);
  .chart-wrapper {
    background: #fff;
    padding: 16px 16px 0;
    margin-bottom: 32px;
  }
}
</style>
