<template>
  <div>
    <div style="margin-left: 10px;margin-top: 10px;display: flex;align-items: center;">
      <el-breadcrumb separator="/" style="font-size: 20px">
        <el-breadcrumb-item v-for="(dd, index) in path" :key="index" class="root_span" @click.native="handleTo"><a>{{ dd }}</a></el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <div style="margin-left: 10px;margin-top: 10px;display: flex;align-items: center;">
      <el-button size="small" type="primary" @click="addFolder"> 新建文件夹</el-button>
      <el-upload
        :action="uploadurl"
        :headers="header"
        :data="uploaddata"
        :on-success="handleSuccess"
        :on-error="handleError"
        style="flex-grow:1;display:flex;justify-content: space-between;margin-left: 10px;align-items: center;"
        class="upload-demo"
        name="share"
        multiple
      >
        <!--<i class="el-icon-upload"></i>-->
        <el-button slot="trigger" size="small" type="primary">上传</el-button>
      </el-upload>
    </div>

    <el-table
      :data="tableData"
      style="width: 100%"
      @cell-mouse-leave="leaveHandle"
      @cell-mouse-enter="enterHandle"
    >
      <el-table-column
        prop="name"
        label="目录"
      >
        <template slot-scope="scope">
          <svg-icon :icon-class="scope.row.isfile?'file':'folder'" />
          <a v-if="!scope.row.isfile" @click="handleEnter(scope.row.name)"><span>{{ scope.row.name }}</span></a>
          <span v-else>{{ scope.row.name }}</span>
          <span v-show="scope.row.name != '..'" style="float: right;margin-right: 10px">
            <a v-if="scope.row.isowner" @click="removeHandle(scope.row.id)"><svg-icon style="margin-right: 10px" icon-class="remove" /></a>
            <a v-if="scope.row.isowner" @click="changeHandle(scope.row)"><svg-icon style="margin-right: 10px" icon-class="rename" /></a>
            <a v-if="scope.row.isfile" :download="scope.row.name" :href="downloadurl + '?id=' + scope.row.id + '&token=' + Token"><svg-icon style="margin-right: 10px" icon-class="download" /></a>
            <!--<svg-icon style="margin-right: 10px" icon-class="moveto" />-->
          </span>
          <span style="clear: both" />
        </template>
      </el-table-column>
      <el-table-column
        prop="size"
        width="200"
        label="大小"
      >
        <template slot-scope="scope">
          <span v-if="scope.row.date!=0">{{ scope.row.size | parseSize }}</span>
        </template>
      </el-table-column>
      <el-table-column
        prop="size"
        width="200"
        label="修改日期"
      >
        <template slot-scope="scope">
          <span v-if="scope.row.date!=0">{{ scope.row.updatetime | parseTime('{y}-{m}-{d} {h}:{i}') }}</span>
        </template>
      </el-table-column>
    </el-table>
    <!--// 创建文件名-->
    <el-dialog :close-on-click-modal="false" :visible.sync="dialogFormVisible" title="文件夹名">
      <el-form>
        <el-form-item label="名称">
          <el-input v-model="form.name" auto-complete="off" />
        </el-form-item>
        <el-form-item label="查看下载">
          <el-radio-group v-model="form.readuser" @change="handleRead">
            <el-radio :label="isReadUser">用戶</el-radio>
            <el-radio :label="!isReadUser">組</el-radio>
          </el-radio-group>
          <el-select v-model="form.readname" placeholder="请选择">
            <el-option
              v-for="(item, index) in readlist"
              :key="index"
              :label="item"
              :value="item"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="查看上传">
          <el-radio-group v-model="form.writeuser" @change="handleRdWr">
            <el-radio :label="isWriteUser">用戶</el-radio>
            <el-radio :label="!isWriteUser">組</el-radio>
          </el-radio-group>
          <el-select v-model="form.writename" placeholder="请选择">
            <el-option
              v-for="(item, index) in rdwrlist"
              :key="index"
              :label="item"
              :value="item"
            />
          </el-select>
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
import { shareList, mkDir, removeFile, renameFile } from '@/api/sharefile'
import { getUsers } from '@/api/get'
import { getGroup } from '@/api/group'
import { getToken } from '@/utils/auth'
export default {
  name: 'ViewList',
  filters: {
    parseSize(size) {
      if (size > 1 << 30) {
        return (size / (1 << 30)).toFixed(2) + 'Gb'
      } else if (size > 1 << 20) {
        return (size / (1 << 20)).toFixed(2) + 'Mb'
      } else if (size > 1 << 10) {
        return (size / (1 << 10)).toFixed(2) + 'Kb'
      } else {
        return size + 'b'
      }
    }
  },
  data() {
    return {
      tableData: [],
      path: ['根目录'],
      readlist: [],
      rdwrlist: [],
      checkread: '',
      checkrdwr: '',
      name: '',
      pwd: '.',
      type: '',
      uploadurl: process.env.VUE_APP_BASE_API + '/share/upload',
      header: {
        'X-Token': getToken()
        // 'Content-Type': 'multipart/form-data'
      },
      dialogFormVisible: false,
      more: false,
      dialogRenameVisible: false,
      oldname: '',
      newname: '',
      downloadname: '',
      urlpwd: '',
      downloadurl: process.env.VUE_APP_BASE_API + '/uploadimg',
      Token: encodeURIComponent(getToken()),
      users: [],
      groups: [],
      realname: '',
      isReadUser: true,
      notReadUser: false,
      isWriteUser: true,
      notWriteUser: false,
      form: {
        id: -1,
        oldname: '',
        filepath: this.pwd,
        name: this.foldername,
        isfile: false,
        size: 0,
        readuser: true,
        writeuser: true,
        readname: '',
        writename: '',
        isShow: false
      }
    }
  },
  computed: {
    uploaddata: function() {
      return {
        dir: this.pwd
      }
    }
  },
  created() {
    this.sharelist()
    this.getuserlist()
    this.getgrouplist()
  },
  methods: {
    // 目录切换目录
    handleTo(e) {
      for (let i = 0; i < this.path.length; i++) {
        if (i === 0) {
          this.pwd = '.'
        } else {
          this.pwd += '/' + this.path[i]
        }

        if (e.toElement.innerText === this.path[i]) {
          this.path = this.path.slice(0, i + 1)
          break
        }
      }
      this.sharelist()
    },
    // 获取用户列表， 赋予权限
    getuserlist() {
      getUsers().then(resp => {
        if (resp.data.code === 0) {
          if (resp.data.users != null) {
            this.users = resp.data.users
          }
        } else {
          this.$message.error(resp.data.msg)
        }
        this.readlist = this.users
        this.rdwrlist = this.users
      })
    },
    // 获取用户列表， 赋予权限
    getgrouplist() {
      getGroup().then(resp => {
        if (resp.data.code === 0) {
          if (resp.data.grouplist != null) {
            this.groups = resp.data.grouplist
          }
        } else {
          this.$message.error(resp.data.msg)
        }
      })
    },
    // 分配权限
    handleRead() {
      if (this.form.readuser) {
        this.readlist = this.users
        this.form.readname = this.realname
      } else {
        this.readlist = this.groups
        this.form.readname = this.readlist[0]
      }
    },
    handleRdWr() {
      if (this.form.writeuser) {
        this.rdwrlist = this.users
        this.form.writename = this.realname
      } else {
        this.rdwrlist = this.groups
        this.form.writename = this.rdwrlist[0]
      }
    },
    // 显示共享的文件
    sharelist() {
      shareList(this.pwd).then(resp => {
        this.realname = resp.data.realname
        if (resp.data.code === 0) {
          this.tableData = resp.data.sharelist
          if (this.tableData !== null) {
            this.tableData.unshift({
              id: -1,
              oldname: '',
              filepath: this.pwd,
              name: '..',
              isfile: false,
              size: 0,
              readuser: true,
              writeuser: true,
              readname: this.realname,
              writename: this.realname,
              isShow: false
            })
          } else {
            this.tableData = [{
              id: -1,
              oldname: '',
              filepath: this.pwd,
              name: '..',
              isfile: false,
              size: 0,
              readuser: true,
              writeuser: true,
              readname: this.realname,
              writename: this.realname,
              isShow: false
            }]
          }
        }
      })
    },
    // 进入文件夹
    handleEnter(dir) {
      // 返回上一级
      if (dir === '..') {
        if (this.pwd === '.') {
          this.path = ['根目录']
          this.$message.warning('已经是根目录了')
          return
        } else {
          // 如果pwd 有多级
          const index = this.pwd.lastIndexOf('/')
          if (index > 0) {
            this.path.pop()
            dir = this.pwd.slice(0, index)
          } else {
            // 如果只有一级，直接退到根
            this.path = ['根目录']
            dir = ''
          }
        }
      } else {
        // 进入一级
        if (this.pwd !== '') {
          this.path.push(dir)
          dir = this.pwd + '/' + dir
        }
      }
      this.pwd = dir
      this.sharelist()
    },
    // 上传文件
    handleSuccess(res, file) {
      if (res.code === 0) {
        this.tableData.push({
          id: res.id,
          name: res.filename,
          isfile: true,
          size: file.size,
          update: file.updatetime,
          isowner: true
          // size: file.
        })
      } else {
        this.$message.error('上传文件失败')
      }
    },
    handleError(error) {
      this.$message.error(error)
    },
    confirm() {
      // 重命名
      this.form.filepath = this.pwd
      if (this.form.id > 0) {
        renameFile(this.form).then(resp => {
          if (resp.data.code === 0) {
            const l = this.tableData.length
            for (let i = 0; i < l; i++) {
              if (this.tableData[i].id === resp.data.id) {
                this.tableData[i] = this.form
              }
            }
          }
        })
      } else {
        // 创建文件夹
        mkDir(this.form).then(resp => {
          if (resp.data.code === 0) {
            this.tableData.push({
              name: this.form.name,
              isowner: true,
              id: resp.data.id,
              oldname: '',
              filepath: this.pwd,
              isfile: false,
              size: 0,
              readuser: this.form.readuser,
              writeuser: this.form.writeuser,
              readname: this.form.readname,
              updatetime: resp.data.updatetime,
              writename: this.form.writename,
              isShow: false
            })
          }
        })
      }
      this.dialogFormVisible = false
    },
    cancel() {
      this.foldername = ''
      this.dialogFormVisible = false
    },
    // 添加文件夹
    addFolder() {
      this.form = {
        id: -1,
        oldname: '',
        filepath: this.pwd,
        name: this.foldername,
        isfile: false,
        size: 0,
        readuser: true,
        writeuser: true,
        readname: this.realname,
        writename: this.realname,
        isShow: false
      }
      this.dialogFormVisible = true
    },
    enterHandle(row, column, cell, event) {
      row.isShow = true
    },
    leaveHandle(row, column, cell, event) {
      row.isShow = false
    },
    removeHandle(id) {
      this.$confirm('此操作将永久删除该文件及其文件夹里面所有文件, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        removeFile(id).then(resp => {
          if (resp.data.code === 0) {
            const l = this.tableData.length
            for (let i = 0; i < l; i++) {
              if (this.tableData[i].id === id) {
                this.tableData.splice(i, 1)
                break
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
    changeHandle(row) {
      this.form = row
      this.dialogFormVisible = true
    }
  }
}
</script>

<style scoped>
  .root_span a {
    text-decoration: underline;
    color: blue
  }
</style>

