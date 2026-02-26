import { api } from 'boot/axios';
import type { Status } from './types/type';

export function getEnvName() {
  return api({
    url: '/env/list',
    method: 'post',
  });
}

export function addEnvName(data: Status) {
  return api({
    url: '/env/add',
    method: 'post',
    data,
  });
}

export function updateEnvName(data: Status) {
  return api({
    url: '/env/update',
    method: 'post',
    data,
  });
}

export function deleteEnvName(id: number) {
  return api({
    url: '/env/delete',
    method: 'get',
    params: { id },
  });
}
