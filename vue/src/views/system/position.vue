<template>
  <div>
    <p class="warn-content">
      职位，在管理的时候，上级有管理下级的权限，级别只有普通员工和管理者，管理者也可以被其他管理者管理（从属于），一个员工只能被一个管理者管理
    </p>
    <el-table
      :data="tableData"
      height="250"
      style="width: 100%">
      <el-table-column
        label="Id"
        width="180">
        <template slot-scope="scope">
          <span >{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="职位名"
        width="180">
        <template slot-scope="scope">
          <span >{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="职位级别"
        width="180">
        <template slot-scope="scope">
          <span >{{ scope.row.level | level }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="从属于"
        width="180">
        <template slot-scope="scope">
          <span >{{ scope.row.hypo }}</span>
        </template>
      </el-table-column>
      <el-table-column width="200" label="操作">
        <template slot-scope="scope">
          <el-button
            size="mini"
            @click="handleUpdate(scope.row)">修改</el-button>
          <el-button
            size="mini"
            type="danger"
            @click="handleDelete(scope.row.id)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <div>
      <el-button type="success" plain style="margin: 20px" @click="addstatus">添加职位</el-button>
    </div>
    <el-dialog :close-on-click-modal="false" :visible.sync="dialogFormVisible" title="职位管理">
      <el-form >
        <el-form-item label="职位名">
          <el-input v-model="form.name"/>
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
          <el-select v-model="form.hypo" placeholder="Select">
            <el-option
              v-for="(hypo, index) in hypos"
              :key="index"
              :label="hypo"
              :value="hypo"/>
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
import { getPositions, addPosition, updatePosition, delPosition, getHypos, PositionsList } from '@/api/position'
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
      tableData: [],
      statuslist: [],
      dialogFormVisible: false,
      status: '',
      levelone: 1,
      leveltwo: 2,
      hypos: [],
      form: {
        id: -1,
        name: '',
        level: 0,
        hypo: ''
      }
    }
  },
  created() {
    this.getlist()
    this.gethypos()
  },
  methods: {
    gethypos() {
      getHypos().then(resp => {
        if (resp.data.statuscode === 0) {
          this.hypos = resp.data.hypos
        }
      })
    },
    getlist() {
      PositionsList().then(resp => {
        if (resp.data.statuscode === 0) {
          if (resp.data.positions != null) {
            this.tableData = resp.data.positions
          }
        }
      })
    },
    confirm() {
      if (this.form.id === -1) {
        addPosition(this.form).then(resp => {
          if (resp.data.statuscode === 0) {
            this.tableData.push({
              id: resp.data.id,
              name: this.form.name,
              hypo: this.form.hypo,
              level: this.form.level
            })
          } else {
            this.$message.error('操作失败')
          }
        })
      } else {
        updatePosition(this.form).then(resp => {
          if (resp.data.statuscode === 0) {
            const l = this.tableData.length
            for (let i = 0; i < l; i++) {
              if (this.tableData[i].id === this.form.id) {
                this.tableData[i].name = this.form.name
              }
            }
          } else {
            this.$message.error('操作失败')
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
        delPosition(id).then(resp => {
          // console
          if (resp.data.statuscode === 21) {
            this.$message.error('此职位有用户在使用')
            return
          }
          if (resp.data.statuscode === 24) {
            this.$message.error('此职位在使用')
            return
          }
          if (resp.data.statuscode === 0) {
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
      this.form = row
      this.dialogFormVisible = true
    }
  }
}
</script>

<style scoped>

</style>
