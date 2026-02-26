<template>
  <q-layout view="hHh Lpr lFf">
    <q-page-container>
      <q-page class="flex flex-center bg-grey-2">
        <q-card class="login-card q-pa-lg" style="width: 100%; max-width: 400px">
          <!-- 标题 -->
          <q-card-section class="text-center">
            <div class="text-h4 text-weight-bold text-primary q-mb-sm">欢迎登录</div>
            <div class="text-subtitle1 text-grey-7">请输入您的账号和密码</div>
          </q-card-section>

          <!-- 表单 -->
          <q-card-section>
            <q-form @submit="onSubmit" class="q-gutter-md">
              <!-- 用户名/邮箱 -->
              <q-input
                v-model="form.email"
                label="邮箱 / 用户名"
                outlined
                dense
                :rules="[(val) => !!val || '请输入邮箱或用户名']"
              >
                <template v-slot:prepend>
                  <q-icon name="person" />
                </template>
              </q-input>

              <!-- 密码 -->
              <q-input
                v-model="form.password"
                label="密码"
                :type="isPwd ? 'password' : 'text'"
                outlined
                dense
                :rules="[(val) => !!val || '请输入密码']"
              >
                <template v-slot:prepend>
                  <q-icon name="lock" />
                </template>
                <template v-slot:append>
                  <q-icon
                    :name="isPwd ? 'visibility_off' : 'visibility'"
                    class="cursor-pointer"
                    @click="isPwd = !isPwd"
                  />
                </template>
              </q-input>

              <!-- 记住我 & 忘记密码 -->
              <div class="row items-center justify-between q-mt-lg">
                <q-checkbox v-model="form.remember" label="记住我" color="primary" />
                <a href="#" class="text-primary text-caption" @click.prevent="onForgotPassword">
                  忘记密码?
                </a>
              </div>

              <!-- 提交按钮 -->
              <div class="q-mt-xl">
                <q-btn
                  type="submit"
                  label="登录"
                  color="primary"
                  class="full-width"
                  size="lg"
                  :loading="loading"
                />
              </div>

              <!-- 注册链接 -->
              <div class="text-center q-mt-lg">
                <span class="text-grey-7">还没有账号?</span>
                <a
                  href="#"
                  class="text-primary text-weight-medium q-ml-sm"
                  @click.prevent="onRegister"
                >
                  立即注册
                </a>
              </div>
            </q-form>
          </q-card-section>

          <!-- 分割线 -->
          <q-card-section v-if="socialLoginEnabled">
            <q-separator class="q-mb-md">
              <span class="text-grey-6 q-px-sm bg-grey-2">或</span>
            </q-separator>

            <!-- 社交登录 -->
            <div class="row justify-center q-gutter-md">
              <q-btn round color="red" icon="fab fa-google" @click="loginWithGoogle" />
              <q-btn round color="blue" icon="fab fa-github" @click="loginWithGithub" />
              <q-btn round color="light-blue" icon="fab fa-weixin" @click="loginWithWechat" />
            </div>
          </q-card-section>
        </q-card>

        <!-- 语言切换 -->
        <div class="absolute-top-right q-pa-md">
          <q-btn round flat icon="translate" @click="toggleLanguage">
            <q-tooltip>切换语言</q-tooltip>
          </q-btn>
        </div>

        <!-- 主题切换 -->
        <!-- <div class="absolute-top-left q-pa-md">
          <q-btn
            round
            flat
            :icon="$q.dark.isActive ? 'light_mode' : 'dark_mode'"
            @click="toggleDarkMode"
          >
            <q-tooltip>切换主题</q-tooltip>
          </q-btn>
        </div> -->
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script setup lang="ts">
import type { Login } from 'src/api/types/type';
import { login } from 'src/api/user';
import { ref, reactive } from 'vue';
import { useQuasar } from 'quasar';
import { useRouter } from 'vue-router';
import {} from 'src/stores/index';
import { LocalStorage } from 'quasar';
const loginForm: Login = reactive({ username: '', password: '' });
const $q = useQuasar();
const router = useRouter();

// 表单数据
const form = reactive({
  email: 'admin',
  password: '12345678',
  remember: false,
});

const isPwd = ref(true);
const loading = ref(false);
const socialLoginEnabled = ref(true); // 是否显示社交登录

// 提交登录
const onSubmit = async () => {
  loading.value = true;
  loginForm.username = form.email;
  loginForm.password = form.password;
  // try {
  // 模拟 API 调用
  // await new Promise(resolve => setTimeout(resolve, 1500))
  loading.value = true;
  await login(loginForm).then(async (resp) => {
    loading.value = false;
    if (resp.data.code == 0) {
      // 登录成功
      $q.notify({
        type: 'positive',
        message: '登录成功',
        position: 'top',
      });
      LocalStorage.set('token', resp.data.data);
      // 跳转到首页
      await router.push('/dashboard');
    }
    return;
  });

  // } catch (error) {
  //   $q.notify({
  //     type: 'negative',
  //     message: error instanceof Error ? error.message : '登录失败',
  //     position: 'top',
  //   });
  // } finally {
  //   loading.value = false;
  // }
};

// 忘记密码
const onForgotPassword = async () => {
  await router.push('/forgot-password');
};

// 注册
const onRegister = async () => {
  await router.push('/register');
};

// 社交登录
const loginWithGoogle = () => {
  $q.notify('Google 登录');
};

const loginWithGithub = () => {
  $q.notify('GitHub 登录');
};

const loginWithWechat = () => {
  $q.notify('微信登录');
};

// 切换语言
const toggleLanguage = () => {
  $q.notify('切换语言功能');
};
</script>

<style scoped>
.login-card {
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.1);
  border-radius: 16px;
}

/* 响应式调整 */
@media (max-width: 600px) {
  .login-card {
    margin: 16px;
    border-radius: 12px;
  }
}
</style>
