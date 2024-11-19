import { defineStore } from 'pinia'

export const useAppStore = defineStore('appStore', {
  state: () => ({
    isLoading: false,
  }),
  actions: {
    showLoading() {
      this.isLoading = true
    },
    hideLoading() {
      this.isLoading = false
    },
  },
})
