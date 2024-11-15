import { api } from '../api'
import type { InputLogin } from './types'

function login(payload: InputLogin) {
  return api('post', '/v1/tokens/authentication', payload)
}

export default {
  login,
}
