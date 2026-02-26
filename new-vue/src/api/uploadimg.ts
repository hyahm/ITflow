import { api } from 'boot/axios';

export function uploadImg(data: FormData) {
  return api({
    url: '/uploadimg',
    method: 'post',
    data,
  });
}
