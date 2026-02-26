import { api } from 'boot/axios';
import type { RequestRole } from './types/type';

export function getRoleGroup() {
  return api({
    url: '/rolegroup/get',
    method: 'post',
  });
}

export function addRole(data: RequestRole) {
  return api({
    url: '/rolegroup/add',
    method: 'post',
    data,
  });
}

export function editRole(data: RequestRole) {
  return api({
    url: '/rolegroup/edit',
    method: 'post',
    data,
  });
}
