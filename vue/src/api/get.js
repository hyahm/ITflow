import request from "@/utils/request";

export function getStatus() {
  return request({
    url: "/get/status",
    method: "post"
  });
}

export function getShowStatus() {
  return request({
    url: "/status/show",
    method: "post"
  });
}

export function getEnv() {
  return request({
    url: "/get/env",
    method: "post"
  });
}

export function getProject() {
  return request({
    url: "/get/project",
    method: "post"
  });
}

export function getProjectKeyName() {
  return request({
    url: "/project/keyname",
    method: "post"
  });
}

export function getMyProject() {
  return request({
    url: "/get/myproject",
    method: "post"
  });
}

// 获取某项目的用户真实姓名
export function getUsers(data) {
  return request({
    url: "/get/user",
    method: "post",
    data
  });
}

// 获取某项目的用户真实姓名
export function getUserKeyNameByProject(data) {
  return request({
    url: "/user/keyname/byproject",
    method: "post",
    data
  });
}

// 获取用户真实姓名
export function getUserKeyName() {
  return request({
    url: "/user/keyname",
    method: "post"
  });
}
// 获取某项目的用户真实姓名
export function getVersionKeyNameByProject(data) {
  return request({
    url: "/version/keyname/byproject",
    method: "post",
    data
  });
}

export function getVersionKeyName() {
  return request({
    url: "/version/keyname",
    method: "post"
  });
}

export function getVersion() {
  return request({
    url: "/get/version",
    method: "post"
  });
}

export function getOs() {
  return request({
    url: "/get/os",
    method: "post"
  });
}

export function getEnvKeyName() {
  return request({
    url: "/env/keyname",
    method: "post"
  });
}

export function getRoles() {
  return request({
    url: "/get/role",
    method: "post"
  });
}

export function getGroup() {
  return request({
    url: "/get/group",
    method: "post"
  });
}

export function getPermStatus() {
  return request({
    url: "/get/permstatus",
    method: "post"
  });
}

export function getThisRole(id) {
  return request({
    url: "/get/thisrole",
    method: "get",
    params: { id }
  });
}

export function getImportants() {
  return request({
    url: "/important/keyname",
    method: "post"
  });
}

export function getLevels() {
  return request({
    url: "/level/keyname",
    method: "post"
  });
}

export function getProjectUser(id) {
  return request({
    url: `/project/user/${id}`,
    method: "get"
  });
}

export function isAdmin() {
  return request({
    url: "/is/admin",
    method: "post"
  });
}

export function getTyp() {
  return request({
    url: "/get/task/typ",
    method: "post"
  });
}
