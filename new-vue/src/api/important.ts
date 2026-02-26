import { api } from 'boot/axios';
import type { Status } from './types/type';

export function getImportant() {
  return api({
    url: '/important/get',
    method: 'post',
  });
}

export function addImportant(data: Status) {
  return api({
    url: '/important/add',
    method: 'post',
    data: data,
  });
}

export function delImportant(id: number) {
  return api({
    url: '/important/del',
    method: 'get',
    params: { id },
  });
}

export function updateImportant(data: Status) {
  return api({
    url: '/important/update',
    method: 'post',
    data: data,
  });
}
