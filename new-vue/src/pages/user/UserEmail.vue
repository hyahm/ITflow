<template>
  <div>
    <el-form class="form-container">
      <div style="height: 20px" />

      <el-form-item label="邮箱:">
        <el-input v-model="email" :maxlength="100" placeholder="请输入标题" style="width: 60%" />
      </el-form-item>

      <div>
        <el-button type="success" style="margin-left: 20px" plain @click="changeEmail"
          >修改</el-button
        >
      </div>
    </el-form>
  </div>
</template>

<script setup lang="ts">
import { myEmail, setEmail } from 'src/api/user';
import { ref } from 'vue';
const email = ref('');
const newpassword = ref('');
const repassword = ref('');

async function getEmailHandle() {
  await myEmail().then((response) => {
    email.value = response.data.email;
  });
}
await getEmailHandle();
async function changeEmail() {
  if (newpassword.value !== repassword.value) {
    // this.$message({
    //   message: '新密码不一致',
    //   type: 'error',
    // });
    return;
  }
  await setEmail(email.value).then(() => {
    // this.$message({
    //   message: '修改密码成功',
    //   type: 'success',
    // });
  });
}
</script>

<style type="text/css">
.el-form-item {
  margin-left: 20px;
}
</style>
