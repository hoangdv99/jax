import { services } from '@/services'
import type {
  Tag,
  Store,
  IGetListCollectionQueryParams,
  Collection,
} from '@/services/store/types'
import { defineStore } from 'pinia'

export const useStoreStore = defineStore('storeStore', {
  state: () => ({
    listTag: <Tag[]>[],
    listStore: <Store[]>[],
    listCollection: <Collection[]>[],
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
    },
    async getListCollection(
      storeId: number,
      queryParams?: IGetListCollectionQueryParams
    ) {
      const res = await services.store.getListCollection(storeId, queryParams)
      if (res.success) {
        const defaultCollection = {
          id: null,
          title: 'All',
          handle: '',
          productCount: null,
        }
        this.listCollection = [
          defaultCollection,
          ...res.data.data,
        ]
      } else {
        this.listCollection = []
      }
      return res
    },
  },
})
