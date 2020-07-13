<template>
  <div>
    <p class="warn-content">
      选择邮件通知的项目，测试邮件是为了验证是否能正确收到邮件通知
    </p>
    <el-form ref="form" label-width="100px" class="form-container" style="margin-top: 40px;" prop="title">
      <!-- <el-switch
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
      /> -->
      <el-switch
        v-model="enable"
        style="display: block;margin: 20px"
        active-color="#13ce66"
        inactive-color="#ff4949"
        active-text="启用"
        inactive-text="禁用"
      />

      <el-form-item
        label="邮箱服务器："
        clearable
        style="width: 500px"
      >
        <el-input v-model="host" type="email" placeholder="请输入邮箱地址" />
      </el-form-item>

      <el-form-item
        label="邮箱地址："
        clearable
        style="width: 500px"
      >
        <el-input v-model="email" type="email" placeholder="请输入邮箱地址" />
      </el-form-item>

      <el-form-item
        label="邮箱密码："
        clearable
        style="width: 500px"
      >
        <el-input v-model="password" type="password" placeholder="请输入邮箱密码" />
      </el-form-item>
      <el-form-item
        label="邮箱端口："
        clearable
        style="width: 500px"
      >
        <el-input v-model="port" type="number" placeholder="请输入邮箱端口" />
      </el-form-item>
      <el-form-item
        label="测试邮箱："
        clearable
        style="width: 500px"
      >
        <el-input v-model="to" type="email" placeholder="请输入接收邮箱" />
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
      email: '',
      password: '',
      port: 25,
      host: '',
      id: 0,
      enable: false,
      // 验证邮箱
      to: ''

    }
  },
  created() {
    this.getemail()
  },
  methods: {
    getemail() {
      getEmailStatus().then(resp => {
        console.log(resp.data)
        if (resp.data.code === 0) {
          this.email = resp.data.email
          this.enable = resp.data.enable
          this.host = resp.data.host
          this.port = resp.data.port
          this.id = resp.data.id
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    handleTest() {
      const data = {
        'host': this.host,
        'enable': this.enable,
        'port': parseInt(this.port),
        'email': this.email,
        'password': this.password,
        'to': this.to
      }
      testEmail(data).then(resp => {
        if (resp.data.code === 0) {
          this.$message.success('发送成功')
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    handleSave() {
      console.log(this.id)
      const data = {
        'id': this.id,
        'host': this.host,
        'enable': this.enable,
        'port': parseInt(this.port),
        'email': this.email,
        'password': this.password
      }
      saveEmail(data).then(resp => {
        if (resp.data.code === 0) {
          this.id = resp.data.id
          this.$message.success('保存成功')
        } else {
          this.$message.error(resp.data.message)
        }
      })
    }
  }
}
</script>

<style scoped>

</style>
