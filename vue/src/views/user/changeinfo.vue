<template>
  <div>
    <p class="warn-content">
      修改用户信息，都只能唯一
    </p>
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
            style="width: 60%;"
          />
        </el-form-item>

        <el-form-item prop="title" label="姓名:">
          <el-input
            v-model="postForm.realname"
            :maxlength="100"
            placeholder="请输入姓名"
            clearable
            style="width: 60%;"
          />
        </el-form-item>

        <el-form-item prop="title" label="邮箱:">
          <el-input
            v-model="postForm.email"
            :maxlength="100"
            placeholder="请输入邮箱"
            clearable
            style="width: 60%;"
          />
        </el-form-item>
        <div>
          <el-button type="success" style="margin-left: 40px" plain @click="handleUpdate">修改</el-button>
        </div>
      </div>
    </el-form>
  </div>
</template>

<script>
import { updateInfo, getInfo } from '@/api/user'
export default {
  name: 'Changeinfo',
  data() {
    return {
      postForm: {
        nickname: '',
        realname: '',
        email: ''
      }
    }
  },
  created() {
    this.getinfo()
  },
  methods: {
    getinfo() {
      getInfo().then(resp => {
        if (resp.data.code === 0) {
          this.postForm = resp.data
        } else {
          this.$message.error(resp.data.msg)
        }
      })
    },
    handleUpdate() {
      updateInfo(this.postForm).then(resp => {
        if (resp.data.code === 0) {
          this.$message.success('修改成功')
        }
      })
    }
  }
}
</script>

