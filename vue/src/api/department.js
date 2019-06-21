import request from '@/utils/request'

export function departmentList() {
  return request({
    url: '/department/list',
    method: 'post'
  })
}

export function addDepartment(data) {
  return request({
    url: '/department/add',
    method: 'post',
    data: data
  })
}

export function editDepartment(data) {
  return request({
    url: '/department/edit',
    method: 'post',
    data: data
  })
}

export function removeDepartment(id) {
  return request({
    url: '/department/remove',
    method: 'get',
    params: { id }
  })
}
