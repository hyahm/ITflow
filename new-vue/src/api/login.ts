import { api } from 'boot/axios'

// export function loginByUsername(data) {
//   return api({
//     url: '/login/login',
//     method: 'post',
//     data: data
//   })
// }

export function logout() {
  return api({
    url: '/login/logout',
    method: 'post'
  })
}

