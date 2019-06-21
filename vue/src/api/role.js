import request from '@/utils/request'

export function roleList() {
  return request({
    url: '/role/list',
    method: 'post'
  })
}

export function addRole(data) {
  return request({
    url: '/role/add',
    method: 'post',
    data: data
  })
}

export function editRole(data) {
  return request({
    url: '/role/edit',
    method: 'post',
    data: data
  })
}

export function removeRole(id) {
  return request({
    url: '/role/remove',
    method: 'get',
    params: { id }
  })
}

export function getRoles() {
  return request({
    url: '/role/get',
    method: 'post'
  })
}

export function roleGroupName() {
  return request({
    url: '/role/groupname',
    method: 'post'
  })
}
