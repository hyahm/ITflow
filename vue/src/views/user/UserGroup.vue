<template>
  <div class="group" style="padding-left: 20px">
    <p class="warn-content">只有创建才能删除， 管理员和创建者可以查看编辑</p>
    <el-table :data="list" height="250" style="width: 100%">
      <el-table-column label="Id" width="180">
        <template slot-scope="scope">
          <span style="margin-left: 10px">{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="组名" width="180">
        <template slot-scope="scope">
          <span style="margin-left: 10px">{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="成员" width="500">
        <template slot-scope="scope">
          <span style="margin-left: 10px">{{ scope.row.uids | toname }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button size="mini" @click="handleUpdate(scope.row)">修改</el-button>
          <el-button size="mini" type="danger" @click="handleDelete(scope.row.id)"
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>
    <el-button
      style="margin-top: 10px; margin-left: 10px"
      type="success"
      @click="handleAdd"
      >添加组</el-button
    >

    <el-dialog
      :close-on-click-modal="false"
      :visible.sync="dialogFormVisible"
      title="平台管理"
    >
      <el-form :model="form">
        <el-form-item label="组名">
          <el-input v-model="form.name" auto-complete="off" />
        </el-form-item>
        <el-form-item label="用户">
          <el-select v-model="form.uids" multiple placeholder="请选择">
            <el-option
              v-for="value in users"
              :key="value.id"
              :label="value.name"
              :value="value.id"
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
  getAllUserKeyName,
  getUserGroups,
  createUserGroups,
  deleteUserGroups,
  updateUserGroup,
} from "@/api/usermanager/usergroup";
import { getUserKeyName } from "@/api/get";
let that;
export default {
  name: "Group",
  filters: {
    toname(uids) {
      let names = [];
      for (let v of uids) {
        names.push(that.userMap.get(v));
      }
      return names.join(", ");
    },
  },
  created() {
    this.getuserkey();
    this.getgroup();
    that = this;
  },
  data() {
    return {
      dialogFormVisible: false,
      list: [],
      form: {
        id: 0,
        name: "",
        uids: [],
      },
      users: [],
      userMap: new Map(),
    };
  },
  activated() {
    this.getuserkey();
  },
  methods: {
    getuserkey() {
      getAllUserKeyName().then((resp) => {
        this.users = resp.data.data;
        for (let v of this.users) {
          this.userMap.set(v.id, v.name);
        }
      });
    },
    getgroup() {
      getUserGroups().then((resp) => {
        this.list = resp.data.data;
      });
    },
    handleAdd() {
      this.form = {
        id: 0,
        name: "",
        uids: [],
      };

      this.dialogFormVisible = true;
    },
    confirm() {
      if (this.form.id > 0) {
        updateUserGroup(this.form).then((_) => {
          const l = this.list.length;
          for (let i = 0; i < l; i++) {
            if (this.list[i].id === this.form.id) {
              this.list[i] = this.form;
            }
          }
          this.$message.success("修改成功");
        });
      } else {
        createUserGroups(this.form).then((resp) => {
          this.list.push({
            id: resp.data.id,
            name: this.form.name,
            uids: this.form.uids,
          });
          this.$message.success("添加用户组成功");
        });
      }
      this.dialogFormVisible = false;
    },
    cancel() {
      this.form = {
        name: "",
        uids: [],
      };
      this.dialogFormVisible = false;
    },
    handleUpdate(row) {
      this.dialogFormVisible = true;
      this.form.id = row.id;
      this.form.uids = row.uids;
      this.form.name = row.name;
    },
    handleDelete(id) {
      this.$confirm("此操作将关闭bug, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          deleteUserGroups(id).then((_) => {
            for (let i in this.list) {
              if (this.list[i].id === id) {
                this.list.splice(i, 1);
                break;
              }
            }
            this.$message.success("删除成功");
            return;
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
