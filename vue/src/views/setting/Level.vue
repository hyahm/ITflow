<template>
  <div>
    <p class="warn-content">
      bug优先级别
    </p>
    <el-table
      :data="tableData"
      height="250"
      style="width: 100%"
    >
      <el-table-column
        label="Id"
        width="180"
      >
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="级别名"
        width="180"
      >
        <template slot-scope="scope">
          <span>{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column width="200" label="操作">
        <template slot-scope="scope">
          <el-button
            size="mini"
            @click="handleUpdate(scope.row)"
          >修改</el-button>
          <el-button
            size="mini"
            type="danger"
            @click="handleDelete(scope.row.id)"
          >删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <div>
      <el-button type="success" plain style="margin: 20px" @click="addstatus">添加重要性</el-button>
    </div>
    <el-dialog :close-on-click-modal="false" :visible.sync="dialogFormVisible" title="优先级别">
      <el-form>
        <el-form-item label="优先级别">
          <el-input v-model="form.name" auto-complete="off" />
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
import { getLevels, addLevel, delLevel, updateLevel } from '@/api/level'
export default {
  name: 'Level',
  data() {
    return {
      tableData: [],
      statuslist: [],
      defaultstatus: '',
      dialogFormVisible: false,
      status: '',
      form: {
        id: -1,
        name: ''
      }
    }
  },
  activated() {
    this.getstatus()
  },
  created() {
    this.getstatus()
  },
  methods: {
    getstatus() {
      getLevels().then(resp => {
        if (resp.data.code === 0) {
          if (resp.data.levels != null) {
            this.tableData = resp.data.levels
            const l = this.tableData.length
            for (let i = 0; i < l; i++) {
              this.statuslist.push(this.tableData[i].name)
            }
          }
        } else {
          this.$message.error(resp.data.msg)
        }
      })
    },
    confirm() {
      if (this.form.id === -1) {
        addLevel(this.form).then(resp => {
          if (resp.data.code === 0) {
            this.tableData.push({
              id: resp.data.id,
              name: this.form.name
            })
          } else {
            this.$message.error(resp.data.msg)
          }
        })
      } else {
        updateLevel(this.form).then(resp => {
          if (resp.data.code === 0) {
            const l = this.tableData.length
            for (let i = 0; i < l; i++) {
              if (this.tableData[i].id === this.form.id) {
                this.tableData[i].name = this.form.name
              }
            }
          } else {
            this.$message.error(resp.data.msg)
          }
        })
      }
      this.dialogFormVisible = false
    },
    cancel() {
      this.dialogFormVisible = false
    },
    handleDelete(id) {
      this.$confirm('此操作将关闭bug, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        delLevel(id).then(resp => {
          if (resp === undefined) {
            return
          }

          if (resp.data.code === 0) {
            const l = this.tableData.length
            for (let i = 0; i < l; i++) {
              if (this.tableData[i].id === id) {
                this.tableData.splice(i, 1)
              }
            }
            this.$message.success('删除成功')
            return
          }
          this.$message.error('操作失败')
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
      })
    },
    addstatus() {
      this.dialogFormVisible = true
      this.form.id = -1
      this.form.name = ''
    },
    handleUpdate(row) {
      this.dialogFormVisible = true
      this.form.id = row.id
      this.form.name = row.name
    }
  }
}
</script>

<style scoped>

</style>
