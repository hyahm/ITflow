import Layout from '@/layout'

const setting = {
    path: '/setting',
    component: Layout,
    redirect: '/setting/list',
    name: 'setting',
    meta: {
        title: '设置中心',
        icon: 'setting',
        roles: ['admin', 'version', 'project', 'env', 'status', 'statusgroup', 'important', 'level']
    },
    children: [{
            path: 'bugstatus',
            component: () =>
                import ('@/views/setting/status'),
            name: 'bugstatus',
            meta: {
                title: 'bug状态',
                icon: 'user',
                roles: ['admin', 'status']
            }
        },

        {
            path: 'versionlist',
            component: () =>
                import ('@/views/setting/versionlist'),
            name: 'versionlist',
            meta: {
                title: '版本列表',
                icon: 'bug',
                roles: ['admin', 'version']
            }
        },
        {
            path: 'buggroup',
            component: () =>
                import ('@/views/setting/bugGroup'),
            name: 'buggroup',
            meta: {
                title: 'bug状态组',
                icon: 'user',
                roles: ['admin', 'statusgroup']
            }
        },

        {
            path: 'addproject',
            component: () =>
                import ('@/views/setting/project'),
            name: 'projectmanager',
            meta: {
                title: '项目管理',
                icon: 'bug',
                roles: ['admin', 'project']
            }
        },
        {
            path: 'important',
            component: () =>
                import ('@/views/setting/Important'),
            name: 'important',
            meta: {
                title: '重要性',
                icon: 'user',
                roles: ['admin', 'important']
            }
        },
        {
            path: 'level',
            component: () =>
                import ('@/views/setting/Level'),
            name: 'level',
            meta: {
                title: '优先级',
                icon: 'user',
                roles: ['admin', 'level']
            }
        },

        {
            path: 'position',
            component: () =>
                import ('@/views/setting/position'),
            name: 'position',
            meta: {
                title: '职位',
                icon: 'user',
                roles: ['admin', 'position']
            }
        },
        {
            path: 'enver',
            component: () =>
                import ('@/views/setting/env'),
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