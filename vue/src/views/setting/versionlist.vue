<template>
  <div class="app-container">
    <p class="warn-content">
      版本管理，有些可能是app的下载地址或者是网页的地址，有一个是备用的
    </p>
    <div class="filter-container"></div>
    <div>
      <el-button type="success" plain style="margin: 20px" @click="add"
        >添加版本</el-button
      >
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
          <span>{{
            scope.row.createtime | parseTime("{y}-{m}-{d} {h}:{i}")
          }}</span>
        </template>
      </el-table-column>
      <el-table-column label="项目名" width="150px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.pid | toName(projectMap) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="版本号" width="90px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.name }}</span>
          <!--<svg-icon v-for="n in +scope.row.importance" icon-class="star" class="meta-item__icon" :key="n"></svg-icon>-->
        </template>
      </el-table-column>
      <el-table-column label="地址一" width="130px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.urlone }}</span>
          <!--<svg-icon v-for="n in +scope.row.importance" icon-class="star" class="meta-item__icon" :key="n"></svg-icon>-->
        </template>
      </el-table-column>
      <el-table-column label="地址二" class-name="status-col" width="150">
        <template slot-scope="scope">
          <span>{{ scope.row.urltwo }}</span>
          <!--<el-tag :type="scope.row.status | statusFilter">{{scope.row.status}}</el-tag>-->
        </template>
      </el-table-column>
      <el-table-column
        label="操作"
        align="center"
        width="230"
        class-name="small-padding fixed-width"
      >
        <template slot-scope="scope">
          <el-button
            size="mini"
            type="success"
            @click="handleModifyStatus(scope.row)"
            >修改
          </el-button>
          <el-button
            v-if="scope.row.status != 'draft'"
            size="mini"
            @click="handleRemove(scope.row, 'draft')"
            >删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog
      :close-on-click-modal="false"
      :visible.sync="dialogFormVisible"
      width="60%"
      title="版本管理"
    >
      <el-form :model="form">
        <el-form-item label-width="100" label="项目名">
          <el-select v-model="form.pid" placeholder="请选择">
            <el-option
              v-for="item in projects"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label-width="100" label="版本号">
          <el-input v-model="form.name" auto-complete="off" />
        </el-form-item>
        <el-form-item label-width="100" label="地址一">
          <el-input v-model="form.urlone" auto-complete="off" />
        </el-form-item>
        <el-form-item label-width="100" label="地址二">
          <el-input v-model="form.urltwo" auto-complete="off" />
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
import {
  getVersion,
  removeVersion,
  updateVersion,
  addVersion
} from "@/api/version";
import { getProjectKeyName } from "@/api/get";
export default {
  name: "Versionlist",
  filters: {
    toName(id, projectMap) {
      return projectMap.get(id);
    }
  },
  data() {
    return {
      list: [],
      dialogFormVisible: false,
      listLoading: false,
      projects: [],
      projectMap: new Map(),
      tableKey: 0,
      listQuery: {
        page: 1,
        limit: 15
      },
      total: 0,
      form: {
        id: 0,
        name: "",
        urltwo: "",
        urlone: "",
        pid: undefined
      }
    };
  },
  activated() {
    this.getproject();
  },
  created() {
    this.getproject();
  },
  methods: {
    getproject() {
      getProjectKeyName().then(resp => {
        this.projects = resp.data.data;
        for (let v of this.projects) {
          this.projectMap.set(v.id, v.name);
        }
        this.getversionlist();
      });
    },
    add() {
      this.form.id = 0;
      this.form.name = "";
      this.form.url = "";
      this.form.bakurl = "";
      this.form.project = "";
      this.dialogFormVisible = true;
    },
    getversionlist() {
      getVersion().then(resp => {
        this.list = resp.data.data;
        console.log(this.list);
        this.total = resp.data.versionlist.length;
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
    handleModifyStatus(row) {
      console.log(row);
      this.dialogFormVisible = true;
      this.form = row;
    },
    confirm() {
      if (this.form.id <= 0) {
        addVersion(this.form)
          .then(response => {
            this.form.id = response.data.id;
            this.form.create_time = response.data.create_time;
            this.list.push(this.form);
            this.$message.success("添加成功");
          })
          .catch();
      } else {
        updateVersion(this.form).then(_ => {
          this.$message.success("修改成功");
          const l = this.list.length;
          for (let i = 0; i < l; i++) {
            if (this.list[i].id === this.form.id) {
              this.list[i].name = this.form.name;
              this.list[i].url = this.form.url;
              this.list[i].bakurl = this.form.bakurl;
              this.list[i].project = this.form.project;
              break;
            }
          }
          this.$message.success("修改成功");
          this.dialogFormVisible = false;
        });
      }

      this.dialogFormVisible = false;
    },
    cancel() {
      this.dialogFormVisible = false;
    },
    handleRemove(row) {
      this.$confirm("此操作将关闭bug, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(() => {
          removeVersion(row.id).then(_ => {
            const l = this.list.length;
            for (let i = 0; i < l; i++) {
              if (this.list[i].id === row.id) {
                this.list.splice(i, 1);
              }
            }
            this.$message.success("删除成功");
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

<style scoped></style>
