<template>
  <div style="padding-left: 20px">
    <p class="warn-content">
      选择可以操作的页面组， 操作的角色由开发者决定, 如果查看的权限没有，
      那么后面的所有权限都被无视
    </p>
    <el-table
      v-loading="listLoading"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%; padding: 10px"
    >
      <el-table-column label="id" align="center" width="50">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>

      <el-table-column label="角色组" width="100px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.name }}</span>
        </template>
      </el-table-column>

      <el-table-column
        label="操作"
        align="center"
        width="230"
        class-name="small-padding fixed-width"
      >
        <template slot-scope="scope">
          <el-button type="primary" size="mini" @click="handleEdit(scope.row)"
            >编辑</el-button
          >
          <el-button type="success" size="mini" @click="handleRemove(scope.row.id)"
            >删除</el-button
          >
          <!--<el-button type="danger" size="mini" @click="handleRemove(scope.row)">{{ $t('list.remove') }}</el-button>-->
        </template>
      </el-table-column>
    </el-table>
    <el-button style="margin: 20px" type="success" size="mini" @click.native="handleAdd"
      >添加角色组</el-button
    >
    <el-dialog
      :visible.sync="dialogVisible"
      :before-close="handleClose"
      :close-on-click-modal="false"
      title="提示"
      width="60%"
    >
      <el-form ref="postForm" class="form-container">
        <el-form-item prop="title" label="角色组名:">
          <el-input
            v-model="form.name"
            :maxlength="100"
            placeholder="请角色组名"
            clearable
          />
        </el-form-item>
        <!-- <el-checkbox-group v-model="perm"> -->
        <div v-for="(item, index) in permlist" :key="index">
          <!-- <label class="name">{{ role.name }}</label> -->
          <!-- <el-checkbox v-model="role.checked" :label="role.name" style="width:150px" @change="changeChecked(role)" /> -->
          <div class="env-title">{{ item.info }}</div>
          <el-checkbox-group :key="index" v-model="item.value" class="rolegroup">
            <el-checkbox v-for="city in item.label" :label="city" :key="city">{{
              city
            }}</el-checkbox>
            <!-- <el-checkbox v-model="item" style="width: 50px" label="read" />
          <el-checkbox v-model="item" style="width: 50px" label="add" />
          <el-checkbox v-model="item" style="width: 50px" label="update" />
          <el-checkbox v-model="item" style="width: 50px" label="remove" /> -->
          </el-checkbox-group>
        </div>
        <!-- </el-checkbox-group> -->
        <!--<el-button type="success" round @click="HandlerAddGroup">添加部门</el-button>-->
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="HandlerAddGroup">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import {
  roleList,
  addRole,
  editRole,
  removeRole,
  getRoleGroupPerm,
  getRoles,
} from "@/api/role";
import { deepClone } from "@/utils";
export default {
  name: "RoleGroup",
  data() {
    return {
      roles: [],
      dialogVisible: false,
      listLoading: false,
      list: [],
      form: {
        id: 0,
        name: "",
        rolelist: [],
      },
      checkboxGroup1: [],
      templateperm: [],
      pages: [],
      permlist: [],
      permValue: new Map(),
      defaultPerm: ["read", "create", "update", "delete"],
    };
  },

  created() {
    this.getroles();
    this.getlist();
    this.permValue.set("read", 1);
    this.permValue.set("create", 2);
    this.permValue.set("update", 4);
    this.permValue.set("delete", 8);
    // this.getTemplate();
  },
  methods: {
    getroles() {
      getRoles().then((resp) => {
        this.pages = resp.data.data;
        console.log(this.pages);
        for (let v of this.pages) {
          this.templateperm.push({
            id: 0,
            rid: v.id,
            label: this.defaultPerm,
            value: [],
            info: v.info,
          });
        }
      });
    },

    // getTemplate() {
    //   // 获取模板
    //   getPermTemplate().then((resp) => {
    //     this.templateperm = resp.data.template;
    //   });
    // },
    handleEdit(row) {
      // 请求拿到权限数据
      // 请求权限
      getRoleGroupPerm(row.id).then((resp) => {
        console.log(resp.data);
        this.permlist = resp.data.data;
        this.form.name = row.name;
        this.form.id = row.id;
      });
      // this.form.id = row.id;
      // this.form.name = row.name;
      // this.form.rolelist = row.rolelist;
      this.dialogVisible = true;
    },
    getlist() {
      roleList().then((resp) => {
        this.list = resp.data.data;
      });
    },
    handleAdd() {
      this.form.id = 0;
      this.form.name = "";
      this.permlist = deepClone(this.templateperm);
      this.dialogVisible = true;
    },
    handleRemove(id) {
      this.$confirm("确认关闭？")
        .then(() => {
          removeRole(id).then(() => {
            const l = this.list.length;
            for (let i = 0; i < l; i++) {
              if (this.list[i].id === id) {
                this.list.splice(i, 1);
              }
            }
            this.$message.success("删除成功");
          });
        })
        .catch((_) => {});
    },
    handleClose() {
      this.dialogVisible = false;
    },
    HandlerAddGroup() {
      if (this.form.name.length < 1) {
        this.$message.error("name no be need");
      }
      this.form.rolelist = [];
      for (let v of this.permlist) {
        let pv = 0;
        for (let row of v.value) {
          pv += this.permValue.get(row);
        }

        this.form.rolelist.push({
          id: v.id,
          pv: pv,
          rid: v.rid,
        });
      }

      // 计算每一行的值， 填充到form里面
      // console.log(this.permlist);
      // return
      if (this.form.id > 0) {
        editRole(this.form).then((_) => {
          // 成功后赋值到源数据
          const l = this.list.length;
          for (let i = 0; i < l; i++) {
            if (this.list[i].id === this.form.id) {
              this.list[i].name = this.form.name;
              this.list[i].rolelist = this.form.rolelist;
            }
          }
          this.$message.success("修改成功");
        });
      } else {
        addRole(this.form).then((resp) => {
          this.list.push({
            id: resp.data.id,
            name: this.form.name,
          });
          this.$message.success("添加成功");
        });
      }

      this.dialogVisible = false;
    },
  },
};
</script>

<style scoped type="text/css">
.env-title {
  display: inline-block;
  width: 120px;
  text-align: right;
}
.rolegroup {
  display: inline-block;
}
label {
  padding: 10px;
}
.form-container > .name {
  padding-right: 30px;
  width: 250px !important;
}
</style>

<style type="text/css">
.form-container > .name {
  padding-right: 30px;
  width: 250px !important;
}
</style>
