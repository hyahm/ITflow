<template>
  <div style="padding-left: 20px">
    <p class="warn-content">增加选择可以改变bug状态的组</p>
    <q-table :rows="list" row-key="name" flat bordered />
    <!-- <el-table
      v-loading="listLoading"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%;padding: 10px"
    >
      <el-table-column label="id" align="center" width="50">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>

      <el-table-column label="项目名" width="100px" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.name }}</span>
        </template>
      </el-table-column>

      <el-table-column label="状态权限" width="500" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.sids | toname(statusMap) }}</span>
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
          <el-button
            type="success"
            size="mini"
            @click="handleRemove(scope.row.id)"
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table> -->
    <el-button style="margin: 20px" type="success" size="mini" @click="handleAdd"
      >添加状态组名</el-button
    >
    <el-dialog
      :close-on-click-modal="false"
      v-model:visible="dialogVisible"
      :before-close="handleClose"
      title="提示"
      width="60%"
    >
      <el-form ref="postForm" class="form-container">
        <el-form-item prop="title" label="状态组名:">
          <el-input v-model="form.name" :maxlength="100" placeholder="请输入状态组名" clearable />
        </el-form-item>
        <el-checkbox-group v-model="form.sids">
          <el-checkbox v-for="status in statuslist" :label="status.id" :key="status.id">{{
            status.name
          }}</el-checkbox>
        </el-checkbox-group>
      </el-form>
      <template v-slot:footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取 消</el-button>
          <el-button type="primary" @click="HandlerAddGroup">确 定</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { getStatus } from 'src/api/get';
import {
  statusGroupList,
  addStatusGroup,
  editStatusGroup,
  // removeStatusGroup,
} from 'src/api/statusgroup';
import type { StatusGroup } from 'src/api/types/type';
import type { KeyName } from 'src/types/response';
import { reactive, ref } from 'vue';

const statuslist = reactive<KeyName[]>([]);
const statusMap = reactive(new Map<number, string>());
const dialogVisible = ref(false);
// let listLoading = ref(false);
const list: StatusGroup[] = reactive([]);
const form: StatusGroup = reactive({
  id: 0,
  name: '',
  sids: [],
});

// function handleEdit(row: StatusGroup) {
//   Object.assign(form, row);
//   dialogVisible.value = true;
// }
async function getlist() {
  await statusGroupList().then((resp) => {
    Object.assign(list, resp.data.data);
  });
}
function handleAdd() {
  Object.assign(form, {
    id: 0,
    name: '',
    sids: [],
  });
  dialogVisible.value = true;
}
// async function handleRemove(id: number) {
//   await removeStatusGroup(id).then((resp) => {
//     const l = list.length;
//     for (let i = 0; i < l; i++) {
//       if ((list[i] as StatusGroup).id === id) {
//         list.splice(i, 1);
//       }
//     }
//   });
// }
async function getstatus() {
  await getStatus().then((resp) => {
    Object.assign(statuslist, resp.data.data);
    for (const v of statuslist) {
      statusMap.set(v.id, v.name);
    }
  });
}
function handleClose() {
  dialogVisible.value = false;
}
async function HandlerAddGroup() {
  if (form.name === '') {
    // this.$message.error('名称不能为空');
    return;
  }
  if (form.id > 0) {
    await editStatusGroup(form).then(() => {
      const l = list.length;
      for (let i = 0; i < l; i++) {
        if ((list[i] as StatusGroup).id === form.id) {
          (list[i] as StatusGroup).name = form.name;
          (list[i] as StatusGroup).sids = form.sids;
        }
      }
      // this.$message.success('修改成功');
    });
  } else {
    await addStatusGroup(form).then((resp) => {
      list.push({
        id: resp.data.id,
        name: form.name,
        sids: form.sids,
      });
      // this.$message.success('添加成功');
    });
  }

  dialogVisible.value = false;
}

// function toname(ids: number[]) {
//   return ids.map((m) => statusMap.get(m)).join(', ');
// }

await getstatus();
await getlist();
</script>

<style scoped type="text/css">
label {
  padding: 10px;
}
</style>
