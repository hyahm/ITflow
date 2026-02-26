<template>
  <div style="padding-left: 20px">
 <q-banner class="bg-orange-1 text-brown q-mb-md rounded-borders">
        <template v-slot:avatar>
          <q-icon name="warning" color="brown" />
        </template>
       bug的所有状态流程， 比如从新建->解决中->测试->完成 等
      </q-banner>
    <q-table title="Treats" :rows="tableData" :columns="columns" row-key="name" />
    <!-- <el-table :data="tableData" height="250" style="width: 100%">
      <el-table-column label="Id" width="180">
        <template slot-scope="scope">
          <span>{{ row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column label="状态名" width="180">
        <template slot-scope="scope">
          <span>{{ scope.name }}</span>
        </template>
      </el-table-column>
      <el-table-column width="200" label="操作">
        <template slot-scope="scope">
          <el-button size="mini" @click="handleUpdate(scope.row)">修改</el-button>
          <el-button size="mini" type="danger" @click="handleDelete(scope.row.id)">删除</el-button>
        </template>
      </el-table-column>
    </el-table> -->
    <div class="q-ma-md">
      <q-btn outline color="positive" label="添加bug状态" @click="addstatus" />
    </div>

    <!-- 状态管理弹窗 -->
    <q-dialog v-model="dialogFormVisible" persistent>
      <q-card style="min-width: 350px">
        <q-card-section class="row items-center q-pb-none">
          <div class="text-h6">状态管理</div>
          <q-space />
          <q-btn icon="close" flat round dense @click="cancel" />
        </q-card-section>

        <q-card-section>
          <q-input v-model="form.name" label="状态" autofocus @keyup.enter="confirm" />
        </q-card-section>

        <q-card-actions align="right" class="q-px-md q-pb-md">
          <q-btn flat label="取消" @click="cancel" />
          <q-btn color="primary" label="确定" @click="confirm" />
        </q-card-actions>
      </q-card>
    </q-dialog>
  </div>
</template>

<script setup lang="ts">
import type { QTableProps } from 'quasar';
import { getStatusList, addStatus, updateStatus } from 'src/api/status';
import type { KeyName } from 'src/types/response';
import { reactive, ref } from 'vue';
let tableData: KeyName[] = reactive([]);
// let statuslist = reactive([]);
const dialogFormVisible = ref(false);
const form: KeyName = {
  id: 0,
  name: '',
};

const columns: QTableProps['columns'] = [
  {
    name: 'id',
    required: true,
    label: 'ID',
    align: 'left',
    field: 'id',
    sortable: true,
  },
  {
    name: 'name',
    required: true,
    label: '名称',
    align: 'left',
    field: 'name',
    sortable: true,
  },
];
async function getstatus() {
  await getStatusList().then((resp) => {
    console.log(resp.data.data);
    tableData = resp.data.data;
  });
}
async function confirm() {
  if (form.id === 0) {
    await addStatus(form).then((resp) => {
      tableData.push({
        id: resp.data.id,
        name: form.name,
      });
      // this.$message.success('添加成功');
    });
  } else {
    await updateStatus(form).then(() => {
      const l = tableData.length;
      for (let i = 0; i < l; i++) {
        if ((tableData[i] as KeyName).id === form.id) {
          (tableData[i] as KeyName).name = form.name;
        }
      }
      // this.$message.success('更新成功');
    });
  }
  dialogFormVisible.value = false;
}
function cancel() {
  dialogFormVisible.value = false;
}
// async function handleDelete(id: number) {
//   await removeStatus(id).then(() => {
//     const l = tableData.length;
//     for (let i = 0; i < l; i++) {
//       if ((tableData[i] as KeyName).id === id) {
//         tableData.splice(i, 1);
//       }
//     }
//   });
// }
function addstatus() {
  dialogFormVisible.value = true;
  form.id = 0;
  form.name = '';
}
// function handleUpdate(row: KeyName) {
//   dialogFormVisible.value = true;
//   form.id = row.id;
//   form.name = row.name;
// }

await getstatus();
</script>
