import Layout from '@/layout'

const docRouter = {
  path: '/doc',
  component: Layout,
  redirect: '/doc/project',
  name: 'doc',
  meta: {
    title: 'doc',
    icon: 'edit'
  },
  children: [
    {
      path: 'project',
      component: () => import('@/views/doc/project'),
      name: 'project',
      meta: {
        title: 'project',
        icon: 'edit'
      }
    },
    {
      path: 'show/:id',
      component: () => import('@/views/doc/show'),
      name: 'showdoc',
      hidden: true,
      meta: {
        title: 'showdoc',
        icon: 'edit'
      }
    }
  ]
}

export default docRouter
