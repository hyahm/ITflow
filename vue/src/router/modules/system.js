import Layout from '@/layout'

const system = {
  path: '/system',
  component: Layout,
  redirect: '/system/list',
  name: 'system',
  meta: {
    title: '系统设置',
    icon: 'setting'
  },
  children: [
    {
      path: 'email',
      component: () => import('@/views/system/email'),
      name: 'email',
      meta: {
        title: 'email',
        icon: 'email',
        roles: ['admin']
      }
    },
    {
      path: 'bugstatus',
      component: () => import('@/views/system/status'),
      name: 'bugstatus',
      meta: {
        title: 'bug状态',
        icon: 'user',
        roles: ['admin', 'status']
      }
    },
    {
      path: 'bugmanager',
      component: () => import('@/views/system/bugmanager'),
      name: 'bugmanager',
      meta: {
        title: 'bug管理',
        icon: 'email',
        roles: ['admin']
      }
    },
    {
      path: 'log',
      component: () => import('@/views/system/Log'),
      name: 'log',
      meta: {
        title: '日志',
        icon: 'email',
        roles: ['admin', 'log']
      }
    },
    {
      path: 'buggroup',
      component: () => import('@/views/system/bugGroup'),
      name: 'buggroup',
      meta: {
        title: 'bug状态组',
        icon: 'user',
        roles: ['admin', 'statusgroup']
      }
    },
    {
      path: 'rolegroup',
      component: () => import('@/views/system/roleGroup'),
      name: 'rolegroup',
      meta: {
        title: '角色组',
        icon: 'user',
        roles: ['admin', 'rolegroup']
      }
    },
    {
      path: 'defaultvalue',
      component: () => import('@/views/system/defaultValue'),
      name: 'defaultvalue',
      meta: {
        title: '默认值',
        icon: 'user',
        roles: ['admin']
      }
    },
    {
      path: 'important',
      component: () => import('@/views/system/Important'),
      name: 'important',
      meta: {
        title: '重要性',
        icon: 'user',
        roles: ['admin', 'important']
      }
    },
    {
      path: 'level',
      component: () => import('@/views/system/Level'),
      name: 'level',
      meta: {
        title: '优先级',
        icon: 'user',
        roles: ['admin', 'level']
      }
    },
    {
      path: 'position',
      component: () => import('@/views/system/position'),
      name: 'position',
      meta: {
        title: '职位',
        icon: 'user',
        roles: ['admin', 'position']
      }
    },
    {
      path: 'usergroup',
      component: () => import('@/views/system/UserGroup'),
      name: 'usergroup',
      meta: {
        title: '用户组',
        icon: 'user',
        roles: ['admin', 'usergroup']
      }
    }
  ]
}

export default system
