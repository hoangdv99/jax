import { services } from '@/services'
import type { InputSignup } from '@/services/account/types'
import { defineStore } from 'pinia'

export const useAccountStore = defineStore('accountStore', {
  actions: {
    async signup(payload: InputSignup) {
      const res = await services.account.signup(payload)
      return res
    },
  },
})
