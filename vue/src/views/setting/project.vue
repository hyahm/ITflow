<template>
  <div style="padding-left: 20px">
    <p class="warn-content">
      如果删除的此项目的某参与者， 在他任务里面还是会显示，只不过没操作权限
    </p>
    <el-table
      :data="tableData"
      fit
      border
      highlight-current-row
      style="width: 100%"
    >
      <el-table-column label="Id" width="180">
        <template slot-scope="scope">
          <span style="margin-left: 10px">{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="项目名" width="180">
        <template slot-scope="scope">
          <span style="margin-left: 10px">{{ scope.row.projectname }}</span>
        </template>
      </el-table-column>

      <el-table-column label="用户组" width="800">
        <template slot-scope="scope">
          <span style="margin-left: 10px">{{ scope.row.groupname }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="200">
        <template slot-scope="scope">
          <el-button size="mini" @click="updatep(scope.row)">修改</el-button>
          <el-button
            size="mini"
            type="danger"
            @click="handleDelete(scope.row.id)"
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>
    <el-button
      style="margin: 20px"
      type="success"
      size="mini"
      @click="addProject"
      >添加项目名</el-button
    >
    <el-dialog
      :close-on-click-modal="false"
      :visible.sync="dialogFormVisible"
      width="60%"
      title="项目管理"
    >
      <el-form :model="form">
        <el-form-item label="项目名">
          <el-input
            v-model="form.projectname"
            width="200"
            auto-complete="off"
          />
        </el-form-item>

        <el-form-item
          style="display: inline-block;width: 300px"
          label="参与者："
        >
          <el-select v-model="form.groupname" placeholder="参与者">
            <el-option
              v-for="(item, index) in usergroups"
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
import { getUserGroupNames } from "@/api/group";
import {
  getProjectName,
  addProjectName,
  updateProjectName,
  deleteProjectName
} from "@/api/project";
import { deepClone } from "@/utils";
export default {
  name: "Addproject",
  data() {
    return {
      dialogFormVisible: false,
      form: {
        projectname: "",
        groupname: "",
        id: 0
      },
      usergroups: [],
      formLabelWidth: "120px",
      tableData: []
    };
  },
  activated() {
    this.getproject();
    this.getusergroup();
  },
  created() {
    this.getusergroup();
    this.getproject();
  },
  methods: {
    getusergroup() {
      getUserGroupNames().then(resp => {
        this.usergroups = resp.data.usergroupnames;
      });
    },
    getproject() {
      getProjectName().then(resp => {
        if (resp.data.code === 0) {
          this.tableData = resp.data.projectlist;
        } else {
          this.$message.error(resp.data.msg);
        }
      });
    },
    addProject() {
      this.form.id = 0;
      this.form.projectname = "";
      this.form.groupname = "";
      this.dialogFormVisible = true;
    },
    handleDelete(id) {
      this.$confirm("此操作将关闭bug, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(() => {
          deleteProjectName(id).then(_ => {
            const fl = this.tableData.length;
            for (let i = 0; i < fl; i++) {
              if (this.tableData[i].id === id) {
                this.tableData.splice(i, 1);
                break;
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
    },
    updatep(row) {
      this.form = row;
      this.dialogFormVisible = true;
    },
    confirm() {
      this.dialogFormVisible = false;
      if (this.form.projectname === "") {
        this.$message.success("至少选择一个名称");
      }
      if (this.form.groupname === "") {
        this.$message.success("请选择用户组");
      }
      if (this.form.id <= 0) {
        addProjectName(this.form).then(resp => {
          this.form.id = resp.data.id;
          this.tableData.push(deepClone(this.form));
          this.$message.success("添加成功");
        });
      } else {
        updateProjectName(this.form).then(resp => {
          if (resp.data.id === 0) {
            this.$message.warning("存在项目名");
            return;
          }
          const fl = this.tableData.length;
          for (let i = 0; i < fl; i++) {
            if (this.tableData[i].id === this.form.id) {
              this.tableData[i] = this.form;
              break;
            }
          }
          this.$message.success("更新成功");
        });
      }
    },
    cancel() {
      this.dialogFormVisible = false;
      this.form.name = "";
      this.form.id = 0;
    }
  }
};
</script>

<style scoped></style>
