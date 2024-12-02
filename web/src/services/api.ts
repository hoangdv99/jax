import { useAppStore } from '@/stores/app'
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

instance.interceptors.request.use(
  config => {
    const appStore = useAppStore()
    appStore.showLoading()
    const token = localStorage.getItem('authToken')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  error => {
    const appStore = useAppStore()
    appStore.hideLoading()
    return Promise.reject(error)
  }
)

instance.interceptors.response.use(
  response => {
    const appStore = useAppStore()
    appStore.hideLoading()
    return response
  },
  error => {
    const appStore = useAppStore()
    appStore.hideLoading()
    return Promise.reject(error)
  }
)
