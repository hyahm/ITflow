import { api } from 'boot/axios'

export function restList() {
  return api({
    url: '/rest/list',
    method: 'post'
  })
}

export function restDel(id) {
  return api({
    url: '/rest/delete',
    method: 'get',
    params: { id }
  })
}

export function restUpdate(data) {
  return api({
    url: '/rest/update',
    method: 'post',
    data: data
  })
}

export function restAdd(data) {
  return api({
    url: '/rest/add',
    method: 'post',
    data: data
  })
}

export function apiList(pid) {
  return api({
    url: '/api/list',
    method: 'get',
    params: { pid }
  })
}

export function apiDel(id) {
  return api({
    url: '/api/delete',
    method: 'get',
    params: { id }
  })
}

export function editOne(id) {
  return api({
    url: '/edit/one',
    method: 'get',
    params: { id }
  })
}

export function apiUpdate(data) {
  return api({
    url: '/api/update',
    method: 'post',
    data: data
  })
}

export function apiAdd(data) {
  return api({
    url: '/api/add',
    method: 'post',
    data: data
  })
}

export function apiOne(id) {
  return api({
    url: '/api/one',
    method: 'get',
    params: { id }
  })
}

export function apiResp(data) {
  return api({
    url: '/api/resp',
    method: 'post',
    data: data
  })
}
