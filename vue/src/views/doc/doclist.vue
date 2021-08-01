<template>
  <div class="app-container">
    <div style="padding: 10px">
      <el-button type="success" @click="handleAdd">
        增加文档
      </el-button>
    </div>

    <el-table
      v-loading="listLoading"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%"
    >
      <el-table-column align="center" label="ID" width="80">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>

      <el-table-column width="180px" align="center" label="添加时间">
        <template slot-scope="scope">
          <span>{{
            scope.row.created | parseTime("{y}-{m}-{d} {h}:{i}")
          }}</span>
        </template>
      </el-table-column>

      <el-table-column width="180px" align="center" label="更新时间">
        <template slot-scope="scope">
          <span>{{ scope.row.uptime | parseTime("{y}-{m}-{d} {h}:{i}") }}</span>
        </template>
      </el-table-column>

      <el-table-column width="120" align="center" label="名称">
        <template slot-scope="scope">
          <a :href="href(scope.row.name)" target="_blank" class="link-type">
            <span>{{ scope.row.name }}</span>
          </a>
        </template>
      </el-table-column>

      <el-table-column width="300" align="center" label="git地址">
        <template slot-scope="scope">
          <span>{{ scope.row.giturl }}</span>
        </template>
      </el-table-column>

      <el-table-column width="120" align="center" label="认证名">
        <template slot-scope="scope">
          <span>{{ scope.row.authname }}</span>
        </template>
      </el-table-column>

      <el-table-column align="left" label="Actions" width="230">
        <template slot-scope="scope">
          <el-button
            type="primary"
            size="small"
            icon="el-icon-update"
            @click="handleUpdate(scope.row.id)"
            >更新</el-button
          >
          <el-button
            type="danger"
            size="small"
            icon="el-icon-delete"
            @click="handleDelete(scope.row.id)"
            >删除</el-button
          >
          <a :data-clipboard-text="scope.row.hook" class="iconfont icon-copy"
            >复制</a
          >
        </template>
      </el-table-column>
    </el-table>

    <el-dialog title="创建文档" :visible.sync="dialogFormVisible">
      <el-form>
        <el-form-item label="子域名：" :label-width="formLabelWidth">
          <el-input
            v-model="form.name"
            autocomplete="off"
            @change="handleDomain"
          />
        </el-form-item>

        <el-form-item label="仓库地址：" :label-width="formLabelWidth">
          <el-input
            v-model="form.giturl"
            placeholder="https://github.com/xxx/xxx.git or git@xxxx.xxxx:xxxx/xxxx.git"
          />
        </el-form-item>

        <el-form-item label="文档根目录：" :label-width="formLabelWidth">
          <el-input v-model="form.dir" placeholder="不填即为当前目录" />
        </el-form-item>

        <el-form-item :label-width="formLabelWidth" label="认证方式：">
          <el-select v-model="form.kid" clearable placeholder="请选择">
            <el-option
              v-for="item in auths"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            >
            </el-option>
          </el-select>
        </el-form-item>
      </el-form>

      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="handleAddDoc">确 定</el-button>
      </div>
    </el-dialog>

    <div class="pagination-container">
      <el-pagination
        :current-page="listQuery.page"
        :page-sizes="[10, 20, 30, 50]"
        :page-size="listQuery.limit"
        :total="total"
        background
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
  </div>
</template>

<script>
import { fetchList, addDoc, dropDoc, checkDomain, updateDoc } from "@/api/doc";
import { getMykeys } from "@/api/key";
import { getToken } from "@/utils/auth";
import Clipboard from "clipboard";
export default {
  name: "DocList",
  data() {
    return {
      host: process.env.VUE_APP_BASE_API,
      list: [],
      data: {},
      headers: {},
      gitchecked: false,
      giturl: "",
      total: 0,
      dialogFormVisible: false,
      auths: [],
      showAuths: ["无认证", "密码认证", "秘钥认证"],
      listLoading: true,
      formLabelWidth: "120px",
      uploadurl: process.env.VUE_APP_BASE_API + "/doc/upload",
      newname: "",
      listQuery: {
        page: 1,
        limit: 20
      },
      form: {
        id: 0,
        name: "",
        giturl: "",
        dir: "",
        kid: 0
      }
    };
  },
  computed: {
    href() {
      return function(name) {
        return (
          process.env.VUE_APP_BASE_API +
          "/docs/" +
          name +
          "/#/?token=" +
          getToken()
        );
      };
    }
  },

  created() {
    this.getMykeyNames();
    this.getList();
  },
  mounted() {
    this.headers = {
      "X-Token": getToken()
    };
    const clipboard = new Clipboard(".icon-copy");
    clipboard.on("success", e => {
      this.$message.success(e.text + " 已复制到剪贴板！");
    });
  },
  methods: {
    getMykeyNames() {
      getMykeys().then(resp => {
        if (resp.data.code === 0) {
          this.auths = resp.data.auths;
        }
      });
    },
    handleAdd() {
      this.dialogFormVisible = true;
    },
    handleUpdate(id) {
      updateDoc(id).then(resp => {
        this.$message.success("更新成功");
      });
    },
    handleDomain(val) {
      // 判断子域名是否被使用
      checkDomain(val).then(() => {});
    },
    handleAuthMethod(val) {
      this.form.auth = val;
    },

    handleAddDoc() {
      if (/^[a-z]{1}[a-z0-9]{1,20}$/.test(this.form.name)) {
        if (this.form.id > 0) {
          return;
        }
        if (this.form.kid === "") {
          this.form.kid = 0;
        }
        addDoc(this.form).then(resp => {
          if (resp.data.code === 0) {
            this.$message.success("添加成功");
            for (let i = 0; i < this.auths.length; i++) {
              if (this.auths[i].id === this.form.kid) {
                resp.data.doc.authname = this.auths[i].name;
                break;
              }
            }
            var h = resp.data.doc.giturl.indexOf("@");
            var domain = "";
            if (h >= 0) {
              const start = resp.data.doc.giturl.indexOf("@");
              const end = resp.data.doc.giturl.indexOf(":");
              domain = resp.data.doc.giturl.slice(start + 1, end);
            } else {
              // https://gitee.com/cander/scs.git
              const start = resp.data.doc.giturl.indexOf(":");
              var newstr = resp.data.doc.giturl.slice(start + 3);
              const end = newstr.indexOf("/");
              domain = newstr.slice(0, end);
            }
            var platform = "";
            switch (domain) {
              case "gitee.com":
                platform = "gitee";
                break;
              case "github.com":
                platform = "github";
                break;
              default:
                platform = "gitlab";
                break;
            }
            resp.data.doc.hook =
              process.env.VUE_APP_BASE_API +
              "/" +
              platform +
              "/" +
              resp.data.doc.name;
            this.list.push(resp.data.doc);
          } else {
            this.$message.error(resp.data.msg);
            this.dialogFormVisible = false;
            return;
          }
        });
      } else {
        this.$message.error("必须是小写英文或数字, 必须英文开头，至少2位");
      }
      this.dialogFormVisible = false;
    },
    getList() {
      this.listLoading = true;
      fetchList(this.listQuery).then(response => {
        const data = response.data;
        this.list = data.doc;
        for (let i = 0; i < this.list.length; i++) {
          var h = this.list[i].giturl.indexOf("@");
          var domain = "";
          if (h >= 0) {
            const start = this.list[i].giturl.indexOf("@");
            const end = this.list[i].giturl.indexOf(":");
            domain = this.list[i].giturl.slice(start + 1, end);
          } else {
            // https://gitee.com/cander/scs.git
            const start = this.list[i].giturl.indexOf(":");
            var newstr = this.list[i].giturl.slice(start + 3);
            const end = newstr.indexOf("/");
            domain = newstr.slice(0, end);
          }
          var platform = "";
          switch (domain) {
            case "gitee.com":
              platform = "gitee";
              break;
            case "github.com":
              platform = "github";
              break;
            default:
              platform = "gitlab";
              break;
          }
          this.list[i].hook =
            process.env.VUE_APP_BASE_API +
            "/" +
            platform +
            "/" +
            this.list[i].name;
        }
        this.total = data.count;
      });
      this.listLoading = false;
    },
    handleSizeChange(val) {
      this.listQuery.limit = val;
      this.getList();
    },
    handleCurrentChange(val) {
      this.listQuery.page = val;
      this.getList();
    },
    handleDelete(id) {
      this.$confirm("此操作将永久删除该文档, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(() => {
          dropDoc(id).then(resp => {
            if (resp.data.code === 0) {
              const l = this.list.length;
              for (let i = 0; i < l; i++) {
                if (this.list[i].id === id) {
                  this.list.splice(i, 1);
                  break;
                }
              }
              this.$message({
                type: "success",
                message: "删除成功!"
              });
            }
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除"
          });
        });
    }
  }
};
</script>

<style scoped>
.edit-input {
  padding-right: 100px;
}
.cancel-btn {
  position: absolute;
  right: 15px;
  top: 10px;
}
.iconfont {
  margin-left: 10px;
}
</style>
