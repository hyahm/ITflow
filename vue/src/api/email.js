import request from '@/utils/request'

export function testEmail(data) {
  return request({
    url: '/email/test',
    method: 'post',
    data: data
  })
}

export function saveEmail(data) {
  return request({
    url: '/email/save',
    method: 'post',
    data: data
  })
}

export function getEmailStatus() {
  return request({
    url: '/email/get',
    method: 'post'
  })
}

