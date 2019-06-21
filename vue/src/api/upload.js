import request from '@/utils/request'

export function uploadImg() {
  return request({
    url: '/get/status',
    method: 'post'
  })
}
