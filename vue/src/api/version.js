import request from '@/utils/request'

export function addVersion(data) {
  return request({
    url: '/version/add',
    method: 'post',
    data: data
  })
}

export function getVersion(data) {
  return request({
    url: '/version/list',
    method: 'post',
    data: data
  })
}

export function removeVersion(id) {
  return request({
    url: '/version/remove',
    method: 'get',
    params: { id }
  })
}

export function updateVersion(data) {
  return request({
    url: '/version/update',
    method: 'post',
    data: data
  })
}
