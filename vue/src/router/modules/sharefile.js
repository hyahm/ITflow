import Layout from '@/views/layout/Layout'

const shareRouter = {
  path: '/share',
  component: Layout,
  // redirect: '/share/view',
  name: 'view',
  meta: {
    title: 'bug',
    icon: 'bug'
  },
  children: [
    {
      path: 'view',
      component: () => import('@/views/sharefile/view'),
      name: 'sharelist',
      meta: {
        title: 'sharefile',
        icon: 'edit'
      }
    }
    // {
    //   path: 'edit/:id(\\d+)',
    //   component: () => import('@/views/bug/edit'),
    //   name: 'editArticle',
    //   meta: { title: 'editArticle', noCache: true },
    //   hidden: true
    // },
    // {
    //   path: 'mybug',
    //   component: () => import('@/views/bug/mybug'),
    //   name: 'articleList',
    //   meta: { title: 'articleList', icon: 'guide' }
    // },
    // {
    //   path: 'allbugs',
    //   component: () => import('@/views/bug/allbugs'),
    //   name: 'bugs',
    //   meta: {
    //     title: 'allbugs',
    //     icon: 'bug',
    //     roles: ['admin', 'see allbug']
    //   }
    // },
    // {
    //   path: 'mytask',
    //   component: () => import('@/views/bug/mytask'),
    //   name: 'mytask',
    //   meta: {
    //     title: 'mytask',
    //     icon: 'bug'
    //   }
    // }
  ]
}

export default shareRouter
