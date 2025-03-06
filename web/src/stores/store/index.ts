import { services } from '@/services'
import type { Tag, Store } from '@/services/store/types'
import { defineStore } from 'pinia'

export const useStoreStore = defineStore('storeStore', {
  state: () => ({
    listTag: <Tag[]>{},
    listStore: <Store[]>[],
  }),
  actions: {
    async getListTag() {
      const res = await services.store.getListTag()
      if (res.success) {
        this.listTag = res.data.tags
      }
      return res
    },
    async getListStore() {
      const res = await services.store.getListStore()
      if (res.success) {
        this.listStore = res.data.data
      }
      return res
    }
  },
})
