import request from '@/utils/request'

export function statusGroupList() {
  return request({
    url: '/statusgroup/list',
    method: 'post'
  })
}

export function addStatusGroup(data) {
  return request({
    url: '/statusgroup/add',
    method: 'post',
    data: data
  })
}

export function editStatusGroup(data) {
  return request({
    url: '/statusgroup/edit',
    method: 'post',
    data: data
  })
}

export function removeStatusGroup(id) {
  return request({
    url: '/statusgroup/remove',
    method: 'get',
    params: { id }
  })
}

export function getStatusGroupName() {
  return request({
    url: '/statusgroup/keyname',
    method: 'post'
  })
}

