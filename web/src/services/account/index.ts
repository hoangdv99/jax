import { api } from '../api'
import type { InputSignin, InputSignup } from './types'

function signup(payload: InputSignup) {
  return api('post', '/v1/users', payload)
}

function signin(payload: InputSignin) {
  return api('post', '/v1/tokens/authentication', payload)
}


export default {
  signup,
  signin,
}
