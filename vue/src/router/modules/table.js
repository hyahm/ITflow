/** When your routing table is too long, you can split it into small modules**/

// import Layout from '@/views/layout/Layout'

// const tableRouter = {
//   path: '/table',
//   component: Layout,
//   redirect: '/table/complex-table',
//   name: 'table',
//   meta: {
//     title: 'Table',
//     icon: 'table',
//     roles: ['admin']
//   },
//   children: [
//     {
//       path: 'complete-table',
//       component: () => import('@/views/table/completeTable'),
//       name: 'completeTable',
//       meta: {
//         title: 'completeTask'
//         // roles: ['admin', 'manager', 'gomanager', 'godevelop', 'phpmanager', 'phpdevelop', 'cocosmanager', 'cocos']
//       }
//     },
//     // {
//     //   path: 'drag-table',
//     //   component: () => import('@/views/table/dragTable'),
//     //   name: 'dragTable',
//     //   meta: { title: 'dragTable' }
//     // },
//     // {
//     //   path: 'inline-edit-table',
//     //   component: () => import('@/views/table/inlineEditTable'),
//     //   name: 'inlineEditTable',
//     //   meta: { title: 'inlineEditTable' }
//     // },
//     {
//       path: 'allTask',
//       component: () => import('@/views/table/allmyTask'),
//       name: 'allmyTask',
//       meta: {
//         title: 'testing'
//       }
//     },
//     {
//       path: 'complex-table',
//       component: () => import('@/views/table/complexTable'),
//       name: 'complexTable',
//       meta: {
//         title: 'complexTable'
//       }
//     }
//   ]
// }
// export default tableRouter
