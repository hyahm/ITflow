<template>
  <Suspense>
    <article-detail :is-edit="false" />
  </Suspense>
</template>

<script lang="ts" setup>
import ArticleDetail from './components/ArticleDetail.vue';
</script>
