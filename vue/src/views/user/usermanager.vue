<template>
  <div class="app-container">
    <p class="warn-content">
      管理员能修改所有信息，管理者只能修改下一级的信息，有些麻烦，此功能最后完善
    </p>
    <div class="filter-container">
      <!--<el-input @keyup.enter.native="handleFilter" style="width: 200px;" class="filter-item" :placeholder="$t('table.title')" v-model="listQuery.title">-->
      <!--</el-input>-->
      <!--<el-select clearable style="width: 90px" class="filter-item" v-model="listQuery.importance" :placeholder="$t('table.importance')">-->
      <!--<el-option v-for="item in importanceOptions" :key="item" :label="item" :value="item">-->
      <!--</el-option>-->
      <!--</el-select>-->
      <!--<el-select clearable class="filter-item" style="width: 130px" v-model="listQuery.type" :placeholder="$t('table.type')">-->
      <!--<el-option v-for="item in  calendarTypeOptions" :key="item.key" :label="item.display_name+'('+item.key+')'" :value="item.key">-->
      <!--</el-option>-->
      <!--</el-select>-->
      <!--<el-select @change='handleFilter' style="width: 140px" class="filter-item" v-model="listQuery.sort">-->
      <!--<el-option v-for="item in sortOptions" :key="item.key" :label="item.label" :value="item.key">-->
      <!--</el-option>-->
      <!--</el-select>-->
      <!--<el-button class="filter-item" type="primary" v-waves icon="el-icon-search" @click="handleFilter">{{$t('table.search')}}</el-button>-->
      <!--<el-button class="filter-item" style="margin-left: 10px;" @click="handleCreate" type="primary" icon="el-icon-edit">{{$t('table.add')}}</el-button>-->
      <!--<el-button class="filter-item" type="primary" :loading="downloadLoading" v-waves icon="el-icon-download" @click="handleDownload">{{$t('table.export')}}</el-button>-->
      <!--<el-checkbox class="filter-item" style='margin-left:15px;' @change='tableKey=tableKey+1' v-model="showReviewer">{{$t('table.reviewer')}}</el-checkbox>-->
    </div>

    <el-table
      :key="tableKey"
      v-loading="listLoading"
      :data="userlist"
      border
      fit
      highlight-current-row
      style="width: 100%;"
    >
      <el-table-column label="id" align="center" width="65">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="日期" width="150px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.createtime | parseTime('{y}-{m}-{d} {h}:{i}') }}</span>
        </template>
      </el-table-column>

      <el-table-column label="真实姓名" width="110px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.realname }}</span>
        </template>
      </el-table-column>
      <el-table-column label="昵称" width="110px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.nickname }}</span>
        </template>
      </el-table-column>

      <el-table-column label="角色组" width="100px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.rolegroup }}</span>
          <!--<svg-icon v-for="n in +scope.row.importance" icon-class="star" class="meta-item__icon" :key="n"></svg-icon>-->
        </template>
      </el-table-column>
      <el-table-column label="状态组" width="110px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.statusgroup }}</span>
        </template>
      </el-table-column>

      <el-table-column label="职位" width="110px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.position }}</span>
        </template>
      </el-table-column>

      <el-table-column label="邮箱" class-name="status-col" width="200">
        <template slot-scope="scope">
          <span>{{ scope.row.email }}</span>
          <!--<el-tag :type="scope.row.status | statusFilter">{{scope.row.status}}</el-tag>-->
        </template>
      </el-table-column>
      <el-table-column label="状态" class-name="status-col" width="60">
        <template slot-scope="scope">
          <span v-if="scope.row.disable==0">启用</span>
          <span v-else>禁用</span>
          <!--<el-tag :type="scope.row.status | statusFilter">{{scope.row.status}}</el-tag>-->
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="300" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button type="primary" size="mini" @click="handleResetPwd(scope.row)">修改密码</el-button>
          <el-button size="mini" type="danger" @click="handlePermission(scope.row)">更新权限
          </el-button>
          <el-button size="mini" type="danger" @click="handleRemove(scope.row)">删除
          </el-button>
          <el-button v-if="scope.row.disable==1" size="mini" type="danger" @click="handleDisable(scope.row)">启用
          </el-button>
          <el-button v-else size="mini" type="danger" @click="handleDisable(scope.row)">禁用
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog
      :visible.sync="dialogVisible"
      :before-close="handleClose"
      title="提示"
      width="30%"
    >
      <el-form ref="postForm" class="form-container" style="padding: 20px" />
      <!--<el-button type="success" round @click="HandlerAddGroup">添加部门</el-button>-->
      <el-form ref="postForm" class="form-container" style="padding: 20px">
        <el-form-item label="昵称">
          <el-input v-model="form.nickname" />
        </el-form-item>
        <el-form-item label="真实姓名">
          <el-input v-model="form.realname" />
        </el-form-item>
        <el-form-item label="邮箱">
          <el-input v-model="form.email" />
        </el-form-item>
        <el-form-item label="状态组">
          <el-select v-model="form.statusgroup" placeholder="Select">
            <el-option
              v-for="(role, index) in statusgrouplist"
              :key="index"
              :label="role"
              :value="role"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="角色组">
          <el-select v-model="form.rolegroup" placeholder="Select">
            <el-option
              v-for="(role, index) in rolegrouplist"
              :key="index"
              :label="role"
              :value="role"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="职位：">
          <el-select v-model="form.position" placeholder="Select">
            <el-option
              v-for="(role, index) in positionlist"
              :key="index"
              :label="role"
              :value="role"
            />
          </el-select>
        </el-form-item>
        <!--<el-button type="success" round @click="HandlerAddGroup">添加部门</el-button>-->
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="cancel">取 消</el-button>
        <el-button type="primary" @click="HandlerUpdateRoles">确 定</el-button>
      </span>
    </el-dialog>

  </div>
</template>

<script>
import { userList, resetPwd, updateUser, userRemove, userDisable } from '@/api/user'
import { getRoleGroup } from '@/api/rolegroup'
import { getStatusGroupName } from '@/api/statusgroup'
import { getPositions } from '@/api/position'
import waves from '@/directive/waves' // 水波纹指令

export default {
  name: 'Usermanager',
  directives: {
    waves
  },
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'info',
        deleted: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      uid: -1,
      rolelist: [],
      dialogVisible: false,
      rolegrouplist: [],
      statusgrouplist: [],
      positionlist: [],
      // users: [],
      tableKey: 0,
      userlist: [],
      admin: false,
      form: {
        id: -1,
        name: ''
      },
      statusgroup: [],
      // total: null,
      listLoading: false,
      sortOptions: [{ label: 'ID Ascending', key: '+id' }, { label: 'ID Descending', key: '-id' }]
    }
  },
  activated() {
    this.getuserList()
  },
  created() {
    this.getuserList()
    this.getroles()
    this.getgrouplist()
  },
  methods: {
    getgrouplist() {
      getRoleGroup().then(resp => {
        if (resp.data.code === 0) {
          this.rolegrouplist = resp.data.roles
        } else {
          this.$message.error(resp.data.msg)
        }
      })
      getStatusGroupName().then(resp => {
        if (resp.data.code === 0) {
          this.statusgrouplist = resp.data.statuslist
        } else {
          this.$message.error(resp.data.msg)
        }
      })
      getPositions().then(resp => {
        if (resp.data.code === 0) {
          this.positionlist = resp.data.positions
        } else {
          this.$message.error(resp.data.msg)
        }
      })
    },
    getroles() {
      getRoleGroup().then(resp => {
        if (resp.data.code === 0) {
          this.rolelist = resp.data.roles
        } else {
          this.$message.error(resp.data.msg)
        }
      })
    },
    cancel() {
      this.dialogVisible = false
    },
    HandlerUpdateRoles() {
      updateUser(this.form).then(resp => {
        if (resp.data.code === 0) {
          const l = this.userlist.length
          for (let i = 0; i < l; i++) {
            if (this.userlist[i].id === this.form.id) {
              this.userlist[i].role = this.form.name
            }
          }
          this.$message.success('修改成功')
          return
        } else {
          this.$message.error(resp.data.msg)
        }
      })
      this.dialogVisible = false
    },
    handleClose() {
      this.dialogVisible = false
    },
    getuserList() {
      userList().then(resp => {
        if (resp.data.code === 0) {
          this.userlist = resp.data.userlist
        } else {
          this.$message.error(resp.data.msg)
        }
      }).catch(err => {
        console.log(err)
      })
    },
    handlePermission(row) {
      this.form = row
      this.dialogVisible = true
    },
    handleRemove(row) {
      this.$confirm('此操作将关闭bug, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        console.log(3333)
        console.log(row.id)
        userRemove(row.id).then(resp => {
          if (resp.data.code === 0) {
            const l = this.userlist.length
            for (let i = 0; i < l; i++) {
              if (this.userlist[i].id === row.id) {
                this.userlist.splice(i, 1)
              }
            }
            this.$message.warning('删除成功')
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
    },
    handleDisable(row) {
      userDisable(row.id).then(resp => {
        if (resp.data.code === 0) {
          const l = this.userlist.length
          for (let i = 0; i < l; i++) {
            if (this.userlist[i].id === row.id) {
              this.userlist[i].disable = Math.abs(this.userlist[i].disable - 1)
              break
            }
          }
        } else {
          this.$message.error(resp.data.msg)
        }
      })
    },
    handleResetPwd(row) {
      this.$prompt('请输入密码', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消'
      }).then(({ value }) => {
        const data = {
          id: row.id,
          newpassword: value
        }
        resetPwd(data).then(resp => {
          if (resp.data.code === 0) {
            this.$message({
              type: 'success',
              message: '你的密码是: ' + value
            })
            return
          }
          this.$message({
            message: '密码重置失败',
            type: 'error'
          })
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '取消输入'
        })
      })
    }
  }
}
</script>
