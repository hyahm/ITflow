import { api } from 'boot/axios';
import type { ReqMyBugFilter } from './types/type';

// export function searchMyBugs(data: ReqMyBugFilter) {
//   return api({
//     url: '/search/mybugs',
//     method: 'post',
//     data: data,
//   });
// }

export function searchAllBugs(data: ReqMyBugFilter) {
  return api({
    url: '/search/allbugs',
    method: 'post',
    data: data,
  });
}

// export function userSearch(name) {
//   return api({
//     url: '/search/user',
//     method: 'get',
//     params: { name },
//   });
// }

// export function bugFilter(data: ReqMyBugFilter) {
//   return api({
//     url: '/search/bugmanager',
//     method: 'post',
//     data: data,
//   });
// }

// export function searchMyTasks(data: ReqMyBugFilter) {
//   return api({
//     url: '/search/mytasks',
//     method: 'post',
//     data: data,
//   });
// }
