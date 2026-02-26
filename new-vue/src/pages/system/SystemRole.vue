<template>
  <div style="padding-left: 20px">
<q-banner class="bg-orange-1 text-brown q-mb-md rounded-borders">
        <template v-slot:avatar>
          <q-icon name="warning" color="brown" />
        </template>
        选择可以操作的页面组， 操作的角色由开发者决定, 如果查看的权限没有， 那么后面的所有权限都被无视
      </q-banner>
    <!-- 添加按钮 -->
    <div class="q-my-md">
      <q-btn color="positive" outline label="添加角色" @click="handleAdd" />
    </div>


    <q-table
      :rows="roles"
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

    <q-dialog
      v-model="dialogVisible"
      persistent
      :maximized="false"
      style="width: 60vw; max-width: 90vw"
    >
      <q-card>
        <q-card-section class="row items-center">
          <div class="text-h6">提示</div>
          <q-space />
          <q-btn icon="close" flat round dense @click="handleClose" />
        </q-card-section>

        <q-card-section>
          <q-form ref="postForm" class="q-gutter-md">
            <!-- 角色组名 -->
            <q-input
              v-model="form.name"
              label="角色名"
              maxlength="100"
              placeholder="请输入角色组名"
              clearable
            />

            <!-- 动态权限组 -->
            <div v-for="(item, index) in roles" :key="index">
              <!-- <div class="env-title q-mb-xs">{{ item.info }}</div> -->
              <!-- <q-option-group
                v-model="item.value"
                :options="item.label.map((c) => ({ label: c, value: c }))"
                type="checkbox"
                class="rolegroup"
                dense
              /> -->
              <!-- {{ item }} -->
            </div>
            <div class="q-pa-md">
              <div class="q-gutter-sm">
                <q-option-group dense :options="options" type="checkbox" v-model="form.perm_ids" />
              </div>
            </div>
          </q-form>
        </q-card-section>

        <q-card-actions align="right">
          <q-btn flat label="取消" @click="dialogVisible = false" />
          <q-btn color="primary" label="确定" @click="HandlerAddGroup" />
        </q-card-actions>
      </q-card>
    </q-dialog>
  </div>
</template>

<script setup lang="ts">
import type { RequestRole, Role } from 'src/api/types/type';
import { useQuasar } from 'quasar';
// import { roleList, addRole, editRole, removeRole, getRoleGroupPerm, getRoles } from 'src/api/role';
import { getPagePerms, addRole,roleList ,delRole} from 'src/api/role';
import { reactive, ref } from 'vue';
import type { KeyName, Option } from 'src/types/response';
// import { docAddUser } from '../../api/doc';
const $q = useQuasar();
const columns = [
  { name: 'id', label: 'Id', align: 'center' as const, field: 'id', sortable: true },
  { name: 'name', label: '角色名', align: 'center' as const, field: 'name', sortable: true },
  { name: 'action', label: '操作', align: 'center' as const, field: '' },
];

const options: Option[] = reactive([]);
// const roles = reactive([]);
const dialogVisible = ref(false);
// const list = reactive([]);
const form: RequestRole = reactive({
  id: 0,
  name: '',
  perm_ids: [],
});
const roles: KeyName[] = reactive([]);
// const checkboxGroup1 = reactive([]);
// const templateperm: Perm[] = reactive([]);
// let pages: Role[] = reactive([]);
// let permlist: Perm[] = reactive([]);
// const permValue = reactive(new Map<string, number>());
// const defaultPerm = ['read', 'create', 'update', 'delete'];
async function getRoleList() {
  await roleList().then((resp) => {
    console.log(resp.data)
    Object.assign(roles, resp.data.data);
    // for (const v of resp.data.data as Role[]) {
    //   options.push({ label: v.info, value: v.id });
    // }
   
    
    // dialogVisible.value = false;
    
  });
}

async function pageList() {
  await getPagePerms().then((resp) => {
    Object.assign(options, []);
    for (const v of resp.data.data as Role[]) {
      options.push({ label: v.info, value: v.id });
    }
    console.log(options)
    dialogVisible.value = false;
    
  });
}

await pageList();
await getRoleList();

function handleUpdate(row: Role) {
  form.id = row.id;
  form.name = row.name;
  dialogVisible.value = true
}

async function handleDelete(id: number) {
  // TODO 调删除接口
  await delRole(id).then(() => {
  const idx = roles.findIndex((v) => v.id === id);
    if (idx > -1) roles.splice(idx, 1);
    
  })
  $q.notify({ type: 'positive', message: '删除成功' });
  
}

// 添加按钮
function handleAdd() {
  form.id = 0;
  form.name = '';
  form.perm_ids = [];
  // permlist = templateperm;
  
  dialogVisible.value = true;
}

function handleClose() {
  dialogVisible.value = false;
}
// 提交添加的按钮
async function HandlerAddGroup() {
  if (form.name == '') {
    $q.notify({ type: 'warning', message: '名称不能为空', position: 'top' });
    return
  }
  await addRole(form).then((resp) => {
     roles.unshift({id: resp.data.id, name: form.name})
    $q.notify({ type: 'positive', message: '创建成功', position: 'top' });
  });
  handleClose()
}
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
.warn-content {
  background-color: burlywood;
}
</style>

<style type="text/css">
.form-container > .name {
  padding-right: 30px;
  width: 250px !important;
}

</style>
