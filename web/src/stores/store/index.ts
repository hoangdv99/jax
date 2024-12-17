import { services } from '@/services'
import type { Tag } from '@/services/store/types'
import { defineStore } from 'pinia'

export const useStoreStore = defineStore('storeStore', {
  state: () => ({
    listTag: <Tag[]>{},
  }),
  actions: {
    async getListTag() {
      const res = await services.store.getListTag()
      if (res.success) {
        this.listTag = res.data.tags
      }
      return res
    },
  },
})
