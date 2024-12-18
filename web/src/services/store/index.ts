import { api } from '../api'
import type { ICreateTagPayload } from './types'

function getListTag() {
  return api('get', '/v1/tags')
}

function createNewTag(payload: ICreateTagPayload) {
  return api('post', '/v1/tag', payload)
}

export default {
  getListTag,
  createNewTag,
}
