<template>
  <div class="sidebar-container">
    <div class="wrapper">
      <img src="../../public/images/logoG.png" alt="" class="logo" />
      <div class="user-menu">
        <div
          class="menu"
          :class="{ '-active': isActive('/') }"
          @click="navigate('/')"
        >
          <i class="pi pi-home"></i>
        </div>
        <div
          class="menu"
          :class="{ '-active': isActive('/stores') }"
          @click="navigate('/stores')"
        >
          <i class="pi pi-list"></i>
        </div>
      </div>
      <Divider />
      <div class="admin-menu">
        <div
          class="menu"
          :class="{ '-active': isActive('/admin/users') }"
          @click="navigate('/admin/users')"
        >
          <i class="pi pi-users"></i>
        </div>
      </div>
    </div>
    <div class="account-menu">
      <div class="menu" @click="logout">
        <i class="pi pi-sign-out"></i>
      </div>
    </div>
  </div>
</template>
<script setup lang="ts">
import router from '@/router'
import { services } from '@/services';
import { Divider } from 'primevue'

defineOptions({
  name: 'Sidebar',
})

function isActive(path: string) {
  return path === router.currentRoute.value.path
}

function navigate(path: string) {
  router.push(path)
}

async function logout() {
  const res = await services.account.logout()
  if (res.success) {
    localStorage.removeItem('authToken')
    localStorage.removeItem('authTokenExpiry')
    router.push({ name: 'Login' })
  }
}
</script>
<style lang="scss" scoped>
.sidebar-container {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  padding: 12px;
  background-color: #fff;
  height: 100vh;
  > .wrapper > .logo {
    margin-bottom: 32px;
    width: 36px;
    height: 36px;
  }
}

.user-menu,
.admin-menu,
.account-menu {
  > .menu {
    display: flex;
    justify-content: center;
    border-radius: 4px;
    padding: 12px 0;
    width: 100%;
    &:hover,
    &.-active {
      color: #10b981;
      background-color: #eaf0ee;
      cursor: pointer;
    }
  }
  > .menu > .pi {
    font-size: 18px;
  }
}
</style>
