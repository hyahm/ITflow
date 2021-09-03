<template>
  <div style="padding-left: 20px">
    <p class="warn-content">bug重要性</p>
    <el-table :data="tableData" height="250" style="width: 100%">
      <el-table-column label="Id" width="180">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="重要程度名" width="180">
        <template slot-scope="scope">
          <span>{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column width="200" label="操作">
        <template slot-scope="scope">
          <el-button size="mini" @click="handleUpdate(scope.row)">修改</el-button>
          <el-button size="mini" type="danger" @click="handleDelete(scope.row.id)"
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>
    <div>
      <el-button type="success" plain style="margin: 20px" @click="addstatus"
        >添加重要程度</el-button
      >
    </div>
    <el-dialog
      :close-on-click-modal="false"
      :visible.sync="dialogFormVisible"
      title="重要程度管理"
    >
      <el-form>
        <el-form-item label="重要程度名">
          <el-input v-model="form.name" auto-complete="off" />
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
  getImportant,
  addImportant,
  delImportant,
  updateImportant,
} from "@/api/important";
export default {
  name: "Important",
  data() {
    return {
      tableData: [],
      defaultstatus: "",
      dialogFormVisible: false,
      status: "",
      form: {
        id: 0,
        name: "",
      },
    };
  },
  activated() {
    this.getimportant();
  },
  created() {
    this.getimportant();
  },
  methods: {
    getimportant() {
      getImportant().then((resp) => {
        this.tableData = resp.data.data;
      });
    },
    confirm() {
      if (this.form.id === 0) {
        addImportant(this.form).then((resp) => {
          this.tableData.push({
            id: resp.data.id,
            name: this.form.name,
          });
          this.$message.success('添加成功')
        });
      } else {
        updateImportant(this.form).then((resp) => {
          const l = this.tableData.length;
          for (let i = 0; i < l; i++) {
            if (this.tableData[i].id === this.form.id) {
              this.tableData[i].name = this.form.name;
            }
          }
          this.$message.success('更新成功')
        });
      }
      this.dialogFormVisible = false;
    },
    cancel() {
      this.dialogFormVisible = false;
    },
    handleDelete(id) {
      this.$confirm("此操作将关闭bug, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          delImportant(id).then((resp) => {
            const l = this.tableData.length;
            for (let i = 0; i < l; i++) {
              if (this.tableData[i].id === id) {
                this.tableData.splice(i, 1);
              }
            }
            this.$message.success("删除成功");
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除",
          });
        });
    },
    addstatus() {
      this.dialogFormVisible = true;
      this.form.id = 0;
      this.form.name = "";
    },
    handleUpdate(row) {
      this.dialogFormVisible = true;
      this.form.id = row.id;
      this.form.name = row.name;
    },
  },
};
</script>

<style scoped></style>
