import axios, { type AxiosResponse } from 'axios'

export interface ApiResponse<T = any> extends AxiosResponse<T> {
  success: boolean
}

const instance = axios.create({
  baseURL: import.meta.env.VITE_API_ENDPOINT,
})

export const api = async (
  methodType: string,
  endpoint: string,
  params: any
) => {
  try {
    const res = await (instance as any)[methodType](endpoint, params)
    return {
      ...res,
      success: true,
    }
  } catch (err: any) {
    return {
      message: err.response?.data?.error,
      success: false,
    }
  }
}
