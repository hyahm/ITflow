import request from '@/utils/request'

export function getAllBugs(query) {
  return request({
    url: '/bug/getallbugs',
    method: 'post',
    data: query
  })
}

export function resumeBug(id) {
  return request({
    url: '/bug/resume',
    method: 'get',
    params: { id }
  })
}

export function getMyBugs(data) {
  return request({
    url: '/bug/mybugs',
    method: 'post',
    data: data
  })
}

export function searchbugs(query) {
  return request({
    url: '/bug/search',
    method: 'post',
    data: query
  })
}

export function changeStatus(query) {
  return request({
    url: '/bug/changestatus',
    method: 'post',
    data: query
  })
}

export function showBug(id) {
  return request({
    url: '/bug/show',
    method: 'get',
    params: { id }
  })
}

export function createBug(data) {
  return request({
    url: '/bug/create',
    method: 'post',
    data: data
  })
}

export function passBug(data) {
  return request({
    url: '/bug/pass',
    method: 'post',
    data
  })
}

export function taskList() {
  return request({
    url: '/task/list',
    method: 'post'
  })
}

export function closeBug(id) {
  return request({
    url: '/bug/close',
    method: 'get',
    params: { id }
  })
}

export function removeBug(id) {
  return request({
    url: '/bug/remove',
    method: 'get',
    params: { id }
  })
}

export function fetchBug(id) {
  return request({
    url: '/bug/edit',
    method: 'get',
    params: { id }
  })
}

export function delBug(id) {
  return request({
    url: '/bug/delete',
    method: 'get',
    params: { id }
  })
}

export function updateBug(data) {
  return request({
    url: '/bug/update',
    method: 'post',
    data
  })
}

