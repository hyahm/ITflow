import request from '@/utils/request'

export function getProjectName() {
  return request({
    url: '/project/list',
    method: 'post'
  })
}

export function addProjectName(data) {
  return request({
    url: '/project/add',
    method: 'post',
    data
  })
}

export function updateProjectName(data) {
  return request({
    url: '/project/update',
    method: 'post',
    data
  })
}

export function deleteProjectName(id) {
  return request({
    url: '/project/delete',
    method: 'get',
    params: { id }
  })
}

