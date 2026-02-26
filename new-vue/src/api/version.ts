import { api } from 'boot/axios';
import type { Version } from './types/type';

export function addVersion(data: Version) {
  return api({
    url: '/version/add',
    method: 'post',
    data: data,
  });
}

export function getVersion() {
  return api({
    url: '/version/list',
    method: 'post',
  });
}

export function removeVersion(id: number) {
  return api({
    url: '/version/remove',
    method: 'get',
    params: { id },
  });
}

export function updateVersion(data: Version) {
  return api({
    url: '/version/update',
    method: 'post',
    data: data,
  });
}
