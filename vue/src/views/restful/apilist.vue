<template>
  <div>
    <el-table
      :data="list"
      fit
      border
      highlight-current-row
      style="width: 100%"
    >
      <el-table-column
        label="Id"
        width="60"
      >
        <template slot-scope="scope">
          <span style="margin-left: 10px">{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column
        label="接口名"
        width="400"
      >
        <template slot-scope="scope">
          <span style="margin-left: 10px"><a :href="'/restful/showapi?id=' + scope.row.id" target="_blank" style="text-decoration:underline;color: blue">{{ scope.row.name }}</a></span>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="150">
        <template slot-scope="scope">
          <el-button
            size="mini"
            @click="update(scope.row)"
          >编辑
          </el-button>
          <el-button
            size="mini"
            type="danger"
            @click="handleDelete(scope.row.id)"
          >删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-button style="margin: 20px" type="success" size="mini" @click="add">添加接口名</el-button>
    <!--dialog-->
    <!--dialog-->
    <!--dialog-->
    <el-dialog :close-on-click-modal="false" :visible.sync="dialogFormVisible" width="60%" title="接口项目名称">
      <el-form :model="form">
        <el-form-item label="接口名">
          <el-input v-model="form.name" width="200" auto-complete="off" />
        </el-form-item>
        <el-form-item label="请求方式">
          <el-select v-model="form.methods" multiple placeholder="请选择">
            <el-option
              v-for="(item, index) in options"
              :key="index"
              :label="item"
              :value="item"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="调用的url">
          <el-input v-model="form.url" width="200" placeholder="只写后面的路径，例如：/login" auto-complete="off" />
        </el-form-item>
        <el-form-item label="请求头">
          <el-select v-model="form.header">
            <el-option v-for="(header, index) in headers" :key="index" :value="header" />
          </el-select>
        </el-form-item>
        <el-form-item label="参数">
          <svg-icon icon-class="add" @click.native="handleAdd" />
          <div v-for="(option, index) in form.opts" :key="index" class="div_opts">
            <el-input :value="option.name" class="key_opts" type="text" placeholder="参数名" @change="handleUpdateName(option.id, $event)" />
            <el-select :value="option.type" class="select_opts" placeholder="类型" @change="handleUpdateType(option.id, $event)">
              <el-option v-for="(t, i) in types" :key="i" :value="t" />
            </el-select>
            <el-select :value="option.need" class="select_opts" placeholder="是否必须" @change="handleUpdateNeed(option.id, $event)">
              <el-option v-for="(n, i) in needs" :key="i" :value="n" />
            </el-select>
            <el-input
              :value="opt.default"
              class="default_opts"
              type="text"
              placeholder="默认值"
              @change="handleUpdateValue(opt.id, $event)"
            />
            <el-input
              :value="opt.info"
              class="info_opts"
              type="text"
              placeholder="说明"
              @change="handleUpdateInfo(opt.id, $event)"
            />
            <svg-icon icon-class="delete" @click.native="handleDel(opt.id)" />
          </div>
        </el-form-item>
        <el-form-item label="请求参数实例">
          <el-input v-model="form.resp" rows="5" placeholder="json" type="textarea" width="200" auto-complete="off" />
        </el-form-item>
        <el-form-item label="回调返回值类型">
          <el-select v-model="form.calltype" clearable placeholder="请选择">
            <el-option
              v-for="(rt, index) in rtls"
              :key="index"
              :label="rt"
              :value="rt"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="回调返回值">
          <el-input v-model="form.result" rows="5" placeholder="只有text和json类型" type="textarea" width="200" auto-complete="off" />
        </el-form-item>
        <el-form-item label="说明">
          <el-input v-model="form.information" width="200" rows="5" type="textarea" placeholder="支持markdown" auto-complete="off" />
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
import { apiList, apiUpdate, apiDel, apiAdd, editOne } from '@/api/restful'
import { typeGet } from '@/api/type'
import { headerGet } from '@/api/header'
// import Icons from "../svg-icons/index"

export default {
  name: 'ProjectList',
  // components: { Icons },
  data() {
    return {
      dialogFormVisible: false,
      selecttype: '',
      form: {
        name: '',
        header: [],
        methods: [],
        url: '',
        resp: '',
        opts: [],
        result: '',
        calltype: 'json',
        id: -1,
        pid: -1,
        information: ''
      },
      list: [],
      options: ['GET', 'POST', 'PUT', 'HEAD', 'DELETE'],
      methods: ['GET'],
      opts: [],
      types: [],
      needs: ['必须', '可选'],
      headers: [],
      rtls: ['text', 'json'],
      opt: {
        id: -1,
        name: '',
        type: '',
        need: true,
        default: '',
        info: ''
      },
      id: -1,
      pid: -1,
      oid: -1
    }
  },
  created() {
    this.GetQueryString()
    this.getrestname()
    this.gettypes()
    this.getheaders()
  },
  methods: {
    getheaders() {
      headerGet().then(resp => {
        if (resp.data.code === 0) {
          if (resp.data.headers !== null) {
            this.headers = resp.data.headers
          }
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    gettypes() {
      typeGet().then(resp => {
        if (resp.data.code === 0) {
          if (resp.data.headers !== null) {
            this.types = resp.data.types
          }
        } else {
          this.$message.error(resp.data.message)
        }
      })
    },
    GetQueryString() {
      this.pid = parseInt(window.location.search.split('=')[1])
    },
    getrestname() {
      this.$nextTick(() => {
        apiList(this.pid).then(resp => {
          if (resp.data.code === 0) {
            if (resp.data.list !== null) {
              this.list = resp.data.list
            }
          } else if (resp.data.code === 14) {
            this.$message({
              type: 'info',
              message: '你没有权限访问'
            })
          } else {
            this.$message({
              type: 'info',
              message: '你访问到月亮去了'
            })
          }
        })
      })
    },
    update(row) {
      editOne(row.id).then(resp => {
        if (resp.data.code === 0) {
          if (resp.data.opts === null) {
            resp.data.opts = []
          }
          this.form = resp.data
        } else {
          this.$message.error(resp.data.message)
        }
      })
      this.dialogFormVisible = true
    },
    handleDelete(id) {
      this.$confirm('此操作将关闭bug, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        apiDel(id).then(resp => {
          if (resp.data.code === 0) {
            for (let i = 0; i < this.list.length; i++) {
              if (this.list[i].id === id) {
                this.list.splice(i, 1)
              }
            }
          } else {
            this.$message.error(resp.data.message)
          }
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
      })
    },
    handleAdd() {
      // 如果是空的，就添加
      if (this.form.opts === null || this.form.opts.length === 0) {
        this.form.opts.push({
          id: -1,
          name: '',
          type: '',
          need: '必须',
          default: '',
          info: ''
        })
        this.oid--
      } else {
        this.opts.id = this.oid
        this.form.opts.push({
          id: this.oid,
          name: '',
          type: '',
          need: '必须',
          default: '',
          info: ''
        })
        this.oid--
      }
    },
    handleDel(id) {
      const l = this.form.opts.length
      for (let i = 0; i < l; i++) {
        if (this.form.opts[i].id === id) {
          this.form.opts.splice(i, 1)
        }
      }
    },
    confirm() {
      if (this.form.name === '') {
        this.$message({
          type: 'info',
          message: '名称不能为空'
        })
      }
      if (this.form.url === '') {
        this.$message({
          type: 'info',
          message: 'url不能为空'
        })
      }
      if (this.form.methods.length === 0) {
        this.$message({
          type: 'info',
          message: 'methods不能为空'
        })
      }
      if (this.form.id > 0) {
        apiUpdate(this.form).then(resp => {
          if (resp.data.code === 0) {
            for (let i = 0; i < this.list.length; i++) {
              if (this.list[i].id === this.form.id) {
                this.list[i].name = this.form.name
              }
            }
          } else {
            this.$message({
              type: 'info',
              message: '修改失败'
            })
          }
        })
      } else {
        apiAdd(this.form).then(resp => {
          if (resp.data.code === 0) {
            this.list.push({
              id: resp.data.id,
              name: this.form.name
            })
          } else {
            this.$message({
              type: 'info',
              message: '添加失败'
            })
          }
        })
      }
      this.dialogFormVisible = false
    },
    add() {
      this.form = {
        name: '',
        methods: [],
        url: '',
        opts: [],
        reflect: '',
        calltype: 'json',
        id: -1,
        pid: this.pid,
        information: ''
      }
      this.dialogFormVisible = true
    },
    cancel() {
      this.form.name = ''
      this.form.id = -1
      this.dialogFormVisible = false
    },
    // opts
    handleUpdateType(id, value) {
      const l = this.form.opts.length
      for (let i = 0; i < l; i++) {
        if (this.form.opts[i].id === id) {
          this.form.opts[i].type = value
        }
      }
    },
    handleUpdateName(id, value) {
      const l = this.form.opts.length
      for (let i = 0; i < l; i++) {
        if (this.form.opts[i].id === id) {
          this.form.opts[i].name = value
        }
      }
    },
    handleUpdateNeed(id, value) {
      const l = this.form.opts.length
      for (let i = 0; i < l; i++) {
        if (this.form.opts[i].id === id) {
          this.form.opts[i].need = value
        }
      }
    },
    handleUpdateValue(id, value) {
      const l = this.form.opts.length
      for (let i = 0; i < l; i++) {
        if (this.form.opts[i].id === id) {
          this.form.opts[i].default = value
        }
      }
    },
    handleUpdateInfo(id, value) {
      const l = this.form.opts.length
      for (let i = 0; i < l; i++) {
        if (this.form.opts[i].id === id) {
          this.form.opts[i].info = value
        }
      }
    }
  }
}
</script>

<style scoped type="text/css">
.default_opts {
  width: 80px;
}
.info_opts {
  width: 300px;
}
.key_opts {
  width: 100px;
}
.select_opts {
  width: 100px;
}
.div_opts {
  padding: 2px 0 2px 0;
}
</style>
