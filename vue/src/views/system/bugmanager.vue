<template>
  <div>
    <p class="warn-content">
      可以改变bug的所有信息
    </p>
    <div class="filter-container">
      <el-input :placeholder="$t('table.id')" v-model="listQuery.id" type="number" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter"/>
      <el-input :placeholder="$t('table.title')" v-model="listQuery.title" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter"/>
      <el-input :placeholder="$t('table.author')" v-model="listQuery.author" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter"/>
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
      <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">{{ $t('table.search') }}</el-button>
      <!--<el-button class="filter-item" style="margin-left: 10px;" @click="handleCreate" type="primary" icon="el-icon-edit">{{$t('table.add')}}</el-button>-->
      <!--<el-button class="filter-item" type="primary" :loading="downloadLoading" v-waves icon="el-icon-download" @click="handleDownload">{{$t('table.export')}}</el-button>-->
      <!--<el-checkbox class="filter-item" style='margin-left:15px;' @change='tableKey=tableKey+1' v-model="showReviewer">{{$t('table.reviewer')}}</el-checkbox>-->
    </div>
    <el-table
      v-loading="listLoading"
      ref="multipleTable"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%">

      <el-table-column :label="$t('table.id')" align="center" width="50">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('table.date')" width="150px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.date | parseTime('{y}-{m}-{d} {h}:{i}') }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('table.project')" width="100px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.projectname }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('table.level')" width="80px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.level }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('table.importance')" width="100px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.importance }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('table.handle')" width="100px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.handle }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('table.status')" align="center" width="110">
        <template slot-scope="scope">
          <span>{{ scope.row.status }}</span>
          <!--<el-select v-model="scope.row.status" style="width: 100px" class="filter-item" @change="changestatus(scope.row)" >-->
          <!--<el-option v-for="(item, index) in statuslist" :key="index" :label="item" :value="item"/>-->
          <!--</el-select>-->
          <!--<el-tag :type="scope.row.status | statusFilter">{{scope.row.status}}</el-tag>-->
        </template>
      </el-table-column>

      <el-table-column :label="$t('table.title')" min-width="300px" align="center">
        <template slot-scope="scope">

          <router-link :to="'/showbug/'+scope.row.id" class="link-type" align="center">
            <span>{{ scope.row.title }}</span>
          </router-link>
        </template>
      </el-table-column>

      <el-table-column :label="$t('table.dustbin')" min-width="60px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.dustbin }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('table.actions')" align="center" width="230" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <router-link :to="'/bug/edit/'+scope.row.id">
            <el-button type="primary" size="mini" >{{ $t('list.edit') }}</el-button>
          </router-link>
          <el-button type="success" size="mini" @click="handleClose(scope.row)">{{ $t('list.close') }}</el-button>
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
          if (resp.data.statuscode === 0) {
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
