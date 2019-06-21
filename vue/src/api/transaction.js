import request from '@/utils/request'

export function fetchList(query) {
  console.log(query)
  return request({
    url: '/transaction/list',
    method: 'get',
    params: query
  })
}
