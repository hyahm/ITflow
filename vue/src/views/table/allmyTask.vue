<template>
  <div class="app-container">
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
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%;min-height:350px;"
    >
      <el-table-column align="center" label="id" width="65">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column width="150px" align="center" label="时间">
        <template slot-scope="scope">
          <span>{{ scope.row.date | parseTime('{y}-{m}-{d} {h}:{i}') }}</span>
        </template>
      </el-table-column>

      <el-table-column width="110px" align="center" label="项目">
        <template slot-scope="scope">
          <span>{{ scope.row.projectname }}</span>
        </template>
      </el-table-column>

      <el-table-column min-width="150px" align="center" label="标题">
        <template slot-scope="scope">
          <router-link class="link-type" :to="'/components/back-to-top/'+scope.row.id">
            <span class="link-type">{{ scope.row.title }}</span>
          </router-link>
          <!--<el-tag>{{scope.row.type | typeFilter}}</el-tag>-->
        </template>
      </el-table-column>
      <el-table-column width="110px" align="center" label="作者">
        <template slot-scope="scope">
          <span>{{ scope.row.author }}</span>
        </template>
      </el-table-column>
      <el-table-column width="80px" align="center" label="优先级">
        <template slot-scope="scope">
          <span>{{ scope.row.level }}</span>
        </template>
      </el-table-column>
      <el-table-column width="80px" align="center" label="重要性">
        <template slot-scope="scope">
          <span>{{ scope.row.importance }}</span>
        </template>
      </el-table-column>
      <el-table-column class-name="status-col" label="状态" width="100">
        <template slot-scope="scope">
          <span>{{ scope.row.status }}</span>
        </template>
      </el-table-column>
    </el-table>

    <div class="pagination-container">
      <el-pagination background :current-page="listQuery.page" :page-sizes="[10,20,30, 50]" :page-size="listQuery.limit" layout="total, sizes, prev, pager, next, jumper" :total="total" @size-change="handleSizeChange" @current-change="handleCurrentChange" />
    </div>

  </div>
</template>

<script>
import { getlist } from '@/api/notcomplete'
import waves from '@/directive/waves' // 水波纹指令
import { parseTime } from '@/utils'
import { getProject } from '@/utils/auth'

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
  name: 'Complextable',
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
        limit: 5,
        importance: undefined,
        title: undefined,
        type: undefined,
        sort: '+id'
      },
      // importanceOptions: [1, 2, 3, 4, 5],
      calendarTypeOptions,
      sortOptions: [{ label: 'ID Ascending', key: '+id' }, { label: 'ID Descending', key: '-id' }],
      statusOptions: ['待领取', '待测试'],
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
      rules: {
        type: [{ required: true, message: 'type is required', trigger: 'change' }],
        timestamp: [{ type: 'date', required: true, message: 'timestamp is required', trigger: 'change' }],
        title: [{ required: true, message: 'title is required', trigger: 'blur' }]
      },
      downloadLoading: false
    }
  },
  activated() {
    this.getList()
  },
  created() {
    this.getList()
    // this.gettaskstatus()
  },
  methods: {
    getList() {
      this.listLoading = true
      getlist(getProject()).then(response => {
        this.list = response.data
        this.listLoading = false
      })
    },
    handleFilter() {
      this.listQuery.page = 1
      this.getList()
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
    }
  }
}
</script>
