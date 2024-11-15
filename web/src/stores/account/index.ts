import { defineStore } from 'pinia'
import type { InputLogin } from '@/services/account/types'
import { AccountService } from '@/services'

export const useAccountStore = defineStore('accountStore', () => {
  async function login(credential: InputLogin) {
    const res = await AccountService.login(credential)
  }
  return {
    login,
  }
})
