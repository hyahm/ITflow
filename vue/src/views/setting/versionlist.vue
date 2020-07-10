<template>
  <div class="app-container">
    <p class="warn-content">
      版本管理，有些可能是app的下载地址或者是网页的地址，有一个是备用的
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
    <div>
      <el-button type="success" plain style="margin: 20px" @click="add">添加版本</el-button>
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

      <el-table-column label="版本" width="90px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.name }}</span>
          <!--<svg-icon v-for="n in +scope.row.importance" icon-class="star" class="meta-item__icon" :key="n"></svg-icon>-->
        </template>
      </el-table-column>
      <el-table-column label="地址一" width="130px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.url }}</span>
          <!--<svg-icon v-for="n in +scope.row.importance" icon-class="star" class="meta-item__icon" :key="n"></svg-icon>-->
        </template>
      </el-table-column>
      <el-table-column label="地址二" class-name="status-col" width="150">
        <template slot-scope="scope">
          <span>{{ scope.row.bakurl }}</span>
          <!--<el-tag :type="scope.row.status | statusFilter">{{scope.row.status}}</el-tag>-->
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="230" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <!--<el-button type="primary" size="mini"  @click="handleUpdate(scope.row)">{{$t('table.pass')}}</el-button>-->
          <el-button size="mini" type="success" @click="handleModifyStatus(scope.row)">修改
          </el-button>
          <!--&lt;!&ndash;v-if="scope.row.status!='published'"&ndash;&gt;-->
          <el-button v-if="scope.row.status!='draft'" size="mini" @click="handleRemove(scope.row,'draft')">删除
          </el-button>
        <!--<el-button  size="mini" type="danger" @click="handleStopStatus(scope.row)">{{ scope.row.stop }}-->
        <!--</el-button>-->
        </template>
      </el-table-column>
    </el-table>

    <!--<div class="pagination-container">-->
    <!--<el-pagination :current-page="listQuery.page" :page-sizes="[10,20,30, 50]" :page-size="listQuery.limit" :total="total" background layout="total, sizes, prev, pager, next, jumper" @size-change="handleSizeChange" @current-change="handleCurrentChange"/>-->
    <!--</div>-->

    <el-dialog :close-on-click-modal="false" :visible.sync="dialogFormVisible" width="60%" title="版本管理">
      <el-form :model="form">
        <el-form-item label-width="100" label="版本名">
          <el-input v-model="form.name" auto-complete="off" />
        </el-form-item>
        <el-form-item label-width="100" label="地址一">
          <el-input v-model="form.url" auto-complete="off" />
        </el-form-item>
        <el-form-item label-width="100" label="地址二">
          <el-input v-model="form.bakurl" auto-complete="off" />
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
import { getVersion, removeVersion, updateVersion, addVersion } from '@/api/version'
export default {
  name: 'Versionlist',
  data() {
    return {
      list: [],
      dialogFormVisible: false,
      listLoading: false,
      tableKey: 0,
      listQuery: {
        page: 1,
        limit: 15
      },
      total: 0,
      form: {
        id: -1,
        name: '',
        url: '',
        bakurl: ''
      }
    }
  },
  created() {
    this.getversionlist()
  },
  methods: {
    add() {
      this.form.id = -1
      this.form.name = ''
      this.form.url = ''
      this.form.bakurl = ''
      this.dialogFormVisible = true
    },
    getversionlist() {
      getVersion().then(resp => {
        console.log(resp.data)
        if (resp.data.code === 0) {
          this.list = resp.data.versionlist
          this.total = resp.data.versionlist.length
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    handleSizeChange(val) {
      this.listQuery.limit = val
      this.getList()
    },
    handleCurrentChange(val) {
      this.listQuery.page = val
      this.getList()
    },
    handleModifyStatus(row) {
      this.dialogFormVisible = true
      this.form.id = row.id
      this.form.name = row.name
      this.form.url = row.url
      this.form.bakurl = row.bakurl
    },
    confirm() {
      if (this.form.id <= 0) {
        addVersion(this.form).then(response => {
          if (response.data.code === 0) {
            var row = this.form
            row.id = response.data.id
            row.date = response.data.updatetime
            this.list.unshift(row)
            this.$message.success('添加成功')
          } else {
            this.$message.error(response.data.message)
          }
        }).catch()
      } else {
        updateVersion(this.form).then(resp => {
          if (resp.data.code === 0) {
            this.$message.success('修改成功')
            const l = this.list.length
            for (let i = 0; i < l; i++) {
              if (this.list[i].id === this.form.id) {
                this.list[i].name = this.form.name
                this.list[i].iphoneurl = this.form.iphone
                this.list[i].notiphoneurl = this.form.noiphone
                break
              }
            }
            this.$message.success('修改成功')
            this.dialogFormVisible = false
            return
          } else {
            this.$message.error(resp.data.message)
          }
        })
      }

      this.dialogFormVisible = false
    },
    cancel() {
      this.dialogFormVisible = false
    },
    handleRemove(row) {
      this.$confirm('此操作将关闭bug, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        removeVersion(row.id).then(resp => {
          if (resp === undefined) {
            return
          }
          // if (resp.data === null) {
          if (resp.data.code === 0) {
            const l = this.list.length
            for (let i = 0; i < l; i++) {
              if (this.list[i].id === row.id) {
                this.list.splice(i, 1)
              }
            }
            this.$message.success('删除成功')
            return
          }
          // }
          this.$message.error('删除失败')
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

<style scoped>

</style>
