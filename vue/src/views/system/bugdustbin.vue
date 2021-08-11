<template>
  <div style="padding-left: 20px">
    <p class="warn-content">
      可以改变bug的所有信息, 必须搜索才有显示
    </p>
    <bug-search :pageType="1" />
  </div>
</template>

<script>
import BugSearch from "@/views/bug/components/Search";
export default {
  name: "Movebug",
  components: {
    BugSearch
  },
  data() {
    return {};
  }
};
</script>

<style scoped></style>
