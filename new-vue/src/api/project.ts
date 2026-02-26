import { api } from 'boot/axios';
import type { ProjectRequsest } from './types/type';

export function getProjectName() {
  return api({
    url: '/project/list',
    method: 'post',
  });
}

export function addProjectName(data: ProjectRequsest) {
  return api({
    url: '/project/add',
    method: 'post',
    data,
  });
}

export function updateProjectName(data: ProjectRequsest) {
  return api({
    url: '/project/update',
    method: 'post',
    data,
  });
}

export function deleteProjectName(id: number) {
  return api({
    url: '/project/delete',
    method: 'get',
    params: { id },
  });
}
