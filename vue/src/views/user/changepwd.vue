<template>
  <div>
    <p class="warn-content">
      修改密码
    </p>
    <el-form class="form-container">
      <div style="height: 20px" />
      <el-form-item label="旧密码:">
        <el-input
          v-model="oldpassword"
          :maxlength="100"
          placeholder="请输入标题"
          type="password"
          clearable
          style="width: 60%;"
        />
      </el-form-item>

      <el-form-item label="新密码:">
        <el-input
          v-model="newpassword"
          :maxlength="100"
          placeholder="请输入标题"
          type="password"
          clearable
          style="width: 60%;"
        />
      </el-form-item>

      <el-form-item label="新密码:">
        <el-input
          v-model="repassword"
          :maxlength="100"
          placeholder="请输入标题"
          type="password"
          clearable
          style="width: 60%;"
        />
      </el-form-item>

      <div>
        <el-button type="success" style="margin-left: 20px" plain @click="changepwd">修改</el-button>
        <el-button type="warning" style="margin-left: 20px" plain @click="clean">清空</el-button>
      </div>
    </el-form>
  </div>
</template>

<script>
import { updatePassword } from '@/api/user'
export default {
  name: 'Changepwd',
  data() {
    return {
      oldpassword: '',
      newpassword: '',
      repassword: ''
    }
  },
  methods: {
    changepwd() {
      if (this.newpassword !== this.repassword) {
        this.$message({
          message: '新密码不一致',
          type: 'error'
        })
        return
      }
      const ch = {
        oldpassword: this.oldpassword,
        newpassword: this.newpassword
      }
      updatePassword(ch).then(response => {
        if (response.data.code === 0) {
          this.$message({
            message: '修改密码成功',
            type: 'success'
          })
          this.clean()
        } else {
          this.$message({
            message: '修改密码失败',
            type: 'error'
          })
        }
      })
    },
    clean() {
      this.oldpassword = ''
      this.newpassword = ''
      this.repassword = ''
    }
  }
}
</script>

<style type="text/css">
.el-form-item {
  margin-left: 20px;
}
</style>
