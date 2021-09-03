import request from '@/utils/request'

export function getEnvName() {
  return request({
    url: '/env/list',
    method: 'post'
  })
}

export function addEnvName(data) {
  return request({
    url: '/env/add',
    method: 'post',
    data
  })
}

export function updateEnvName(data) {
  return request({
    url: '/env/update',
    method: 'post',
    data
  })
}

export function deleteEnvName(id) {
  return request({
    url: '/env/delete',
    method: 'get',
    params: { id }
  })
}

