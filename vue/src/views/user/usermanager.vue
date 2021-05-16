<template>
  <div class="app-container">
    <div class="filter-container">
      <p class="warn-content">
        管理员能修改所有信息，管理者只能修改下一级的信息或者创建下一级的用户账号
      </p>
    </div>

    <el-table
      :key="tableKey"
      v-loading="listLoading"
      :data="userlist"
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

      <el-table-column label="真实姓名" width="110px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.realname }}</span>
        </template>
      </el-table-column>
      <el-table-column label="昵称" width="110px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.nickname }}</span>
        </template>
      </el-table-column>

      <el-table-column label="职位" width="110px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.position }}</span>
        </template>
      </el-table-column>

      <el-table-column label="邮箱" class-name="status-col" width="200">
        <template slot-scope="scope">
          <span>{{ scope.row.email }}</span>
          <!--<el-tag :type="scope.row.status | statusFilter">{{scope.row.status}}</el-tag>-->
        </template>
      </el-table-column>
      <el-table-column label="状态" class-name="status-col" width="60">
        <template slot-scope="scope">
          <span v-if="scope.row.disable == 0">启用</span>
          <span v-else>禁用</span>
          <!--<el-tag :type="scope.row.status | statusFilter">{{scope.row.status}}</el-tag>-->
        </template>
      </el-table-column>
      <el-table-column
        label="操作"
        align="center"
        width="400"
        class-name="small-padding fixed-width"
      >
        <template slot-scope="scope">
          <el-button
            type="primary"
            size="mini"
            @click="handleResetPwd(scope.row)"
            >修改密码</el-button
          >
          <el-button
            size="mini"
            type="danger"
            @click="handlePermission(scope.row)"
            >更改信息
          </el-button>
          <el-button size="mini" type="danger" @click="handleRemove(scope.row)"
            >删除
          </el-button>
          <el-button
            v-if="scope.row.disable == 1"
            size="mini"
            type="danger"
            @click="handleDisable(scope.row)"
            >启用
          </el-button>
          <el-button
            v-else
            size="mini"
            type="danger"
            @click="handleDisable(scope.row)"
            >禁用
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog
      :close-on-click-modal="false"
      :visible.sync="dialogVisible"
      :before-close="handleClose"
      title="提示"
      width="30%"
    >
      <el-form ref="postForm" />
      <!--<el-button type="success" round @click="HandlerAddGroup">添加部门</el-button>-->
      <el-form ref="postForm">
        <el-form-item label="昵称">
          <el-input v-model="form.nickname" />
        </el-form-item>
        <el-form-item label="真实姓名">
          <el-input v-model="form.realname" />
        </el-form-item>
        <el-form-item label="邮箱">
          <el-input v-model="form.email" />
        </el-form-item>

        <el-form-item label="职位：">
          <el-select v-model="form.position" placeholder="Select">
            <el-option
              v-for="(role, index) in positionlist"
              :key="index"
              :label="role"
              :value="role"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="cancel">取 消</el-button>
        <el-button type="primary" @click="HandlerUpdateRoles">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import {
  userList,
  resetPwd,
  updateUser,
  userRemove,
  userDisable
} from "@/api/user";
import { getRoleGroup } from "@/api/rolegroup";
import { getStatusGroupName } from "@/api/statusgroup";
import { getPositions } from "@/api/position";

export default {
  name: "Usermanager",
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: "success",
        draft: "info",
        deleted: "danger"
      };
      return statusMap[status];
    }
  },
  data() {
    return {
      uid: -1,
      rolelist: [],
      dialogVisible: false,
      rolegrouplist: [],
      statusgrouplist: [],
      positionlist: [],
      tableKey: 0,
      userlist: [],
      admin: false,
      form: {
        id: -1,
        name: ""
      },
      statusgroup: [],
      listLoading: false,
      sortOptions: [
        { label: "ID Ascending", key: "+id" },
        { label: "ID Descending", key: "-id" }
      ]
    };
  },
  activated() {
    this.getuserList();
  },
  created() {
    this.getuserList();
    this.getgrouplist();
  },
  methods: {
    getgrouplist() {
      getRoleGroup().then(resp => {
        this.rolegrouplist = resp.data.rolelist;
      });
      getStatusGroupName().then(resp => {
        this.statusgrouplist = resp.data.names;
      });
      getPositions().then(resp => {
        this.positionlist = resp.data.positions;
      });
    },

    cancel() {
      this.dialogVisible = false;
    },
    HandlerUpdateRoles() {
      updateUser(this.form).then(_ => {
        const l = this.userlist.length;
        for (let i = 0; i < l; i++) {
          if (this.userlist[i].id === this.form.id) {
            this.userlist[i].role = this.form.name;
          }
        }
        this.$message.success("修改成功");
      });
      this.dialogVisible = false;
    },
    handleClose() {
      this.dialogVisible = false;
    },
    getuserList() {
      userList().then(resp => {
        this.userlist = resp.data.userlist;
      });
    },
    handlePermission(row) {
      this.form = row;
      this.dialogVisible = true;
    },
    handleRemove(row) {
      this.$confirm("此操作将关闭bug, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(() => {
          userRemove(row.id).then(_ => {
            const l = this.userlist.length;
            for (let i = 0; i < l; i++) {
              if (this.userlist[i].id === row.id) {
                this.userlist.splice(i, 1);
              }
            }
            this.$message.warning("删除成功");
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除"
          });
        });
    },
    handleDisable(row) {
      userDisable(row.id).then(_ => {
        const l = this.userlist.length;
        for (let i = 0; i < l; i++) {
          if (this.userlist[i].id === row.id) {
            this.userlist[i].disable = Math.abs(this.userlist[i].disable - 1);
            break;
          }
        }
      });
    },
    handleResetPwd(row) {
      this.$prompt("请输入密码", "提示", {
        cancelButtonText: "取消",
        confirmButtonText: "确定"
      })
        .then(({ value }) => {
          const data = {
            id: row.id,
            newpassword: value
          };
          resetPwd(data).then(_ => {
            this.$message({
              type: "success",
              message: "你的密码是: " + value
            });
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "取消输入"
          });
        });
    }
  }
};
</script>
