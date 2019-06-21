import Cookies from 'js-cookie'
import { JSEncrypt } from 'jsencrypt'
import g from '@/config/config.js'

const TokenKey = 'X-Token'
// const ProjectName = 'Project-Name'
// 40分钟
const inFifteenMinutes = new Date(new Date().getTime() + g.expirament * 60 * 1000)
export function getToken() {
  return Cookies.get(TokenKey)
}

export function setToken(token) {
  return Cookies.set(TokenKey, encrypt(token), {
    expires: inFifteenMinutes
  })
}

export function setTimeout() {
  return Cookies.set(TokenKey, getToken(TokenKey), {
    expires: inFifteenMinutes
  })
}

export function removeToken() {
  return Cookies.remove(TokenKey)
}

export function encrypt(key) {
  const jes = new JSEncrypt()
  jes.setPublicKey(g.pubkey)
  return jes.encrypt(key)
}
