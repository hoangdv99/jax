import { api } from '../api'
import type { InputSignin, InputSignup, UserActivationPayload } from './types'

function signup(payload: InputSignup) {
  return api('post', '/users', payload)
}

function signin(payload: InputSignin) {
  return api('post', '/login', payload)
}

function resendActivationToken(payload: UserActivationPayload) {
  return api('post', '/tokens/activation', payload)
}

function logout() {
  return api('post', '/logout')
}

function getCurrentUser(token: string) {
  return api('get', `/v1/current-user`, { token })
}

function getListUser() {
  return api('get', '/v1/users')
}

function getUserStores(userId: number) {
  return api('get', `/v1/user/${userId}/stores`)
}

export default {
  signup,
  signin,
  resendActivationToken,
  logout,
  getCurrentUser,
  getListUser,
  getUserStores,
}
