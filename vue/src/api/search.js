import request from '@/utils/request'

export function searchMyBugs(data) {
  return request({
    url: '/search/mybugs',
    method: 'post',
    data: data
  })
}

export function searchAllBugs(data) {
  return request({
    url: '/search/allbugs',
    method: 'post',
    data: data
  })
}

export function userSearch(name) {
  return request({
    url: '/search/user',
    method: 'get',
    params: { name }
  })
}

export function bugFilter(data) {
  return request({
    url: '/search/bugmanager',
    method: 'post',
    data: data
  })
}

export function searchMyTasks(data) {
  return request({
    url: '/search/mytasks',
    method: 'post',
    data: data
  })
}
