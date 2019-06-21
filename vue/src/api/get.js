import request from '@/utils/request'

export function getStatus() {
  return request({
    url: '/get/status',
    method: 'post'
  })
}

export function getShowStatus() {
  return request({
    url: '/status/show',
    method: 'post'
  })
}

export function getEnv() {
  return request({
    url: '/get/env',
    method: 'post'
  })
}

export function getProject() {
  return request({
    url: '/get/project',
    method: 'post'
  })
}

export function getUsers() {
  return request({
    url: '/get/user',
    method: 'post'
  })
}

export function getVersion() {
  return request({
    url: '/get/version',
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
    url: '/get/role',
    method: 'post'
  })
}

export function getGroup() {
  return request({
    url: '/get/group',
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
    url: '/get/thisrole',
    method: 'get',
    params: { id }
  })
}

export function getImportants() {
  return request({
    url: '/get/importants',
    method: 'post'
  })
}

export function getLevels() {
  return request({
    url: '/get/levels',
    method: 'post'
  })
}
