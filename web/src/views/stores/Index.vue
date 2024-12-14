<template>
  <div class="list-store-page">
    <Button
      icon="pi pi-plus"
      label="Add new store"
      @click="showStoreDialog = true"
    />
    <StoreDialog :visible="showStoreDialog" @hide="showStoreDialog = false" />
  </div>
</template>
<script setup lang="ts">
import { Button } from 'primevue'
import { ref } from 'vue'
import StoreDialog from '@/components/stores/StoreDialog.vue'

defineOptions({
  name: 'ListStore',
})

const showStoreDialog = ref(false)
</script>
