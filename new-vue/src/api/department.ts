import { api } from 'boot/axios';
import type { StatusGroup } from './types/type';

// export function departmentList() {
//   return api({
//     url: '/department/list',
//     method: 'post'
//   })
// }

export function addDepartment(data: StatusGroup) {
  return api({
    url: '/department/add',
    method: 'post',
    data: data,
  });
}

export function editDepartment(data: StatusGroup) {
  return api({
    url: '/department/edit',
    method: 'post',
    data: data,
  });
}

export function removeDepartment(id: number) {
  return api({
    url: '/department/remove',
    method: 'get',
    params: { id },
  });
}
