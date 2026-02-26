import type { RouteRecordRaw } from 'vue-router';
// import Layout from 'src/layouts/MainLayout.vue';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [{ path: '', component: () => import('pages/IndexPage.vue') }],
  },

  {
    path: '/',
    component: () => import('src/layouts/MainLayout.vue'),
    redirect: 'dashboard',
    meta: {
      title: '主面板',
    },
    children: [
      {
        path: 'dashboard',
        component: () => import('src/pages/dashboard/admin/AdminIndex.vue'),
        name: 'dashboard',
        meta: { title: '面板', icon: 'dashboard', noCache: true },
      },
    ],
  },
  {
    path: '/login',
    component: () => import('src/pages/login/LoginIndex.vue'),
    meta: {
      hidden: true,
    },
  },

  {
    path: '/bug',
    component: () => import('src/layouts/MainLayout.vue'),
    redirect: '/bug/allbugs',
    name: 'bug',
    meta: {
      title: '任务管理',
      icon: 'bi-bug',
    },
    children: [
      {
        path: 'create',
        component: () => import('src/pages/bug/BugCreate.vue'),
        name: 'createArticle',
        meta: {
          title: '创建任务',
          icon: 'bi-file-earmark-plus-fill',
        },
      },
      {
        path: 'edit',
        component: () => import('src/pages/bug/BugEdit.vue'),
        name: 'editArticle',
        meta: { title: '编辑任务', noCache: true, icon: 'bi-file-code' },
      },
      {
        path: 'mybug',
        component: () => import('src/pages/bug/MyBug.vue'),
        name: 'articleList',
        meta: { title: '我的创建', icon: 'bi-bug' },
      },
      {
        path: 'allbugs',
        component: () => import('src/pages/bug/AllBugs.vue'),
        name: 'bugs',
        meta: {
          title: '所有任务',
          icon: 'bi-bug',
        },
      },
      {
        path: 'mytask',
        component: () => import('src/pages/bug/MyTask.vue'),
        name: 'mytask',
        meta: {
          title: '我的任务',
          icon: 'bi-bug',
        },
      },
    ],
  },
  {
    path: '/system',
    component: () => import('src/layouts/MainLayout.vue'),
    redirect: '/system/list',
    name: 'system',
    meta: {
      title: '系统设置',
      icon: 'settings',
      roles: ['admin', 'log'],
    },
    children: [
      {
        path: 'email',
        component: () => import('src/pages/system/SystemEmail.vue'),
        name: 'email',
        meta: {
          title: 'email',
          icon: 'email',
          roles: ['admin'],
        },
      },

      {
        path: 'dustbin',
        component: () => import('src/pages/system/BugDustbin.vue'),
        name: 'dustbin',
        meta: {
          title: '任务垃圾箱',
          icon: 'email',
          roles: ['admin'],
        },
      },

      {
        path: 'rolegroup',
        component: () => import('src/pages/system/SystemRole.vue'),
        name: 'rolegroup',
        meta: {
          title: '角色组',
          icon: 'group',
          roles: ['admin'],
        },
      },
      {
        path: 'log',
        component: () => import('src/pages/system/SystemLog.vue'),
        name: 'log',
        meta: {
          title: '日志',
          icon: 'notifications',
          roles: ['admin', 'log'],
        },
      },

      {
        path: 'defaultvalue',
        component: () => import('src/pages/system/DefaultValue.vue'),
        name: 'defaultvalue',
        meta: {
          title: '默认值',
          icon: 'disabled_by_default',
          roles: ['admin'],
        },
      },
    ],
  },
  {
    path: '/user',

    redirect: '/user/list',
    name: 'userManager',
    component: () => import('layouts/MainLayout.vue'),
    meta: {
      title: '用户管理',
      icon: 'person',
    },
    children: [
      {
        path: 'add',
        component: () => import('src/pages/user/AddUser.vue'),
        name: 'adduser',
        meta: {
          title: '添加用户',
          icon: 'edit',
          roles: ['admin', 'user'],
        },
      },
      // {
      //   path: 'changepwd',
      //   component: () => import('src/pages/user/ChangePwd.vue'),
      //   name: 'changepwd',
      //   meta: {
      //     title: '修改密码',
      //     noCache: true,
      //     icon: 'edit',
      //   },
      // },
      {
        path: 'usermanager',
        component: () => import('src/pages/user/UserManager.vue'),
        name: 'usermanager',
        meta: {
          title: '用户管理',
          icon: 'group',
          roles: ['admin', 'user'],
        },
      },
      // {
      //   path: 'usergroup',
      //   component: () => import('src/pages/user/UserGroup.vue'),
      //   name: 'usergroup',
      //   meta: {
      //     title: '用户组',
      //     icon: 'user',
      //   },
      // },
      // {
      //   path: 'myemail',
      //   component: () => import('src/pages/user/UserEmail.vue'),
      //   name: 'myemail',
      //   meta: {
      //     title: '修改邮箱',
      //     icon: 'email',
      //   },
      // },
      {
        path: 'uploadhead',
        component: () => import('src/pages/user/headImg.vue'),
        name: 'uploadhead',
        meta: {
          title: '头像管理',
          icon: 'image',
        },
      },
    ],
  },
  // {
  //   path: '/charts',

  //   redirect: 'noredirect',
  //   name: 'charts',
  //   meta: {
  //     title: 'charts',
  //     icon: 'chart',
  //   },
  //   children: [
  //     {
  //       path: 'keyboard',
  //       component: () => import('src/pages/charts/keyboard.vue'),
  //       name: 'keyboardChart',
  //       meta: { title: 'keyboardChart', noCache: true },
  //     },
  //     {
  //       path: 'line',
  //       component: () => import('src/pages/charts/line.vue'),
  //       name: 'lineChart',
  //       meta: { title: 'lineChart', noCache: true },
  //     },
  //     {
  //       path: 'mixchart',
  //       component: () => import('src/pages/charts/mixChart.vue'),
  //       name: 'mixChart',
  //       meta: { title: 'mixChart', noCache: true },
  //     },
  //   ],
  // },
  {
    path: '/setting',

    redirect: '/setting/list',
    component: () => import('layouts/MainLayout.vue'),
    name: 'setting',
    meta: {
      title: '设置中心',
      icon: 'miscellaneous_services',
      roles: ['admin', 'version', 'project', 'env', 'status', 'statusgroup', 'important', 'level'],
    },
    children: [
      {
        path: 'bugstatus',
        component: () => import('src/pages/setting/SettingStatus.vue'),
        name: 'bugstatus',
        meta: {
          title: 'bug状态',
          icon: 'group',
          roles: ['admin', 'status'],
        },
      },

      // {
      //   path: 'versionlist',
      //   component: () => import('src/pages/setting/VersionList.vue'),
      //   name: 'versionlist',
      //   meta: {
      //     title: '版本列表',
      //     icon: 'bug',
      //     roles: ['admin', 'version'],
      //   },
      // },
      // {
      //   path: "buggroup",
      //   component: () => import("pages/setting/bugGroup"),
      //   name: "buggroup",
      //   meta: {
      //     title: "bug状态组",
      //     icon: "user",
      //     roles: ["admin", "statusgroup"]
      //   }
      // },

      {
        path: 'project',
        component: () => import('src/pages/setting/SettingProject.vue'),
        name: 'projectmanager',
        meta: {
          title: '项目管理',
          icon: 'work',
          roles: ['admin', 'project'],
        },
      },
      {
        path: 'important',
        component: () => import('src/pages/setting/SettingImportant.vue'),
        name: 'important',
        meta: {
          title: '重要性',
          icon: 'important_devices',
          roles: ['admin', 'important'],
        },
      },
      {
        path: 'level',
        component: () => import('src/pages/setting/SettingLevel.vue'),
        name: 'level',
        meta: {
          title: '优先级',
          icon: 'filter_list',
          roles: ['admin', 'level'],
        },
      },

      {
        path: 'position',
        component: () => import('src/pages/setting/SettingPosition.vue'),
        name: 'position',
        meta: {
          title: '职位',
          icon: 'apartment',
          roles: ['admin', 'position'],
        },
      },
      {
        path: 'enver',
        component: () => import('src/pages/setting/SettingEnv.vue'),
        name: '环境管理',
        meta: {
          title: '环境管理',
          icon: 'web_stories',
          roles: ['admin', 'env'],
        },
      },
    ],
  },
  // {
  //   path: '/showbug',

  //   redirect: 'noredirect',
  //   children: [
  //     {
  //       path: ':id',
  //       component: () => import('src/pages/bug/BugShow.vue'),
  //       name: 'showbug',
  //       meta: { title: 'backToTop' },
  //     },
  //   ],
  // },
  // Always leave this as last one,
  // but you can also remove it

  {
    path: '/:catchAll(.*)*',
    redirect: '/login',
    // component: () => import('src/pages/ErrorNotFound.vue'),
  },
];

export default routes;
