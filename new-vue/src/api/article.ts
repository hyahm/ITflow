import { api } from "boot/axios";

// export function fetchList(query) {
//   return api({
//     url: "/article/list",
//     method: "get",
//     params: query
//   });
// }

export function fetchArticle(id: number) {
  return api({
    url: "/article/detail",
    method: "get",
    params: { id }
  });
}

// export function testtest(data) {
//   return api({
//     url: "/test/test",
//     method: "post",
//     data
//   });
// }

// export function fetchPv(pv) {
//   return api({
//     url: "/article/pv",
//     method: "get",
//     params: { pv }
//   });
// }

// export function createArticle(data) {
//   return api({
//     url: "/article/create",
//     method: "post",
//     data
//   });
// }

// export function updateArticle(data) {
//   return api({
//     url: "/article/update",
//     method: "post",
//     data
//   });
// }
