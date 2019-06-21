import request from '@/utils/request'

export function createuser(data) {
  return request({
    url: '/setting/createuser',
    method: 'post',
    data: data
  })
}

export function getroles(data) {
  return request({
    url: '/setting/getroles',
    method: 'post',
    data: data
  })
}

export function chpwd(data) {
  return request({
    url: '/setting/chpwd',
    method: 'post',
    data: data
  })
}

export function getuserlist() {
  return request({
    url: '/setting/userlist',
    method: 'post'
  })
}

export function resetpwd(id) {
  return request({
    url: '/setting/resetpwd',
    method: 'get',
    params: { id }
  })
}

