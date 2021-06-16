import Layout from '@/layout'

// 用户添加和管理的权限   usermanager

const docRoute = {
    path: '/doc',
    component: Layout,
    name: 'docManager',
    meta: {
        title: '文档管理',
        icon: 'setting',
        roles: ['admin', 'user']
    },
    children: [{
            path: 'doc',
            component: () =>
                import ('@/views/doc/doclist'),
            name: 'keys',
            meta: {
                title: '文档管理',
                icon: 'edit',

            }
        },
        {
            path: 'keys',
            component: () =>
                import ('@/views/doc/keys'),
            name: 'keylist',
            meta: {
                title: '秘钥管理',
                noCache: true,
                icon: 'edit'
            }
        }
    ]
}

export default docRoute