import request from '@/utils/request'

export function getImportant() {
  return request({
    url: '/important/get',
    method: 'post'
  })
}

export function addImportant(data) {
  return request({
    url: '/important/add',
    method: 'post',
    data: data
  })
}

export function delImportant(id) {
  return request({
    url: '/important/del',
    method: 'get',
    params: { id }
  })
}

export function updateImportant(data) {
  return request({
    url: '/important/update',
    method: 'post',
    data: data
  })
}
