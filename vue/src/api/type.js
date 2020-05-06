import request from '@/utils/request'

export function typeList() {
  return request({
    url: '/type/list',
    method: 'post'
  })
}

export function typeDel(id) {
  return request({
    url: '/type/delete',
    method: 'get',
    params: { id }
  })
}

export function typeUpdate(data) {
  return request({
    url: '/type/update',
    method: 'post',
    data: data
  })
}

export function typeAdd(data) {
  return request({
    url: '/type/add',
    method: 'post',
    data: data
  })
}

export function typeGet() {
  return request({
    url: '/type/get',
    method: 'get'
  })
}

export function getUpdateType() {
  return request({
    url: '/type/updateget',
    method: 'get'
  })
}
