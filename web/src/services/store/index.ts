import { api } from '../api'
import type {
  ICreateTagPayload,
  IUpdateTagPayload,
  ICreateStorePayload,
  IUpdateStorePayload,
} from './types'

function getListTag() {
  return api('get', '/v1/tags')
}

function createNewTag(payload: ICreateTagPayload) {
  return api('post', '/v1/tag', payload)
}

function updateTag(id: number, payload: IUpdateTagPayload) {
  return api('patch', `/v1/tag/${id}`, payload)
}

function deleteTag(id: number) {
  return api('delete', `/v1/tag/${id}`)
}

function addStore(payload: ICreateStorePayload) {
  return api('post', '/v1/store', payload)
}

function getListStore() {
  return api('get', '/v1/stores')
}

function updateStore(id: number, payload: IUpdateStorePayload) {
  return api('put', `/v1/store/${id}`, payload)
}

function deleteStore(id: number) {
  return api('delete', `/v1/store/${id}`)
}

export default {
  getListTag,
  createNewTag,
  updateTag,
  deleteTag,
  addStore,
  getListStore,
  updateStore,
  deleteStore,
}
