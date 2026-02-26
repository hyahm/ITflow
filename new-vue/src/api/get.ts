import { api } from 'boot/axios';
import type { RequestProject } from './types/type';

export function getStatus() {
  return api({
    url: '/get/status',
    method: 'post',
  });
}

export function getShowStatus() {
  return api({
    url: '/status/show',
    method: 'post',
  });
}

export function getEnv() {
  return api({
    url: '/get/env',
    method: 'post',
  });
}

export function getProject() {
  return api({
    url: '/get/project',
    method: 'post',
  });
}

export function getProjectKeyName() {
  return api({
    url: '/project/keyname',
    method: 'post',
  });
}

// export function getProjectKeyName() {
//   return api({
//     url: "/get/myproject",
//     method: "post"
//   });
// }

// 获取某项目的用户真实姓名
export function getUsers() {
  return api({
    url: '/get/user',
    method: 'post',
  });
}

// 获取某项目的用户真实姓名
export function getUserKeyNameByProject(data: RequestProject) {
  return api({
    url: '/user/keyname/byproject',
    method: 'post',
    data,
  });
}

// 获取用户真实姓名
export function getUserKeyName() {
  return api({
    url: '/user/keyname',
    method: 'post',
  });
}
// 获取某项目的用户真实姓名
export function getVersionKeyNameByProject(data: RequestProject) {
  return api({
    url: '/version/keyname/byproject',
    method: 'post',
    data,
  });
}

export function getVersionKeyName() {
  return api({
    url: '/version/keyname',
    method: 'post',
  });
}

export function getVersion() {
  return api({
    url: '/get/version',
    method: 'post',
  });
}

export function getOs() {
  return api({
    url: '/get/os',
    method: 'post',
  });
}

export function getRoles() {
  return api({
    url: '/get/role',
    method: 'post',
  });
}

export function getGroup() {
  return api({
    url: '/get/group',
    method: 'post',
  });
}

export function getPermStatus() {
  return api({
    url: '/get/permstatus',
    method: 'post',
  });
}

export function getThisRole(id: number) {
  return api({
    url: '/get/thisrole',
    method: 'get',
    params: { id },
  });
}

export function getImportants() {
  return api({
    url: '/important/keyname',
    method: 'post',
  });
}

export function getLevels() {
  return api({
    url: '/level/keyname',
    method: 'post',
  });
}

export function getProjectUser(id: number) {
  return api({
    url: `/project/user/${id}`,
    method: 'get',
  });
}

export function isAdmin() {
  return api({
    url: '/is/admin',
    method: 'post',
  });
}

export function getTyp() {
  return api({
    url: '/get/task/typ',
    method: 'post',
  });
}
