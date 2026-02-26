<template>
  <div class="q-pa-md app-container">
    <!-- 提示 -->
    <q-banner class="bg-warning text-white q-mb-md rounded-borders">
      版本管理，有些可能是 App 的下载地址或者是网页的地址，有一个是备用的
    </q-banner>

    <!-- 添加按钮 -->
    <q-btn color="positive" outline label="添加版本" class="q-mb-md" @click="add" />

    <!-- 表格 -->
    <q-table :rows="list" :columns="columns" row-key="id" flat bordered :loading="listLoading">
      <!-- 自定义列：日期 -->
      <template #body-cell-createtime="props">
        <q-td :props="props">
          {{ parseTime(props.row.createtime) }}
        </q-td>
      </template>

      <!-- 自定义列：项目名 -->
      <template #body-cell-pid="props">
        <q-td :props="props">
          {{ toName(props.row.pid, projectMap) }}
        </q-td>
      </template>

      <!-- 自定义列：操作 -->
      <template #body-cell-action="props">
        <q-td :props="props">
          <q-btn
            flat
            size="sm"
            color="positive"
            label="修改"
            @click="handleModifyStatus(props.row)"
          />
          <q-btn
            v-if="props.row.status !== 'draft'"
            flat
            size="sm"
            label="删除"
            @click="handleRemove(props.row, 'draft')"
          />
        </q-td>
      </template>
    </q-table>

    <!-- 弹窗：版本编辑 -->
    <q-dialog v-model="dialogFormVisible" persistent>
      <q-card style="width: 60vw; max-width: 90vw">
        <q-card-section>
          <div class="text-h6">版本管理</div>
        </q-card-section>

        <q-card-section>
          <q-form class="q-gutter-md">
            <q-select
              v-model="form.pid"
              :options="projects"
              label="项目名"
              option-value="id"
              option-label="name"
              emit-value
              map-options
            />
            <q-input v-model="form.name" label="版本号" maxlength="50" />
            <q-input
              v-model="text"
              type="textarea"
              label="备注"
              placeholder="请输入多行内容..."
              :rows="10"
              autogrow
              maxlength="500"
              counter
            />
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
import { ref, reactive } from 'vue';

/* ---------- 数据 ---------- */
const listLoading = ref(false);
const list = ref<Version[]>([]); // 表格数据
const dialogFormVisible = ref(false);
const form = reactive<Version>({
  id: 0,
  pid: null,
  name: '',
  urlone: '',
  urltwo: '',
});

const text = ref(''); // 备注文本

/* 项目下拉 */
const projects = ref<Project[]>([
  { id: 1, name: 'App-Android' },
  { id: 2, name: 'Web-Admin' },
]);
const projectMap = ref<Record<number, string>>({
  1: 'App-Android',
  2: 'Web-Admin',
});

/* 表格列定义 */
const columns = [
  { name: 'id', label: 'ID', align: 'center' as const, field: 'id' },
  { name: 'createtime', label: '日期', align: 'center' as const, field: 'createtime' },
  { name: 'pid', label: '项目名', align: 'center' as const, field: 'pid' },
  { name: 'name', label: '版本号', align: 'center' as const, field: 'name' },
  { name: 'urlone', label: '地址一', align: 'center' as const, field: 'row => row.urlone' },
  { name: 'urltwo', label: '地址二', align: 'center' as const, field: 'row => row.urltwo' },
  { name: 'action', label: '操作', align: 'center' as const, field: '' },
];

/* ---------- 方法 ---------- */
function parseTime(ts: number | string) {
  // 简化：用 dayjs 或 date-fns 均可
  return new Date(ts).toLocaleString();
}
function toName(pid: number, map: Record<number, string>) {
  return map[pid] || '';
}

function add() {
  form.id = 0;
  form.pid = null;
  form.name = '';
  form.urlone = '';
  form.urltwo = '';
  dialogFormVisible.value = true;
}
function handleModifyStatus(row: Version) {
  Object.assign(form, row);
  dialogFormVisible.value = true;
}
function handleRemove(row: Version, status: string) {
  console.log('删除', row.id, status);
  // TODO 调接口
}
function cancel() {
  dialogFormVisible.value = false;
}
function confirm() {
  console.log('保存版本', form);
  // TODO 调接口
  cancel();
}

/* ---------- 类型 ---------- */
interface Version {
  id: number;
  pid: number | null;
  name: string;
  urlone: string;
  urltwo: string;
  createtime?: number | string;
  status?: string;
}
interface Project {
  id: number;
  name: string;
}
</script>

<style scoped>
.app-container {
  max-width: 1200px;
}
</style>
