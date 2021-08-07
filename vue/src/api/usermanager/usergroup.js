import request from "@/utils/request";

// 获取所有用户keyname
export function getAllUserKeyName() {
  return request({
    url: "/alluser/keyname",
    method: "post"
  });
}

// 获取所有用户组
export function getUserGroups() {
  return request({
    url: "/usergroup/list",
    method: "post"
  });
}

// 修改用户组
export function updateUserGroup(data) {
  return request({
    url: "/usergroup/update",
    method: "post",
    data
  });
}

// 创建用户组
export function createUserGroups(data) {
  return request({
    url: "/usergroup/create",
    method: "post",
    data
  });
}

// 删除用户组
export function deleteUserGroups(id) {
  return request({
    url: `/usergroup/delete?id=${id}`,
    method: "get"
  });
}
