<template>
  <div class="group">
    <p class="warn-content">
      共享文件和文档权限控制组
    </p>
    <el-table
      :data="list"
      height="250"
      style="width: 100%;"
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
        label="组名"
        width="180"
      >
        <template slot-scope="scope">
          <span style="margin-left: 10px">{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="成员"
        width="500"
      >
        <template slot-scope="scope">
          <span style="margin-left: 10px">{{ scope.row.users }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作">
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
    <el-button style="margin-top: 10px;margin-left: 10px" type="success" @click="handleAdd">添加组</el-button>
    <el-dialog :visible.sync="dialogFormVisible" title="平台管理">
      <el-form :model="form">
        <el-form-item label="组名">
          <el-input v-model="form.name" auto-complete="off" />
        </el-form-item>
        <el-form-item label="用户">
          <el-select v-model="form.users" multiple placeholder="请选择">
            <el-option
              v-for="(user, index) in users"
              :key="index"
              :label="user"
              :value="user"
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
import { getGroup, addGroup, updateGroup, delGroup } from '@/api/group'
import { getUsers } from '@/api/get'
export default {
  name: 'Group',
  data() {
    return {
      dialogFormVisible: false,
      list: null,
      form: {
        id: -1,
        name: '',
        users: []
      },
      users: []
    }
  },
  mounted() {
    this.getgroup()
    this.getuser()
  },
  methods: {
    getuser() {
      getUsers().then(resp => {
        if (resp.data.userlist === null) {
          this.users = []
        } else {
          this.users = resp.data.users
        }
      })
    },
    getgroup() {
      getGroup().then(resp => {
        if (resp.data.code === 0) {
          this.list = resp.data.grouplist
        }
      })
    },
    handleAdd() {
      this.form = {
        id: -1,
        name: '',
        users: []
      }
      this.dialogFormVisible = true
    },
    confirm() {
      if (this.form.id > 0) {
        updateGroup(this.form).then(resp => {
          if (resp.data.code === 0) {
            const l = this.list.length
            for (let i = 0; i < l; i++) {
              if (this.list[i].id === this.form.id) {
                this.list[i].name = this.form.name
                this.list[i].users = this.form.users
              }
            }
            this.$message.success('修改成功')
            return
          } else {
            this.$message.error(resp.data.msg)
          }
        })
      } else {
        addGroup(this.form).then(resp => {
          if (resp.data.code === 0) {
            this.list.push({
              id: resp.data.id,
              name: this.form.name,
              users: this.form.users
            })
            this.$message.success('添加用户组成功')
          } else {
            this.$message.error(resp.data.msg)
          }
        })
      }
      this.dialogFormVisible = false
    },
    cancel() {
      this.form = {
        name: '',
        users: []
      }
      this.dialogFormVisible = false
    },
    handleUpdate(row) {
      this.dialogFormVisible = true
      this.form.id = row.id
      this.form.users = row.users
      this.form.name = row.name
    },
    handleDelete(id) {
      this.$confirm('此操作将关闭bug, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        delGroup(id).then(resp => {
          if (resp.data.code === 23) {
            this.$message.error('用户组在使用')
            return
          }
          if (resp.data.code === 0) {
            const l = this.list.length
            for (let i = 0; i < l; i++) {
              if (this.list[i].id === id) {
                this.list.splice(i, 1)
                break
              }
            }
            this.$message.success('删除成功')
            return
          } else {
            this.$message.error(resp.data.msg)
          }
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
      })
    }
  }
}
</script>

