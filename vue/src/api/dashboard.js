import request from '@/utils/request'

export function getlist() {
  return request({
    url: '/dashboard/usercount',
    method: 'post'
  })
}

export function fetchArticle(id) {
  return request({
    url: '/article/detail',
    method: 'get',
    params: { id }
  })
}

export function fetchPv(pv) {
  return request({
    url: '/article/pv',
    method: 'get',
    params: { pv }
  })
}

export function getprojectlist() {
  return request({
    url: '/dashboard/projectcount',
    method: 'post'
  })
}

