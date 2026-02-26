import { api } from 'boot/axios';
import type { ChangePasswod, User, Login } from './types/type';

export function login(data: Login) {
  return api({
    url: '/user/login',
    method: 'post',
    data,
  });
}

export function createUser(data: User) {
  return api({
    url: '/user/create',
    method: 'post',
    data,
  });
}

export function getInfo() {
  return api({
    url: '/user/info',
    method: 'get',
  });
}

export function logout() {
  return api({
    url: '/user/logout',
    method: 'post',
  });
}

export function updatePassword(data: ChangePasswod) {
  return api({
    url: '/password/update',
    method: 'post',
    data,
  });
}

export function userList() {
  return api({
    url: '/user/list',
    method: 'post',
  });
}

export function resetPwd(data: ChangePasswod) {
  return api({
    url: '/password/reset',
    method: 'post',
    data,
  });
}

export function userRemove(id: number) {
  return api({
    url: '/user/remove',
    method: 'get',
    params: { id },
  });
}

export function updateUser(data: User) {
  return api({
    url: '/user/update',
    method: 'post',
    data,
  });
}

export function getEmail() {
  return api({
    url: '/email/get',
    method: 'post',
  });
}

export function setEmail(email: string) {
  return api({
    url: '/email/update',
    method: 'get',
    params: { email },
  });
}

export function userDisable(id: number) {
  return api({
    url: '/user/disable',
    method: 'get',
    params: { id },
  });
}

export function myEmail() {
  return api({
    url: '/my/email',
    method: 'post',
  });
}
