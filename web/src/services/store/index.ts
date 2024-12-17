import { api } from '../api'

function getListTag() {
  return api('get', '/v1/tags')
}

export default {
  getListTag,
}
