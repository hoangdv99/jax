import { services } from '@/services'
import type { InputSignup, InputSignin, User } from '@/services/account/types'
import { defineStore } from 'pinia'

export const useAccountStore = defineStore('accountStore', {
  state: () => ({
    currentUser: <User>{},
    listUser: <User[]>[],
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
    async getListUser() {
      const res = await services.account.getListUser()
      if (res.success) {
        this.listUser = res.data.users
      }
    },
    async getUserStores(userId: number) {
      const res = await services.account.getUserStores(userId)
      if (res.success) {
        const userIndex = this.listUser.findIndex(user => user.id === userId)
        if (userIndex !== -1) {
          this.listUser[userIndex] = {
            ...this.listUser[userIndex],
            stores: res.data.data,
          }
        }
      }
    },
  },
})
