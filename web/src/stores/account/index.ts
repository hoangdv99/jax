import { services } from '@/services'
import type { InputSignup, InputSignin } from '@/services/account/types'
import { defineStore } from 'pinia'

export const useAccountStore = defineStore('accountStore', {
  actions: {
    async signup(payload: InputSignup) {
      const res = await services.account.signup(payload)
      return res
    },
    async signin(payload: InputSignin) {
      const res = await services.account.signin(payload)
      return res
    }
  },
})
