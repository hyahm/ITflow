<template>
  <div>
    <el-table :data="tableData" height="250" style="width: 100%">
      <el-table-column label="Id" width="180">
        <template slot-scope="scope">
          <span style="margin-left: 10px">{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="项目名" width="180">
        <template slot-scope="scope">
          <span style="margin-left: 10px">{{ scope.row.osname }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作">
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
    <el-button @click="addp">添加运行平台</el-button>
    <el-dialog
      :close-on-click-modal="false"
      :visible.sync="dialogFormVisible"
      title="平台管理"
    >
      <el-form :model="form">
        <el-form-item :label-width="formLabelWidth" label="平台">
          <el-input v-model="form.osname" auto-complete="off" />
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
import { getOSName, addOSName, updateOSName, deleteOSName } from "@/api/os";
export default {
  name: "Runos",
  data() {
    return {
      dialogFormVisible: false,
      form: {
        osname: "",
        delivery: false,
        id: 0
      },
      formLabelWidth: "120px",
      tableData: []
    };
  },
  created() {
    this.getosname();
  },
  methods: {
    getosname() {
      getOSName().then(resp => {
        this.tableData = resp.data.oslist;
      });
    },
    addp() {
      this.form.id = 0;
      this.dialogFormVisible = true;
    },
    handleDelete(id) {
      deleteOSName(id).then(resp => {
        const fl = this.tableData.length;
        for (let i = 0; i < fl; i++) {
          if (this.tableData[i].id === id) {
            this.tableData.splice(i, 1);
            break;
          }
        }
      });
    },
    updatep(row) {
      this.dialogFormVisible = true;
      this.form.id = row.id;
      this.form.osname = row.osname;
    },
    confirm() {
      this.dialogFormVisible = false;
      if (this.form.id === 0) {
        addOSName(this.form.osname).then(resp => {
          this.tableData.push({
            id: resp.data.id,
            osname: this.form.osname
          });
        });
      } else {
        updateOSName(this.form).then(resp => {
          const fl = this.form.length;
          for (let i = 0; i < fl; i++) {
            if (this.tableData[i].id === this.form.id) {
              this.tableData[i].osname = this.form.osname;
              break;
            }
          }
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
