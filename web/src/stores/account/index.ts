import { services } from '@/services'
import type { InputSignup, InputSignin, User } from '@/services/account/types'
import { defineStore } from 'pinia'

export const useAccountStore = defineStore('accountStore', {
  state: () => ({
    currentUser: <User>{},
  }),
  actions: {
    async signup(payload: InputSignup) {
      const res = await services.account.signup(payload)
      return res
    },
    async signin(payload: InputSignin) {
      const res = await services.account.signin(payload)
      return res
    },
    setCurrentUser(user: User) {
      this.currentUser = user
    },
  },
})
