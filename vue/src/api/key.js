import request from '@/utils/request'

export function fetchList() {
    return request({
        url: '/keys/list',
        method: 'post'
    })
}


export function addKey(data) {
    return request({
        url: '/keys/add',
        method: 'post',
        data
    })
}

export function delKey(id) {
    return request({
        url: "/keys/delete",
        method: 'get',
        params: { id }
    })
}


export function checkName(name) {
    return request({
        url: '/keys/check/name',
        method: 'get',
        params: { name }
    })
}


export function getMykeys() {
    return request({
        url: '/keys/get/me',
        method: 'post'
    })
}