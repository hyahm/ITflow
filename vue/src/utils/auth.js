import Cookies from 'js-cookie'

const TokenKey = 'Admin-Token'
const inFifteenMinutes = new Date(new Date().getTime() + 120 * 60 * 1000 * 24)
export function getToken() {
  return Cookies.get(TokenKey)
}

export function setToken(token) {
  return Cookies.set(TokenKey, token)
}

export function removeToken() {
  return Cookies.remove(TokenKey)
}

export function setTimeout() {
  return Cookies.set(TokenKey, getToken(TokenKey), {
    expires: inFifteenMinutes
  })
}
