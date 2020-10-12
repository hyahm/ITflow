<template>
  <div>
    <p class="warn-content">
      某些选项的默认值
    </p>
    <div style="margin: 10px 0 10px 10px"> bug创建时的状态:
      <el-select v-model="form.created" placeholder="Select" @change="handleCreated">
        <el-option
          v-for="(status, index) in statuslist"
          :key="index"
          :label="status"
          :value="status"
        />
      </el-select>
    </div>
    <div style="margin: 10px 0 10px 10px"> bug完成时的状态:
      <el-select v-model="form.completed" placeholder="Select" @change="handleCompleted">
        <el-option
          v-for="(status, index) in statuslist"
          :key="index"
          :label="status"
          :value="status"
        />
      </el-select>
    </div>
    <!-- <div style="margin: 10px 0 10px 10px"> 严重级别:
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
      form: {
        created: '',
        completed: ''
        // defaultlevel: ''
      },

      statuslist: [],
      importants: [],
      levels: []
    }
  },
  created() {
    this.getdefaultstatus()
    this.getstatuslist()
  },
  methods: {
    getlevels() {
      getLevels().then(resp => {
        if (resp.data.code === 0) {
          this.levels = resp.data.levels
        } else {
          this.$message.error(resp.data.msg)
        }
      })
    },
    getdefaultlevel() {
      level().then(resp => {
        if (resp.data.code === 0) {
          this.form.defaultlevel = resp.data.defaultlevel
        } else {
          this.$message.error(resp.data.msg)
        }
      })
    },
    getimportantlist() {
      getImportants().then(resp => {
        if (resp.data.code === 0) {
          this.importants = resp.data.importants
        } else {
          this.$message.error(resp.data.msg)
        }
      })
    },
    getdefaultimportant() {
      important().then(resp => {
        if (resp.data.code === 0) {
          this.form.defaultimportant = resp.data.defaultimportant
        } else {
          this.$message.error(resp.data.msg)
        }
      })
    },
    getdefaultstatus() {
      status().then(resp => {
        if (resp.data.code === 0) {
          this.form.created = resp.data.created
          this.form.completed = resp.data.completed
        } else {
          this.$message.error(resp.data.msg)
        }
      })
    },
    getstatuslist() {
      getStatus().then(resp => {
        if (resp.data.code === 0) {
          this.statuslist = resp.data.statuslist
        } else {
          this.$message.error(resp.data.msg)
        }
      })
    },
    handleCompleted(e) {
      this.form.completed = e
    },
    handleCreated(e) {
      this.form.created = e
    },
    handleChangeLevel(e) {
      this.form.defaultlevel = e
    },
    handleSave() {
      save(this.form).then(resp => {
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
