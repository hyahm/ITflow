import { api } from 'boot/axios';
import type { Auth } from './types/type';

export function fetchList() {
  return api({
    url: '/keys/list',
    method: 'post',
  });
}

export function addKey(data: Auth) {
  return api({
    url: '/keys/add',
    method: 'post',
    data,
  });
}

export function delKey(id: number) {
  return api({
    url: '/keys/delete',
    method: 'get',
    params: { id },
  });
}

export function checkName(name: string) {
  return api({
    url: '/keys/check/name',
    method: 'get',
    params: { name },
  });
}

export function getMykeys() {
  return api({
    url: '/keys/get/me',
    method: 'post',
  });
}
