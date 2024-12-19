import { api } from '../api'
import type { ICreateTagPayload, IUpdateTagPayload } from './types'

function getListTag() {
  return api('get', '/v1/tags')
}

function createNewTag(payload: ICreateTagPayload) {
  return api('post', '/v1/tag', payload)
}

function updateTag(id: number, payload: IUpdateTagPayload) {
  return api('patch', `/v1/tag/${id}`, payload)
}

export default {
  getListTag,
  createNewTag,
  updateTag,
}
