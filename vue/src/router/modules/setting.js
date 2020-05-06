import Layout from '@/layout'

const setting = {
  path: '/setting',
  component: Layout,
  redirect: '/setting/list',
  name: 'setting',
  meta: {
    title: 'bug设置',
    icon: 'setting',
    roles: ['admin', 'create version', 'create project', 'create env']
  },
  children: [
    {
      path: 'addversion',
      component: () => import('@/views/setting/pubversion'),
      name: 'pubversion',
      meta: {
        title: '发布版本',
        icon: 'guide',
        roles: ['admin', 'version']
      }
    },
    {
      path: 'versionlist',
      component: () => import('@/views/setting/versionlist'),
      name: 'versionlist',
      meta: {
        title: '版本列表',
        icon: 'bug',
        roles: ['admin', 'version']
      }
    },
    {
      path: 'addproject',
      component: () => import('@/views/setting/project'),
      name: 'projectmanager',
      meta: {
        title: '项目管理',
        icon: 'bug',
        roles: ['admin', 'project']
      }
    },
    {
      path: 'enver',
      component: () => import('@/views/setting/env'),
      name: '环境管理',
      meta: {
        title: '环境管理',
        icon: 'bug',
        roles: ['admin', 'env']
      }
    }
  ]
}

export default setting
