import Layout from '@/layout'

const bugRouter = {
    path: '/bug',
    component: Layout,
    redirect: '/bug/list',
    name: 'bug',
    meta: {
        title: '任务管理',
        icon: 'bug'
    },
    children: [{
            path: 'create',
            component: () =>
                import ('@/views/bug/create'),
            name: 'createArticle',
            meta: {
                title: '创建任务',
                icon: 'edit'
            }
        },
        {
            path: 'edit/:id(\\d+)',
            component: () =>
                import ('@/views/bug/edit'),
            name: 'editArticle',
            meta: { title: '编辑任务', noCache: true },
            hidden: true
        },
        {
            path: 'mybug',
            component: () =>
                import ('@/views/bug/mybug'),
            name: 'articleList',
            meta: { title: '我的任务', icon: 'guide' }
        },
        {
            path: 'allbugs',
            component: () =>
                import ('@/views/bug/allbugs'),
            name: 'bugs',
            meta: {
                title: '所有任务',
                icon: 'bug'
            }
        },
        {
            path: 'mytask',
            component: () =>
                import ('@/views/bug/mytask'),
            name: 'mytask',
            meta: {
                title: '我的任务',
                icon: 'bug'
            }
        }
    ]
}

export default bugRouter