<template>
  <div>

    <el-form class="form-container">
      <div style="height: 20px" />

      <el-form-item style="margin-bottom: 40px;" label="版本号:">
        <el-input
          v-model="versionlist.version"
          :maxlength="100"
          placeholder="请输入版本号"
          clearable
          style="width: 60%;"
        />
      </el-form-item>

      <el-form-item style="margin-bottom: 40px;" label="地址一:">
        <el-input
          v-model="versionlist.iphoneurl"
          :maxlength="100"
          placeholder="请输入iphone下载地址"
          clearable
          style="width: 60%;"
        />
      </el-form-item>

      <el-form-item style="margin-bottom: 40px;" label="地址二:">
        <el-input
          v-model="versionlist.notiphoneurl"
          :maxlength="100"
          placeholder="请输入非iphone下载地址"
          clearable
          style="width: 60%;"
        />
      </el-form-item>

      <div>
        <el-button type="success" plain style="margin-left: 20px" @click="addversion">增加</el-button>
        <el-button type="warning" plain style="margin-left: 20px" @click="clean">清空</el-button>
      </div>
    </el-form>
  </div>
</template>

<script>
import { addVersion } from '@/api/version'
export default {
  name: 'Pubversion',
  data() {
    return {
      versionlist: {
        version: '',
        iphoneurl: '',
        notiphoneurl: ''
      },
      projectnames: [],
      runenvs: [],
      platforms: []
    }
  },
  created() {
  },
  methods: {
    clean() {
      this.versionlist = {
        version: '',
        iphoneurl: '',
        notiphoneurl: ''
      }
    },
    addversion() {
      addVersion(this.versionlist).then(response => {
        if (response.data.code === 1 && response.data.id !== 0) {
          this.$message.success('版本存在')
        } else if (response.data.code === 0) {
          this.$message.success('添加成功')
        } else {
          this.$message.success('错误码' + response.data.code)
        }
      })
    }
  }
}
</script>

<style type="text/css">
  .el-form-item {
    margin-left: 20px;
  }
</style>
