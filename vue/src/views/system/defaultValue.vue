<template>
  <div>
    <p class="warn-content">
      某些选项的默认值
    </p>
    <div style="margin: 10px 0 10px 10px"> 新建bug的状态:
      <el-select v-model="defaultstatus" placeholder="Select" @change="handleChangeStatus">
        <el-option
          v-for="(status, index) in statuslist"
          :key="index"
          :label="status"
          :value="status"
        />
      </el-select>
    </div>
    <!-- <div style="margin: 10px 0 10px 10px"> 重要程度:
      <el-select v-model="form.defaultimportant" placeholder="Select" @change="handleChangeimportants">
        <el-option
          v-for="(important, index) in importants"
          :key="index"
          :label="important"
          :value="important"
        />
      </el-select>
    </div>
    <div style="margin: 10px 0 10px 10px"> 严重级别:
      <el-select v-model="form.defaultlevel" placeholder="Select" @change="handleChangeLevel">
        <el-option
          v-for="(important, index) in levels"
          :key="index"
          :label="important"
          :value="important"
        />
      </el-select>
    </div> -->
    <el-button type="primary" plain @click="handleSave">保存</el-button>
  </div>
</template>

<script>
import { status, save, important, level } from '@/api/defaultvalue'
import { getStatus, getImportants, getLevels } from '@/api/get'
export default {
  name: 'DefaultValue',
  data() {
    return {
      // form: {
      //   defaultstatus: '',
      //   defaultimportant: '',
      //   defaultlevel: ''
      // },
      defaultstatus: '',
      statuslist: [],
      importants: [],
      levels: []
    }
  },
  created() {
    this.getdefaultstatus()
    this.getstatuslist()
    // this.getimportantlist()
    // this.getlevels()
    // this.getdefaultimportant()
    // this.getdefaultlevel()
  },
  methods: {
    getlevels() {
      getLevels().then(resp => {
        if (resp.data.code === 0) {
          if (resp.data.levels !== null) {
            this.levels = resp.data.levels
          }
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    getdefaultlevel() {
      level().then(resp => {
        if (resp.data.code === 0) {
          this.form.defaultlevel = resp.data.defaultlevel
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    getimportantlist() {
      getImportants().then(resp => {
        if (resp.data.code === 0) {
          if (resp.data.importantlist !== null) {
            this.importants = resp.data.importants
          }
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    getdefaultimportant() {
      important().then(resp => {
        if (resp.data.code === 0) {
          this.form.defaultimportant = resp.data.defaultimportant
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    getdefaultstatus() {
      status().then(resp => {
        console.log(resp.data)
        if (resp.data.code === 0) {
          this.defaultstatus = resp.data.defaultstatus
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    getstatuslist() {
      getStatus().then(resp => {
        if (resp.data.code === 0) {
          if (resp.data.statuslist !== null) {
            this.statuslist = resp.data.statuslist
          }
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    handleChangeimportants(e) {
      this.form.defaultimportant = e
    },
    handleChangeStatus(e) {
      this.defaultstatus = e
    },
    handleChangeLevel(e) {
      this.form.defaultlevel = e
    },
    handleSave() {
      const data = {
        'defaultstatus': this.defaultstatus
      }
      save(data).then(resp => {
        console.log(resp.data)
        if (resp.data.code === 0) {
          this.$message.success('保存成功')
        } else {
          this.$message.error(resp.data.code)
        }
      })
    }
  }
}
</script>

<style scoped>

</style>
