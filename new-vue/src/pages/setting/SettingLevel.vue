<template>
  <div class="q-pl-md">
    <!-- 提示 -->
    <q-banner class="bg-warning text-white q-mb-md rounded-borders"> bug优先级别 </q-banner>

    <!-- 表格 -->
    <q-table
      :rows="tableData"
      :columns="columns"
      row-key="id"
      flat
      bordered
      style="max-height: 250px"
    >
      <template #body-cell-id="props">
        <q-td :props="props">
          {{ props.row.id }}
        </q-td>
      </template>

      <template #body-cell-name="props">
        <q-td :props="props">
          {{ props.row.name }}
        </q-td>
      </template>

      <template #body-cell-action="props">
        <q-td :props="props">
          <q-btn flat size="sm" color="primary" label="修改" @click="handleUpdate(props.row)" />
          <q-btn flat size="sm" color="negative" label="删除" @click="handleDelete(props.row.id)" />
        </q-td>
      </template>
    </q-table>

    <!-- 添加按钮 -->
    <q-btn class="q-mt-md" color="positive" label="添加优先级别" @click="addstatus" />

    <!-- 弹窗：优先级别编辑 -->
    <q-dialog v-model="dialogFormVisible" persistent>
      <q-card style="width: 60vw; max-width: 90vw">
        <q-card-section>
          <div class="text-h6">优先级别</div>
        </q-card-section>

        <q-card-section>
          <q-form class="q-gutter-md">
            <q-input v-model="form.name" label="优先级别" maxlength="50" />
          </q-form>
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
import { reactive, ref, onMounted } from 'vue';
import { useQuasar } from 'quasar';
import { getLevels, addLevel, updateLevel } from 'src/api/level';
import type { KeyName } from 'src/types/response';

/* ---------- 数据 ---------- */
const $q = useQuasar();

const columns = [
  { name: 'id', label: 'ID', align: 'left' as const, field: 'id', sortable: true },
  { name: 'name', label: '级别名', align: 'left' as const, field: 'name', sortable: true },
  { name: 'action', label: '操作', align: 'left' as const, field: '' },
];

const tableData = ref<KeyName[]>([]);
const statuslist = ref<string[]>([]);

const dialogFormVisible = ref(false);
const form = reactive<KeyName>({ id: 0, name: '' });

/* ---------- 方法 ---------- */
async function getstatus() {
  const resp = await getLevels();
  tableData.value = resp.data.data;
  statuslist.value = tableData.value.map((v) => v.name);
}

async function confirm() {
  if (form.id === 0) {
    // 新增
    const resp = await addLevel(form);
    tableData.value.push({ id: resp.data.id, name: form.name });
    $q.notify({ type: 'positive', message: '添加成功' });
  } else {
    // 更新
    await updateLevel(form);
    const idx = tableData.value.findIndex((v) => v.id === form.id);
    if (idx > -1) (tableData.value[idx] as KeyName).name = form.name;
    $q.notify({ type: 'positive', message: '更新成功' });
  }
  dialogFormVisible.value = false;
}

function cancel() {
  dialogFormVisible.value = false;
}

function addstatus() {
  form.id = 0;
  form.name = '';
  dialogFormVisible.value = true;
}

function handleUpdate(row: KeyName) {
  form.id = row.id;
  form.name = row.name;
  dialogFormVisible.value = true;
}

function handleDelete(id: number) {
  // TODO 调删除接口
  const idx = tableData.value.findIndex((v) => v.id === id);
  if (idx > -1) tableData.value.splice(idx, 1);
  $q.notify({ type: 'positive', message: '删除成功' });
}

/* ---------- 生命周期 ---------- */
onMounted(() => getstatus());
</script>

<style scoped></style>
