<template>
  <div>
    <p class="warn-content">
      请求参数key的类型
    </p>
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
        label="类型名"
        width="400"
      >
        <template slot-scope="scope">
          <span style="margin-left: 10px">{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="250">
        <template slot-scope="scope">
          <el-button
            size="mini"
            @click="update(scope.row)"
          >修改
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
    <el-button style="margin: 20px" type="success" size="mini" @click="add">添加类型名</el-button>
    <el-dialog :close-on-click-modal="false" :visible.sync="dialogFormVisible" width="60%" title="类型名称">
      <el-form :model="form">
        <el-form-item label="类型名">
          <el-input v-model="form.name" width="200" auto-complete="off" />
        </el-form-item>
        <el-form-item>
          <el-radio-group v-model="form.checktype" @change="handleChange">
            <el-radio :label="1">数组</el-radio>
            <el-radio :label="2">对象</el-radio>
          </el-radio-group>
          <div v-if="form.checktype === 1">
            <el-select v-model="form.listtype" placeholder="请选择">
              <el-option
                v-for="(type, index) in types"
                :key="index"
                :label="type"
                :value="type"
              />
            </el-select>
          </div>
          <div v-if="form.checktype === 2">
            <svg-icon icon-class="add" @click.native="handleAdd" />
            <div v-for="(opt, index) in form.opts" :key="index" class="div_opts">
              <el-input :value="opt.name" class="key_opts" type="text" placeholder="参数名" @change="handleUpdateName(opt.id, $event)" />
              <el-select :value="opt.type" class="select_opts" placeholder="类型" @change="handleUpdateType(opt.id, $event)">
                <el-option v-for="(t, i) in types" :key="i" :value="t" />
              </el-select>
              <el-select :value="opt.need" class="select_opts" placeholder="是否必须" @change="handleUpdateNeed(opt.id, $event)">
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
          </div>
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
import { typeList, typeUpdate, typeDel, typeAdd } from '@/api/type'
import { typeGet } from '@/api/type'
export default {
  name: 'TypeList',
  data() {
    return {
      dialogFormVisible: false,
      form: {
        name: '',
        id: -1,
        checktype: 0,
        opts: [],
        listtype: ''
      },
      list: [],
      types: [],
      oid: -1,
      needs: ['必须', '可选']
    }
  },
  created() {
    this.getrestname()
    // this.gettypes()
  },
  methods: {
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
    },
    handleUpdateNeed(id, value) {
      const l = this.form.opts.length
      for (let i = 0; i < l; i++) {
        if (this.form.opts[i].id === id) {
          this.form.opts[i].need = value
        }
      }
    },
    handleUpdateType(id, value) {
      const l = this.form.opts.length
      for (let i = 0; i < l; i++) {
        if (this.form.opts[i].id === id) {
          this.form.opts[i].type = value
        }
      }
    },
    handleAdd() {
      // 如果是空的，就添加
      this.form.opts.push({
        id: this.oid,
        name: '',
        type: '',
        need: '必须',
        default: '',
        info: ''
      })
      this.oid--
    },
    handleUpdateName(id, value) {
      const l = this.form.opts.length
      for (let i = 0; i < l; i++) {
        if (this.form.opts[i].id === id) {
          this.form.opts[i].name = value
        }
      }
    },
    gettypes() {
      typeGet().then(resp => {
        if (resp.data.code === 0) {
          if (resp.data.headers !== null) {
            this.types = resp.data.types
          }
        } else {
          this.$message.error(resp.data.msg)
        }
      })
    },
    handleChange(e) {
      console.log(e)
    },
    getrestname() {
      typeList().then(resp => {
        if (resp.data.code === 0) {
          if (resp.data.list !== null) {
            this.list = resp.data.list
            const l = this.list.length
            for (let i = 0; i < l; i++) {
              if (this.list[i].opts === null) {
                this.list[i].opts = []
              }
            }
          }
        } else {
          this.$message.error(resp.data.msg)
        }
      })
    },
    update(row) {
      this.form = JSON.parse(JSON.stringify(row))
      typeGet().then(resp => {
        if (resp.data.code === 0) {
          if (resp.data.headers !== null) {
            this.types = resp.data.types
            const l = this.types.length
            for (let i = 0; i < l; i++) {
              if (this.types[i] === this.form.name) {
                this.types.splice(i, 1)
              }
            }
          }
        } else {
          this.$message.error(resp.data.msg)
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
        typeDel(id).then(resp => {
          if (resp.data.code === 0) {
            for (let i = 0; i < this.list.length; i++) {
              if (this.list[i].id === id) {
                this.list.splice(i, 1)
              }
            }
          } else {
            this.$message.error(resp.data.msg)
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
        typeUpdate(this.form).then(resp => {
          if (resp.data.code === 0) {
            if (this.form.checktype === 1) {
              const data = JSON.parse(JSON.stringify(this.form))
              data.id = resp.data.id
              const l = this.list.length
              for (let i = 0; i < l; i++) {
                if (this.list[i].id === data.id) {
                  this.list.splice(i, 1, data)
                }
              }
            }
            if (this.form.checktype === 2) {
              console.log(resp.data)
              const data = this.form
              data.id = resp.data.id
              data.opts = resp.data.opts
              const l = this.list.length

              for (let i = 0; i < l; i++) {
                if (this.list[i].id === data.id) {
                  this.list.splice(i, 1, data)
                }
              }
            }
          }
        })
      } else {
        typeAdd(this.form).then(resp => {
          if (resp.data.code === 0) {
            if (this.form.checktype === 1) {
              const data = this.form
              data.id = resp.data.id
              this.list.push(data)
            }
            if (this.form.checktype === 2) {
              const data = this.form
              data.id = resp.data.id
              data.opts = resp.data.opts
              this.list.push(data)
            }
          }
        })
      }
      this.dialogFormVisible = false
    },
    add() {
      this.form = {
        name: '',
        id: -1,
        checktype: 1,
        opts: [],
        listtype: ''
      }
      this.gettypes()
      this.dialogFormVisible = true
    },
    cancel() {
      this.dialogFormVisible = false
    }
  }
}
</script>

<style scoped>
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
