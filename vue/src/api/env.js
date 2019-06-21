import request from '@/utils/request'

export function getEnvName() {
  return request({
    url: '/env/list',
    method: 'post'
  })
}

export function addEnvName(name) {
  return request({
    url: '/env/add',
    method: 'get',
    params: { name }
  })
}

export function updateEnvName(data) {
  return request({
    url: '/env/update',
    method: 'post',
    data: data
  })
}

export function deleteEnvName(id) {
  return request({
    url: '/env/delete',
    method: 'get',
    params: { id }
  })
}

