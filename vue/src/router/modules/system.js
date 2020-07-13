import Layout from '@/layout'

const system = {
  path: '/system',
  component: Layout,
  redirect: '/system/list',
  name: 'system',
  meta: {
    title: '系统设置',
    icon: 'setting',
    roles: ['admin', 'log']
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
      path: 'bugmanager',
      component: () => import('@/views/system/bugdustbin'),
      name: 'bugmanager',
      meta: {
        title: 'bug垃圾箱',
        icon: 'email',
        roles: ['admin']
      }
    },

    {
      path: 'rolegroup',
      component: () => import('@/views/system/roleGroup'),
      name: 'rolegroup',
      meta: {
        title: '角色组',
        icon: 'user',
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
      path: 'defaultvalue',
      component: () => import('@/views/system/defaultValue'),
      name: 'defaultvalue',
      meta: {
        title: '默认值',
        icon: 'user',
        roles: ['admin']
      }
    }

  ]
}

export default system
