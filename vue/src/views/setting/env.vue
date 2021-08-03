<template>
  <div style="padding-left: 20px">
    <p class="warn-content">
      正常的，一般至少有个测试环境，和一个生产环境
    </p>
    <el-table :data="tableData" height="250" style="width: 100%">
      <el-table-column label="Id" width="180">
        <template slot-scope="scope">
          <span style="margin-left: 10px">{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="环境名" width="180">
        <template slot-scope="scope">
          <span style="margin-left: 10px">{{ scope.row.envname }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="180">
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
      @click="handleAdd"
      >添加环境</el-button
    >
    <el-dialog
      :close-on-click-modal="false"
      :visible.sync="dialogFormVisible"
      width="60%"
      title="运行环境"
    >
      <el-form :model="form">
        <el-form-item label="环境名">
          <el-input v-model="form.envname" auto-complete="off" />
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
  getEnvName,
  addEnvName,
  updateEnvName,
  deleteEnvName
} from "@/api/env";

export default {
  name: "Env",
  data() {
    return {
      dialogFormVisible: false,
      form: {
        envname: "",
        delivery: false,
        id: 0
      },
      formLabelWidth: "120px",
      tableData: []
    };
  },
  created() {
    this.getenvname();
  },
  methods: {
    getenvname() {
      getEnvName().then(resp => {
        this.tableData = resp.data.envlist;
      });
    },
    handleAdd() {
      this.form.id = 0;
      this.form.envname = "";
      this.dialogFormVisible = true;
    },
    handleDelete(id) {
      this.$confirm("此操作将关闭bug, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(() => {
          deleteEnvName(id).then(resp => {
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
      this.dialogFormVisible = true;
      this.form.id = row.id;
      this.form.envname = row.envname;
    },
    confirm() {
      this.dialogFormVisible = false;
      if (this.form.id === 0) {
        addEnvName(this.form.envname).then(resp => {
          this.tableData.push({
            id: resp.data.id,
            envname: this.form.envname
          });
          this.$message.success("添加成功");
        });
      } else {
        updateEnvName(this.form).then(resp => {
          const fl = this.tableData.length;
          for (let i = 0; i < fl; i++) {
            if (this.tableData[i].id === this.form.id) {
              this.tableData[i].envname = this.form.envname;
              break;
            }
          }
          this.$message.success("修改成功");
        });
      }
    },
    cancel() {
      this.dialogFormVisible = false;
      this.form.envname = "";
      this.form.id = 0;
    }
  }
};
</script>

<style scoped></style>
