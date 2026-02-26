import { api } from 'boot/axios';
import type { Email } from './types/type';

export function testEmail(data: Email) {
  return api({
    url: '/email/test',
    method: 'post',
    data,
  });
}

export function saveEmail(data: Email) {
  return api({
    url: '/email/save',
    method: 'post',
    data,
  });
}

export function getEmailStatus() {
  return api({
    url: '/email/get',
    method: 'post',
  });
}
