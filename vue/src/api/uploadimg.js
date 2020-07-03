import request from '@/utils/request'

export function uploadImg(data) {
  return request({
    url: '/uploadimg',
    method: 'post',
    data
  })
}
