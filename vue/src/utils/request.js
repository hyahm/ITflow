import axios from 'axios'
import { Message } from 'element-ui'
import { getToken, setTimeout, removeToken } from '@/utils/auth'
import store from '@/store'

// create an axios instance
console.log(process.env.VUE_APP_BASE_API)
const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_API, // url = base url + request url
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 5000 // request timeout
})

// request interceptor
service.interceptors.request.use(
  config => {
    // do something before request is sent

    if (store.getters.token) {
      config.headers['Authorization'] = 'Bearer ' + getToken()
      setTimeout()
    }
    return config
  },
  error => {
    // do something with request error
    Message({
      message: error.message,
      type: 'error',
      duration: 5 * 1000
    })
    // for debug
    return Promise.reject(error)
  }
)

// response interceptor
service.interceptors.response.use(
  /**
   * If you want to get http information such as headers or status
   * Please return  response => response
  */

  /**
   * Determine the request status by custom code
   * Here is just an example
   * You can also judge the status by HTTP Status Code
   */
  response => {
    const url = response.request.responseURL
    if (url.indexOf('/login/login') < 0 && url.indexOf('/api/resp') < 0) {
      if (response.data.code === 20) {
        removeToken()
        location.href = '/login'
      } else if (response.data.code === 10) {
        Message({
          message: response.data.message,
          type: 'error',
          duration: 5 * 1000
        })
        return
      }
    }
    return response
  },
  error => {
    ('err' + error) // for debug
    Message({
      message: error.message,
      type: 'error',
      duration: 5 * 1000
    })
    return Promise.reject(error)
  }
)

export default service
