import request from '@/utils/request'

export function login(data) {
    return request({
        url: '/user/login',
        method: 'post',
        data
    })
}

export function createUser(data) {
    return request({
        url: '/user/create',
        method: 'post',
        data
    })
}

export function getInfo(token) {
    return request({
        url: '/user/info',
        method: 'get',
        params: { token }
    })
}

export function logout() {
    return request({
        url: '/user/logout',
        method: 'post'
    })
}

export function updatePassword(data) {
    return request({
        url: '/password/update',
        method: 'post',
        data
    })
}

export function userList(data) {
    return request({
        url: '/user/list',
        method: 'post',
        data
    })
}

export function resetPwd(data) {
    return request({
        url: '/password/reset',
        method: 'post',
        data
    })
}

export function userRemove(id) {
    return request({
        url: '/user/remove',
        method: 'get',
        params: { id }
    })
}

export function updateUser(data) {
    return request({
        url: '/user/update',
        method: 'post',
        data
    })
}

export function getEmail() {
    return request({
        url: '/email/get',
        method: 'post'
    })
}

export function setEmail(data) {
    return request({
        url: '/email/update',
        method: 'post',
        data
    })
}

export function userDisable(id) {
    return request({
        url: '/user/disable',
        method: 'get',
        params: { id }
    })
}