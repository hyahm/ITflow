import { defineBoot } from '#q-app/wrappers';
import axios, { type AxiosInstance } from 'axios';
import { LocalStorage, Notify } from 'quasar';

declare module 'vue' {
  interface ComponentCustomProperties {
    $axios: AxiosInstance;
    $api: AxiosInstance;
  }
}

// Be careful when using SSR for cross-request state pollution
// due to creating a Singleton instance here;
// If any client changes this (global) instance, it might be a
// good idea to move this instance creation inside of the
// "export default () => {}" function below (which runs individually
// for each client)
const api = axios.create({ baseURL: 'http://192.168.60.129:10001' });

api.interceptors.request.use((config) => {
  const token = LocalStorage.getItem('token') as string;
  if (token) config.headers.Authorization = `Bearer ${token}`;
  return config;
});

/* 响应拦截：全局错误 + 自动退登 */
api.interceptors.response.use(
  (res) => {
    if (res.data.code === 401) {
      LocalStorage.remove('token');
      Notify.create({ type: 'negative', message: '登录已失效' });
      // 这里拿不到 router，用 window 硬跳或下面把 router 传进来
      console.log('negative');
      window.location.replace('/login');
    }
    if (res.data.code != 0) {
      Notify.create({ type: 'error', message: res.data.msg });
      // 这里拿不到 router，用 window 硬跳或下面把 router 传进来
      // window.location.replace('/login');
      // return Promise.reject(new Error(res.data.msg));
    }

    return res;
  },
  (err) => {
    return Promise.reject(new Error(err));
  },
);

export default defineBoot(({ app }) => {
  // for use inside Vue files (Options API) through this.$axios and this.$api

  app.config.globalProperties.$axios = axios;
  // ^ ^ ^ this will allow you to use this.$axios (for Vue Options API form)
  //       so you won't necessarily have to import axios in each vue file

  app.config.globalProperties.$api = api;
  // ^ ^ ^ this will allow you to use this.$api (for Vue Options API form)
  //       so you can easily perform requests against your app's API
});

export { api };
