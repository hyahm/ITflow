import request from '@/utils/request'

export function getRoleGroup() {
  return request({
    url: '/rolegroup/get',
    method: 'post'
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

