import { api } from 'boot/axios';
import type { Bug, RequestPass, ChangeStatus, ReqMyBugFilter } from './types/type';

export function getAllBugs(query: ReqMyBugFilter) {
  return api({
    url: '/bug/getallbugs',
    method: 'post',
    data: query,
  });
}

export function resumeBug(id: number) {
  return api({
    url: '/bug/resume',
    method: 'get',
    params: { id },
  });
}

export function getMyBugs(data: ReqMyBugFilter) {
  return api({
    url: '/bug/mybugs',
    method: 'post',
    data: data,
  });
}

// export function searchbugs(query: ReqMyBugFilter) {
//   return api({
//     url: '/bug/search',
//     method: 'post',
//     data: query
//   })
// }

export function changeStatus(query: ChangeStatus) {
  return api({
    url: '/bug/changestatus',
    method: 'post',
    data: query,
  });
}

export function showBug(id: number) {
  return api({
    url: '/bug/show',
    method: 'get',
    params: { id },
  });
}

export function createBug(data: Bug) {
  return api({
    url: '/bug/create',
    method: 'post',
    data: data,
  });
}

export function passBug(data: RequestPass) {
  return api({
    url: '/bug/pass',
    method: 'post',
    data,
  });
}

export function completeBug(data: Bug) {
  return api({
    url: '/bug/complete',
    method: 'post',
    data,
  });
}

export function receiveBug(data: Bug) {
  return api({
    url: '/bug/receive',
    method: 'post',
    data,
  });
}

export function taskList() {
  return api({
    url: '/task/list',
    method: 'post',
  });
}

export function closeBug(id: number) {
  return api({
    url: '/bug/close',
    method: 'get',
    params: { id },
  });
}

export function removeBug(id: number) {
  return api({
    url: '/bug/remove',
    method: 'get',
    params: { id },
  });
}

export function fetchBug(id: number) {
  return api({
    url: '/bug/edit',
    method: 'get',
    params: { id },
  });
}

export function delBug(id: number) {
  return api({
    url: '/bug/delete',
    method: 'get',
    params: { id },
  });
}

export function updateBug(data: Bug) {
  return api({
    url: '/bug/update',
    method: 'post',
    data,
  });
}
