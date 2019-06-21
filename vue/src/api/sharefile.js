import request from '@/utils/request'

export function shareList(path) {
  return request({
    url: '/share/list',
    method: 'get',
    params: { path }
  })
}

export function mkDir(path) {
  return request({
    url: '/share/mkdir',
    method: 'post',
    data: path
  })
}

export function removeFile(id) {
  return request({
    url: '/share/remove',
    method: 'get',
    params: { id }
  })
}

export function renameFile(path) {
  return request({
    url: '/share/rename',
    method: 'post',
    data: path
  })
}

export function downloadFile(path) {
  return request({
    url: '/share/down',
    method: 'post',
    data: path
  })
}
