import { defineStore } from 'pinia'
import { services } from '@/services'
import type {
  Tag,
  Store,
  Collection,
  Product,
  IGetListCollectionQueryParams,
  IGetCollectionProductsQueryParams,
  IGetProductsQueryParams,
} from '@/services/store/types'

export const useStoreStore = defineStore('storeStore', {
  state: () => ({
    listTag: <Tag[]>[],
    listStore: <Store[]>[],
    listCollection: <Collection[]>[],
    listProduct: <Product[]>[],
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
        if (!queryParams?.page || queryParams.page === 1) {
          const defaultCollection = {
            id: null,
            title: 'All',
            handle: '',
            productCount: null,
          }
          this.listCollection = [defaultCollection, ...res.data.data]
        } else {
          this.listCollection = [...this.listCollection, ...res.data.data]
        }
      } else {
        this.listCollection = []
      }
      return res
    },
    async getCollectionProducts(
      queryParams: IGetCollectionProductsQueryParams
    ) {
      const res = await services.store.getCollectionProducts(queryParams)
      if (res.success) {
        if (!queryParams.page || queryParams.page === 1) {
          this.listProduct = res.data.products
        } else {
          this.listProduct = [...this.listProduct, ...res.data.products]
        }
      }
      return res
    },
    async getProducts(queryParams: IGetProductsQueryParams) {
      const res = await services.store.getProducts(queryParams)
      if (res.success) {
        if (!queryParams.page || queryParams.page === 1) {
          this.listProduct = res.data.products
        } else {
          this.listProduct = [...this.listProduct, ...res.data.products]
        }
      }
      return res
    },
  },
})
