<template>
  <div>
    <p class="warn-content">
      职位，在管理的时候，上级有管理下级的权限，级别只有普通员工和管理者，管理者也可以被其他管理者管理（从属于），一个员工只能被一个管理者管理
    </p>
    <el-table
      :data="tableData"
      height="600"
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
        label="职位名"
        width="180"
      >
        <template slot-scope="scope">
          <span>{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="职位级别"
        width="180"
      >
        <template slot-scope="scope">
          <span>{{ scope.row.level | level }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="从属于"
        width="180"
      >
        <template slot-scope="scope">
          <span>{{ scope.row.hyponame }}</span>
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
            @click="handleDelete(scope.row)"
          >删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <div>
      <el-button type="success" plain style="margin: 20px" @click="addposition">添加职位</el-button>
    </div>
    <el-dialog :close-on-click-modal="false" :visible.sync="dialogFormVisible" title="职位管理">
      <el-form>
        <el-form-item label="职位名">
          <el-input v-model="form.name" />
        </el-form-item>
      </el-form>
      <el-form>
        <el-radio-group v-model="form.level">
          <el-radio :label="levelone">管理者</el-radio>
          <el-radio :label="leveltwo">普通员工</el-radio>
        </el-radio-group>
      </el-form>
      <el-form style="margin-top: 10px">
        <!--从属于哪个管理者-->
        <el-form-item label="从属于">
          <el-select v-model="form.hyponame" clearable placeholder="Select">
            <el-option
              v-for="(hypo, index) in manager"
              :key="index"
              :label="hypo.name"
              :value="hypo.name"
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
import { addPosition, updatePosition, delPosition, getHypos, PositionsList } from '@/api/position'
export default {
  name: 'Position',
  filters: {
    level: function(value) {
      switch (value) {
        case 1:
          return '管理者'
        default:
          return '普通员工'
      }
    }
  },
  data() {
    return {
      changeName: '',
      tableData: [],
      statuslist: [],
      dialogFormVisible: false,
      status: '',
      levelone: 1,
      leveltwo: 2,
      manager: [],
      form: {
        id: -1,
        name: '',
        level: 0,
        hyponame: ''
      }
    }
  },
  created() {
    this.getlist()
  },
  methods: {
    changeHypo(e) {
      console.log(e)
    },
    gethypos(id) {
      getHypos(id).then(resp => {
        if (resp.data.code === 0) {
          this.manager = resp.data.hypos
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    getlist() {
      PositionsList().then(resp => {
        if (resp.data.code === 0) {
          this.tableData = resp.data.positions
          console.log(this.tableData)
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    confirm() {
      if (this.form.id === -1) {
        addPosition(this.form).then(resp => {
          if (resp.data.code === 0) {
            this.tableData.push({
              id: resp.data.id,
              name: this.form.name,
              hyponame: this.form.hyponame,
              level: this.form.level
            })

            if (this.form.level === 1) {
              this.manager.push({
                'id': this.form.id,
                'name': this.form.hyponame
              })
            }
          } else {
            this.$message.error(resp.data.message)
          }
        })
      } else {
        updatePosition(this.form).then(resp => {
          if (resp.data.code === 0) {
            const l = this.tableData.length
            for (let i = 0; i < l; i++) {
              if (this.tableData[i].id === this.form.id) {
                this.tableData[i].name = this.form.name
                break
              }
              for (let i = 0; i < this.manager.length; i++) {
                if (this.manager[i] === this.changeName) {
                  this.manager[i] = this.form.name
                  return
                }
              }
            }
          } else {
            this.$message.error(resp.data.message)
          }
        })
      }
      this.dialogFormVisible = false
    },
    cancel() {
      this.dialogFormVisible = false
    },
    handleDelete(row) {
      this.$confirm('此操作将关闭bug, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        delPosition(row.id).then(resp => {
          // console
          if (resp === undefined) {
            return
          }

          if (resp.data.code === 0) {
            const l = this.tableData.length
            for (let i = 0; i < l; i++) {
              if (this.tableData[i].id === row.id) {
                this.tableData.splice(i, 1)
                break
              }
            }
            for (let i = 0; i < this.manager.length; i++) {
              if (this.manager[i] === row.name) {
                this.manager.splice(i, 1)
                break
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
    addposition() {
      this.dialogFormVisible = true
      this.form.id = -1
      this.form.level = 0
      this.form.hyponame = ''
      this.form.name = ''
    },
    handleUpdate(row) {
      this.form = row
      this.changeName = row.name
      this.dialogFormVisible = true
      this.gethypos(row.id)
    }
  }
}
</script>

<style scoped>

</style>
