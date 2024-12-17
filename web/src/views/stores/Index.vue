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
import { Button, useToast } from 'primevue'
import { onMounted, ref } from 'vue'
import StoreDialog from '@/components/stores/StoreDialog.vue'
import { useStoreStore } from '@/stores/store'

defineOptions({
  name: 'ListStore',
})
const toast = useToast()
const storeStore = useStoreStore()

const showStoreDialog = ref(false)

onMounted(async () => {
  const res = await storeStore.getListTag()
  if (!res.success) {
    toast.add({
      severity: 'error',
      summary: 'Error',
      detail: 'Internal server error',
    })
  }
})
</script>
