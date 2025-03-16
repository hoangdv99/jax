<template>
  <main class="home-page">
    <Filters :stores="storeStore.listStore" :tags="storeStore.listTag" />
  </main>
</template>
<script lang="ts" setup>
import { onBeforeMount, onMounted } from 'vue'
import Filters from '@/components/home/Filters.vue'
import { services } from '../../services'
import router from '../../router'
import { useAccountStore } from '../../stores/account'
import { useStoreStore } from '@/stores/store'
import { USER_STATUS } from '../../constants/user'

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

const storeStore = useStoreStore()

onMounted(async () => {
  storeStore.getListTag()
  storeStore.getListStore()
})
</script>
