import { api } from '../api'
import type { InputSignin, InputSignup, UserActivationPayload } from './types'

function signup(payload: InputSignup) {
  return api('post', '/users', payload)
}

function signin(payload: InputSignin) {
  return api('post', '/tokens/authentication', payload)
}

function resendActivationToken(payload: UserActivationPayload) {
  return api('post', '/tokens/activation', payload)
}

function logout() {
  return api('post', '/logout', {})
}

export default {
  signup,
  signin,
  resendActivationToken,
  logout,
}
