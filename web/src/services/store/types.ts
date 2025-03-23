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
  productCount: number,
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
