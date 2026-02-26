<template>
  <div class="app-container">
   

    <bug-search :pageType="2" />
  </div>
</template>

<script setup lang="ts">
import BugSearch from 'src/pages/bug/components/BugSearch.vue';
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
.edit-input {
  padding-right: 100px;
}
.cancel-btn {
  position: absolute;
  right: 15px;
  top: 10px;
}
</style>
