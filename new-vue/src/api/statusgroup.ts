import { api } from 'boot/axios';
import type { StatusGroup } from './types/type';

export function statusGroupList() {
  return api({
    url: '/statusgroup/list',
    method: 'post',
  });
}

export function addStatusGroup(data: StatusGroup) {
  return api({
    url: '/statusgroup/add',
    method: 'post',
    data: data,
  });
}

export function editStatusGroup(data: StatusGroup) {
  return api({
    url: '/statusgroup/edit',
    method: 'post',
    data: data,
  });
}

export function removeStatusGroup(id: number) {
  return api({
    url: '/statusgroup/remove',
    method: 'get',
    params: { id },
  });
}

export function getStatusGroupName() {
  return api({
    url: '/statusgroup/keyname',
    method: 'post',
  });
}
