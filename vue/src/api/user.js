import request from '@/utils/request'

export function createUser(data) {
  return request({
    url: '/user/create',
    method: 'post',
    data: data
  })
}

export function updatePassword(data) {
  return request({
    url: '/password/update',
    method: 'post',
    data: data
  })
}

export function userList() {
  return request({
    url: '/user/list',
    method: 'post'
  })
}

export function resetPwd(data) {
  return request({
    url: '/password/reset',
    method: 'post',
    data: data
  })
}

export function updateInfo(data) {
  return request({
    url: '/info/update',
    method: 'post',
    data: data
  })
}

export function getInfo() {
  return request({
    url: '/get/info',
    method: 'post'
  })
}

export function updateUser(data) {
  return request({
    url: '/user/update',
    method: 'post',
    data: data
  })
}

export function userRemove(id) {
  return request({
    url: '/user/remove',
    method: 'get',
    params: { id }
  })
}

export function userDisable(id) {
  return request({
    url: '/user/disable',
    method: 'get',
    params: { id }
  })
}
