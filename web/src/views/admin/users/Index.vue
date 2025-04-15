<template>
  <div class="list-user-page">
    <DataTable
      v-model:expandedRows="expandedRows"
      :value="accountStore.listUser"
      data-key="id"
      @row-expand="onRowExpand"
    >
      <Column expander></Column>
      <Column field="id" header="ID"></Column>
      <Column field="email" header="Email"></Column>
      <Column field="status" header="Status">
        <template #body="slotProps">
          <Tag
            :value="getStatusLabel(slotProps.data.status)"
            rounded
            :severity="getStatusSeverity(slotProps.data.status)"
            class="tag"
          />
        </template>
      </Column>
      <template #expansion="slotProps">
        <div v-if="!slotProps.data.stores || !slotProps.data.stores.length">
          No data
        </div>
        <DataTable v-else :value="slotProps.data.stores" class="store-table">
          <Column field="url" header="Store Url">
            <template #body="slotProps">
              <a :href="slotProps.data.url" target="_blank">{{
                slotProps.data.url
              }}</a>
            </template>
          </Column>
          <Column field="platform" header="Platform">
            <template #body="slotProps">
              <div class="platform">
                <img
                  :src="getPlatformIcon(slotProps.data.platform)"
                  :alt="getPlatformData(slotProps.data.platform)?.name"
                  class="icon"
                />
                <span>{{
                  getPlatformData(slotProps.data.platform)?.name
                }}</span>
              </div>
            </template>
          </Column>
          <Column field="tags" header="Tags">
            <template #body="slotProps">
              <Tag
                v-for="tag in slotProps.data.tags"
                :key="tag.id"
                :value="tag.name"
                rounded
                severity="secondary"
                class="tag"
              />
            </template>
          </Column>
        </DataTable>
      </template>
    </DataTable>
  </div>
</template>
<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { DataTable, Column, Tag, type DataTableRowExpandEvent } from 'primevue'
import { useAccountStore } from '@/stores/account'
import { PLATFORM } from '@/constants/store'
import { USER_STATUS } from '@/constants/user'
import shopifyIcon from '@/assets/icons/shopify.svg'
import woocommerceIcon from '@/assets/icons/woocommerce.svg'
import shopbaseIcon from '@/assets/icons/shopbase.svg'

defineOptions({
  name: 'AdminListUser',
})

const accountStore = useAccountStore()

const expandedRows = ref({})

onMounted(() => {
  accountStore.getListUser()
})

const platformIcons = {
  shopify: shopifyIcon,
  woocommerce: woocommerceIcon,
  shopbase: shopbaseIcon,
}
const getPlatformIcon = (platform: string) => {
  return platformIcons[platform as keyof typeof platformIcons] || ''
}

const getPlatformData = (key: string) => {
  return Object.values(PLATFORM).find(platform => platform.key === key)
}

const onRowExpand = (event: DataTableRowExpandEvent) => {
  accountStore.getUserStores(event.data.id)
}

const getStatusLabel = (value: number) => {
  const status = Object.values(USER_STATUS).find(
    status => status.value === value
  )
  return status ? status.label : ''
}

const getStatusSeverity = (value: number) => {
  if (value === USER_STATUS.ACTIVE.value) {
    return 'primary'
  } else {
    return 'secondary'
  }
}
</script>
<style lang="scss" scoped>
.store-table {
  .tag {
    margin-right: 8px;
    font-weight: 400;
  }

  .platform {
    display: flex;
    align-items: center;
    gap: 6px;
  }

  .icon {
    width: 32px;
    height: 16px;
  }
}
</style>
