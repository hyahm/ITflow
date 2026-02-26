<template>
  <div>
    <p class="warn-content">修改用户信息，都只能唯一</p>
    <el-form ref="postForm" :model="postForm" class="form-container">
      <div class="createPost-main-container">
        <!--<el-col :span="24" >-->
        <div style="height: 30px" />
        <el-form-item prop="title" label="昵称:">
          <el-input
            v-model="postForm.nickname"
            :maxlength="100"
            placeholder="姓名首字母"
            clearable
            style="width: 60%"
          />
        </el-form-item>

        <el-form-item prop="title" label="姓名:">
          <el-input
            v-model="postForm.realname"
            :maxlength="100"
            placeholder="请输入姓名"
            clearable
            style="width: 60%"
          />
        </el-form-item>

        <el-form-item prop="title" label="邮箱:">
          <el-input
            v-model="postForm.email"
            :maxlength="100"
            placeholder="请输入邮箱"
            clearable
            style="width: 60%"
          />
        </el-form-item>
        <div>
          <el-button type="success" style="margin-left: 40px" plain @click="handleUpdate"
            >修改</el-button
          >
        </div>
      </div>
    </el-form>
  </div>
</template>

<script setup lang="ts">
import { updateUser, getInfo } from 'src/api/user';
import type { User } from 'src/api/types/type';
import { reactive } from 'vue';
const postForm: User = reactive({
  nickname: '',
  realname: '',
  email: '',
  id: 0,
  password: '*****',
  headimg: '',
  disable: false,
  position_id: 0,
});
async function getinfo() {
  await getInfo().then((resp) => {
    console.log(resp.data)
    Object.assign(postForm, resp.data);
  });
}
async function handleUpdate() {
  await updateUser(postForm).then(() => {
    // this.$message.success('修改成功');
  });
}

await getinfo();
</script>
