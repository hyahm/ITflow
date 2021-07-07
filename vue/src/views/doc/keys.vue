<template>
  <div class="app-container">
    <div style="padding: 10px">
      <el-button
        type="success"
        @click="handleAdd"
      >
        增加认证
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
      <el-table-column
        align="center"
        label="ID"
        width="80"
      >
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>

      <el-table-column
        width="180px"
        align="center"
        label="添加时间"
      >
        <template slot-scope="scope">
          <span>{{ scope.row.created | parseTime('{y}-{m}-{d} {h}:{i}') }}</span>
        </template>
      </el-table-column>

      <el-table-column
        width="180px"
        align="center"
        label="更新时间"
      >
        <template slot-scope="scope">
          <span v-if="scope.row.uptime>0">{{ scope.row.uptime | parseTime('{y}-{m}-{d} {h}:{i}') }}</span>
          <span v-else>无更新</span>
        </template>
      </el-table-column>

      <el-table-column
        width="120"
        align="center"
        label="认证名"
      >
        <template slot-scope="scope">
          <span>{{ scope.row.name }}</span>
        </template>
      </el-table-column>

      <el-table-column
        width="120"
        align="center"
        label="认证方式"
      >
        <template slot-scope="scope">
          <span>{{ auths[scope.row.typ] }}</span>
        </template>
      </el-table-column>

      <el-table-column
        align="left"
        label="Actions"
        width="200"
      >
        <template slot-scope="scope">
          <el-button
            type="primary"
            size="small"
            icon="el-icon-update"
            @click="handleUpdate(scope.row)"
          >更新</el-button>
          <el-button
            type="danger"
            size="small"
            icon="el-icon-delete"
            @click="handleDeleteKey(scope.row.id)"
          >删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog
      title="创建文档"
      :visible.sync="dialogFormVisible"
    >
      <el-form>
        <el-form-item
          label="认证名："
          :label-width="formLabelWidth"
        >
          <el-input
            v-model="form.name"
            autocomplete="off"
            @change="handleDomain"
          />
        </el-form-item>
        <el-form-item :label-width="formLabelWidth">
          <el-radio-group
            v-model="form.typ"
            @change="handleAuthMethod"
          >
            <el-radio :label="user">账号密码认证</el-radio>
            <el-radio :label="key">key认证</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item
          v-if="form.typ==2"
          label="私钥："
          :label-width="formLabelWidth"
        >
          <el-input
            type="textarea"
            v-model="form.pri"
            placeholder=""
          />
        </el-form-item>
        <el-form-item
          v-if="form.typ==2"
          label="公钥："
          :label-width="formLabelWidth"
        >
          <el-input
            type="textarea"
            v-model="form.pub"
            placeholder=""
          />
        </el-form-item>
        <el-form-item
          label="用户："
          :label-width="formLabelWidth"
        >
          <el-input
            v-model="form.user"
            placeholder=""
          />
        </el-form-item>
        <el-form-item
          label="密码："
          :label-width="formLabelWidth"
        >
          <el-input
            v-model="form.password"
            placeholder=""
          />
        </el-form-item>
      </el-form>

      <div
        slot="footer"
        class="dialog-footer"
      >
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button
          type="primary"
          @click="handleAddDoc"
        >确 定</el-button>
      </div>
    </el-dialog>

    <div class="pagination-container">
      <el-pagination
        :current-page="listQuery.page"
        :page-sizes="[10,20,30, 50]"
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
import { fetchList, addKey, delKey, checkName } from "@/api/key";
import { getToken } from "@/utils/auth";
export default {
  name: "KeysList",
  data() {
    return {
      list: [],
      user: 1,
      key: 2,
      data: {},
      headers: {},
      giturl: "",
      total: 0,
      dialogFormVisible: false,
      auths: ["", "用户认证", "key认证"],
      listLoading: true,
      formLabelWidth: "120px",
      newname: "",
      listQuery: {
        page: 1,
        limit: 20,
      },
      form: {
        id: 0,
        name: "",
        pri: "",
        pub: "",
        typ: 1,
        user: "",
        created: 0,
        uptime: 0,
        password: "",
      },
    };
  },
  created() {
    this.getList();
  },
  mounted() {
    this.headers = {
      "X-Token": getToken(),
    };
  },
  methods: {
    handleAdd() {
      this.form = {
        id: 0,
        name: "",
        pri: "",
        pub: "",
        typ: 1,
        user: "",
        created: 0,
        uptime: 0,
        password: "",
      };
      this.dialogFormVisible = true;
    },
    handleUpdate(row) {
      console.log(row);

      this.form = row;
      this.dialogFormVisible = true;
    },
    handleDomain(val) {
      // 判断子域名是否被使用
      checkName(val).then((resp) => {
        if (resp.data.code !== 0) {
          this.$message.error("name 重复");
        }
      });
    },
    handleAuthMethod(val) {
      console.log(val);
      this.form.auth = val;
    },

    handleAddDoc() {
      console.log(this.form);
      this.form.typ = parseInt(this.form.typ);
      addKey(this.form).then((resp) => {
        if (this.form.id > 0) {
          for (let i = 0; i < this.list.length; i++) {
            if (this.list[i].id == this.form.id) {
              this.list[i] = this.form;
              this.list[i].uptime = resp.data.time;
              break;
            }
          }
        } else {
          this.form.id = resp.data.id;
          this.form.created = resp.data.time;
          console.log(this.form)
          this.list.push(this.form);
        }
      });

      this.dialogFormVisible = false;
      this.$message.success("更新成功")
    },
    getList() {
      this.listLoading = true;
      fetchList(this.listQuery).then((response) => {
        const data = response.data;
        console.log(data)
        this.list = data.keys;
        this.total = data.total;
        this.listLoading = false;
      });
    },
    handleSizeChange(val) {
      this.listQuery.limit = val;
      this.getList();
    },
    handleCurrentChange(val) {
      this.listQuery.page = val;
      this.getList();
    },
    handleDeleteKey(id) {
      this.$confirm("此操作将永久删除该文档, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          delKey(id).then((resp) => {
            console.log(resp.data);
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
                message: "删除成功!",
              });
            }
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除",
          });
        });
    },
  },
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
</style>