<template>
  <div style="padding-left: 20px">
    <p class="warn-content">
      职位，在管理的时候，上级有管理下级的权限，级别只有普通员工和管理者，管理者也可以被其他管理者管理（从属于），一个员工只能被一个管理者管理
    </p>
    <div>
      <el-button type="success" plain style="margin: 20px" @click="addposition"
        >添加职位</el-button
      >
    </div>
    <el-table :data="tableData" height="600" style="width: 100%">
      <el-table-column label="Id" width="180">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="职位名" width="180">
        <template slot-scope="scope">
          <span>{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="职位级别" width="180">
        <template slot-scope="scope">
          <span>{{ scope.row.level | level }}</span>
        </template>
      </el-table-column>
      <el-table-column label="从属于" width="180">
        <template slot-scope="scope">
          <span>{{ scope.row.hyponame }}</span>
        </template>
      </el-table-column>

      <el-table-column label="角色组" width="100px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.rolegroup }}</span>
        </template>
      </el-table-column>
      <el-table-column label="状态组" width="110px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.statusgroup }}</span>
        </template>
      </el-table-column>
      <el-table-column width="200" label="操作">
        <template slot-scope="scope">
          <el-button size="mini" @click="handleUpdate(scope.row)"
            >修改</el-button
          >
          <el-button size="mini" type="danger" @click="handleDelete(scope.row)"
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>

    <el-dialog
      :close-on-click-modal="false"
      :visible.sync="dialogFormVisible"
      title="职位管理"
    >
      <el-form>
        <el-form-item label="管理层：">
          <el-radio-group v-model="form.level">
            <el-radio :label="levelone">管理者</el-radio>
            <el-radio :label="leveltwo">普通员工</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="职位名：" label-width="100">
          <el-input v-model="form.name" />
        </el-form-item>

        <el-form-item style="margin-bottom: 40px;" label="状态组:">
          <el-select
            v-model="form.statusgroup"
            class="filter-item"
            style="width: 130px"
          >
            <el-option
              v-for="(s, index) in statusgroups"
              :key="index"
              :label="s"
              :value="s"
            />
          </el-select>
        </el-form-item>
        <el-form-item style="margin-bottom: 40px;" label="角色组:">
          <el-select
            v-model="form.rolegroup"
            class="filter-item"
            style="width: 130px"
          >
            <el-option
              v-for="(role, index) in rolegroups"
              :key="index"
              :label="role"
              :value="role"
            />
          </el-select>
        </el-form-item>
        <!-- <el-form style="margin-top: 10px"> -->
        <!--从属于哪个管理者-->
        <el-form-item label="从属于:">
          <el-select v-model="form.hyponame" clearable placeholder="Select">
            <el-option
              v-for="(hypo, index) in manager"
              :key="index"
              :label="hypo"
              :value="hypo"
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
import {
  addPosition,
  updatePosition,
  delPosition,
  getHypos,
  PositionsList
} from "@/api/position";
import { getRoleGroup } from "@/api/rolegroup";
import { getStatusGroupName } from "@/api/statusgroup";
export default {
  name: "Position",
  filters: {
    level: function(value) {
      switch (value) {
        case 1:
          return "管理者";
        default:
          return "普通员工";
      }
    }
  },
  data() {
    return {
      changeName: "",
      tableData: [],
      statuslist: [],
      statusgroups: [],
      rolegroups: [],
      dialogFormVisible: false,
      status: "",
      levelone: 1,
      leveltwo: 2,
      manager: [],
      form: {
        id: 0,
        name: "",
        level: 2,
        hyponame: "",
        rolegroup: "",
        statusgroup: ""
      }
    };
  },
  created() {
    this.getlist();
    this.getrolegroups();
    this.getstatusgroups();
  },
  methods: {
    getrolegroups() {
      getRoleGroup().then(resp => {
        this.rolegroups = resp.data.data;
      });
    },
    handleGetHypos(id) {
      getHypos(id).then(resp => {
        for (let i = 0; i < resp.data.length; i++) {
          this.manager.push(resp.data.hypos[i].name);
        }
      });
    },
    getlist() {
      PositionsList().then(resp => {
        this.tableData = resp.data.data;
        for (let i = 0; i < this.tableData.length; i++) {
          if (this.tableData[i].level === 1) {
            this.manager.push(this.tableData[i].name);
          }
        }
      });
    },
    confirm() {
      if (this.form.id === 0) {
        addPosition(this.form).then(resp => {
          this.tableData.push({
            id: resp.data.id,
            name: this.form.name,
            hyponame: this.form.hyponame,
            level: this.form.level
          });

          if (this.form.level === 1) {
            this.manager.push(this.form.name);
          }
        });
      } else {
        updatePosition(this.form).then(resp => {
          const l = this.tableData.length;
          for (let i = 0; i < l; i++) {
            if (this.tableData[i].id === this.form.id) {
              this.tableData[i].name = this.form.name;
              break;
            }
            for (let i = 0; i < this.manager.length; i++) {
              if (this.manager[i] === this.changeName) {
                this.manager[i] = this.form.name;
                return;
              }
            }
          }
        });
      }
      this.dialogFormVisible = false;
    },
    cancel() {
      this.dialogFormVisible = false;
    },
    getstatusgroups() {
      getStatusGroupName().then(resp => {
        this.statusgroups = resp.data.names;
      });
    },
    handleDelete(row) {
      this.$confirm("此操作将关闭bug, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(() => {
          delPosition(row.id).then(resp => {
            const l = this.tableData.length;
            for (let i = 0; i < l; i++) {
              if (this.tableData[i].id === row.id) {
                this.tableData.splice(i, 1);
                break;
              }
            }
            for (let i = 0; i < this.manager.length; i++) {
              if (this.manager[i] === row.name) {
                this.manager.splice(i, 1);
                break;
              }
            }
            this.$message.success(resp.data.msg);
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除"
          });
        });
    },
    addposition() {
      this.dialogFormVisible = true;
      this.form.id = 0;
      this.form.level = 2;
      this.form.hyponame = "";
      this.form.name = "";
      this.form.rolegroup = "";
      this.form.statusgroup = "";
    },
    handleUpdate(row) {
      this.form = row;
      this.changeName = row.name;

      this.dialogFormVisible = true;
      this.handleGetHypos(row.id);
    }
  }
};
</script>

<style scoped></style>
