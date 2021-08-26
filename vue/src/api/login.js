import request from '@/utils/request'

export function loginByUsername(data) {
  return request({
    url: '/login/login',
    method: 'post',
    data: data
  })
}

export function logout() {
  return request({
    url: '/login/logout',
    method: 'post'
  })
}

