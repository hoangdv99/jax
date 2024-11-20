<template>
  <div v-if="appStore.isLoading" class="spinner-overlay">
    <ProgressSpinner strokeWidth="4" />
  </div>
  <Toast />
  <RouterView />
</template>
<script lang="ts" setup>
import ProgressSpinner from 'primevue/progressspinner'
import { Toast } from 'primevue'
import { useAppStore } from './stores/app'

const appStore = useAppStore()
</script>
<style lang="scss" scoped>
.spinner-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.3);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}
</style>
