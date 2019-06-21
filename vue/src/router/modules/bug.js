import Layout from '@/views/layout/Layout'

const bugRouter = {
  path: '/bug',
  component: Layout,
  redirect: '/bug/list',
  name: 'bug',
  meta: {
    title: 'bug',
    icon: 'bug'
  },
  children: [
    {
      path: 'create',
      component: () => import('@/views/bug/create'),
      name: 'createArticle',
      meta: {
        title: 'createArticle',
        icon: 'edit'
      }
    },
    {
      path: 'edit/:id(\\d+)',
      component: () => import('@/views/bug/edit'),
      name: 'editArticle',
      meta: { title: 'editArticle', noCache: true },
      hidden: true
    },
    {
      path: 'mybug',
      component: () => import('@/views/bug/mybug'),
      name: 'articleList',
      meta: { title: 'articleList', icon: 'guide' }
    },
    {
      path: 'allbugs',
      component: () => import('@/views/bug/allbugs'),
      name: 'bugs',
      meta: {
        title: 'allbugs',
        icon: 'bug'
      }
    },
    {
      path: 'mytask',
      component: () => import('@/views/bug/mytask'),
      name: 'mytask',
      meta: {
        title: 'mytask',
        icon: 'bug'
      }
    }
  ]
}

export default bugRouter
