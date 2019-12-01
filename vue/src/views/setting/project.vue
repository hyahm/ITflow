<template>
  <div>
    <p class="warn-content">
      一个公司可能不止一个项目
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
      <el-table-column label="操作" width="200">
        <template slot-scope="scope">
          <el-button
            size="mini"
            @click="updatep(scope.row)"
          >修改</el-button>
          <el-button
            size="mini"
            type="danger"
            @click="handleDelete(scope.row.id)"
          >删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-button style="margin: 20px" type="success" size="mini" @click="addp">添加项目名</el-button>
    <el-dialog :visible.sync="dialogFormVisible" width="60%" title="项目管理">
      <el-form :model="form">
        <el-form-item label="项目名">
          <el-input v-model="form.projectname" width="200" auto-complete="off" />
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
import { getProjectName, addProjectName, updateProjectName, deleteProjectName } from '@/api/project'
export default {
  name: 'Addproject',
  data() {
    return {
      dialogFormVisible: false,
      form: {
        projectname: '',
        delivery: false,
        id: -1
      },
      formLabelWidth: '120px',
      tableData: []
    }
  },
  created() {
    this.getprojectname()
  },
  methods: {
    getprojectname() {
      getProjectName().then(resp => {
        if (resp.data.code === 0) {
          if (resp.data.projectlist === null) {
            this.$message.info('no data')
            return
          }
          this.tableData = resp.data.projectlist
        } else {
          this.$message.error(resp.data.msg)
        }
      })
    },
    addp() {
      this.form.id = -1
      this.dialogFormVisible = true
    },
    handleDelete(id) {
      deleteProjectName(id).then(resp => {
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
          this.$message.error(resp.data.msg)
        }
      })
    },
    updatep(row) {
      this.dialogFormVisible = true
      this.form.id = row.id
      this.form.projectname = row.projectname
    },
    confirm() {
      this.dialogFormVisible = false
      if (this.form.id === -1) {
        addProjectName(this.form.projectname).then(resp => {
          if (resp.data.code === 0) {
            this.tableData.push({
              id: resp.data.id,
              projectname: this.form.projectname
            })
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
                this.tableData[i].projectname = this.form.projectname
                break
              }
            }
            this.$message.success('更新成功')
          } else {
            this.$message.error(resp.data.msg)
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
