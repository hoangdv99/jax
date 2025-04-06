export type Tag = {
  id: number
  name: string
}

export type Store = {
  id: number | null
  url: string
  tags: Tag[]
}

export type Collection = {
  id: number | null
  title: string
  handle: string
  productCount: number
}

export type Product = {
  id: number
  storeUrl: string
  price: string
  images: string[]
  createdDate: string
  sourcePostUrl: string
  featureMedia: number
}

export interface ICreateTagPayload {
  name: string
}

export interface IUpdateTagPayload {
  name: string
}

export interface ICreateStorePayload {
  url: string
  tagIds: number[]
}

export interface IUpdateStorePayload {
  tagIds: number[]
}

export interface IGetListCollectionQueryParams {
  limit: number
  page: number
}

export interface IGetCollectionProductsQueryParams {
  storeId: number
  collectionId: number
  handle: string
  page?: number
  limit?: number
}

export interface IGetProductsQueryParams {
  storeIds: number[]
  page?: number
  limit?: number
}
