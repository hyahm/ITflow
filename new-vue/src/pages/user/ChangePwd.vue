<template>
  <div>
    <p class="warn-content" style="padding-left: 20px">修改密码</p>
    <el-form class="form-container">
      <div style="height: 20px" />
      <el-form-item label="旧密码:">
        <el-input
          v-model="form.oldpassword"
          :maxlength="100"
          placeholder="请输入标题"
          type="password"
          clearable
          style="width: 60%"
        />
      </el-form-item>

      <el-form-item label="新密码:">
        <el-input
          v-model="form.newpassword"
          :maxlength="100"
          placeholder="请输入标题"
          type="password"
          clearable
          style="width: 60%"
        />
      </el-form-item>

      <el-form-item label="新密码:">
        <el-input
          v-model="form.repassword"
          :maxlength="100"
          placeholder="请输入标题"
          type="password"
          clearable
          style="width: 60%"
        />
      </el-form-item>

      <div>
        <el-button type="success" style="margin-left: 20px" plain @click="changepwd"
          >修改</el-button
        >
        <el-button type="warning" style="margin-left: 20px" plain @click="clean">清空</el-button>
      </div>
    </el-form>
  </div>
</template>

<script setup lang="ts">
import { updatePassword } from 'src/api/user';
import { reactive } from 'vue';
const form = reactive({
  oldpassword: '',
  newpassword: '',
  repassword: '',
});
function clean() {
  form.oldpassword = '';
  form.newpassword = '';
  form.repassword = '';
}
async function changepwd() {
  if (form.newpassword !== form.repassword) {
    return;
  }
  const ch = {
    oldpassword: form.oldpassword,
    newpassword: form.newpassword,
  };
  await updatePassword(ch).then(() => {
    clean();
  });
}
</script>

<style type="text/css">
.el-form-item {
  margin-left: 20px;
}
</style>
