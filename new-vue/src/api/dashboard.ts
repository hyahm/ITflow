import { api } from 'boot/axios'

export function getlist() {
  return api({
    url: '/dashboard/usercount',
    method: 'post'
  })
}

export function fetchArticle(id: number) {
  return api({
    url: '/article/detail',
    method: 'get',
    params: { id }
  })
}

export function fetchPv(pv: number) {
  return api({
    url: '/article/pv',
    method: 'get',
    params: { pv }
  })
}

export function getprojectlist() {
  return api({
    url: '/dashboard/projectcount',
    method: 'post'
  })
}

export function getBugCount() {
  return api({
    url: '/dashboard/bugcount',
    method: 'post'
  })
}
