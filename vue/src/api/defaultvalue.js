import request from '@/utils/request'

export function status() {
  return request({
    url: '/default/status',
    method: 'post'
  })
}

export function important() {
  return request({
    url: '/default/important',
    method: 'post'
  })
}

export function level() {
  return request({
    url: '/default/level',
    method: 'post'
  })
}

export function save(data) {
  return request({
    url: '/default/save',
    method: 'post',
    data: data
  })
}
