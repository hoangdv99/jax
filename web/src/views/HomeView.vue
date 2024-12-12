<template>
  <main>Hello world!</main>
  <button @click="logout">Logout</button>
</template>
<script lang="ts" setup>
import { services } from '@/services'
import router from '@/router'

async function logout() {
  const res = await services.account.logout()
  if (res.success) {
    localStorage.removeItem('authToken')
    localStorage.removeItem('authTokenExpiry')
    router.push({ name: 'Login' })
  }
}
</script>
