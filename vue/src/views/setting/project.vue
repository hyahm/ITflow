<template>
  <div>
    <p class="warn-content">
      如果删除的此项目的某参与者， 在他任务里面还是会显示，只不过没操作权限
    </p>
    <el-table
      :data="tableData"
      fit
      border
      highlight-current-row
      style="width: 100%"
    >
      <el-table-column
        label="Id"
        width="180"
      >
        <template slot-scope="scope">
          <span style="margin-left: 10px">{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="项目名"
        width="180"
      >
        <template slot-scope="scope">
          <span style="margin-left: 10px">{{ scope.row.projectname }}</span>
        </template>

      </el-table-column>

      <el-table-column
        label="参与者"
        width="800"
      >

        <template slot-scope="scope">
          <span style="margin-left: 10px">{{ scope.row.selectuser }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="200">
        <template slot-scope="scope">
          <el-button
            size="mini"
            @click="updatep(scope.row)"
          >修改</el-button>
          <el-button
            size="mini"
            type="danger"
            @click="handleDelete(scope.row.selectuser)"
          >删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-button style="margin: 20px" type="success" size="mini" @click="addProject">添加项目名</el-button>
    <el-dialog :close-on-click-modal="false" :visible.sync="dialogFormVisible" width="60%" title="项目管理">
      <el-form :model="form">
        <el-form-item label="项目名">
          <el-input v-model="form.projectname" width="200" auto-complete="off" />
        </el-form-item>

        <el-form-item style="display: inline-block;width: 300px" label="参与者：">
          <el-select v-model="form.selectuser" multiple placeholder="参与者">
            <el-option
              v-for="(item, index) in users"
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
import { getUsers } from '@/api/get'
import { getProjectName, addProjectName, updateProjectName, deleteProjectName } from '@/api/project'
export default {
  name: 'Addproject',
  data() {
    return {
      dialogFormVisible: false,
      form: {
        projectname: '',
        selectuser: [],
        id: -1
      },
      users: [],
      formLabelWidth: '120px',
      tableData: []
    }
  },
  created() {
    this.getuser()
    this.getprojectname()
  },
  methods: {
    getuser() {
      getUsers().then(resp => {
        if (resp.data.code === 0) {
          if (resp.data.users !== null) {
            this.users = resp.data.users
          } else {
            this.$message.error(resp.data.message)
          }
        }
      }).catch(err => {
        this.$message.error(err)
      })
    },
    getprojectname() {
      getProjectName().then(resp => {
        console.log(resp.data)
        if (resp.data.code === 0) {
          this.tableData = resp.data.projectlist
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    addProject() {
      this.form.id = -1
      this.form.projectname = ''
      this.form.selectuser = []
      this.dialogFormVisible = true
    },
    handleDelete(id) {
      deleteProjectName(id).then(resp => {
        if (resp === undefined) {
          return
        }
        if (resp.data.code === 0) {
          const fl = this.tableData.length
          for (let i = 0; i < fl; i++) {
            if (this.tableData[i].id === id) {
              this.tableData.splice(i, 1)
              break
            }
          }
          this.$message.success('删除成功')
          return
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    updatep(row) {
      this.form = row
      this.dialogFormVisible = true
    },
    confirm() {
      this.dialogFormVisible = false
      if (this.form.projectname === '') {
        this.$message.success('至少选择一个名称')
      }
      if (this.form.selectuser.length === 0) {
        this.$message.success('至少选择一个参与者')
      }
      if (this.form.id <= 0) {
        addProjectName(this.form).then(resp => {
          if (resp.data.code === 0) {
            this.form.id = resp.data.id
            this.tableData.push(this.form)
            this.$message.success('添加成功')
          } else {
            this.$message.error('添加错误')
          }
        })
      } else {
        updateProjectName(this.form).then(resp => {
          if (resp.data.id === 0) {
            this.$message.warning('存在项目名')
            return
          }
          if (resp.data.code === 0) {
            const fl = this.tableData.length
            for (let i = 0; i < fl; i++) {
              if (this.tableData[i].id === this.form.id) {
                this.tableData[i] = this.form
                break
              }
            }
            this.$message.success('更新成功')
          } else {
            this.$message.error(resp.data.message)
          }
        })
      }
    },
    cancel() {
      this.dialogFormVisible = false
      this.form.name = ''
      this.form.id = -1
    }
  }
}
</script>

<style scoped>

</style>
