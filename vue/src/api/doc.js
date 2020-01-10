import request from '@/utils/request'

export function addProject() {
  return request({
    url: '/doc/add/project',
    method: 'post'
  })
}

export function addDepartment(data) {
  return request({
    url: '/doc/delete/project',
    method: 'post',
    data: data
  })
}
