<template>
  <div>
    <p class="warn-content">
      选择邮件通知的项目，测试邮件是为了验证是否能正确收到邮件通知
    </p>
    <el-form ref="form" label-width="100px" class="form-container" style="margin-top: 40px;" prop="title">
      <el-switch
        v-model="email.createuser"
        style="display: block;margin: 20px"
        active-color="#13ce66"
        inactive-color="#ff4949"
        active-text="启用创建用户通知"
        inactive-text="禁用通知"
      />
      <el-switch
        v-model="email.createbug"
        active-color="#13ce66"
        inactive-color="#ff4949"
        style="display: block;margin: 20px"
        active-text="启用创建bug通知"
        inactive-text="禁用通知"
      />
      <el-switch
        v-model="email.passbug"
        style="display: block;margin: 20px"
        active-color="#13ce66"
        inactive-color="#ff4949"
        active-text="启用转交bug通知"
        inactive-text="禁用通知"
      />
      <el-form-item
        label="邮箱地址："
        clearable
        style="width: 500px"
      >
        <el-input v-model="email.emailaddr" type="email" placeholder="请输入邮箱地址" />
      </el-form-item>
      <el-form-item
        label="邮箱密码："
        clearable
        style="width: 500px"
      >
        <el-input v-model="email.password" type="password" placeholder="请输入邮箱密码" />
      </el-form-item>
      <el-form-item
        label="邮箱端口："
        clearable
        style="width: 500px"
      >
        <el-input v-model="email.port" type="number" placeholder="请输入邮箱端口" />
      </el-form-item>
      <el-form-item
        label="测试邮箱："
        clearable
        style="width: 500px"
      >
        <el-input v-model="email.to" type="email" placeholder="请输入测试邮箱" />
      </el-form-item>
      <el-button style="margin-left: 40px" type="primary" @click="handleTest">验证</el-button>
      <el-button style="margin-left: 40px" type="primary" @click="handleSave">保存</el-button>
    </el-form>
  </div>
</template>

<script>
import { testEmail, saveEmail, getEmailStatus } from '@/api/email'
export default {
  name: 'Email',
  data() {
    return {
      email: {
        emailaddr: '',
        password: '',
        port: 25,
        to: '',
        createuser: true,
        createbug: true,
        passbug: true,
        id: -1
      }
    }
  },
  created() {
    this.getemail()
  },
  methods: {
    getemail() {
      getEmailStatus().then(resp => {
        if (resp.data.code === 0) {
          this.email = resp.data
        } else {
          this.$message.error(resp.data.msg)
        }
      })
    },
    handleTest() {
      testEmail(this.email).then(resp => {
        if (resp.data.code === 0) {
          this.$message.success('发送成功')
        } else {
          this.$message.success(resp.data.msg)
        }
      })
    },
    handleSave() {
      this.email.port = parseInt(this.email.port)
      saveEmail(this.email).then(resp => {
        if (resp.data.code === 0) {
          this.$message.success('保存成功')
        } else {
          this.$message.error(resp.data.msg)
        }
      })
    }
  }
}
</script>

<style scoped>

</style>
