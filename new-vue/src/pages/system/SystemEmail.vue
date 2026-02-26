<template>
  <div class="q-pa-md q-pl-lg">
    <q-form ref="formRef" class="q-gutter-md" style="max-width: 500px">
      <q-banner class="bg-orange-1 text-brown q-mb-md rounded-borders">
        <template v-slot:avatar>
          <q-icon name="warning" color="brown" />
        </template>
        设置并启用后 只有创建用户，创建bug，转交bug才会有邮件通知
      </q-banner>

      <q-item tag="label" v-ripple class="rounded-borders q-mb-sm">
        <q-item-section>
          <q-item-label>邮件通知功能</q-item-label>
          <q-item-label caption>{{ form.enable ? '已启用' : '已禁用' }}</q-item-label>
        </q-item-section>
        <q-item-section avatar>
          <q-toggle
            v-model="form.enable"
            color="positive"
            keep-color
            :label="form.enable ? '启用' : '禁用'"
            left-label
          />
        </q-item-section>
      </q-item>

      <q-input
        v-model="form.host"
        label="邮箱服务器"
        clearable
        placeholder="请输入邮箱服务器地址"
        outlined
        dense
      />

      <q-input
        v-model="form.nickname"
        label="昵称"
        clearable
        placeholder="请输入昵称不填默认邮箱的名称"
        outlined
        dense
      />

      <q-input
        v-model="form.email"
        label="邮箱地址"
        type="email"
        clearable
        placeholder="请输入邮箱地址"
        outlined
        dense
      />

      <q-input
        v-model="form.password"
        label="邮箱密码"
        type="password"
        clearable
        placeholder="请输入邮箱密码"
        outlined
        dense
      />

      <q-input
        v-model.number="form.port"
        label="邮箱端口"
        type="number"
        clearable
        placeholder="请输入邮箱端口"
        outlined
        dense
      />

      <q-input
        v-model="form.to"
        label="测试邮箱"
        type="email"
        clearable
        placeholder="请输入接收邮箱"
        outlined
        dense
      />

      <div class="q-gutter-md q-mt-md q-ml-lg">
        <q-btn color="primary" label="验证" @click="handleTest" unelevated />
        <q-btn color="primary" label="保存" @click="handleSave" unelevated />
      </div>
    </q-form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { testEmail, saveEmail, getEmailStatus } from 'src/api/email';
import type { Email } from 'src/api/types/type';
import { useQuasar } from 'quasar';

const $q = useQuasar();
const formRef = ref();

const form = ref<Email>({
  email: '',
  password: '',
  port: 25,
  host: '',
  id: 0,
  enable: false,
  nickname: '',
  to: '',
});

async function getEmail() {
  await getEmailStatus().then((resp) => {
    const data = resp.data;
    form.value = {
      ...form.value,
      email: data.email,
      enable: data.enable,
      host: data.host,
      port: data.port,
      nickname: data.nickname,
      id: data.id,
      password: data.password,
    };
  });
}
await getEmail();
function validateBeforeTest(): boolean {
  const rules = [
    { field: form.value.host, msg: '邮箱服务器不能为空' },
    { field: form.value.to, msg: '测试邮箱不能为空' },
    { field: form.value.email, msg: '邮箱账号不能为空' },
    { field: form.value.password, msg: '邮箱密码不能为空' },
  ];

  for (const v of rules) {
    if (!v.field || v.field === '') {
      $q.notify({
        color: 'negative',
        message: v.msg,
        position: 'top',
        timeout: 2000,
      });
      return false;
    }
  }
  return true;
}

async function handleTest() {
  if (!validateBeforeTest()) return;

  if (form.value.port === 0) {
    form.value.port = 25;
  }

  await testEmail(form.value).then(() => {
    $q.notify({
      color: 'positive',
      message: '测试邮件发送成功',
      position: 'top',
      timeout: 2000,
    });
  });
}

async function handleSave() {
  await saveEmail(form.value).then((resp) => {
    form.value.id = resp.data.id;
    $q.notify({
      color: 'positive',
      message: '保存成功',
      position: 'top',
      timeout: 2000,
    });
  });
}
</script>

<style scoped>
.form-container :deep(.q-field__control) {
  background: white;
}
</style>
