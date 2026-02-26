<template>
  <div class="q-pl-md">
    <!-- 提示 -->
    <q-banner class="bg-warning text-white q-mb-md rounded-borders">
      正常的，一般至少有个测试环境，和一个生产环境
    </q-banner>

    <!-- 表格 -->
    <q-table
      :rows="tableData"
      :columns="columns"
      row-key="id"
      flat
      bordered
      style="max-height: 250px"
    >
      <!-- 自定义单元格：ID -->
      <template #body-cell-id="props">
        <q-td :props="props">
          <span class="q-ml-sm">{{ props.row.id }}</span>
        </q-td>
      </template>

      <!-- 自定义单元格：操作 -->
      <template #body-cell-action="props">
        <q-td :props="props">
          <q-btn flat size="sm" color="primary" label="修改" @click="updatep(props.row)" />
          <q-btn flat size="sm" color="negative" label="删除" @click="handleDelete(props.row.id)" />
        </q-td>
      </template>
    </q-table>

    <!-- 添加按钮 -->
    <q-btn class="q-mt-md" color="positive" size="sm" label="添加环境" @click="handleAdd" />

    <!-- 弹窗：编辑/新增 -->
    <q-dialog v-model="dialogFormVisible" persistent>
      <q-card style="width: 60vw; max-width: 90vw">
        <q-card-section>
          <div class="text-h6">运行环境</div>
        </q-card-section>

        <q-card-section>
          <q-input v-model="form.name" label="环境名" maxlength="50" clearable />
        </q-card-section>

        <q-card-actions align="right">
          <q-btn flat label="取消" @click="cancel" />
          <q-btn color="primary" label="确定" @click="confirm" />
        </q-card-actions>
      </q-card>
    </q-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue';
import type { Status } from 'src/api/types/type';
import { getEnvName, addEnvName, deleteEnvName } from 'src/api/env';
import type { KeyName } from 'src/types/response';
import { useQuasar } from 'quasar';

const $q = useQuasar();
/* ---------- 数据 ---------- */
const tableData: KeyName[] = reactive([]);

const dialogFormVisible = ref(false);
const form: Status = reactive({ id: 0, name: '' });

/* 表格列定义 */
const columns = [
  { name: 'id', label: 'ID', align: 'left' as const, field: 'id' },
  { name: 'name', label: '环境名', align: 'left' as const, field: 'name' },
  { name: 'action', label: '操作', align: 'left' as const, field: '' },
];

/* ---------- 方法 ---------- */

async function envList() {
  await getEnvName().then((resp) => {
    console.log(resp.data.data);
    Object.assign(tableData, resp.data.data);
  });
}
function handleAdd() {
  form.name = '';
  dialogFormVisible.value = true;
}
function updatep(row: Status) {
  form.name = row.name;
  dialogFormVisible.value = true;
}
async function handleDelete(id: number) {
  await deleteEnvName(id).then(() => {
    for (let i = 0; i < tableData.length; i++) {
      if ((tableData[i] as KeyName).id == id) {
        tableData.splice(i, 1);
      }
    }
  });
  $q.notify({ type: 'positive', message: '已删除', position: 'top' });
}
function cancel() {
  dialogFormVisible.value = false;
}
async function confirm() {
  console.log('保存环境', form.name);
  await addEnvName(form).then((resp) => {
    tableData.push({ id: resp.data.id, name: form.name });
  });
  $q.notify({ type: 'positive', message: '已删除', position: 'top' });
  cancel();
}

await envList();
</script>

<style scoped>
/* 自定义样式 */
</style>
