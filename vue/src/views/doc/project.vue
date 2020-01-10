<template>
  <div class="chart-container">
    <div class="project-container">
      <el-card v-for="(project, index) in projects" :key="index" class="box-card">
        <span class="setting"><svg-icon icon-class="setting" /></span>
        <div style="clear: both" />
        <div class="div_box_card"><a :href="'/doc/show/' + project.id">{{ project.name }}</a></div>
      </el-card>
      <el-card class="add-card">
        <span @click="one">添加项目+</span>
      </el-card>
    </div>
    <el-dialog :close-on-click-modal="false" :visible.sync="dialogFormVisible" width="60%" title="接口项目名称">
      <el-form :model="form">
        <el-form-item label="项目名">
          <el-input v-model="form.name" width="200" auto-complete="off" />
        </el-form-item>
        <el-form-item label="权限">
          <el-radio-group v-model="form.auth" @change="handleAuth">
            <el-radio :label="!perm">无</el-radio>
            <el-radio :label="perm">授权</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item v-if="form.auth" label="只能查看">
          <el-radio-group v-model="form.readuser" @change="handleRead">
            <el-radio :label="perm">用戶</el-radio>
            <el-radio :label="!perm">組</el-radio>
          </el-radio-group>
          <el-select v-model="form.readname" placeholder="请选择">
            <el-option
              v-for="(item, index) in readlist"
              :key="index"
              :label="item"
              :value="item"
            />
          </el-select>
        </el-form-item>
        <el-form-item v-if="form.auth" label="查看编辑">
          <el-radio-group v-model="form.edituser" @change="handleRdWr">
            <el-radio :label="perm">用戶</el-radio>
            <el-radio :label="!perm">組</el-radio>
          </el-radio-group>
          <el-select v-model="form.editname" placeholder="请选择">
            <el-option
              v-for="(item, index) in rdwrlist"
              :key="index"
              :label="item"
              :value="item"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancel">取 消</el-button>
        <el-button type="primary" @click="confirm">确 定</el-button>
      </div>
    </el-dialog>
  </div>

</template>

<script>
import { addProject } from '@/api/doc'
export default {
  name: 'Project',
  data() {
    return {
      projects: [
        { name: 'aaa', id: 0 },
        { name: 'bbb', id: 1 },
        { name: 'ccc', id: 2 }
      ],
      dialogFormVisible: false,
      form: {
        readuser: true,
        edituser: true,
        readname: '',
        editname: '',
        auth: false,
        name: '',
        owner: '',
        id: -1
      },
      rdwrlist: []
    }
  },
  methods: {
    cancel() {

    },
    confirm() {

    },
    one() {
      console.log(12222)
      this.projects.push(new Date())
      console.log(this.projects)
    },
    addproject() {
      addProject().then(resp => {
        console.log(resp.data)
      })
    }
  }
}
</script>

<style>
.el-card__body {
  padding: 0px !important;
}
</style>

<style scoped>
.char-container{
  position: relative;
  padding:20px;
  width: 100%;
  height:85vh;
}
.box-card {
  padding: 0px;
  width: 180px;
  font-weight: 700;
  color: #777;
  height: 100px;
   margin: 20px;
}

.add-card {
  cursor: pointer;
  width: 180px;
  padding: 20px;
   height: 100px;
   font-weight: 700;
  margin: 20px;
}
.project-container {
  margin: auto;
  display: flex;
  /* flex-direction:column; */
  flex-wrap: wrap ;
  max-width: 700px;
  margin-top: 30px;
}
.setting {
  float: right;
}
.div_box_card {
  padding: 20px
}
</style>

