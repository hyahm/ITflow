<template>
  <div class="add-user">
<q-banner class="bg-orange-1 text-brown q-mb-md rounded-borders">
        <template v-slot:avatar>
          <q-icon name="warning" color="brown" />
        </template>
         状态组是相对于 bug
      管理的状态的，角色组是共享文件夹和接口文档的权限，都是必须项，姓名、邮箱、姓名都必须是唯一值
      </q-banner>
    <q-form class="form-container q-gutter-md">
      <q-input
        v-model="postForm.nickname"
        label="昵称"
        maxlength="100"
        placeholder="姓名首字母"
        clearable
        style="max-width: 60%"
      />

      <q-input
        v-model="postForm.realname"
        label="姓名"
        maxlength="100"
        placeholder="请输入姓名"
        clearable
        style="max-width: 60%"
      />

      <q-input
        v-model="postForm.email"
        type="email"
        label="邮箱"
        maxlength="100"
        placeholder="请输入邮箱"
        autocomplete="off"
        clearable
        style="max-width: 60%"
      />

      <q-input
        v-model="postForm.password"
        type="password"
        label="密码"
        maxlength="100"
        placeholder="请输入密码"
        autocomplete="new-password"
        clearable
        style="max-width: 60%"
      />

      <q-input
        v-model="repassword"
        type="password"
        label="确认密码"
        maxlength="100"
        placeholder="请再次输入密码"
        autocomplete="off"
        clearable
        style="max-width: 60%"
      />

      <q-select
        v-model="postForm.position_id"
        :options="positions"
        label="职位名"
        option-value="id"
        option-label="name"
        emit-value
        map-options
        style="max-width: 60%"
      />

      <div class="row q-gutter-sm q-mt-sm">
        <q-btn color="positive" outline label="添加" @click="addUser" />
        <q-btn color="warning" outline label="清空" @click="clean" />
      </div>
    </q-form>
  </div>
</template>

<script setup lang="ts">
import { createUser } from 'src/api/user';
import { positionsList } from 'src/api/position';
import type { User } from 'src/api/types/type';
import { reactive, ref } from 'vue';
import type { KeyName } from 'src/types/response';
import { useQuasar } from 'quasar';
const $q = useQuasar();
let postForm: User = reactive({
  nickname: '',
  email: '',
  password: '',
  headimg: '',
  realname: '',
  position_id: null,
  id: 0,
  disable: false,
});

const repassword = ref('');
const positions: KeyName[] = reactive([]);

async function addUser() {
  postForm.id = 0;
  const regEmail = /^.*@.+\.[A-Za-zd]{2,5}$/;
  if (!regEmail.test(postForm.email)) {
     $q.notify({ type: 'warning', message: '邮箱不对', position: 'top' });
    return;
  }
  if (postForm.nickname == '') {
     $q.notify({ type: 'warning', message: '昵称不能为空', position: 'top' });
    return;
  }
   if (postForm.realname == '') {
     $q.notify({ type: 'warning', message: '姓名不能为空', position: 'top' });
    return;
  }
  if (repassword.value !== postForm.password) {
     $q.notify({ type: 'warning', message: '2次的密码不对', position: 'top' });
    return;
  }
  if (!postForm.position_id) {
    $q.notify({ type: 'warning', message: '职位没有选', position: 'top' });
    return;
  }

  await createUser(postForm).then(() => {
    $q.notify({ type: 'positive', message: '创建成功', position: 'top' });
    clean();
  });
}
function clean() {
  Object.assign(postForm, {
    nickname: '',
    email: '',
    password: '',
    headimg: '',
    realname: '',
    position_id: null,
    id: 0,
    disable: false,
  });
  repassword.value = '';
}
async function GetPosistions() {
  await positionsList().then((resp) => {
    Object.assign(positions, resp.data.data);
  });
}

await GetPosistions();
</script>

<style type="text/css">
.form-container .createPost-main-container .el-form-item {
  padding: 0px 15px 0px 15px;
}
</style>

<style scoped>
.add-user {
  margin: 20px;
}
</style>
