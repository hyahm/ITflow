<template>
  <div class="app-container">
    <p class="warn-content">
      我接到的bug，可以转交和改变状态， 选择器的状态是显示所选的状态，永久保存，多页面生效
    </p>
    <div class="filter-container">
      <el-input v-model="listQuery.title" placeholder="标题" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter" />
      <el-select v-model="listQuery.level" placeholder="优先级" clearable style="width: 90px" class="filter-item">
        <el-option v-for="(item, index) in levels" :key="index" :label="item" :value="item" />
      </el-select>
      <el-select v-model="listQuery.project" placeholder="项目" clearable class="filter-item" style="width: 130px">
        <el-option v-for="(item, index) in projectnames" :key="index" :label="item" :value="item" />
      </el-select>
      <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">搜索</el-button>
      <el-dropdown :hide-on-click="false" :show-timeout="100" trigger="click" style="vertical-align: top;">
        <el-button plain>
          状态({{ statuslength }})
          <i class="el-icon-caret-bottom el-icon--right" />
        </el-button>
        <el-dropdown-menu slot="dropdown" class="no-border">
          <el-checkbox-group v-model="checkstatus" style="padding-left: 15px;" @change="HandleChange">
            <el-checkbox v-for="(item, index) in platformsOptions" :key="index" :label="item">
              {{ item }}
            </el-checkbox>
          </el-checkbox-group>
        </el-dropdown-menu>
      </el-dropdown>
    </div>

    <el-table
      :key="tableKey"
      v-loading="listLoading"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%;min-height:350px;"
    >
      <el-table-column label="id" align="center" width="65">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="时间" width="150px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.date | parseTime('{y}-{m}-{d} {h}:{i}') }}</span>
        </template>
      </el-table-column>

      <el-table-column label="项目" width="110px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.projectname }}</span>
        </template>
      </el-table-column>

      <el-table-column label="标题" min-width="150px" align="center">
        <template slot-scope="scope">
          <router-link :to="'/showbug/'+scope.row.id" class="link-type">
            <span class="link-type">{{ scope.row.title }}</span>
          </router-link>
          <!--<el-tag>{{scope.row.type | typeFilter}}</el-tag>-->
        </template>
      </el-table-column>
      <el-table-column label="作者" width="110px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.author }}</span>
        </template>
      </el-table-column>
      <el-table-column label="优先级" width="80px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.level }}</span>
          <!--<svg-icon v-for="n in +scope.row.importance" icon-class="star" class="meta-item__icon" :key="n"></svg-icon>-->
        </template>
      </el-table-column>
      <el-table-column label="重要性" width="80px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.important }}</span>
          <!--<svg-icon v-for="n in +scope.row.importance" icon-class="star" class="meta-item__icon" :key="n"></svg-icon>-->
        </template>
      </el-table-column>
      <el-table-column label="状态" class-name="status-col" width="120">
        <template slot-scope="scope">
          <el-select v-model="scope.row.status" class="filter-item" placeholder="修改状态" @change="changestatus(scope.row)">
            <el-option v-for="(item, index) in statuslist" :key="index" :label="item" :value="item" />
          </el-select>
          <!--<el-tag :type="scope.row.status | statusFilter">{{scope.row.status}}</el-tag>-->
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="130" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button type="primary" size="mini" @click="handlePass(scope.row)">转交</el-button>
        </template>
      </el-table-column>
    </el-table>

    <div class="pagination-container">
      <el-pagination :current-page="listQuery.page" :page-sizes="[10]" :page-size="listQuery.limit" :total="total" background layout="total, sizes, prev, pager, next, jumper" @size-change="handleSizeChange" @current-change="handleCurrentChange" />
    </div>

    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
      <el-form ref="dataForm" :rules="rules" :model="temp" label-position="left" label-width="70px" style="width: 400px; margin-left:50px;">
        <el-form-item label="状态:">
          <el-select v-model="temp.status" class="filter-item" placeholder="Please select">
            <el-option v-for="(item, index) in platformsOptions" :key="index" :label="item" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item style="margin-bottom: 40px;" label="任务给：">
          <el-select
            v-model="temp.selectusers"
            filterable
            multiple
            allow-create
            default-first-option
            placeholder="请选择指定的用户"
          >
            <el-option
              v-for="(item, index) in users"
              :key="index"
              :label="item"
              :value="item"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="说明:">
          <el-input v-model="temp.remark" :autosize="{ minRows: 2, maxRows: 4}" type="textarea" placeholder="Please input" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <!--<el-button v-if="dialogStatus=='create'" type="primary" @click="createData">{{$t('table.confirm')}}</el-button>-->
        <el-button type="primary" @click="updateData">确认</el-button>
      </div>
    </el-dialog>

  </div>
</template>

<script>
import { getUsers, getPermStatus, getProject, getStatus, getShowStatus, getLevels } from '@/api/get'
import { passBug, changeStatus } from '@/api/bugs'
import { searchMyTasks } from '@/api/search'
import { statusFilter } from '@/api/status'
import waves from '@/directive/waves' // 水波纹指令
import { parseTime } from '@/utils'
// import { getProject } from '@/utils/auth'

const calendarTypeOptions = [
  { key: 'CN', display_name: 'China' },
  { key: 'US', display_name: 'USA' },
  { key: 'JP', display_name: 'Japan' },
  { key: 'EU', display_name: 'Eurozone' }
]

// arr to obj ,such as { CN : "China", US : "USA" }
const calendarTypeKeyValue = calendarTypeOptions.reduce((acc, cur) => {
  acc[cur.key] = cur.display_name
  return acc
}, {})

export default {
  name: 'Mytask',
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
    },
    typeFilter(type) {
      return calendarTypeKeyValue[type]
    }
  },
  data() {
    return {
      users: [],
      tableKey: 0,
      list: null,
      total: null,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 10,
        importance: undefined,
        title: undefined,
        type: undefined,
        sort: '+id',
        status: []
      },
      calendarTypeOptions,
      sortOptions: [{ label: 'ID Ascending', key: '+id' }, { label: 'ID Descending', key: '-id' }],
      showReviewer: false,
      temp: {
        id: undefined,
        remark: '',
        status: '新建',
        selectusers: ''
      },
      changeaction: {
        id: 0,
        status: '',
        action: ''
      },
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: 'Edit',
        create: 'Create'
      },
      dialogPvVisible: false,
      platformsOptions: [],
      statuslist: [],
      pvData: [],
      stop: {
        id: undefined,
        stop: '',
        status: ''
      },
      rules: {
        type: [{ required: true, message: 'type is required', trigger: 'change' }],
        timestamp: [{ type: 'date', required: true, message: 'timestamp is required', trigger: 'change' }],
        title: [{ required: true, message: 'title is required', trigger: 'blur' }]
      },
      downloadLoading: false,
      levels: [],
      projectnames: [],
      checkstatus: [],
      statuslength: 0
    }
  },
  activated() {
    this.getstatus()
    this.getlevels()
    this.handleFilter()
  },
  mounted() {
    this.getspuser()
  },
  created() {
    this.getstatus()
    this.getmystatus()
    this.handleFilter()
    this.getpname()
    // this.gettaskstatus()
  },
  methods: {
    getlevels() {
      getLevels().then(resp => {
        if (resp.data.code === 0) {
          this.levels = resp.data.levels
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    getstatus() {
      getStatus().then(resp => {
        if (resp.data.code === 0) {
          this.platformsOptions = resp.data.statuslist
        } else {
          this.$message.error(resp.data.message)
        }
      })
      getPermStatus().then(resp => {
        if (resp.data.code === 0) {
          this.statuslist = resp.data.statuslist
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    HandleChange() {
      const data = {
        checkstatus: this.checkstatus
      }
      this.statuslength = this.checkstatus.length
      statusFilter(data).then(resp => {
        if (resp.data.code === 0) {
          this.statuslength = this.checkstatus.length
          this.listLoading = true
          this.handleFilter()
          // searchMyTasks(this.listQuery).then(respfilter => {
          //   console.log(respfilter)
          //   if (respfilter.data.code === 0) {
          //     this.list = respfilter.data.articlelist
          //     this.total = respfilter.data.total
          //     this.listQuery.page = respfilter.data.page
          //   }
          // })
          this.listLoading = false
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    getpname() {
      getProject().then(resp => {
        if (resp.data.code === 0) {
          this.projectnames = resp.data.projectlist
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    getmystatus() {
      // 需要显示的状态
      getShowStatus().then(resp => {
        if (resp.data.code === 0) {
          if (resp.data.checkstatus !== null) {
            this.checkstatus = resp.data.checkstatus
            this.statuslength = this.checkstatus.length
          }
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    getspuser() {
      getUsers().then(resp => {
        if (resp.data.code === 0) {
          this.users = resp.data.users
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    handleFilter() {
      this.listLoading = true
      searchMyTasks(this.listQuery).then(resp => {
        this.list = resp.data.articlelist
        this.total = resp.data.total
        this.listQuery.page = resp.data.page
      })
      this.listLoading = false
    },

    handleSizeChange(val) {
      this.listQuery.limit = val
      this.handleFilter()
    },
    handleCurrentChange(val) {
      this.listQuery.page = val
      this.handleFilter()
    },
    handleModifyStatus(row) {
      this.changeaction.id = row.id
      this.changeaction.status = row.status
      this.changeaction.action = row.action
    },
    handleCreate() {
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    handlePass(row) {
      this.temp.id = parseInt(row.id) // copy obj
      this.temp.status = row.status
      this.temp.selectusers = row.handle
      this.dialogFormVisible = true
    },
    updateData() {
      passBug(this.temp).then(resp => {
        if (resp.data.code === 0) {
          const data = resp.data
          this.temp.remark = ''
          this.temp.status = data.status
          this.temp.selectusers = ''
          this.$message({
            message: '操作成功',
            type: 'success'
          })
        } else {
          this.$message.error(resp.data.message)
        }
      })
      this.dialogFormVisible = false
    },
    handleDelete(row) {
      this.$notify({
        title: '成功',
        message: '删除成功',
        type: 'success',
        duration: 2000
      })
      const index = this.list.indexOf(row)
      this.list.splice(index, 1)
    },
    handleDownload() {
      this.downloadLoading = true
      import('@/vendor/Export2Excel').then(excel => {
        const tHeader = ['timestamp', 'title', 'type', 'importance', 'status']
        const filterVal = ['timestamp', 'title', 'type', 'importance', 'status']
        const data = this.formatJson(filterVal, this.list)
        excel.export_json_to_excel({
          header: tHeader,
          data,
          filename: 'table-list'
        })
        this.downloadLoading = false
      })
    },
    changestatus(row) {
      const param = {
        id: row.id,
        status: row.status
      }
      changeStatus(param).then(response => {
        if (response.data.code === 0) {
          this.$notify({
            title: '成功',
            message: '修改成功',
            type: 'success'
          })
        } else {
          this.$notify({
            title: '失败',
            message: '操作失败',
            type: 'error'
          })
        }
      })
    },
    formatJson(filterVal, jsonData) {
      return jsonData.map(v => filterVal.map(j => {
        if (j === 'timestamp') {
          return parseTime(v[j])
        } else {
          return v[j]
        }
      }))
    }
  }
}
</script>
