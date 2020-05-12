import request from '@/utils/request'

export function getRoles() {
  return request({
    url: '/roles/get',
    method: 'get'
  })
}

export function addRole(data) {
  return request({
    url: '/rolegroup/add',
    method: 'post',
    data
  })
}

export function editRole(data) {
  return request({
    url: '/rolegroup/edit',
    method: 'post',
    data
  })
}

export function removeRole(id) {
  return request({
    url: '/rolegroup/remove',
    method: 'get',
    params: { id }
  })
}

export function roleList() {
  return request({
    url: `/rolegroup/list`,
    method: 'post'
  })
}

export function roleGroupName() {
  return request({
    url: `/role/list`,
    method: 'post'
  })
}
