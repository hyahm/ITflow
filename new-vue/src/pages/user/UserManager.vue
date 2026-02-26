<template>
  <div class="q-pa-md">
    <q-banner class="bg-orange-1 text-brown q-mb-md rounded-borders">
      <template v-slot:avatar>
        <q-icon name="info" color="brown" />
      </template>
      管理员能修改所有信息，管理者只能修改下一级的信息或者创建下一级的用户账号
    </q-banner>

    <q-table
      :rows="userlist"
      :columns="columns"
      row-key="id"
      :loading="listLoading"
      flat
      bordered
      binary-state-sort
      style="width: 100%"
    >
      <template v-slot:body-cell-createtime="props">
        <q-td :props="props">
          {{ props.row.createtime }}
        </q-td>
      </template>

      <template v-slot:body-cell-position="props">
        <q-td :props="props">
          {{ toJobName(props.row.jid) }}
        </q-td>
      </template>

 

      <template v-slot:body-cell-actions="props">
        <q-td :props="props" class="q-gutter-xs">
          <q-btn color="primary" size="sm" label="修改密码" @click="handleResetPwd()" unelevated />
          <q-btn
            color="warning"
            size="sm"
            label="更改信息"
            @click="handleChangeInfo(props.row)"
            unelevated
          />
          <q-btn
            color="negative"
            size="sm"
            label="删除"
            @click="handleRemove(props.row.id)"
            unelevated
          />
          <q-btn
            v-if="props.row.disable === 1"
            color="positive"
            size="sm"
            label="启用"
            @click="handleDisable(props.row)"
            unelevated
          />
          <q-btn
            v-else
            color="negative"
            size="sm"
            label="禁用"
            @click="handleDisable(props.row)"
            unelevated
          />
        </q-td>
      </template>
    </q-table>

    <q-dialog v-model="dialogVisible" persistent transition-show="scale" transition-hide="scale">
      <q-card style="min-width: 400px">
        <q-card-section>
          <div class="text-h6">编辑用户信息</div>
        </q-card-section>

        <q-card-section class="q-pt-none">
          <q-form ref="postForm" class="q-gutter-md">
            <q-input v-model="form.nickname" label="昵称" outlined dense />
            <q-input v-model="form.realname" label="真实姓名" outlined dense />
            <q-input v-model="form.email" label="邮箱" type="email" outlined dense />
            <q-select
              v-model="form.position_id"
              :options="positionOptions"
              label="职位"
              option-value="id"
              option-label="name"
              emit-value
              map-options
              outlined
              dense
            />
          </q-form>
        </q-card-section>

        <q-card-actions align="right" class="q-pr-md q-pb-md">
          <q-btn flat label="取消" color="grey" v-close-popup @click="cancel" />
          <q-btn flat label="确定" color="primary" @click="HandlerUpdate" />
        </q-card-actions>
      </q-card>
    </q-dialog>


        <q-dialog v-model="changePasswordFormdialogVisible" persistent transition-show="scale" transition-hide="scale">
       <q-card class="q-pa-md form-container">
    <!-- 间距替代原 div 占位 -->
    <div class="q-mb-sm" style="height: 20px"></div>

    <!-- 旧密码输入项 -->
     <div class="q-mb-md">
      <label class="text-sm q-mb-xs block">旧密码：</label>
      <q-input
        v-model="changePasswordForm.oldpassword"
        maxlength="100"
        placeholder="请输入旧密码"
        type="password"
        clearable
        style="width: 60%"
        dense
      />
    </div>
  

    <!-- 新密码输入项 -->
     <div class="q-mb-md">
      <label class="text-sm q-mb-xs block">新密码：</label>
      <q-input
        v-model="changePasswordForm.newpassword"
        maxlength="100"
        placeholder="请输入新密码"
        type="password"
        clearable
        style="width: 60%"
        dense
      />
    </div>
 

    <!-- 确认新密码输入项 -->
       <div class="q-mb-md">
      <label class="text-sm q-mb-xs block">确认新密码：</label>
      <q-input
        v-model="changePasswordForm.repassword"
        maxlength="100"
        placeholder="请输入新密码"
        type="password"
        clearable
        style="width: 60%"
        dense
      />
    </div>
  

    <!-- 按钮区域 -->
    <div class="q-mt-md">
      <q-btn
        label="修改"
        color="positive"
        outlined
        class="q-ml-sm"
        @click="changepwd"
      />
      <q-btn
        label="关闭"
        color="warning"
        outlined
        class="q-ml-sm"
        @click="clean"
      />
    </div>
  </q-card>
    </q-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue';
import { useQuasar, date } from 'quasar';
import { userList, resetPwd, updateUser, userRemove, userDisable } from 'src/api/user';
import type { KeyName } from 'src/types/response';
import { getPositionManager } from 'src/api/position';
import type { ChangePasswod, User } from 'src/api/types/type';

const $q = useQuasar();

const listLoading = ref(false);
const positionlist = ref<KeyName[]>([]);
const positionMap = reactive(new Map<number, string>());
const positionOptions = ref<KeyName[]>([]);

const form = reactive<User>({
  id: 0,
  nickname: '',
  password: '',
  email: '',
  headimg: '',
  realname: '',
  disable: false,
  position_id: 0,
});

const userlist = ref<User[]>([]);
const dialogVisible = ref(false);

const columns = [
  { name: 'id', label: 'ID', field: 'id', align: 'center' as const },
  {
    name: 'created',
    label: '日期',
    field: 'created',
    align: 'center' as const,
    format: (val: string) => {
      if (!val) return '-';
      return date.formatDate(new Date(val), 'YYYY-MM-DD HH:mm:ss');
    },
  },
  {
    name: 'realname',
    label: '真实姓名',
    field: 'realname',
    align: 'center' as const,
  },
  {
    name: 'nickname',
    label: '昵称',
    field: 'nickname',
    align: 'center' as const,
  },
  {
    name: 'position',
    label: '职位',
    field: 'jid',
    align: 'center' as const,
  },
  { name: 'email', label: '邮箱', field: 'email', align: 'left' as const },
  {
    name: 'status',
    label: '状态',
    field: 'disable',
    align: 'center' as const,
    format: (val: boolean) => {
      return val?'禁用':'启用'
    },
  },
  {
    name: 'actions',
    label: '操作',
    field: '',
    align: 'center' as const,
    width: '400px',
  },
];

const toJobName = (id: number) => {
  return positionMap.get(id) || '-';
};

const getUserList = async () => {
  listLoading.value = true;
  // try {
    const position = await getPositionManager();
    positionlist.value = position.data.data;
    positionOptions.value = position.data.data;

    positionMap.clear();
    for (const v of positionlist.value) {
      positionMap.set(v.id, v.name);
    }

    const resp = await userList();
    userlist.value = resp.data.data;
  // } catch () {
  //   $q.notify({
  //     color: 'negative',
  //     message: '获取用户列表失败',
  //     position: 'top',
  //     timeout: 2000,
  //   });
  // } finally {
    listLoading.value = false;
  // }
};

const cancel = () => {
  dialogVisible.value = false;
};

const HandlerUpdate = async () => {
  await updateUser(form);
  $q.notify({
    color: 'positive',
    message: '修改成功',
    position: 'top',
    timeout: 2000,
  });
  dialogVisible.value = false;
  await getUserList();
};

const handleChangeInfo = (row: User) => {
  console.log(row)
  Object.assign(form, row);
  dialogVisible.value = true;
};

const handleRemove = async (id: number) => {
  await userRemove(id);
  $q.notify({
    color: 'positive',
    message: '删除成功',
    position: 'top',
    timeout: 2000,
  });
  await getUserList();
};

const handleDisable = async (row: User) => {
  await userDisable(row.id).then(async () => {
    $q.notify({
      color: 'positive',
      message: row.disable ? '已启用' : '已禁用',
      position: 'top',
      timeout: 2000,
    });
    await getUserList();
  });
};

const handleResetPwd = () => {
  changePasswordFormdialogVisible.value = true
  // const pwdData: ChangePasswod = {
  //   oldpassword: '',
  //   newpassword: '',
  // };
  // await resetPwd(pwdData).then(() => {
  //   $q.notify({
  //     color: 'positive',
  //     message: '密码修改成功',
  //     position: 'top',
  //     timeout: 2000,
  //   });
  // });
};

await getUserList();


const changePasswordFormdialogVisible = ref(false)

// 表单数据模型
const changePasswordForm = ref({
  oldpassword: '', // 旧密码
  newpassword: '', // 新密码
  repassword: ''   // 确认新密码
});

// 修改密码方法（需自行补充接口逻辑）
const changepwd =async  () => {
  // 1. 简单校验（可根据需求扩展）
  if (!changePasswordForm.value.oldpassword) {
    return alert('请输入旧密码');
  }
  if (!changePasswordForm.value.newpassword) {
    return alert('请输入新密码');
  }
  if (changePasswordForm.value.newpassword !== changePasswordForm.value.repassword) {
    return alert('两次输入的新密码不一致');
  }

  // 2. 调用修改密码接口（示例，需替换为你的实际接口）
  console.log('提交修改密码请求：', changePasswordForm.value);
  // 示例：await api.changePassword(form.value);
  const pwdData: ChangePasswod = {
    oldpassword: '',
    newpassword: '',
  };
  await resetPwd(pwdData).then(() => {
    $q.notify({
      color: 'positive',
      message: '密码修改成功',
      position: 'top',
      timeout: 2000,
    });
  });
};

// 清空表单方法
const clean = () => {
  changePasswordForm.value = {
    oldpassword: '',
    newpassword: '',
    repassword: ''
  };
  changePasswordFormdialogVisible.value = false
};
</script>

<style scoped>
.q-table th {
  font-weight: 500;
}
</style>
