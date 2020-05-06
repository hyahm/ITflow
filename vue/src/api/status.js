import request from '@/utils/request'

export function addStatus(data) {
  return request({
    url: '/status/add',
    method: 'post',
    data: data
  })
}

export function getStatusList() {
  return request({
    url: '/status/list',
    method: 'post'
  })
}

export function removeStatus(id) {
  return request({
    url: '/status/remove',
    method: 'get',
    params: { id }
  })
}

export function updateStatus(data) {
  return request({
    url: '/status/update',
    method: 'post',
    data: data
  })
}

export function statusFilter(data) {
  return request({
    url: '/status/filter',
    method: 'post',
    data: data
  })
}

export function showStatus() {
  return request({
    url: '/status/show',
    method: 'post'
  })
}

export function getStatusName() {
  return request({
    url: '/status/groupname',
    method: 'post'
  })
}

