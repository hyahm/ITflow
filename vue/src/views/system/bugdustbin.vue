<template>
  <div>
    <p class="warn-content">
      可以改变bug的所有信息, 必须搜索才有显示
    </p>
    <div class="filter-container">
      <el-input v-model="listQuery.title" placeholder="标题" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter" />
      <el-select v-model="listQuery.level" placeholder="级别" clearable style="width: 90px" class="filter-item" @keyup.enter.native="handleFilter">
        <el-option v-for="(item, index) in levels" :key="index" :label="item" :value="item" @keyup.enter.native="handleFilter" />
      </el-select>
      <el-select v-model="listQuery.project" placeholder="项目名" clearable class="filter-item" style="width: 130px">
        <el-option v-for="(item, index) in projectnames" :key="index" :label="item" :value="item" />
      </el-select>
      <!-- <el-input v-model="listQuery.id" placeholder="id" type="number" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter" />
      <el-input v-model="listQuery.title" placeholder="标题" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter" />
      <el-input v-model="listQuery.author" placeholder="作者" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter" /> -->
      <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">搜索</el-button>
      <el-dropdown :hide-on-click="false" :show-timeout="100" trigger="click" style="vertical-align: top;">
        <el-button plain>
          状态({{ statuslength }})
          <i class="el-icon-caret-bottom el-icon--right" />
        </el-button>
        <el-dropdown-menu slot="dropdown" class="no-border">
          <el-checkbox-group v-model="showstatus" style="padding-left: 15px;" @change="HandleChange">
            <el-checkbox v-for="(item, index) in allStatus" :key="index" :label="item">
              {{ item }}
            </el-checkbox>
          </el-checkbox-group>
        </el-dropdown-menu>
      </el-dropdown>
    </div>
    <el-table
      ref="multipleTable"
      v-loading="listLoading"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%"
    >

      <el-table-column label="id" align="center" width="50">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>

      <el-table-column label="日期" width="150px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.date | parseTime('{y}-{m}-{d} {h}:{i}') }}</span>
        </template>
      </el-table-column>

      <el-table-column label="项目" width="100px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.projectname }}</span>
        </template>
      </el-table-column>

      <el-table-column label="优先级" width="80px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.level }}</span>
        </template>
      </el-table-column>

      <el-table-column label="重要性" width="100px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.important }}</span>
        </template>
      </el-table-column>

      <el-table-column label="处理者" width="100px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.handle }}</span>
        </template>
      </el-table-column>

      <el-table-column label="状态" align="center" width="110">
        <template slot-scope="scope">
          <span>{{ scope.row.status }}</span>
          <!--<el-select v-model="scope.row.status" style="width: 100px" class="filter-item" @change="changestatus(scope.row)" >-->
          <!--<el-option v-for="(item, index) in statuslist" :key="index" :label="item" :value="item"/>-->
          <!--</el-select>-->
          <!--<el-tag :type="scope.row.status | statusFilter">{{scope.row.status}}</el-tag>-->
        </template>
      </el-table-column>

      <el-table-column label="标题" min-width="300px" align="center">
        <template slot-scope="scope">

          <router-link :to="'/showbug/'+scope.row.id" class="link-type" align="center">
            <span>{{ scope.row.title }}</span>
          </router-link>
        </template>
      </el-table-column>

      <!-- <el-table-column label="垃圾箱" min-width="60px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.dustbin }}</span>
        </template>
      </el-table-column> -->

      <el-table-column label="操作" align="center" width="230" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <router-link :to="'/bug/edit/'+scope.row.id">
            <el-button type="primary" size="mini">编辑</el-button>
          </router-link>
          <el-button type="success" size="mini" @click="resume(scope.row.id)">恢复</el-button>
          <!--<el-button type="danger" size="mini" @click="handleRemove(scope.row)">{{ $t('list.remove') }}</el-button>-->
        </template>
      </el-table-column>
    </el-table>
    <div class="pagination-container">
      <el-pagination
        :current-page="listQuery.page"
        :page-sizes="[10,15,20,30]"
        :page-size="listQuery.limit"
        :total="total"
        background
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
    <!--<el-pagination-->
    <!--:current-page="currentPage4"-->
    <!--:page-sizes="[100, 200, 300, 400]"-->
    <!--:page-size="100"-->
    <!--:total="total"-->
    <!--layout="total, sizes, prev, pager, next, jumper"-->
    <!--@size-change="handleSizeChange"-->
    <!--@current-change="handleCurrentChange">-->
    <!--</el-pagination>-->
  </div>
</template>

<script>
import waves from '@/directive/waves' // 水波纹指令
import { bugFilter } from '@/api/search' // 水波纹指令
import { resumeBug } from '@/api/bugs'
import { statusFilter } from '@/api/status'
import { getProject, getLevels, getStatus, getPermStatus } from '@/api/get'
export default {
  name: 'Movebug',
  directives: {
    waves
  },
  data() {
    return {
      projectnames: [],
      levels: [],
      list: [],
      listLoading: false,
      listQuery: {
        page: 1,
        limit: 15,
        level: '',
        project: '',
        title: ''
      },
      showstatus: [],
      total: 0,
      allStatus: [],
      statuslength: 0
    }
  },
  created() {
    this.getpname()
    this.getlevels()
    this.getstatus()
    this.handleFilter()
  },
  methods: {
    getstatus() {
      getStatus().then(resp => {
        if (resp.data.code === 0) {
          this.allStatus = resp.data.statuslist
        } else {
          this.$message.error(resp.data.message)
        }
      })
      // 可以修改的权限
      getPermStatus().then(resp => {
        if (resp.data.code === 0) {
          this.showstatus = resp.data.statuslist
          this.statuslength = this.showstatus.length
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    getlevels() {
      getLevels().then(resp => {
        if (resp.data.code === 0) {
          this.levels = resp.data.levels
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
    HandleChange() {
      const data = {
        checkstatus: this.showstatus
      }
      statusFilter(data).then(resp => {
        if (resp.data.code === 0) {
          this.statuslength = this.showstatus.length
          this.handleFilter()
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    handleFilter() {
      bugFilter(this.listQuery).then(resp => {
        if (resp.data.code === 0) {
          this.list = resp.data.articlelist
          this.total = resp.data.total
          this.listQuery.page = resp.data.page
        }
      })
    },
    handleSizeChange(val) {
      this.listQuery.limit = val
      this.handleFilter()
    },
    handleCurrentChange(val) {
      this.listQuery.page = val
      this.handleFilter()
    },
    resume(id) {
      resumeBug(id).then(resp => {
        if (resp.data.code === 0) {
          const l = this.list.length
          for (var i = 0; i < l; i++) {
            if (this.list[i].id === id) {
              this.list.splice(i, 1)
              this.$message.success('恢复成功')
              return
            }
          }
        } else {
          this.$message.error(resp.data.message)
        }
      })
    }
  }
}
</script>

<style scoped>

</style>
