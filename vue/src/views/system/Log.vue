<template>
  <div id="log" style="padding-left: 20px">
    <p class="warn-content">
      操作日志
    </p>
    <div class="filter-container">
      <div style="margin-left: 10px">
        <el-select v-model="listQuery.classify" placeholder="分类" clearable class="filter-item">
          <el-option v-for="(item, index) in classifys" :key="index" :label="item" :value="item" />
        </el-select>

        <el-date-picker
          v-model="value2"
          type="daterange"
          value-format="timestamp"
          unlink-panels
          range-separator="至"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
          @change="changeDate"
        />

        <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">搜索</el-button>

      </div>

    </div>
    <el-table
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

      <el-table-column label="日期" width="150px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.exectime | parseTime('{y}-{m}-{d} {h}:{i}') }}</span>
        </template>
      </el-table-column>

      <el-table-column label="分类" width="150" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.classify }}</span>
        </template>
      </el-table-column>
      <el-table-column label="ip" width="150" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.ip }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作者" width="150" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.username }}</span>
        </template>
      </el-table-column>
      <!--<el-table-column min-width="150px" align="center" :label="$t('table.title')">-->
      <!--<template slot-scope="scope">-->
      <!--<router-link class="link-type" :to="'/components/back-to-top/'+scope.row.id">-->
      <!--<span class="link-type" >{{scope.row.title}}</span>-->
      <!--</router-link>-->
      <!--&lt;!&ndash;<el-tag>{{scope.row.type | typeFilter}}</el-tag>&ndash;&gt;-->
      <!--</template>-->
      <!--</el-table-column>-->
      <el-table-column label="操作" width="400" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.action }}</span>
        </template>
      </el-table-column>
      <!--<el-table-column width="110px" v-if='showReviewer' align="center" :label="$t('table.reviewer')">-->
      <!--<template slot-scope="scope">-->
      <!--<span style='color:red;'>{{scope.row.reviewer}}</span>-->
      <!--</template>-->
      <!--</el-table-column>-->

    </el-table>
    <div>
      <el-pagination
        background
        :current-page="listQuery.page"
        :page-sizes="[listQuery.limit]"
        layout="total, sizes, prev, pager, next"
        :total="count"
        @current-change="handleCurrentChange"
      />
    </div>

  </div>
</template>

<script>
import waves from '@/directive/waves' // 水波纹指令
import { searchLog, logClassify } from '@/api/log'
export default {
  name: 'Log',
  directives: {
    waves
  },
  data() {
    return {
      list: [],
      listLoading: false,
      listQuery: {
        page: 1,
        limit: 20,
        classify: '',
        starttime: 0,
        endtime: 0
      },
      count: 10,
      classifys: [],
      value2: ''
    }
  },
  activated() {
    this.classifylist()
  },
  created() {
    this.handleFilter()
    this.classifylist()
  },
  methods: {
    classifylist() {
      logClassify().then(resp => {
        this.classifys = resp.data
      })
    },
    handleCurrentChange(val) {
      this.listQuery.page = val
      this.handleFilter()
    },
    changeDate(e) {
      if (e.length === 0) {
        this.listQuery.starttime = 0
        this.listQuery.endtime = 0
      } else {
        this.listQuery.starttime = e[0] / 1000
        this.listQuery.endtime = e[1] / 1000
      }
    },

    handleFilter() {
      searchLog(this.listQuery).then(resp => {
        if (resp.data.code === 0) {
          this.list = resp.data.loglist
          this.count = resp.data.count
          this.listQuery.page = resp.data.page
        } else {
          this.$message.error(resp.data.msg)
        }
      })
    }
  }
}
</script>

<style scoped>

</style>
<style>
  #log .el-date-editor .el-range-separator{
    width: 7%;
  }
  .filter-container .filter-item {
    display: inline-fix;
    vertical-align: middle;
    margin-bottom: 3px;
}
</style>
