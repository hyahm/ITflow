import { api } from 'boot/axios';

import type { Status, User } from './types/type';
export function addStatus(data: Status) {
  return api({
    url: '/status/add',
    method: 'post',
    data: data,
  });
}

export function getStatusList() {
  return api({
    url: '/status/list',
    method: 'post',
  });
}

export function removeStatus(id: number) {
  return api({
    url: '/status/remove',
    method: 'get',
    params: { id },
  });
}

export function updateStatus(data: Status) {
  return api({
    url: '/status/update',
    method: 'post',
    data,
  });
}

export function statusFilter(data: User) {
  return api({
    url: '/status/save',
    method: 'post',
    data,
  });
}

export function showStatus() {
  return api({
    url: '/status/show',
    method: 'post',
  });
}

export function getStatusName() {
  return api({
    url: '/status/groupname',
    method: 'post',
  });
}
