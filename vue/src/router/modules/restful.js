import Layout from '@/layout'

const restful = {
  path: '/restful',
  component: Layout,
  redirect: '/restful/list',
  name: 'restful',
  meta: {
    title: 'rest接口',
    icon: 'http'
  },
  children: [
    {
      path: 'restfulnlist',
      component: () => import('@/views/restful/apiproject'),
      name: 'restfullist',
      meta: {
        title: '项目管理',
        icon: 'list'
      }
    },
    {
      path: 'apilist',
      component: () => import('@/views/restful/apilist'),
      name: 'apilist',
      hidden: true,
      meta: {
        title: '接口列表',
        icon: 'list'
      }
    },
    {
      path: 'type',
      component: () => import('@/views/restful/TypeList'),
      name: 'type',
      meta: {
        title: '类型',
        icon: 'list'
      }
    },
    {
      path: 'showapi',
      component: () => import('@/views/restful/ShowApi'),
      name: 'showapi',
      hidden: true,
      meta: {
        title: '显示接口',
        icon: 'list'
      }
    },
    {
      path: 'header',
      component: () => import('@/views/restful/Header'),
      name: 'header',
      meta: {
        title: '请求头管理',
        icon: 'user'
      }
    }
  ]
}

export default restful
