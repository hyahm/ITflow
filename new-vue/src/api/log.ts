import { api } from 'boot/axios';
import type { SearchLog } from './types/type';
// export function getLog(data) {
//   return api({
//     url: '/log/list',
//     method: 'post',
//     data: data
//   })
// }

export function searchLog(data: SearchLog) {
  return api({
    url: '/search/log',
    method: 'post',
    data: data,
  });
}

export function logClassify() {
  return api({
    url: '/log/classify',
    method: 'post',
  });
}
