import Layout from '@/layout'

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
        title: '创建bug',
        icon: 'edit'
      }
    },
    {
      path: 'edit/:id(\\d+)',
      component: () => import('@/views/bug/edit'),
      name: 'editArticle',
      meta: { title: '编辑bug', noCache: true },
      hidden: true
    },
    {
      path: 'mybug',
      component: () => import('@/views/bug/mybug'),
      name: 'articleList',
      meta: { title: '我的bug', icon: 'guide' }
    },
    {
      path: 'allbugs',
      component: () => import('@/views/bug/allbugs'),
      name: 'bugs',
      meta: {
        title: '所有bug',
        icon: 'bug'
      }
    },
    {
      path: 'mytask',
      component: () => import('@/views/bug/mytask'),
      name: 'mytask',
      meta: {
        title: '我的任务',
        icon: 'bug'
      }
    }
  ]
}

export default bugRouter
