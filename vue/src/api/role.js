import request from '@/utils/request'

export function getRoutes() {
  return request({
    url: '/routes',
    method: 'get'
  })
}

export function getRoles() {
  return request({
    url: '/role/get',
    method: 'get'
  })
}

export function addRole(data) {
  return request({
    url: '/role/add',
    method: 'post',
    data
  })
}

export function editRole(id, data) {
  return request({
    url: '/role/edit',
    method: 'post',
    data
  })
}

export function removeRole(id) {
  return request({
    url: '/role/remove',
    method: 'get',
    params: { id }
  })
}

export function roleList() {
  return request({
    url: `/role/list`,
    method: 'post'
  })
}

export function roleGroupName() {
  return request({
    url: `/role/list`,
    method: 'post'
  })
}
