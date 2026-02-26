import { api } from 'boot/axios';
import type { Status } from './types/type';

export function getLevels() {
  return api({
    url: '/level/get',
    method: 'post',
  });
}

export function addLevel(data: Status) {
  return api({
    url: '/level/add',
    method: 'post',
    data: data,
  });
}

export function delLevel(id: number) {
  return api({
    url: '/level/del',
    method: 'get',
    params: { id },
  });
}

export function updateLevel(data: Status) {
  return api({
    url: '/level/update',
    method: 'post',
    data: data,
  });
}
