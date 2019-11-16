<template>
  <div>
    <p class="warn-content">
      一个项目可能使用头一个请求头，这是一个模板
    </p>
    <el-table
      :data="list"
      fit
      border
      highlight-current-row
      style="width: 100%">
      <el-table-column
        label="Id"
        width="60">
        <template slot-scope="scope">
          <span style="margin-left: 10px">{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="类型名"
        width="400">
        <template slot-scope="scope">
          <span style="margin-left: 10px">{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="250">
        <template slot-scope="scope">
          <el-button
            size="mini"
            @click="update(scope.row)">修改
          </el-button>
          <el-button
            size="mini"
            type="danger"
            @click="handleDelete(scope.row.id)">删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-button style="margin: 20px" type="success" size="mini" @click="add">添加请求头</el-button>
    <el-dialog :close-on-click-modal="false" :visible.sync="dialogFormVisible" width="60%" title="类型名称">
      <el-form :model="form">
        <el-form-item label="请求头名">
          <el-input v-model="form.name" width="200" auto-complete="off"/>
        </el-form-item>
        <el-form-item label="请求头">
          <svg-icon icon-class="add" @click.native="handleAdd"/>
          <div v-for="(opt, index) in form.hhids" :key="index" class="div_opts" >
            <el-input v-model="opt.key" class="key_opts" type="text" placeholder="key"/>
            <el-input v-model="opt.value" class="value_opts" type="text" placeholder="value" />
            <svg-icon icon-class="delete" @click.native="handleDel(opt.id)"/>
          </div>
        </el-form-item>
        <el-form-item label="说明">
          <el-input v-model="form.remark" type="textarea" width="200" auto-complete="off"/>
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
import { headerList, headerUpdate, headerDel, headerAdd } from '@/api/header'

export default {
  name: 'Header',
  data() {
    return {
      dialogFormVisible: false,
      form: {
        name: '',
        id: -1,
        hhids: [],
        remark: ''
      },
      list: [],
      index: -1,
    }
  },
  created() {
    this.getlist()
  },
  methods: {
    handleDel(id) {
      const l = this.form.hhids.length
      for (let i = 0; i < l; i++) {
        if (this.form.hhids[i].id === id) {
          this.form.hhids.splice(i, 1)
        }
      }
    },
    handleAdd() {
      this.form.hhids.push({
        id: this.index--,
        key: '',
        value: ''
      })
    },
    getlist() {
      headerList().then(resp => {
        if (resp.data.code === 0) {
          if (resp.data.headers !== null) {
            this.list = resp.data.headers
          }
        }
      })
    },
    update(row) {
      this.form = row
      if (row.hhids == null || row.hhids.length === 0) {
        this.form.hhids = []
      }
      this.dialogFormVisible = true
    },
    handleDelete(id) {
      this.$confirm('此操作将删除此项目, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        headerDel(id).then(resp => {
          if (resp.data.code === 0) {
            this.list.hhids = resp.data.headerlist
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
      if (this.form.id > 0) {
        headerUpdate(this.form).then(resp => {
          if (resp.data.code === 0) {

            for (let i = 0; i < this.list.length; i++) {
              if (this.list[i].id === this.form.id) {
                this.list[i].name = this.form.name
                this.list[i].remark = this.form.remark
                this.list[i].hhids = resp.data.headerlist
              }
            }
            this.$message.success('修改成功')
          } else {
            this.$message.error('修改失败')
          }

        })
      } else {
        headerAdd(this.form).then(resp => {
          if (resp.data.code === 0) {
            this.list.push({
              id: resp.data.id,
              name: this.form.name,
              hhids: this.form.hhids,
              remark: this.form.remark
            })
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
        name: '',
        id: this.index--,
        hhids: [],
        remark: ''
      },
      this.dialogFormVisible = true
    },
    cancel() {
      this.dialogFormVisible = false
    }
  }
}
</script>

<style scoped type="text/css">
.key_opts {
  float: left;
  width: 40%;
}
  .value_opts {
    float: left;
    margin-left: 5px;
    width: 55%;
  }
</style>
