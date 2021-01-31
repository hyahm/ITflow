import Layout from '@/layout'

// 用户添加和管理的权限   usermanager

const userRoute = {
    path: '/user',
    component: Layout,
    redirect: '/user/list',
    name: 'userManager',
    meta: {
        title: '用户管理',
        icon: 'setting'
    },
    children: [{
            path: 'add',
            component: () =>
                import ('@/views/user/adduser'),
            name: 'adduser',
            meta: {
                title: '添加用户',
                icon: 'edit',
                roles: ['admin', 'user']
            }
        },
        {
            path: 'changepwd',
            component: () =>
                import ('@/views/user/changepwd'),
            name: 'changepwd',
            meta: {
                title: '修改密码',
                noCache: true,
                icon: 'edit'
            }
        },
        {
            path: 'usermanager',
            component: () =>
                import ('@/views/user/usermanager'),
            name: 'usermanager',
            meta: {
                title: '用户管理',
                icon: 'user',
                roles: ['admin', 'user']
            }
        },
        {
            path: 'usergroup',
            component: () =>
                import ('@/views/user/UserGroup'),
            name: 'usergroup',
            meta: {
                title: '用户组',
                icon: 'user'
            }
        },
        {
            path: 'myemail',
            component: () =>
                import ('@/views/user/email'),
            name: 'myemail',
            meta: {
                title: '修改邮箱',
                icon: 'email'
            }
        },
        {
            path: 'uploadhead',
            component: () =>
                import ('@/views/user/headImg'),
            name: 'uploadhead',
            meta: {
                title: '头像管理',
                icon: 'user'
            }
        }
    ]
}

export default userRoute