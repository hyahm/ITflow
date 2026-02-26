<template>
  <el-form>
    <el-form-item label="Name">
      <el-input v-model.trim="user.name" />
    </el-form-item>
    <el-form-item label="Email">
      <el-input v-model.trim="user.email" />
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="submit">Update</el-button>
    </el-form-item>
  </el-form>
</template>

<script setup lang="ts">
import { toRefs } from 'vue';

const { user } = toRefs(defineProps(['user']));
// props: {
//   user: {
//     type: Object,
//     default: () => {
//       return {
//         name: '',
//         email: '',
//       };
//     },
//   },
// },
function submit() {
  // this.$message({
  //   message: 'User information has been updated successfully',
  //   type: 'success',
  //   duration: 5 * 1000,
  // });
}
// };
</script>
