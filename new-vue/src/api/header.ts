import { api } from 'boot/axios'

export function headerList() {
  return api({
    url: '/header/list',
    method: 'post'
  })
}

export function headerDel(id: number) {
  return api({
    url: '/header/del',
    method: 'get',
    params: { id }
  })
}

// export function headerUpdate(data) {
//   return api({
//     url: '/header/update',
//     method: 'post',
//     data: data
//   })
// }

// export function headerAdd(data) {
//   return api({
//     url: '/header/add',
//     method: 'post',
//     data: data
//   })
// }

// export function headerGet() {
//   return api({
//     url: '/header/get',
//     method: 'post'
//   })
// }

