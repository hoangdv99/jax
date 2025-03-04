<template>
  <div class="list-store-page">
    <Button
      icon="pi pi-plus"
      label="Add new store"
      class="button"
      @click="showStoreDialog = true"
    />
    <StoreDialog :visible="showStoreDialog" @hide="showStoreDialog = false" />
    <StoreTable class="table" />
  </div>
</template>
<script setup lang="ts">
import { Button } from 'primevue'
import { onMounted, ref } from 'vue'
import StoreDialog from '@/components/stores/StoreDialog.vue'
import StoreTable from '@/components/stores/StoreTable.vue'
import { useStoreStore } from '@/stores/store'

defineOptions({
  name: 'ListStore',
})
const storeStore = useStoreStore()

const showStoreDialog = ref(false)

onMounted(async () => {
  await storeStore.getListTag()
  await storeStore.getListStore()
})
</script>
<style lang="scss" scoped>
.list-store-page {
  > .button {
    margin-bottom: 16px;
  }
  > .table {
    width: 100%;
  }
}
</style>
