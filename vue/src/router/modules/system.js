import Layout from '@/views/layout/Layout'

const system = {
  path: '/system',
  component: Layout,
  redirect: '/system/list',
  name: 'system',
  meta: {
    title: 'system',
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
        title: 'bugStatus',
        icon: 'user',
        roles: ['admin', 'status']
      }
    },
    {
      path: 'bugmanager',
      component: () => import('@/views/system/bugmanager'),
      name: 'bugmanager',
      meta: {
        title: 'bugmanager',
        icon: 'email',
        roles: ['admin']
      }
    },
    {
      path: 'log',
      component: () => import('@/views/system/Log'),
      name: 'log',
      meta: {
        title: 'log',
        icon: 'email',
        roles: ['admin', 'log']
      }
    },
    {
      path: 'buggroup',
      component: () => import('@/views/system/bugGroup'),
      name: 'buggroup',
      meta: {
        title: 'buggroup',
        icon: 'user',
        roles: ['admin', 'statusgroup']
      }
    },
    {
      path: 'rolegroup',
      component: () => import('@/views/system/roleGroup'),
      name: 'rolegroup',
      meta: {
        title: 'rolegroup',
        icon: 'user',
        roles: ['admin', 'rolegroup']
      }
    },
    {
      path: 'defaultvalue',
      component: () => import('@/views/system/defaultValue'),
      name: 'defaultvalue',
      meta: {
        title: 'defaultvalue',
        icon: 'user',
        roles: ['admin']
      }
    },
    {
      path: 'important',
      component: () => import('@/views/system/Important'),
      name: 'important',
      meta: {
        title: 'important',
        icon: 'user',
        roles: ['admin', 'important']
      }
    },
    {
      path: 'level',
      component: () => import('@/views/system/Level'),
      name: 'level',
      meta: {
        title: 'level',
        icon: 'user',
        roles: ['admin', 'level']
      }
    },
    {
      path: 'position',
      component: () => import('@/views/system/position'),
      name: 'position',
      meta: {
        title: 'position',
        icon: 'user',
        roles: ['admin', 'position']
      }
    },
    {
      path: 'usergroup',
      component: () => import('@/views/system/UserGroup'),
      name: 'usergroup',
      meta: {
        title: 'usergroup',
        icon: 'user',
        roles: ['admin', 'usergroup']
      }
    }
  ]
}

export default system
