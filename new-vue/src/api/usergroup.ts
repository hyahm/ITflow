import { api } from 'boot/axios';
import type { UserGroup } from './types/type';

// 获取所有用户keyname
export function getAllUserKeyName() {
  return api({
    url: '/alluser/keyname',
    method: 'post',
  });
}

// 获取所有用户组
export function getUserGroups() {
  return api({
    url: '/usergroup/list',
    method: 'post',
  });
}

// 修改用户组
export function updateUserGroup(data: UserGroup) {
  return api({
    url: '/usergroup/update',
    method: 'post',
    data,
  });
}

// 创建用户组
export function createUserGroups(data: UserGroup) {
  return api({
    url: '/usergroup/create',
    method: 'post',
    data,
  });
}

// 删除用户组
export function deleteUserGroups(id: number) {
  return api({
    url: `/usergroup/delete?id=${id}`,
    method: 'get',
  });
}
