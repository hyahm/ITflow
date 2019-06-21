import request from '@/utils/request'

export function PositionsList() {
  return request({
    url: '/position/list',
    method: 'post'
  })
}

export function getPositions() {
  return request({
    url: '/get/positions',
    method: 'post'
  })
}

export function addPosition(data) {
  return request({
    url: '/position/add',
    method: 'post',
    data: data
  })
}

export function delPosition(id) {
  return request({
    url: '/position/del',
    method: 'get',
    params: { id }
  })
}

export function updatePosition(data) {
  return request({
    url: '/position/update',
    method: 'post',
    data: data
  })
}

export function getHypos() {
  return request({
    url: '/get/hypos',
    method: 'post'
  })
}
