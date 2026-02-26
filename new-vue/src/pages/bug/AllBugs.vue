<template>
  <div class="app-container">
    <div class="filter-container">
      <bug-search :pageType="3" />
    </div>
  </div>
</template>

<script lang="ts" setup>
import BugSearch from 'src/pages/bug/components/BugSearch.vue';
</script>
