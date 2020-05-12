<template>
  <div>
    <p class="warn-content">
      选择可以操作的页面组， 操作的角色由开发者决定
    </p>
    <el-table v-loading="listLoading" :data="list" border fit highlight-current-row style="width: 100%;padding: 10px">

      <el-table-column label="id" align="center" width="50">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>

      <el-table-column label="角色组" width="100px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.name }}</span>
        </template>
      </el-table-column>

      <el-table-column label="角色组" width="500" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.rolelist }}</span>
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
    <el-button :close-on-click-modal="false" style="margin: 20px" type="success" size="mini" @click="handleAdd">添加角色组</el-button>
    <el-dialog
      :visible.sync="dialogVisible"
      :before-close="handleClose"
      title="提示"
      width="60%"
    >
      <el-form ref="postForm" class="form-container">
        <el-form-item prop="title" label="角色组名:">
          <el-input
            v-model="form.name"
            :maxlength="100"
            placeholder="请角色组名"
            clearable
          />
        </el-form-item>
        <el-checkbox-group v-model="form.rolelist">
          <div v-for="(role, index) in roles" :key="index">
            <el-checkbox :label="role" />
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
import { roleList, addRole, editRole, removeRole, getRoles } from '@/api/role'
export default {
  name: 'RoleGroup',
  data() {
    return {
      roles: [],
      dialogVisible: false,
      listLoading: false,
      list: [],
      form: {
        id: -1,
        name: '',
        rolelist: []
      }

    }
  },
  created() {
    this.getroles()
    this.getlist()
  },
  methods: {
    handleEdit(row) {
      this.form = row
      this.dialogVisible = true
    },
    getlist() {
      roleList().then(resp => {
        if (resp.data.code === 0) {
          if (resp.data.datalist === null) {
            this.list = []
          } else {
            this.list = resp.data.datalist
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
        rolelist: []
      }
      this.dialogVisible = true
    },
    handleRemove(id) {
      this.$confirm('确认关闭？')
        .then(_ => {
          removeRole(id).then(resp => {
            if (resp.data.code === 0) {
              const l = this.list.length
              for (let i = 0; i < l; i++) {
                if (this.list[i].id === id) {
                  this.list.splice(i, 1)
                }
              }
              this.$message.success('删除成功')
              return
            } else {
              this.$message.error(resp.data.msg)
            }
          })
        })
        .catch(_ => {})
    },
    getroles() {
      getRoles().then(resp => {
        if (resp.data.code === 0) {
          this.roles = resp.data.roles
        } else {
          this.$message.error(resp.data.msg)
        }
      })
    },
    handleClose() {
      this.dialogVisible = false
    },
    HandlerAddGroup() {
      if (this.form.name.length < 1) {
        this.$message.error('name no be need')
      }
      if (this.form.id > 0) {
        console.log(this.form)
        editRole(this.form).then(resp => {
          if (resp.data.code === 0) {
            const l = this.list.length
            for (let i = 0; i < l; i++) {
              if (this.list[i].id === this.id) {
                this.list[i].name = this.name
                this.list[i].rolelist = this.rolelist
              }
            }
            this.$message.success('修改成功')
          } else {
            this.$message.error(resp.data.msg)
          }
        })
      } else {
        console.log(this.form)
        addRole(this.form).then(resp => {
          if (resp.data.code === 0) {
            this.list.push({
              id: resp.data.id,
              name: this.form.name,
              rolelist: this.form.rolelist
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
