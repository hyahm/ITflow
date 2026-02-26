import { api } from 'boot/axios';
import type { RequestRole } from './types/type';

export function getPagePerms() {
  return api({
    url: '/page/perm/list',
    method: 'get',
  });
}

export function addRoleGroup(data: RequestRole) {
  return api({
    url: '/rolegroup/add',
    method: 'post',
    data,
  });
}

export function getRoleGroupPerm(id: number) {
  return api({
    url: `/rolegroup/perm/get?id=${id}`,
    method: 'get',
  });
}

export function editRole(data: RequestRole) {
  return api({
    url: '/rolegroup/edit',
    method: 'post',
    data,
  });
}

export function getPermTemplate(data: RequestRole) {
  return api({
    url: '/rolegroup/template',
    method: 'post',
    data,
  });
}

export function removeRole(id: number) {
  return api({
    url: '/rolegroup/remove',
    method: 'get',
    params: { id },
  });
}

// export function roleList() {
//   return api({
//     url: `/rolegroup/list`,
//     method: 'post',
//   });
// }

export function roleList() {
  return api({
    url: `/role/list`,
    method: 'post',
  });
}

export function addRole(data: RequestRole) {
  return api({
    url: `/role/add`,
    method: 'post',
    data,
  });
}

export function delRole(id: number) {
  return api({
    url: `/role/delete`,
    method: 'get',
    params: { id },
  });
}