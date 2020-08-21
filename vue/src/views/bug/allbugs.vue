<template>
  <div class="app-container">
    <p class="warn-content">
      所有的bug，按道理是不能修改状态的，还不清楚， 选择器的状态是显示所选的状态，永久保存，多页面生效
    </p>
    <div class="filter-container">
      <el-input v-model="listQuery.title" placeholder="标题" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter" />
      <el-select v-model="listQuery.level" placeholder="级别" clearable style="width: 90px" class="filter-item">
        <el-option v-for="(item, index) in levels" :key="index" :label="item" :value="item" />
      </el-select>
      <el-select v-model="listQuery.project" placeholder="项目名" clearable class="filter-item" style="width: 130px">
        <el-option v-for="(item, index) in projectnames" :key="index" :label="item" :value="item" />
      </el-select>
      <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">搜索</el-button>
      <el-dropdown :hide-on-click="false" :show-timeout="100" trigger="click" style="vertical-align: top;">
        <el-button plain>
          状态({{ statuslength }})
          <i class="el-icon-caret-bottom el-icon--right" />
        </el-button>
        <el-dropdown-menu slot="dropdown" class="no-border">
          <el-checkbox-group v-model="listQuery.status" style="padding-left: 15px;" @change="HandleChange">
            <el-checkbox v-for="(status, index) in allstatus" :key="index" :label="status">
              {{ status }}
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
      style="width: 100%;"
    >
      <el-table-column label="id" align="center" width="65">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="日期" width="150px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.date | parseTime('{y}-{m}-{d} {h}:{i}') }}</span>
        </template>
      </el-table-column>

      <el-table-column label="项目名" width="110px" align="center">
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
      <!--<el-table-column width="110px" v-if='showReviewer' align="center" :label="$t('table.reviewer')">-->
      <!--<template slot-scope="scope">-->
      <!--<span style='color:red;'>{{scope.row.reviewer}}</span>-->
      <!--</template>-->
      <!--</el-table-column>-->
      <el-table-column label="级别" width="80px" class-name="status-col">
        <template slot-scope="scope">
          <span>{{ scope.row.level }}</span>
          <!--<svg-icon v-for="n in +scope.row.importance" icon-class="star" class="meta-item__icon" :key="n"></svg-icon>-->
        </template>
      </el-table-column>
      <el-table-column label="重要性" width="80px" class-name="status-col">
        <template slot-scope="scope">
          <span>{{ scope.row.important }}</span>
          <!--<svg-icon v-for="n in +scope.row.importance" icon-class="star" class="meta-item__icon" :key="n"></svg-icon>-->
        </template>
      </el-table-column>

      <el-table-column label="环境" class-name="status-col" width="100">
        <template slot-scope="scope">
          <span>{{ scope.row.env }}</span>
          <!--<el-tag :type="scope.row.status | statusFilter">{{scope.row.env}}</el-tag>-->
        </template>
      </el-table-column>
      <el-table-column label="处理人" align="center" width="300">
        <template slot-scope="scope">
          <span>{{ scope.row.handle }}</span>
          <!--<span v-if="scope.row.handle" class="link-type" @click='handleFetchPv(scope.row.pageviews)'>{{scope.row.pageviews}}</span>-->
          <!--<span v-else>0</span>-->
        </template>
      </el-table-column>
      <el-table-column label="状态" align="center" width="230" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <span>{{ scope.row.status }}</span>
          <!-- <el-select v-model="scope.row.status" style="width: 200px" class="filter-item" placeholder="修改状态" @change="changestatus(scope.row)" /> -->
        </template>
      </el-table-column>
    </el-table>

    <div class="pagination-container">
      <el-pagination
        :current-page="listQuery.page"
        :pager-count="11"
        :page-sizes="[15]"
        :page-size="listQuery.limit"
        :total="total"
        background
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>

  </div>
</template>

<script>
import { changeStatus } from '@/api/bugs'
import { statusFilter, showStatus } from '@/api/status'
import { searchAllBugs } from '@/api/search'
import { getProject, getPermStatus, getStatus, getLevels } from '@/api/get'
import waves from '@/directive/waves' // 水波纹指令
import { parseTime } from '@/utils'

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
  name: 'Allbugs',
  directives: {
    waves
  },
  filters: {
    // statusFilter(status) {
    //   const statusMap = {
    //     published: 'success',
    //     draft: 'info',
    //     deleted: 'danger'
    //   }
    //   return statusMap[status]
    // },
    typeFilter(type) {
      return calendarTypeKeyValue[type]
    }
  },
  data() {
    return {
      users: [],
      tableKey: 0,
      list: null,
      total: 0,
      listLoading: true,
      importanceOptions: [],
      listQuery: {
        page: 1,
        limit: 15,
        level: '',
        project: '',
        title: '',
        status: []
      },
      // importanceOptions: [1, 2, 3, 4, 5],
      calendarTypeOptions,
      sortOptions: [{ label: 'ID Ascending', key: '+id' }, { label: 'ID Descending', key: '-id' }],
      // statusOptions: ['待领取', '待测试'],
      showReviewer: false,
      temp: {
        id: undefined,
        // importance: 1,
        remark: '',
        // timestamp: new Date(),
        // title: '',
        // type: '',
        status: '待领取',
        selectusers: []
        // status: 'published'
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
      pvData: [],
      stop: {
        id: undefined,
        stop: '',
        status: ''
      },
      projectnames: [],
      rules: {
        type: [{ required: true, message: 'type is required', trigger: 'change' }],
        timestamp: [{ type: 'date', required: true, message: 'timestamp is required', trigger: 'change' }],
        title: [{ required: true, message: 'title is required', trigger: 'blur' }]
      },
      downloadLoading: false,
      filterstatus: [],
      levels: [],
      checkstatus: [],
      statuslength: 0,
      allstatus: []
    }
  },
  activated() {
    this.getstatus()
    this.getmystatus()
    this.getlevels()
    this.getprojectname()
  },
  created() {
    this.getstatus()
    this.getmystatus()
    this.getlevels()
    this.getList()
    this.getprojectname()
    // this.gettaskstatus()
  },
  methods: {
    getlevels() {
      getLevels().then(resp => {
        if (resp.data.code === 0) {
          this.levels = resp.data.levels
        } else {
          this.$message.error(resp.data.msg)
        }
      })
    },
    HandleChange() {
      console.log(this.listQuery.status)
      const data = {
        checkstatus: this.checkstatus
      }
      statusFilter(data).then(resp => {
        console.log(resp.data)
        if (resp.data.code === 0) {
          this.listLoading = true
          searchAllBugs(this.listQuery).then(resp => {
            if (resp.data.code === 0) {
              this.list = resp.data.articlelist
              this.total = resp.data.total
              this.listQuery.page = resp.data.page
            }
          })
          this.listLoading = false
        } else {
          this.$message.error(resp.data.msg)
        }
      })
    },
    getprojectname() {
      // const now = new Date().getTime()
      getProject().then(resp => {
        if (resp.data.code === 0) {
          this.projectnames = resp.data.projectlist
        } else {
          this.$message.error(resp.data.msg)
        }
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
    getList() {
      this.listLoading = true
      searchAllBugs(this.listQuery).then(resp => {
        if (resp.data.code === 0) {
          this.total = resp.data.total
          this.list = resp.data.articlelist
        } else {
          this.$message.error(resp.data.msg)
        }
        this.listLoading = false
      })
    },
    handleFilter() {
      this.listLoading = true
      this.listQuery.page = 1
      if (this.listQuery.level === '' && this.listQuery.title === '' && this.listQuery.status === '' && this.listQuery.project === '') {
        this.getList()
      } else {
        searchAllBugs(this.listQuery).then(resp => {
          if (resp.data.code === 0) {
            this.total = resp.data.total
            this.list = resp.data.articlelist
          } else {
            this.$message.error(resp.data.msg)
          }
          this.listLoading = false
        })
      }
    },
    handleSizeChange(val) {
      this.listQuery.limit = val
      this.getList()
    },
    handleCurrentChange(val) {
      this.listQuery.page = val
      this.getList()
    },
    handleCreate() {
      this.resetTemp()
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
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
    formatJson(filterVal, jsonData) {
      return jsonData.map(v => filterVal.map(j => {
        if (j === 'timestamp') {
          return parseTime(v[j])
        } else {
          return v[j]
        }
      }))
    },
    getmystatus() {
      // 允许改变的状态
      getPermStatus().then(resp => {
        if (resp.data.code === 0) {
          // this.checkstatus = resp.data.statuslist
        }
      })
      // 过滤的状态
      showStatus().then(resp => {
        if (resp.data.code === 0) {
          this.checkstatus = resp.data.checkstatus
          this.statuslength = this.checkstatus.length
        } else {
          this.$message.error(resp.data.msg)
        }
      })
    },
    getstatus() {
      getStatus().then(resp => {
        console.log(resp.data)
        if (resp.data.code === 0) {
          this.allstatus = resp.data.statuslist
        } else {
          this.$message.error(resp.data.msg)
        }
      })
    }
  }
}
</script>
