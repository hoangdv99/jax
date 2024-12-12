<template>
  <main>Homepage</main>
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
</script>
