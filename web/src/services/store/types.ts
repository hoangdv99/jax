export type Tag = {
  id: number
  name: string
}

export type Store = {
  id: number | null
  url: string
  tags: Tag[]
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
