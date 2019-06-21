import request from '@/utils/request'

export function headerList() {
  return request({
    url: '/header/list',
    method: 'post'
  })
}

export function headerDel(id) {
  return request({
    url: '/header/del',
    method: 'get',
    params: { id }
  })
}

export function headerUpdate(data) {
  return request({
    url: '/header/update',
    method: 'post',
    data: data
  })
}

export function headerAdd(data) {
  return request({
    url: '/header/add',
    method: 'post',
    data: data
  })
}

export function headerGet() {
  return request({
    url: '/header/get',
    method: 'post'
  })
}

