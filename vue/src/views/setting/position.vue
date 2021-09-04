<template>
  <div style="padding-left: 20px">
    <p class="warn-content">
      职位，在管理的时候，上级有管理下级的权限，普通管理者创建的职位默认为自己的下级
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
          <span>{{ scope.row.hypo | toManager(managerMap) }}</span>
        </template>
      </el-table-column>

      <el-table-column label="角色组" width="100px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.rolegroup | toRoleName(roleMap) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="状态组" width="110px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.statusgroup | toStatusName(statusMap) }}</span>
        </template>
      </el-table-column>
      <el-table-column width="200" label="操作">
        <template slot-scope="scope">
          <el-button size="mini" @click="handleUpdate(scope.row)">修改</el-button>
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

        <el-form-item style="margin-bottom: 40px" label="状态组:">
          <el-select v-model="form.statusgroup" class="filter-item" style="width: 130px">
            <el-option
              v-for="statusgroup in statusgroups"
              :key="statusgroup.id"
              :label="statusgroup.name"
              :value="statusgroup.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item style="margin-bottom: 40px" label="角色组:">
          <el-select v-model="form.rolegroup" class="filter-item" style="width: 130px">
            <el-option
              v-for="role in rolegroups"
              :key="role.id"
              :label="role.name"
              :value="role.id"
            />
          </el-select>
        </el-form-item>
        <!-- <el-form style="margin-top: 10px"> -->
        <!--从属于哪个管理者-->
        <el-form-item label="从属于:" v-if="isadmin">
          <el-select v-model="form.hypo" clearable placeholder="Select">
            <el-option
              v-for="item in manager"
              :key="item.id"
              :label="item.name"
              :value="item.id"
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
  PositionsList,
  getManagerKeyName,
} from "@/api/position";
import { isAdmin } from "@/api/get";
import { getRoleGroup } from "@/api/rolegroup";
import { getStatusGroupName } from "@/api/statusgroup";
export default {
  name: "Position",
  filters: {
    level: function (value) {
      switch (value) {
        case 1:
          return "管理者";
        default:
          return "普通员工";
      }
    },
    toStatusName(sid, statusMap) {
      return statusMap.get(sid);
    },
    toRoleName(rid, roleMap) {
      return roleMap.get(rid);
    },
    toManager(mid, managerMap) {
      return managerMap.get(mid);
    },
  },
  data() {
    return {
      changeName: "",
      tableData: [],
      statuslist: [],
      statusgroups: [],
      statusMap: new Map(),
      roleMap: new Map(),
      rolegroups: [],
      dialogFormVisible: false,
      status: "",
      levelone: 1,
      leveltwo: 2,
      isadmin: false,
      manager: [],
      managerMap: new Map(),
      form: {
        id: 0,
        name: "",
        level: 2,
        hypo: undefined,
        rolegroup: undefined,
        statusgroup: undefined,
      },
    };
  },
  created() {
    this.checkAdmin();
    this.init();
  },
  methods: {
    checkAdmin() {
      isAdmin().then((resp) => {
        this.isadmin = resp.data.admin;
      });
    },
    async init() {
      const role = await getRoleGroup();
      this.rolegroups = role.data.data;
      for (let v of this.rolegroups) {
        this.roleMap.set(v.id, v.name);
      }

      const status = await getStatusGroupName();
      this.statusgroups = status.data.data;

      for (let v of this.statusgroups) {
        this.statusMap.set(v.id, v.name);
      }

      const mg = await getManagerKeyName();
      this.manager = mg.data.data;
      for (let v of this.manager) {
        this.managerMap.set(v.id, v.name);
      }
      this.getlist();
    },

    getlist() {
      PositionsList().then((resp) => {
        this.tableData = resp.data.data;
      });
    },
    confirm() {
      if (!this.form.name) {
        this.$message.error('必须添加职位名称')
        return
      }
       if (!this.form.statusgroup) {
        this.$message.error('必须添加状态组')
        return
      }
       if (!this.form.rolegroup) {
        this.$message.error('必须添加角色组')
        return
      }
      if (this.form.id === 0) {
        addPosition(this.form).then((resp) => {
          this.tableData.push({
            id: resp.data.id,
            name: this.form.name,
            hypo: this.form.hypo,
            level: this.form.level,
            rolegroup: this.form.rolegroup,
            statusgroup: this.form.statusgroup
          });

          if (this.form.level === 1) {
            this.manager.push(this.form.name);
            this.managerMap.set(resp.data.id, this.form.name);
          }
          this.$message.success('添加成功')

        });
      } else {
        updatePosition(this.form).then((resp) => {
          this.tableData.map((m) => {
            if (m.id === this.form.id) {
              m = this.form;
            }
            return m;
          });

          if (this.form.level === 1) {
            for (let i = 0; i < this.manager.length; i++) {
              if (this.manager[i] === this.changeName) {
                this.manager[i] = this.form.name;
                return;
              }
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

    handleDelete(row) {
      this.$confirm("此操作将关闭bug, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          delPosition(row.id).then((resp) => {
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
            message: "已取消删除",
          });
        });
    },
    addposition() {
      this.dialogFormVisible = true;
      this.form.id = 0;
      this.form.level = 2;
      this.form.hypo = undefined;
      this.form.name = "";
      this.form.rolegroup = undefined;
      this.form.statusgroup = undefined;
    },
    handleUpdate(row) {
      this.form = row;
      this.changeName = row.name;
      this.dialogFormVisible = true;
    },
  },
};
</script>

<style scoped></style>
