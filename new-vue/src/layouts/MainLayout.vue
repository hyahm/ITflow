<template>
  <q-layout view="hHh Lpr lFf">
    <!-- 页头 -->
    <q-header elevated>
        <q-toolbar class="justify-between">
        <!-- 左侧：菜单按钮 + 应用名称 -->
        <div class="flex items-center">
          <q-btn flat dense round icon="menu" @click="leftDrawerOpen = !leftDrawerOpen" />
          <q-toolbar-title class="ml-2">应用名称</q-toolbar-title>
        </div>

        <!-- 右侧：头像 + 退出按钮 -->
        <div class="flex items-center gap-2">
          <!-- 头像组件：可点击展开下拉（可选） -->
          <q-avatar
            size="40px"
            class="cursor-pointer"
            @click="toggleUserDropdown"
          >
            <!-- 优先显示用户头像，无则显示默认占位图 -->
            <q-img
              :src="user.headImg"
              alt="用户头像"
              fit="cover"
            />
            <!-- 纯文字占位（备选）：<span class="text-white text-xs">用户</span> -->
          </q-avatar>

          <!-- 退出按钮 -->
          <q-btn
            flat
            dense
            label="退出"
            icon="logout"
            color="negative"
            @click="handleLogout"
          />

          <!-- 可选：用户下拉菜单（如需更多操作如个人中心） -->
          <q-menu v-model="userDropdownOpen" anchor="top right" self="top right">
            <q-list dense style="min-width: 150px;">
              <q-item clickable @click="goToProfile">
                <q-item-section avatar>
                  <q-icon name="person" />
                </q-item-section>
                <q-item-section>个人中心</q-item-section>
              </q-item>
              <q-item clickable @click="handleLogout" color="negative">
                <q-item-section avatar>
                  <q-icon name="logout" />
                </q-item-section>
                <q-item-section>退出登录</q-item-section>
              </q-item>
            </q-list>
          </q-menu>
        </div>
      </q-toolbar>
    </q-header>

    <!-- 左侧抽屉 -->
    <q-drawer v-model="leftDrawerOpen" bordered content-class="bg-grey-1">
      <!-- <q-list>
        <div v-for="(value, index) in routes" :key="index">
          <q-item clickable :to="value.path">
            <q-item-section avatar>
              <q-icon :name="value.icon" />
            </q-item-section>
            <q-item-section>
              <q-item-label>{{ value.title }}</q-item-label>
            </q-item-section>
            <q-item v-for="(child, x) in value.children" clickable :to="child.path" :key="x">
              <q-item-section avatar>
                <q-icon :name="child.name" />
              </q-item-section>
              <q-item-section>
                <q-item-label>{{ child.title }}</q-item-label>
              </q-item-section>
            </q-item>
          </q-item>
        </div>
      </q-list> -->
      <!-- 有子路由：用可展开项 -->
      <div v-for="(value, index) in routes" :key="index">
        <q-expansion-item
          v-if="value.children.length > 1"
          :icon="value.icon"
          :label="value.title"
          expand-separator
          dense
          dense-toggle
          icon-size="16px"
        >
          <q-item
            v-for="(child, cIdx) in value.children"
            :key="cIdx"
            clickable
            :to="child.path"
            :inset-level="0.5"
            dense
            dense-toggle
            icon-size="16px"
          >
            <q-item-section avatar>
              <q-icon :name="child.icon" />
            </q-item-section>
            <q-item-section>
              <q-item-label>{{ child.title }}</q-item-label>
            </q-item-section>
          </q-item>
        </q-expansion-item>

        <!-- 无子路由：单条普通项 -->
        <q-item v-else clickable :to="value.path">
          <q-item-section avatar>
            <q-icon :name="value.icon" />
          </q-item-section>
          <q-item-section>
            <q-item-label>{{ value.title }}</q-item-label>
          </q-item-section>
        </q-item>
      </div>
    </q-drawer>

    <!-- 页面容器 -->
    <q-page-container>
      <!-- QPage 必须放在这里 -->
      <Suspense>
        <router-view />
      </Suspense>
    </q-page-container>
  </q-layout>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue';
import type { RouteRecordRaw } from 'vue-router';
import { useRouter } from 'vue-router';
import { useQuasar } from 'quasar'
import { useUserStore } from 'src/stores/user'



import {  getInfo } from 'src/api/user';
async function getinfo() {
  await getInfo().then((resp) => {
    user.SetHead(resp.data.data.avatar)
    // Object.assign(postForm, resp.data);
  });
}


const userAvatar = ref(''); // 用户头像地址（可从接口/本地存储获取）
const userDropdownOpen = ref(false); // 用户下拉菜单显隐
const router = useRouter();
const $q = useQuasar()
const user = useUserStore()
await getinfo()
interface RouteLabel {
  title: string;
  path: string;
  name: string;
  icon: string;
  // hidden: boolean;  // 不显示的不会进来
  children: RouteLabel[];
}
// const routes = router.options.routes;
const routes: RouteLabel[] = reactive([]);

function validMeta(meta: object | undefined): boolean {
  if (meta == undefined) {
    // 如果没有meta标签的不做处理
    return false;
  }
  if ('hidden' in meta && meta.hidden) {
    // 如果是隐藏的就跳过
    return false;
  }
  return true;
}

// 新增：切换用户下拉菜单
const toggleUserDropdown = () => {
  userDropdownOpen.value = !userDropdownOpen.value;
};

function getString(key: string, meta: object | undefined, val: string): string {
  if (meta == undefined) {
    // 如果没有meta标签的不做处理
    return val;
  }
  if (key in meta) {
    // 如果是隐藏的就跳过
    return (meta as Record<string, string>)[key] as string;
  }
  return val;
}

const joinPath = (...parts: string[]) =>
  '/' +
  parts
    .map((p) => p.replace(/^\/+|\/+$/g, '')) // 去掉每段头尾 /
    .filter(Boolean) // 去掉空串
    .join('/');

function c() {
  for (const route of router.options.routes) {
    const meta = route.meta;
    if (!validMeta(meta)) {
      continue;
    }
    const rt: RouteLabel = {
      title: getString('title', meta, ''),
      path: (route.redirect as string) || route.path,
      name: route.name as string,
      icon: getString('icon', meta, 'home'),
      children: [],
    };
    if (route.children!.length == 1) {
      rt.path = route.redirect as string;
    } else if (route.children!.length > 1) {
      for (const child of route.children as RouteRecordRaw[]) {
        // 最多2层
        const childMeta = child.meta;
        if (!validMeta(childMeta)) {
          continue;
        }
        const childRT: RouteLabel = {
          title: getString('title', childMeta, '标签'),
          path: joinPath(route.path, child.path),
          name: child.name as string,
          icon: getString('icon', childMeta, 'home'),
          children: [],
        };
        rt.children.push(childRT);
      }
    }

    routes.push(rt);
  }
}
c();


// 独立封装异步退出逻辑（带错误捕获）
const handleLogout = async () => {
  try {
    // 1. 清除用户信息
    // localStorage.removeItem('token');
    // localStorage.removeItem('avatar');
    userAvatar.value = '';

    // 2. 提示退出成功
    $q.notify({
      type: 'positive',
      message: '退出登录成功！',
      icon: 'check_circle',
    });

    // 3. 跳转登录页（异步操作）
    await router.push('/login');
  } catch (error) {
    // 捕获异步操作错误
    $q.notify({
      type: 'negative',
      message: `退出失败：${(error as Error).message}`,
      icon: 'error',
    });
  }
};

// 新增：跳转到个人中心（可选）
const goToProfile = async () => {
  userDropdownOpen.value = false; // 关闭下拉菜单
  await router.push('/profile'); // 替换为你的个人中心路径
};


const leftDrawerOpen = ref(true);
</script>
