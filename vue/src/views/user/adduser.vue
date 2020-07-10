<template>
  <div>
    <p class="warn-content">
      状态组是相对于bug管理的状态的，角色组是共享文件夹和接口文档的权限，都是必须项，姓名，邮箱，姓名都必须是唯一值
    </p>
    <el-form ref="postForm" :model="postForm" class="form-container">
      <div class="createPost-main-container">
        <!--<el-col :span="24" >-->
        <div style="height: 30px" />
        <el-form-item style="margin-bottom: 40px;" prop="title" label="昵称:">
          <el-input
            v-model="postForm.nickname"
            :maxlength="100"
            placeholder="姓名首字母"
            clearable
            style="width: 60%;"
          />
        </el-form-item>

        <el-form-item style="margin-bottom: 40px;" prop="title" label="姓名:">
          <el-input
            v-model="postForm.realname"
            :maxlength="100"
            placeholder="请输入姓名"
            clearable
            style="width: 60%;"
          />
        </el-form-item>

        <el-form-item style="margin-bottom: 40px;" prop="title" label="邮箱:">
          <el-input
            v-model="postForm.email"
            :maxlength="100"
            placeholder="请输入邮箱"
            clearable
            style="width: 60%;"
          />
        </el-form-item>

        <el-form-item style="margin-bottom: 40px;" prop="title" label="密码:">
          <el-input
            v-model="postForm.password"
            :maxlength="100"
            placeholder="请输入密码"
            type="password"
            clearable
            style="width: 60%;"
          />
        </el-form-item>

        <el-form-item style="margin-bottom: 40px;" prop="title" label="确认:">
          <el-input
            v-model="postForm.repassword"
            :maxlength="100"
            placeholder="请输入密码"
            type="password"
            clearable
            style="width: 60%;"
          />
        </el-form-item>

        <el-form-item style="margin-bottom: 40px;" prop="title" label="状态组:">
          <el-select v-model="postForm.statusgroup" class="filter-item" style="width: 130px">
            <el-option v-for="(status, index) in statusgroups" :key="index" :label="status" :value="status" />
          </el-select>
        </el-form-item>
        <el-form-item style="margin-bottom: 40px;" prop="title" label="角色组:">
          <el-select v-model="postForm.rolegroup" class="filter-item" style="width: 130px">
            <el-option v-for="(role, index) in rolegroups" :key="index" :label="role" :value="role" />
          </el-select>
        </el-form-item>
        <el-form-item style="margin-bottom: 40px;" prop="title" label="职位名:">
          <el-select v-model="postForm.position" placeholder="Select">
            <el-option
              v-for="(position, index) in positions"
              :key="index"
              :label="position"
              :value="position"
            />
          </el-select>
        </el-form-item>
        <div>
          <el-button type="success" style="margin-left: 10px" plain @click="adduser">添加</el-button>
          <el-button type="warning" style="margin-left: 20px" plain @click="clean">清空</el-button>
        </div>
        <div style="margin-bottom: 30px;height: 30px" />
      </div>
    </el-form>
  </div>
</template>

<script>
import { createUser } from '@/api/user'
import { getStatusGroupName } from '@/api/statusgroup'
import { getRoleGroup } from '@/api/rolegroup'
import { getPositions } from '@/api/position'
export default {
  name: 'Adduser',
  data() {
    return {
      postForm: {
        nickname: '',
        email: '',
        password: '',
        repassword: '',
        realname: '',
        rolegroup: '',
        statusgroup: '',
        level: 2
      },
      number: {
        one: 1,
        two: 2
      },
      positions: [],
      checkAll: false,
      rolegroups: [],
      statusgroups: [],
      checkedRoles: [],
      isIndeterminate: true,
      groups: null
    }
  },
  created() {
    this.getrolegroups()
    this.getpositions()
    this.getstatusgroups()
  },
  methods: {
    getpositions() {
      getPositions().then(resp => {
        if (resp.data.code === 0) {
          if (resp.data.positions !== null) {
            this.positions = resp.data.positions
          }
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    getstatusgroups() {
      getStatusGroupName().then(resp => {
        if (resp.data.code === 0) {
          if (resp.data.names.length > 0) {
            this.statusgroups = resp.data.names
          }
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    getrolegroups() {
      getRoleGroup().then(resp => {
        if (resp.data.code === 0) {
          if (resp.data.rolelist !== null) {
            this.rolegroups = resp.data.rolelist
          }
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    adduser() {
      const regEmail = /^.*@.+\.[A-Za-zd]{2,5}$/
      if (this.postForm.email === '') {
        this.$message({
          message: '请输入邮箱',
          type: 'error'
        })
        return
      }
      if (!regEmail.test(this.postForm.email)) {
        this.$message({
          message: '邮箱格式正确',
          type: 'error'
        })
        return
      }
      if (this.postForm.repassword !== this.postForm.password) {
        this.$message({
          message: '2次密码不对',
          type: 'error'
        })
        return
      }
      if (this.postForm.statusgroup === '') {
        this.$message({
          message: '请选择状态组',
          type: 'error'
        })
        return
      }
      if (this.postForm.rolegroups === '') {
        this.$message({
          message: '请选择角色组',
          type: 'error'
        })
        return
      }
      if (this.postForm.position === '') {
        this.$message({
          message: '请选择职位',
          type: 'error'
        })
        return
      }
      createUser(this.postForm).then(resp => {
        if (resp.data.code === 0) {
          this.$message({
            message: '添加用户成功',
            type: 'success'
          })
          this.clean()
        } else {
          this.$message({
            message: '添加用户失败',
            type: 'error'
          })
        }
      })
    },
    clean() {
      this.postForm = {
        nickname: '',
        email: '',
        password: '',
        repassword: '',
        rolegroup: '',
        statusgroup: ''
      }
    }
  }
}
</script>

<style type="text/css" >
  .form-container .createPost-main-container .el-form-item {
    padding: 0px 15px 0px 15px;
  }
</style>
