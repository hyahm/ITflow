import Layout from '@/views/layout/Layout'

const restful = {
  path: '/restful',
  component: Layout,
  redirect: '/restful/list',
  name: 'restful',
  meta: {
    title: 'restful',
    icon: 'http'
  },
  children: [
    {
      path: 'restfulnlist',
      component: () => import('@/views/restful/apiproject'),
      name: 'restfullist',
      meta: {
        title: 'restfullist',
        icon: 'list'
      }
    },
    {
      path: 'apilist',
      component: () => import('@/views/restful/apilist'),
      name: 'apilist',
      hidden: true,
      meta: {
        title: 'apilist',
        icon: 'list'
      }
    },
    {
      path: 'type',
      component: () => import('@/views/restful/TypeList'),
      name: 'type',
      meta: {
        title: 'type',
        icon: 'list'
      }
    },
    {
      path: 'showapi',
      component: () => import('@/views/restful/ShowApi'),
      name: 'showapi',
      hidden: true,
      meta: {
        title: 'showapi',
        icon: 'list'
      }
    },
    {
      path: 'header',
      component: () => import('@/views/restful/Header'),
      name: 'header',
      meta: {
        title: 'header',
        icon: 'user'
      }
    }
  ]
}

export default restful
