<template>
  <div class="q-pl-md">
    <!-- 提示 -->
    <q-banner class="bg-warning text-white q-mb-md rounded-borders">
      如果删除的此项目的某参与者，在他任务里面还是会显示，只不过没操作权限
    </q-banner>

    <!-- 表格 -->
    <q-table :rows="tableData" :columns="columns" row-key="id" flat bordered>
      <template #body-cell-id="props">
        <q-td :props="props">
          <span class="q-ml-sm">{{ props.row.id }}</span>
        </q-td>
      </template>

      <template #body-cell-name="props">
        <q-td :props="props">
          <span class="q-ml-sm">{{ props.row.name }}</span>
        </q-td>
      </template>

      <template #body-cell-ugid="props">
        <q-td :props="props">
          {{ toUserGroupName(props.row.ugid) }}
        </q-td>
      </template>

      <template #body-cell-action="props">
        <q-td :props="props">
          <q-btn flat size="sm" color="primary" label="修改" @click="updateProject(props.row)" />
          <q-btn flat size="sm" color="negative" label="删除" @click="handleDelete(props.row.id)" />
        </q-td>
      </template>
    </q-table>

    <!-- 添加按钮 -->
    <q-btn class="q-mt-md" color="positive" label="添加项目名" @click="addProject" />

    <!-- 弹窗：项目编辑 -->
    <q-dialog v-model="dialogFormVisible" persistent>
      <q-card style="width: 60vw; max-width: 90vw">
        <q-card-section>
          <div class="text-h6">项目管理</div>
        </q-card-section>

        <q-card-section>
          <q-form class="q-gutter-md">
            <q-input v-model="form.name" label="项目名" maxlength="50" />

            <q-select
              filled
              v-model="form.uids"
              label="成员"
              option-value="id"
              option-label="name"
              multiple
              emit-value
              map-options
              :options="userOptions"
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
import { reactive, ref } from 'vue';
import { useQuasar } from 'quasar';
import {
  getProjectName,
  addProjectName,
  updateProjectName,
  deleteProjectName,
} from 'src/api/project';
import type { KeyName } from 'src/types/response';
import type { ProjectRequsest } from 'src/api/types/type';
import { getUserKeyName } from 'src/api/get';

const columns = [
  { name: 'id', label: 'ID', align: 'left' as const, field: 'id' },
  { name: 'name', label: '项目名', align: 'left' as const, field: 'name' },
  {
    name: 'users',
    label: '用户成员',
    align: 'left' as const,
    field: (row: ProjectList) => row.user_info?.map((u) => u.name).join(',') || null,
  },
  { name: 'action', label: '操作', align: 'left' as const, field: '' },
];

interface ProjectList {
  id: number;
  name: string;
  user_info: KeyName[] | null;
}

/* ---------- 数据 ---------- */
const $q = useQuasar();
const dialogFormVisible = ref(false);
// 让 form 是 reactive 对象，后续 v-model 直接改内部属性
const form = reactive<ProjectRequsest>({
  uids: null,
  name: '',
  id: 0,
  uid: 0,
});

const userOptions = reactive<KeyName[]>([]);
const tableData = ref<ProjectList[]>([]);

const usergroupMap = new Map<number, string>();

/* ---------- 方法 ---------- */
function toUserGroupName(id: number) {
  return usergroupMap.get(id) ?? '';
}

async function getusergroup() {
  await getUserKeyName().then((resp) => {
    Object.assign(userOptions, resp.data.data);
    console.log(resp.data.data)
    for (const v of resp.data.data) {
      usergroupMap.set(v.id, v.name);
    }
    form.uids = [];
  });

  // usergroupMap.clear();
  // usergroups.value.forEach((v) => usergroupMap.set(v.id, v.name));
}

async function getproject() {
  await getProjectName().then((resp) => {
    console.log(resp.data.data);
    tableData.value = resp.data.data;
  });

  // tableData.value = resp.data.data;
}

function addProject() {
  form.id = 0;
  form.name = '';
  form.uids = null;
  dialogFormVisible.value = true;
}

async function handleDelete(id: number) {
  await deleteProjectName(id);
  const idx = tableData.value.findIndex((v) => v.id === id);
  if (idx > -1) tableData.value.splice(idx, 1);
  $q.notify({ type: 'positive', message: '删除成功' });
}

function updateProject(row: ProjectList) {
  form.id = row.id;
  form.name = row.name;
  form.uids = [];
  row.user_info?.map((u) => {
    if (form.uids) {
      form.uids.push(u.id);
    }
  });

  dialogFormVisible.value = true;
}

async function confirm() {
  dialogFormVisible.value = false;
  if (form.id <= 0) {
    // 新增
    const resp = await addProjectName(form);
    form.id = resp.data.id;
    const td = { id: form.id, name: form.name, user_info: [] };
    for (const uid of form.uids || []) {
      const user = userOptions.find((u) => u.id === uid);
      if (user) {
        (td.user_info as KeyName[]).push(user);
      }
    }
    tableData.value.push(td);
    $q.notify({ type: 'positive', message: '添加成功' });
  } else {
    // 更新
    const resp = await updateProjectName(form);
    if (resp.data.id === 0) {
      $q.notify({ type: 'warning', message: '存在项目名' });
      return;
    }
    const idx = tableData.value.findIndex((v) => v.id === form.id);
    const td = { id: form.id, name: form.name, user_info: [] };
    for (const uid of form.uids || []) {
      const user = userOptions.find((u) => u.id === uid);
      if (user) {
        (td.user_info as KeyName[]).push(user);
      }
    }
    tableData.value[idx] = td;

    $q.notify({ type: 'positive', message: '更新成功' });
  }
}

function cancel() {
  dialogFormVisible.value = false;
  form.name = '';
  form.id = 0;
  form.uids = null;
}
await getproject();
await getusergroup();
/* ---------- 生命周期 ---------- */
</script>

<style scoped></style>
