<template>
  <div class="group" style="padding-left: 20px">
    <p class="warn-content">只有创建才能删除， 管理员和创建者可以查看编辑</p>
    <el-table :data="list" height="250" style="width: 100%">
      <el-table-column label="Id" width="180">
        <template v-slot="scope">
          <span style="margin-left: 10px">{{ scope.row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="组名" width="180">
        <template v-slot="scope">
          <span style="margin-left: 10px">{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="成员" width="500">
        <template v-slot="scope">
          <span style="margin-left: 10px">{{ toname(scope.row.uids) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template v-slot="scope">
          <el-button size="mini" @click="handleUpdate(scope.row)">修改</el-button>
          <el-button size="mini" type="danger" @click="handleDelete(scope.row.id)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-button style="margin-top: 10px; margin-left: 10px" type="success" @click="handleAdd"
      >添加组</el-button
    >

    <el-dialog :close-on-click-modal="false" v-model:visible="dialogFormVisible" title="平台管理">
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
      <template v-slot:footer>
        <div class="dialog-footer">
          <el-button @click="cancel">取 消</el-button>
          <el-button type="primary" @click="confirm">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import {
  getAllUserKeyName,
  getUserGroups,
  createUserGroups,
  deleteUserGroups,
  updateUserGroup,
} from 'src/api/usergroup';
// import { getUserKeyName } from 'src/api/get';
import { reactive, ref } from 'vue';
// import variables from 'styles/variables.scss';
import type { KeyName } from 'src/types/response';
// import { DefaultValue } from '../../api/types/type';
import type { UserGroup } from 'src/api/types/type';
// import { laAws } from '@quasar/extras/line-awesome';

const dialogFormVisible = ref(false);
const list: UserGroup[] = reactive([]);
let form: UserGroup = reactive({
  id: 0,
  name: '',
  uids: [],
  uid: 0,
});
let users: KeyName[] = reactive([]);
const userMap = reactive(new Map());
async function getuserkey() {
  await getAllUserKeyName().then((resp) => {
    users = resp.data.data;
    for (const v of users) {
      userMap.set(v.id, v.name);
    }
  });
}
async function getgroup() {
  await getUserGroups().then((resp) => {
    Object.assign(list, resp.data.data);
  });
}
const handleAdd = () => {
  form = {
    id: 0,
    name: '',
    uids: [],
    uid: 0,
  };

  dialogFormVisible.value = true;
};
async function confirm() {
  if (form.id > 0) {
    await updateUserGroup(form).then((response) => {
      if (response.data.code === 0) {
        for (let i = 0; i < list.length; i++) {
          if ((list[i] as UserGroup).id == form.id) {
            (list[i] as UserGroup).uids = form.uids;
            break;
          }
        }
        // this.$message.success('修改成功');
      }
      // else {
      //   this.$message.error(response.data.msg);
      // }
    });
  } else {
    await createUserGroups(form).then((resp) => {
      list.push({
        id: resp.data.id,
        name: form.name,
        uids: form.uids,
        uid: 0,
      });
      // this.$message.success('添加用户组成功');
    });
  }
  dialogFormVisible.value = false;
}
const cancel = () => {
  form = {
    name: '',
    uids: [],
    id: 0,
    uid: 0,
  };
  dialogFormVisible.value = false;
};
const handleUpdate = (row: UserGroup) => {
  dialogFormVisible.value = true;
  form.id = row.id;
  form.uids = row.uids;
  form.name = row.name;
};

const toname = (uids: number[]) => {
  const names = [];
  for (const v of uids) {
    names.push(userMap.get(v));
  }
  return names.join(', ');
};
async function handleDelete(id: number) {
  // this.$confirm('此操作将关闭bug, 是否继续?', '提示', {
  //   confirmButtonText: '确定',
  //   cancelButtonText: '取消',
  //   type: 'warning',
  // })
  //   .then(() => {
  await deleteUserGroups(id).then(() => {
    for (let i = 0; i < list.length; i++) {
      if ((list[i] as UserGroup).id === id) {
        list.splice(i, 1);
        break;
      }
    }
    // this.$message.success('删除成功');
    return;
  });
  // })
  // .catch(() => {
  //   this.$message({
  //     type: 'info',
  //     message: '已取消删除',
  //   });
  // });
}

await getuserkey();
await getgroup();
</script>
