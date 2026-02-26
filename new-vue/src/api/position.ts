import { api } from 'boot/axios';
import type { Position } from './types/type';

export function positionsList() {
  return api({
    url: '/position/list',
    method: 'post',
  });
}

export function getPositions() {
  return api({
    url: '/get/positions',
    method: 'post',
  });
}

export function getPositionManager() {
  return api({
    url: '/position/manager',
    method: 'get',
  });
}

export function getManagerKeyName() {
  return api({
    url: '/manager/keyname',
    method: 'post',
  });
}

export function addPosition(data: Position) {
  return api({
    url: '/position/add',
    method: 'post',
    data: data,
  });
}

export function delPosition(id: number) {
  return api({
    url: '/position/del',
    method: 'get',
    params: { id },
  });
}

export function updatePosition(data: Position) {
  return api({
    url: '/position/update',
    method: 'post',
    data: data,
  });
}
