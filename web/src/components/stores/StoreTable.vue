<template>
  <div class="store-table">
    <DataTable :value="props.stores" class="table">
      <Column field="url" header="Store URL">
        <template #body="slotProps">
          <a :href="slotProps.data.url" target="_blank">{{ slotProps.data.url }}</a>
        </template>
      </Column>
      <Column field="platform" header="Platform">
        <template #body="slotProps">
          <div class="platform">
            <img :src="getPlatformIcon(slotProps.data.platform)" :alt="getPlatformData(slotProps.data.platform)?.name"
              class="icon">
            <span>{{ getPlatformData(slotProps.data.platform)?.name }}</span>
          </div>
        </template>
      </Column>
      <Column field="tags" header="Tags">
        <template #body="slotProps">
          <Tag v-for="tag in slotProps.data.tags" :key="tag.id" :value="tag.name" rounded severity="secondary"
            class="tag" />
        </template>
      </Column>
      <Column>
        <template #body="slotProps">
          <div class="action-wrapper">
            <Button icon="pi pi-pencil" variant="text" severity="secondary" @click="onUpdate(slotProps.data)"></Button>
            <Button icon="pi pi-trash" variant="text" severity="danger" @click="onDelete(slotProps.data.id)"></Button>
          </div>
        </template>
      </Column>
      <template #empty>
        <p class="no-data">No data</p>
      </template>
    </DataTable>
    <StoreDialog :originalStore="selectedStore" :visible="showStoreDialog" @hide="showStoreDialog = false" />
  </div>
</template>
<script setup lang="ts">
import { DataTable, Column, Button, Tag } from 'primevue'
import { useConfirm } from 'primevue/useconfirm'
import { useToast } from 'primevue/usetoast'
import { ref } from 'vue'
import StoreDialog from './StoreDialog.vue'
import type { Store } from '@/services/store/types'
import { PLATFORM } from '@/constants/store'
import shopifyIcon from '@/assets/icons/shopify.svg'
import woocommerceIcon from '@/assets/icons/woocommerce.svg'
import shopbaseIcon from '@/assets/icons/shopbase.svg'
import { services } from '@/services'
import { useStoreStore } from '@/stores/store'

defineOptions({
  name: 'StoreTable',
})

const storeStore = useStoreStore()

const platformIcons = {
  shopify: shopifyIcon,
  woocommerce: woocommerceIcon,
  shopbase: shopbaseIcon,
}

const getPlatformIcon = (platform: string) => {
  return platformIcons[platform as keyof typeof platformIcons] || ''
}

const confirm = useConfirm()
const toast = useToast()

const props = defineProps({
  stores: {
    type: Array,
    default: () => [],
  },
})

const showStoreDialog = ref(false)
const selectedStore = ref()

function onUpdate(store: Store) {
  selectedStore.value = store
  showStoreDialog.value = true
}

function onDelete(storeId: number) {
  confirm.require({
    header: 'Confirmation',
    message: 'Do you want to delete this store?',
    rejectLabel: 'Cancel',
    icon: 'pi pi-info-circle',
    rejectProps: {
      label: 'Cancel',
      severity: 'secondary',
      outlined: true,
    },
    acceptProps: {
      label: 'Delete',
      severity: 'danger',
    },
    accept: async () => {
      const res = await services.store.deleteStore(storeId)
      if (res.success) {
        storeStore.getListStore()
        toast.add({
          severity: 'success',
          summary: 'Success',
          detail: 'Added store',
        })
        selectedStore.value = null
      } else {
        toast.add({
          severity: 'error',
          summary: 'Error',
          detail: res.message?.name || 'Internal server error',
        })
      }
    },
  })
}

function getPlatformData(key: string) {
  return Object.values(PLATFORM).find((platform) => platform.key === key)
}

</script>
<style lang="scss" scoped>
.store-table {
  .tag {
    margin-right: 8px;
    font-weight: 400;
  }

  .no-data {
    text-align: center;
  }

  .platform {
    display: flex;
    align-items: center;
    gap: 6px;
  }

  .icon {
    width: 16px;
    height: 16px;
  }
}

.action-wrapper {
  display: flex;
  gap: 16px;
  justify-content: flex-end;
}
</style>
