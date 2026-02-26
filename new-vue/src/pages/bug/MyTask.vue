<template>
  <div class="app-container">

    <div class="filter-container">
      <bug-search :pageType="4" />
    </div>
  </div>
</template>

<script setup lang="ts">
import BugSearch from 'src/pages/bug/components/BugSearch.vue';
</script>
