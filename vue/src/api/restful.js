import request from '@/utils/request'

export function restList() {
  return request({
    url: '/rest/list',
    method: 'post'
  })
}

export function restDel(id) {
  return request({
    url: '/rest/delete',
    method: 'get',
    params: { id }
  })
}

export function restUpdate(data) {
  return request({
    url: '/rest/update',
    method: 'post',
    data: data
  })
}

export function restAdd(data) {
  return request({
    url: '/rest/add',
    method: 'post',
    data: data
  })
}

export function apiList(pid) {
  return request({
    url: '/api/list',
    method: 'get',
    params: { pid }
  })
}

export function apiDel(id) {
  return request({
    url: '/api/delete',
    method: 'get',
    params: { id }
  })
}

export function editOne(id) {
  return request({
    url: '/edit/one',
    method: 'get',
    params: { id }
  })
}

export function apiUpdate(data) {
  return request({
    url: '/api/update',
    method: 'post',
    data: data
  })
}

export function apiAdd(data) {
  return request({
    url: '/api/add',
    method: 'post',
    data: data
  })
}

export function apiOne(id) {
  return request({
    url: '/api/one',
    method: 'get',
    params: { id }
  })
}

export function apiResp(data) {
  return request({
    url: '/api/resp',
    method: 'post',
    data: data
  })
}
