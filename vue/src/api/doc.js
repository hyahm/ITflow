import request from '@/utils/request'

export function fetchList(data) {
    return request({
        url: '/doc/list',
        method: 'post',
        data
    })
}

export function addDoc(data) {
    return request({
        url: '/doc/create',
        method: 'post',
        data
    })
}

export function dropDoc(id) {
    return request({
        url: '/doc/drop',
        method: 'get',
        params: { id }
    })
}

export function showDocFiles(id) {
    return request({
        url: '/doc/files/show',
        method: 'get',
        params: { id }
    })
}

export function checkDomain(name) {
    return request({
        url: '/doc/check/name',
        method: 'get',
        params: { name }
    })
}

export function getContent(id, name) {
    return request({
        url: '/doc/getcontent',
        method: 'get',
        params: { id, name }
    })
}

export function saveFile(id, name, content) {
    return request({
        url: '/doc/file/save',
        method: 'get',
        params: { id, name, content }
    })
}

export function createFile(id, name) {
    return request({
        url: '/doc/file/create',
        method: 'get',
        params: { id, name }
    })
}

export function deleteFile(id, name) {
    return request({
        url: '/doc/file/delete',
        method: 'get',
        params: { id, name }
    })
}

export function docDownload(id) {
    return request({
        url: '/doc/download',
        method: 'get',
        params: { id }
    })
}

export function docUpload(id) {
    return request({
        url: '/doc/upload',
        method: 'get',
        params: { id }
    })
}

export function docUser(id) {
    return request({
        url: '/doc/user/list',
        method: 'get',
        params: { id }
    })
}

export function docAddUser(id, name) {
    return request({
        url: '/doc/user/add',
        method: 'get',
        params: { id, name }
    })
}

export function docDelUser(id, name) {
    return request({
        url: '/doc/user/del',
        method: 'get',
        params: { id, name }
    })
}

export function addGit(data) {
    return request({
        url: '/doc/add/git',
        method: 'post',
        data: data
    })
}

export function pullGit(id) {
    return request({
        url: '/doc/git/pull',
        method: 'get',
        params: { id }
    })
}

export function updateDoc(id) {
    return request({
        url: '/doc/update',
        method: 'get',
        params: { id }
    })
}