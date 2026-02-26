import { api } from 'boot/axios';
import type { Doc } from './types/type';

export function fetchList(data: Doc) {
  return api({
    url: '/doc/list',
    method: 'post',
    data,
  });
}

export function addDoc(data: Doc) {
  return api({
    url: '/doc/create',
    method: 'post',
    data,
  });
}

export function dropDoc(id: number) {
  return api({
    url: '/doc/drop',
    method: 'get',
    params: { id },
  });
}

export function showDocFiles(id: number) {
  return api({
    url: '/doc/files/show',
    method: 'get',
    params: { id },
  });
}

export function checkDomain(name: string) {
  return api({
    url: '/doc/check/name',
    method: 'get',
    params: { name },
  });
}

export function getContent(id: number, name: string) {
  return api({
    url: '/doc/getcontent',
    method: 'get',
    params: { id, name },
  });
}

export function saveFile(id: number, name: string, content: string) {
  return api({
    url: '/doc/file/save',
    method: 'get',
    params: { id, name, content },
  });
}

export function createFile(id: number, name: string) {
  return api({
    url: '/doc/file/create',
    method: 'get',
    params: { id, name },
  });
}

export function deleteFile(id: number, name: string) {
  return api({
    url: '/doc/file/delete',
    method: 'get',
    params: { id, name },
  });
}

export function docDownload(id: number) {
  return api({
    url: '/doc/download',
    method: 'get',
    params: { id },
  });
}

export function docUpload(id: number) {
  return api({
    url: '/doc/upload',
    method: 'get',
    params: { id },
  });
}

export function docUser(id: number) {
  return api({
    url: '/doc/user/list',
    method: 'get',
    params: { id },
  });
}

export function docAddUser(id: number, name: string) {
  return api({
    url: '/doc/user/add',
    method: 'get',
    params: { id, name },
  });
}

export function docDelUser(id: number, name: string) {
  return api({
    url: '/doc/user/del',
    method: 'get',
    params: { id, name },
  });
}

// export function addGit(data) {
//     return api({
//         url: '/doc/add/git',
//         method: 'post',
//         data: data
//     })
// }

// export function pullGit(id) {
//     return api({
//         url: '/doc/git/pull',
//         method: 'get',
//         params: { id }
//     })
// }

export function updateDoc(id: number) {
  return api({
    url: '/doc/update',
    method: 'get',
    params: { id },
  });
}
