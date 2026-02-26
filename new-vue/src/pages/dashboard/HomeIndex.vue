<template>
  <div class="dashboard-container">
    <component :is="currentRole" />
  </div>
</template>

<script lang="ts" setup>

import {ref} from 'vue'
  const currentRole =  ref('adminDashboard')
</script>
