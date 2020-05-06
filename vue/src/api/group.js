import request from '@/utils/request'

export function getStatus() {
  return request({
    url: '/group/status',
    method: 'post'
  })
}

export function getMyStatus() {
  return request({
    url: '/group/mystatus',
    method: 'post'
  })
}

export function getEnv() {
  return request({
    url: '/group/env',
    method: 'post'
  })
}

export function getProject() {
  return request({
    url: '/group/project',
    method: 'post'
  })
}

export function getUsers() {
  return request({
    url: '/group/user',
    method: 'post'
  })
}

export function getVersion() {
  return request({
    url: '/group/version',
    method: 'post'
  })
}

export function getOs() {
  return request({
    url: '/get/os',
    method: 'post'
  })
}

export function getRoles() {
  return request({
    url: '/group/role',
    method: 'post'
  })
}

export function getGroup() {
  return request({
    url: '/group/get',
    method: 'post'
  })
}

export function getPermStatus() {
  return request({
    url: '/get/permstatus',
    method: 'post'
  })
}

export function getThisRole(id) {
  return request({
    url: '/group/thisrole',
    method: 'get',
    params: { id }
  })
}

export function addGroup(data) {
  return request({
    url: '/group/add',
    method: 'post',
    data: data
  })
}

export function delGroup(id) {
  return request({
    url: '/group/del',
    method: 'get',
    params: { id }
  })
}

export function updateGroup(data) {
  return request({
    url: '/group/update',
    method: 'post',
    data: data
  })
}

