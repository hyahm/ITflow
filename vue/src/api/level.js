import request from '@/utils/request'

export function getLevels() {
  return request({
    url: '/level/get',
    method: 'post'
  })
}

export function addLevel(data) {
  return request({
    url: '/level/add',
    method: 'post',
    data: data
  })
}

export function delLevel(id) {
  return request({
    url: '/level/del',
    method: 'get',
    params: { id }
  })
}

export function updateLevel(data) {
  return request({
    url: '/level/update',
    method: 'post',
    data: data
  })
}
