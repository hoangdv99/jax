import { api } from '../api'
import type { InputSignup } from './types'

function signup(payload: InputSignup) {
  return api('post', '/v1/users', payload)
}

export default {
  signup,
}
