<template>
  <div>
    <p class="warn-content">
      可以改变bug的所有信息, 必须搜索才有显示
    </p>
    <div class="filter-container">
      <el-input v-model="listQuery.id" placeholder="id" type="number" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter" />
      <el-input v-model="listQuery.title" placeholder="标题" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter" />
      <el-input v-model="listQuery.author" placeholder="作者" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter" />
      <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">搜索</el-button>
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
          <span>{{ scope.row.importance }}</span>
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

      <el-table-column label="垃圾箱" min-width="60px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.dustbin }}</span>
        </template>
      </el-table-column>

      <el-table-column label="操作" align="center" width="230" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <router-link :to="'/bug/edit/'+scope.row.id">
            <el-button type="primary" size="mini">编辑</el-button>
          </router-link>
          <el-button type="success" size="mini" @click="handleClose(scope.row)">关闭</el-button>
          <!--<el-button type="danger" size="mini" @click="handleRemove(scope.row)">{{ $t('list.remove') }}</el-button>-->
        </template>
      </el-table-column>
    </el-table>
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
export default {
  name: 'Movebug',
  directives: {
    waves
  },
  data() {
    return {
      list: null,
      listLoading: false,
      listQuery: {
        page: 1,
        limit: 15,
        id: undefined,
        title: '',
        author: ''
      },
      total: 0
    }
  },
  created() {
    // this.getstatus()
  },
  methods: {
    // getstatus() {
    //   getStatus().then(resp => {
    //     this.
    //   })
    // },
    handleFilter() {
      if (this.listQuery.id > 0 || this.listQuery.title.length > 0 || this.listQuery.author.length > 0) {
        bugFilter(this.listQuery).then(resp => {
          if (resp.data.code === 0) {
            this.list = resp.data.articlelist
          }
        })
      } else {
        this.$message.warning('必须添加搜索条件')
      }
    }
  }
}
</script>

<style scoped>

</style>
