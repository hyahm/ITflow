<template>
  <div>
    <p class="warn-content">
      选择可以操作的页面组， 操作的角色由开发者决定, 如果查看的权限没有， 那么后面的所有权限都被无视
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

      <!-- <el-table-column label="角色组" width="500" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.rolelist }}</span>
        </template>
      </el-table-column> -->

      <el-table-column label="操作" align="center" width="230" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button type="primary" size="mini" @click="handleEdit(scope.row)">编辑</el-button>
          <el-button type="success" size="mini" @click="handleRemove(scope.row.id)">删除</el-button>
          <!--<el-button type="danger" size="mini" @click="handleRemove(scope.row)">{{ $t('list.remove') }}</el-button>-->
        </template>
      </el-table-column>
    </el-table>
    <el-button style="margin: 20px" type="success" size="mini" @click="handleAdd">添加角色组</el-button>
    <el-dialog
      :visible.sync="dialogVisible"
      :before-close="handleClose"
      :close-on-click-modal="false"
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
        <!-- <el-checkbox-group v-model="perm"> -->
        <div v-for="(role, index) in form.rolelist" :key="index">
          <!-- <label class="name">{{ role.name }}</label> -->
          <!-- <el-checkbox v-model="role.checked" :label="role.name" style="width:150px" @change="changeChecked(role)" /> -->
          <el-checkbox v-model="role.select" style="width:50px" label="select" @change="changeChecked(role)" />
          <el-checkbox v-model="role.add" :disabled="!role.select" style="width:50px" label="add" />
          <el-checkbox v-model="role.update" :disabled="!role.select" style="width:50px" label="update" />
          <el-checkbox v-model="role.remove" :disabled="!role.select" style="width:50px" label="remove" />
          <label style="padding-left: 100px;width:50px">{{ role.info }}</label>
        </div>
        <!-- </el-checkbox-group> -->
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

import { roleList, addRole, editRole, removeRole, getPermTemplate } from '@/api/role'
import { deepClone } from '@/utils'
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
      },
      templateperm: []
    }
  },

  created() {
    // this.getroles()
    this.getlist()
    this.getTemplate()
  },
  methods: {

    changeChecked(row) {
      if (!row.select) {
        row.add = false
        row.remove = false
        row.update = false
      }
    },
    getTemplate() {
      // 获取模板
      getPermTemplate().then(resp => {
        this.templateperm = resp.data
      })
    },
    handleEdit(row) {
      this.form.id = row.id
      this.form.name = row.name
      this.form.rolelist = row.rolelist
      this.dialogVisible = true
    },
    getlist() {
      roleList().then(resp => {
        if (resp.data.code === 0) {
          this.list = resp.data.rolelist
          console.log(this.list)
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    handleAdd() {
      this.form.id = -1
      this.form.name = ''
      this.form.rolelist = deepClone(this.templateperm)
      this.dialogVisible = true
    },
    handleRemove(id) {
      this.$confirm('确认关闭？')
        .then(_ => {
          removeRole(id).then(resp => {
            if (resp === undefined) {
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
            } else {
              this.$message.error(resp.data.message)
            }
          })
        })
        .catch(_ => {})
    },
    handleClose() {
      this.dialogVisible = false
    },
    HandlerAddGroup() {
      if (this.form.name.length < 1) {
        this.$message.error('name no be need')
      }
      // return
      if (this.form.id > 0) {
        editRole(this.form).then(resp => {
          // 成功后赋值到源数据
          if (resp.data.code === 0) {
            const l = this.list.length
            for (let i = 0; i < l; i++) {
              if (this.list[i].id === this.form.id) {
                this.list[i].name = this.form.name
                this.list[i].rolelist = this.form.rolelist
              }
            }
            this.$message.success('修改成功')
          } else {
            this.$message.error(resp.data.message)
          }
        })
      } else {
        addRole(this.form).then(resp => {
          if (resp.data.code === 0) {
            this.list.push({
              id: resp.data.id,
              name: this.form.name,
              rolelist: this.form.rolelist
            })
            this.$message.success('添加成功')
          } else {
            this.$message.error(resp.data.message)
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
.form-container > .name {
  padding-right: 30px;
  width:250px !important;
}
</style>

<style  type="text/css">
.form-container > .name {
  padding-right: 30px;
  width:250px !important;
}
</style>
