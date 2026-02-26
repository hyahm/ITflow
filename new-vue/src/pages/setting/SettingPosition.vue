<template>
  <div style="padding: 20px">
    <!-- 提示 -->
      <q-table
      :rows="tableData"
      :columns="columns"
      row-key="id"
      flat
      bordered
      style="max-height: 500px"
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
    <q-btn class="q-mt-md" color="positive" label="添加职位" @click="addImportant" />

    <!-- 弹窗：编辑/新增 -->
    <q-dialog v-model="dialogFormVisible" persistent>
      <q-card flat bordered>
        <q-card-section>
          <div class="text-h6 q-mb-md">添加职位</div>

          <!-- 提示 -->
          <q-banner class="bg-warning text-white q-mb-md rounded-borders">
            上级职位拥有管理下级的权限；普通管理者创建的职位默认为自己的下级。
          </q-banner>

          <!-- 表单 -->
          <q-form ref="formRef" class="q-gutter-md" @submit="onSubmit">
            <!-- 职位名 -->
            <q-input
              v-model="form.name"
              label="职位名称 *"
              :rules="[(val) => !!val || '请输入职位名称']"
              maxlength="50"
              clearable
            />

            <!-- 管理层级 -->
            <q-option-group
              v-model="form.level"
              :options="levelOptions"
              type="radio"
              label="管理层级 *"
              :rules="[(val: string) => val !== null || '请选择层级']"
            />

            <!-- 角色组 -->
            <q-select
              v-model="form.role_id"
              :options="rolesOptions"
              label="角色 *"
              option-value="id"
              option-label="name"
              emit-value
              map-options
              clearable
              :rules="[(val) => !!val || '请选择角色组']"
            />

            <!-- 从属于 -->
            <q-select
              v-model="form.hypo"
              :options="managerOptions"
              label="从属于（上级职位）"
              option-value="id"
              option-label="name"
              emit-value
              map-options
              clearable
            >
              <template #prepend>
                <q-icon name="account_tree" />
              </template>
            </q-select>

            <!-- 按钮 -->
            <div>
              <q-btn label="保存" type="submit" color="primary" />
              <q-btn
                label="重置"
                type="reset"
                color="secondary"
                flat
                class="q-ml-sm"
                @click="onReset"
              />
               <q-btn
                label="取消"
                color="secondary"
                flat
                class="q-ml-sm"
                @click="close"
              />
              </div>
          </q-form>
        </q-card-section>
      </q-card>
    </q-dialog>
 
  </div>
</template>

<script setup lang="ts">
import { reactive,ref } from 'vue';
import { addPosition, getPositionManager,positionsList } from 'src/api/position';
import { roleList } from 'src/api/role';
import type { KeyName } from 'src/types/response';
import type { Position } from 'src/api/types/type';
import { useQuasar } from 'quasar';
const $q = useQuasar();

/* -------------- 数据 -------------- */

const form: Position = reactive({
  id: 0,
  name: '',
  level: 0,
  role_id: null,
  hypo: null,
});
const columns = [
  { name: 'id', label: 'Id', align: 'center' as const, field: 'id', sortable: true },
  { name: 'name', label: '职位名', align: 'center' as const, field: 'name', sortable: true },
  { name: 'action', label: '操作', align: 'center' as const, field: '' },
];
/* 下拉选项 */
const levelOptions = [
  { label: '管理者', value: 1 },
  { label: '普通员工', value: 0 },
];


const rolesOptions: KeyName[] = reactive([]);
const managerOptions: KeyName[] = reactive([]);
const dialogFormVisible = ref(false)
async function GetRoles() {
  await roleList().then((resp) => {
    Object.assign(rolesOptions, resp.data.data);
  });
}
const tableData:KeyName[] = reactive([])
function handleUpdate(row: KeyName) {
  form.id = row.id;
  form.name = row.name;
  dialogFormVisible.value = true;
}

function addImportant() {
   dialogFormVisible.value = true;
}


async function list() {
   await positionsList().then(resp => {
    console.log(resp.data)
    Object.assign(tableData,resp.data.data)
  })
}

await list()

function handleDelete(id: number) {
  // TODO 调删除接口
  const idx = tableData.findIndex((v) => v.id === id);
  if (idx > -1) tableData.splice(idx, 1);
  $q.notify({ type: 'positive', message: '删除成功' });
}
/* -------------- 方法 -------------- */
async function onSubmit() {
  await addPosition(form).then((resp) => {
    tableData.unshift({
      id: resp.data.id,
      name: form.name,
    })
    $q.notify({ type: 'positive', message: '创建成功', position: 'top' });
  });
  close()
}

async function initManager() {
  await getPositionManager().then((resp) => {
    Object.assign(managerOptions, resp.data.data);
  });
}

function close() {
   dialogFormVisible.value = false;
}

function onReset() {
  form.name = '';
  form.level = 0;
  form.role_id = null;
  form.hypo = null;
}
await GetRoles();
await initManager();
/* -------------- 初始化 -------------- */
</script>

<style scoped></style>
