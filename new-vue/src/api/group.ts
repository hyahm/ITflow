import { api } from "boot/axios";

export function getStatus() {
  return api({
    url: "/group/status",
    method: "post"
  });
}

export function getMyStatus() {
  return api({
    url: "/group/mystatus",
    method: "post"
  });
}

export function getEnv() {
  return api({
    url: "/group/env",
    method: "post"
  });
}

export function getProject() {
  return api({
    url: "/group/project",
    method: "post"
  });
}

export function getUsers() {
  return api({
    url: "/group/user",
    method: "post"
  });
}

export function getVersion() {
  return api({
    url: "/group/version",
    method: "post"
  });
}

export function getOs() {
  return api({
    url: "/get/os",
    method: "post"
  });
}

export function getGroupRoles() {
  return api({
    url: "/group/role",
    method: "post"
  });
}

export function getGroup() {
  return api({
    url: "/group/get",
    method: "post"
  });
}

export function getUserGroupNames() {
  return api({
    url: "/usergroup/keyname",
    method: "post"
  });
}

export function getPermStatus() {
  return api({
    url: "/get/permstatus",
    method: "post"
  });
}

export function getThisRole(id: number) {
  return api({
    url: "/group/thisrole",
    method: "get",
    params: { id }
  });
}

// export function addGroup(data) {
//   return api({
//     url: "/group/add",
//     method: "post",
//     data: data
//   });
// }

export function delGroup(id: number) {
  return api({
    url: "/group/del",
    method: "get",
    params: { id }
  });
}

// export function updateGroup(data) {
//   return api({
//     url: "/group/update",
//     method: "post",
//     data: data
//   });
// }
