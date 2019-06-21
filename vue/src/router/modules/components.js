/** When your routing table is too long, you can split it into small modules**/

import Layout from '@/views/layout/Layout'

const componentsRouter = {
  path: '/showbug',
  component: Layout,
  redirect: 'noredirect',
  hidden: true,
  children: [
    {
      path: ':id',
      component: () => import('@/views/bug/show'),
      name: 'showbug',
      meta: { title: 'backToTop' }
    }
  ]
}

export default componentsRouter
