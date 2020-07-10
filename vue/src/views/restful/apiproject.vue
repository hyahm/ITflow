<template>
  <div>
    <p class="warn-content">
      添加接口的项目
    </p>
    <el-table
      :data="list"
      fit
      border
      highlight-current-row
      style="width: 100%"
    >
      <el-table-column
        label="Id"
        width="60"
      >
        <template slot-scope="scope">
          <span style="margin-left: 10px">{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="项目名"
        width="400"
      >
        <template slot-scope="scope">
          <a :href="'/restful/apilist?id=' + scope.row.id" target="_blank" style="text-decoration:underline;color: blue">{{ scope.row.name }}</a>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="150">
        <template slot-scope="scope">
          <el-button
            size="mini"
            @click="update(scope.row)"
          >修改
          </el-button>
          <el-button
            size="mini"
            type="danger"
            @click="handleDelete(scope.row.id)"
          >删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-button style="margin: 20px" type="success" size="mini" @click="add">添加项目名</el-button>
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
import { restList, restUpdate, restDel, restAdd } from '@/api/restful'
import { getUsers } from '@/api/get'
import { getGroup } from '@/api/group'

export default {
  name: 'ProjectList',
  data() {
    return {
      isReadUser: true,
      // notReadUser: false,
      isWriteUser: true,
      // notWriteUser: false,
      users: [],
      realname: 'admin',
      groups: [],
      perm: true,
      readwrite: '',
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
      list: []
    }
  },
  created() {
    this.getrestname()

    this.getuserlist()
    this.getgrouplist()
  },
  methods: {
    getuserlist() {
      getUsers().then(resp => {
        if (resp.data.code === 0) {
          if (resp.data.users !== null) {
            this.users = resp.data.users
          }
        } else {
          this.$message.error(resp.data.message)
        }
        this.readlist = this.users
        this.rdwrlist = this.users
      })
    },
    getgrouplist() {
      getGroup().then(resp => {
        if (resp.data.code === 0) {
          if (resp.data.grouplist !== null) {
            this.groups = resp.data.grouplist
          }
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    handleAuth() {
      if (this.form.auth) {
        if (this.form.readname === '') {
          this.form.readuser = true
          this.form.readname = this.realname
        }
        if (this.form.editname === '') {
          this.form.edituser = true
          this.form.editname = this.realname
        }
      }
    },
    handleRead() {
      if (this.form.readuser) {
        this.readlist = this.users
        this.form.readname = this.realname
      } else {
        this.readlist = this.groups
        this.form.readname = this.readlist[0]
      }
    },
    handleRdWr() {
      if (this.form.edituser) {
        this.rdwrlist = this.users
        this.form.editname = this.realname
      } else {
        this.rdwrlist = this.groups
        this.form.editname = this.rdwrlist[0]
      }
    },
    getrestname() {
      restList().then(resp => {
        if (resp.data.code === 0) {
          if (resp.data.list !== null) {
            this.list = resp.data.list
          }
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    update(row) {
      this.form = JSON.parse(JSON.stringify(row))
      this.form.owner = this.realname
      this.dialogFormVisible = true
    },
    handleDelete(id) {
      this.$confirm('此操作将删除此项目, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        restDel(id).then(resp => {
          if (resp.data.code === 0) {
            for (let i = 0; i < this.list.length; i++) {
              if (this.list[i].id === id) {
                this.list.splice(i, 1)
              }
            }
          }
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
      })
    },
    confirm() {
      if (this.form.name === '') {
        this.$message({
          type: 'info',
          message: '名称不能为空'
        })
      }
      if (this.form.id > 0) {
        restUpdate(this.form).then(resp => {
          if (resp.data.code === 0) {
            for (let i = 0; i < this.list.length; i++) {
              if (this.list[i].id === this.form.id) {
                this.list.splice(i, 1, this.form)
              }
            }
            this.$message.success('添加成功')
          } else {
            this.$message.error('添加失败')
          }
        })
      } else {
        restAdd(this.form).then(resp => {
          if (resp.data.code === 0) {
            const data = this.form
            data.id = resp.data.id
            this.list.push(data)
            this.$message.success('添加成功')
          } else {
            this.$message.error('添加失败')
          }
        })
      }

      this.dialogFormVisible = false
    },
    add() {
      this.form = {
        owner: this.realname,
        readuser: true,
        edituser: true,
        readname: this.realname,
        editname: this.realname,
        auth: false,
        name: '',
        id: -1
      }
      this.dialogFormVisible = true
    },
    cancel() {
      this.dialogFormVisible = false
    }
  }
}
</script>

<style scoped>

</style>
