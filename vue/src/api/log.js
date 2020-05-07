import request from '@/utils/request'

export function getLog(data) {
  return request({
    url: '/log/list',
    method: 'post',
    data: data
  })
}

export function searchLog(data) {
  return request({
    url: '/search/log',
    method: 'post',
    data: data
  })
}

export function logClassify() {
  return request({
    url: '/log/classify',
    method: 'post'
  })
}
