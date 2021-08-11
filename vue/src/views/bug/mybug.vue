<template>
  <div class="app-container">
    <p class="warn-content">
      如果没看到需要的bug, 请勾选需要显示的状态 ,永久保存，多页面生效
    </p>

    <bug-search :pageType="2" />
  </div>
</template>

<script>
import BugSearch from "@/views/bug/components/Search";

export default {
  name: "ArticleList",
  components: {
    BugSearch
  },

  data() {
    return {};
  }
};
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
