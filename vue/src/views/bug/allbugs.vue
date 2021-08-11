<template>
  <div class="app-container">
    <p class="warn-content">
      如果没看到需要的bug, 请勾选需要显示的状态 ,永久保存，多页面生效
    </p>
    <div class="filter-container">
      <bug-search :pageType="3" />
    </div>
  </div>
</template>

<script>
import BugSearch from "@/views/bug/components/Search";

export default {
  name: "Allbugs",
  components: {
    BugSearch
  }
};
</script>
