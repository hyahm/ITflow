import Layout from '@/layout'

const userRoute = {
  path: '/user',
  component: Layout,
  redirect: '/user/list',
  name: 'userManager',
  meta: {
    title: '用户管理',
    icon: 'setting'
  },
  children: [
    {
      path: 'add',
      component: () => import('@/views/user/adduser'),
      name: 'adduser',
      meta: {
        title: '添加用户',
        icon: 'edit',
        roles: ['admin', 'user']
      }
    },
    {
      path: 'changepwd',
      component: () => import('@/views/user/changepwd'),
      name: 'changepwd',
      meta: {
        title: '修改密码',
        noCache: true,
        icon: 'edit'
      }
    },
    {
      path: 'usermanager',
      component: () => import('@/views/user/usermanager'),
      name: 'usermanager',
      meta: {
        title: '用户管理',
        icon: 'user',
        roles: ['admin', 'usermanager']
      }
    },
    {
      path: 'uploadhead',
      component: () => import('@/views/user/headImg'),
      name: 'uploadhead',
      meta: {
        title: '头像管理',
        icon: 'user'
      }
    }
    // {
    //   path: 'changeinfo',
    //   component: () => import('@/views/user/changeinfo'),
    //   name: 'changeinfo',
    //   meta: {
    //     title: 'changeinfo',
    //     icon: 'user'
    //   }
    // }
  ]
}

export default userRoute
