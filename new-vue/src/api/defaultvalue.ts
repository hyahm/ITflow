import { api } from 'boot/axios';
import type { DefaultValue } from './types/type';

export function defaultValue() {
  return api({
    url: '/default/status',
    method: 'post',
  });
}

export function important() {
  return api({
    url: '/default/important',
    method: 'post',
  });
}

export function level() {
  return api({
    url: '/default/level',
    method: 'post',
  });
}

export function save(data: DefaultValue) {
  return api({
    url: '/default/save',
    method: 'post',
    data: data,
  });
}
