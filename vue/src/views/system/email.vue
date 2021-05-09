<template>
  <div style="padding-left: 20px">
   
    <el-form ref="form" label-width="100px" class="form-container"  prop="title"> 
      <p class="warn-content">
      设置并启用后 只有创建用户，创建bug，转交bug才会有邮件通知
    </p>
       <el-form-item
   
        style="width: 500px"    >
          <el-switch
        v-model="enable"
        style="display: block;"
        active-color="#13ce66"
        inactive-color="#ff4949"
        active-text="启用"
        inactive-text="禁用"
      />
         </el-form-item>
    

      <el-form-item
        label="邮箱服务器："
        clearable
        style="width: 500px"    >
        <el-input v-model="host" type="email" placeholder="请输入邮箱地址" />
      </el-form-item>

      <el-form-item
        label="昵称："
        clearable
        style="width: 500px"  >
        <el-input v-model="nickname" type="text" placeholder="请输入昵称不填默认邮箱的名称" />
      </el-form-item>

      <el-form-item
        label="邮箱地址："
        clearable
        style="width: 500px"      >
        <el-input v-model="email" type="text" placeholder="请输入邮箱地址" />
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
        <el-input v-model.number="port" type="number" placeholder="请输入邮箱端口" />
      </el-form-item>
      <el-form-item
        label="测试邮箱："
        clearable
        style="width: 500px"
      >
        <el-input v-model="to" type="text" placeholder="请输入接收邮箱" />
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
      nickname: '',
      to: ''

    }
  },
  created() {
    this.getemail()
  },
  methods: {
    getemail() {
      getEmailStatus().then(resp => {
        console.log(resp.data.email)
        this.email = resp.data.email
        this.enable = resp.data.enable
        this.host = resp.data.host
        this.port = resp.data.port
        this.nickname = resp.data.nickname
        this.id = resp.data.id
        this.password = resp.data.password
      })
    },
    handleTest() {
      const rules = [
        {
          filed: this.host,
          msg: 'host不能为空'
        },
         {
          filed: this.to,
          msg: '收件人不能为空'
        },
        {
          filed: this.email,
          msg: '邮箱账号不能为空'
        },
        {
          filed: this.password,
          msg: '邮箱地址不能为空'
        }
      ]
      if (this.port === 0) {
        this.port = 25
      }
      for (let v of rules) {
        if (v.filed === "") {
          this.$message.error(v.msg)
          return
        }
      }
      const data = {
        'host': this.host,
        'enable': this.enable,
        'port': parseInt(this.port),
        'email': this.email,
        'password': this.password,
        'nickname': this.nickname,
        'to': this.to
      }
      testEmail(data).then(_ => {
        this.$message.success('发送成功')
      })
    },
    handleSave() {
      const data = {
        'id': this.id,
        'host': this.host,
        'enable': this.enable,
        'port': parseInt(this.port),
        'email': this.email,
        'password': this.password,
        'nickname': this.nickname
      }
      saveEmail(data).then(resp => {
          this.id = resp.data.id
          this.$message.success('保存成功')
      })
    }
  }
}
</script>

<style scoped>

</style>
