import request from '@/utils/request'

export function getProjectName() {
  return request({
    url: '/project/list',
    method: 'post'
  })
}

export function addProjectName(name) {
  return request({
    url: '/project/add',
    method: 'post',
    params: { name }
  })
}

export function updateProjectName(data) {
  return request({
    url: '/project/update',
    method: 'post',
    data: data
  })
}

export function deleteProjectName(id) {
  return request({
    url: '/project/delete',
    method: 'get',
    params: { id }
  })
}

