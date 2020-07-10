<template>
  <div>
    <p class="warn-content">
      增加选择可以改变bug状态的组
    </p>
    <el-table v-loading="listLoading" :data="list" border fit highlight-current-row style="width: 100%;padding: 10px">

      <el-table-column label="id" align="center" width="50">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>

      <el-table-column label="项目名" width="100px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.name }}</span>
        </template>
      </el-table-column>

      <el-table-column label="状态权限" width="500" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.bugstatuslist }}</span>
        </template>
      </el-table-column>

      <el-table-column label="操作" align="center" width="230" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button type="primary" size="mini" @click="handleEdit(scope.row)">编辑</el-button>
          <el-button type="success" size="mini" @click="handleRemove(scope.row.id)">删除</el-button>
          <!--<el-button type="danger" size="mini" @click="handleRemove(scope.row)">{{ $t('list.remove') }}</el-button>-->
        </template>
      </el-table-column>
    </el-table>
    <el-button style="margin: 20px" type="success" size="mini" @click="handleAdd">添加状态组名</el-button>
    <el-dialog
      :close-on-click-modal="false"
      :visible.sync="dialogVisible"
      :before-close="handleClose"
      title="提示"
      width="60%"
    >
      <el-form ref="postForm" class="form-container">
        <el-form-item prop="title" label="状态组名:">
          <el-input
            v-model="form.name"
            :maxlength="100"
            placeholder="请输入状态组名"
            clearable
          />
        </el-form-item>
        <el-checkbox-group v-model="form.checklist">
          <div v-for="(status, index) in statuslist" :key="index">
            <el-checkbox :label="status" />
          </div>
        </el-checkbox-group>
        <!--<el-button type="success" round @click="HandlerAddGroup">添加部门</el-button>-->
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="HandlerAddGroup">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import { getStatus } from '@/api/get'
import { statusGroupList, addStatusGroup, editStatusGroup, removeStatusGroup } from '@/api/statusgroup'
export default {
  name: 'BugGroup',
  data() {
    return {

      statuslist: [],
      dialogVisible: false,
      listLoading: false,
      list: [],
      form: {
        id: -1,
        name: '',
        checklist: []
      }

    }
  },
  created() {
    this.getstatus()
    this.getlist()
  },
  methods: {
    handleEdit(row) {
      this.form.id = row.id
      this.form.name = row.name
      if (row.bugstatuslist === null) {
        this.form.checklist = []
      } else {
        this.form.checklist = row.bugstatuslist
      }

      this.dialogVisible = true
    },
    getlist() {
      statusGroupList().then(resp => {
        if (resp.data.code === 0) {
          if (resp.data.departmentlist !== null) {
            this.list = resp.data.departmentlist
          }
        } else {
          this.$message.error(resp.data.msg)
        }
      })
    },
    handleAdd() {
      this.form = {
        id: -1,
        name: '',
        checklist: []
      }
      this.dialogVisible = true
    },
    handleRemove(id) {
      this.$confirm('确认关闭？')
        .then(_ => {
          removeStatusGroup(id).then(resp => {
            if (resp.data.code === 23) {
              this.$message.warning('删除失败，此状态组有用户在使用')
              return
            }
            if (resp.data.code === 0) {
              const l = this.list.length
              for (let i = 0; i < l; i++) {
                if (this.list[i].id === id) {
                  this.list.splice(i, 1)
                }
              }
              this.$message.success('删除成功')
              return
            }
            this.$message.success('删除失败,错误码：' + resp.data.msg)
          })
        })
        .catch(_ => {})
    },
    getstatus() {
      getStatus().then(resp => {
        if (resp.data.code === 0) {
          this.statuslist = resp.data.statuslist
        }
      })
    },
    handleClose() {
      this.dialogVisible = false
    },
    HandlerAddGroup() {
      if (this.form.name === '') {
        this.$message.error('名称不能为空')
        return
      }
      if (this.form.id > 0) {
        editStatusGroup(this.form).then(resp => {
          if (resp.data.code === 0) {
            const l = this.list.length
            for (let i = 0; i < l; i++) {
              if (this.list[i].id === this.form.id) {
                this.list[i].name = this.form.name
                this.list[i].bugstatuslist = this.form.checklist
              }
            }
            this.$message.success('修改成功')
          } else {
            this.$message.error(resp.data.msg)
          }
        })
      } else {
        addStatusGroup(this.form).then(resp => {
          if (resp.data.code === 0) {
            this.list.push({
              id: resp.data.id,
              name: this.form.name,
              bugstatuslist: this.form.checklist
            })
            this.$message.success('添加成功')
          } else {
            this.$message.error(resp.data.msg)
          }
        })
      }

      this.dialogVisible = false
    }
  }
}
</script>

<style scoped type="text/css">
label {
  padding: 10px;
}
</style>
