<template>
  <main>Hello world!</main>
  <button @click="logout">Logout</button>
</template>
<script lang="ts" setup>
import { onBeforeMount } from 'vue'
import { services } from '@/services'
import router from '@/router'
import { useAccountStore } from '@/stores/account'
import { USER_STATUS } from '@/constants/user'

onBeforeMount(async () => {
  const token = localStorage.getItem('authToken')
  if (token) {
    const res = await services.account.getCurrentUser(token)
    if (res.success) {
      const accountStore = useAccountStore()
      accountStore.setCurrentUser(res.data.currentUser)
      if (accountStore.currentUser.status === USER_STATUS.WAITING_ACTIVATION) {
        router.push({ name: 'UserActivation' })
      } else if (
        accountStore.currentUser.status === USER_STATUS.WAITING_APPROVAL
      ) {
        router.push({ name: 'WaitlistNotification' })
      }
    }
  }
})

async function logout() {
  const res = await services.account.logout()
  if (res.success) {
    localStorage.removeItem('authToken')
    localStorage.removeItem('authTokenExpiry')
    router.push({ name: 'Login' })
  }
}
</script>
