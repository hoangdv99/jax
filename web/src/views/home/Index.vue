<template>
  <main class="home-page">
    <Filters
      :selected-store="selectedStore"
      :selected-tags="selectedTags"
      :selected-collection="selectedCollection"
      :stores="storeStore.listStore"
      :collection-page="collectionPage"
      :tags="storeStore.listTag"
      :collections="storeStore.listCollection"
      :all-collections-fetched="allCollectionsFetched"
      class="filters"
      @get-list-collection="storeStore.getListCollection"
      @update-collection-page="updateCollectionPage"
      @update-selected-store="selectedStore = $event"
      @update-selected-collection="selectedCollection = $event"
    />
    <ProductList
      :products="storeStore.listProduct"
      :platform="selectedStore?.platform"
      class="products"
      @load-more-products="onLoadMoreProducts"
    />
  </main>
</template>
<script lang="ts" setup>
import { onBeforeMount, onMounted, ref, watch } from 'vue'
import Filters from '@/components/home/Filters.vue'
import ProductList from '@/components/home/ProductList.vue'
import { services } from '@/services'
import router from '@/router'
import { useAccountStore } from '@/stores/account'
import { useStoreStore } from '@/stores/store'
import { USER_STATUS } from '@/constants/user'
import type { Tag, Store, Collection } from '@/services/store/types'
import { COLLECTION_PAGE_LIMIT, PRODUCT_PAGE_LIMIT } from '@/constants/store'

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

const selectedStore = ref<Store | undefined>(
  (storeStore?.listStore[0] as Store) || undefined
)
const selectedTags = ref<Tag[]>([])
const selectedCollection = ref<Collection | undefined>(
  (storeStore.listCollection[0] as Collection) || undefined
)
const collectionPage = ref(1)
const productPage = ref(1)
const allProductsFetched = ref(false)
const allCollectionsFetched = ref(false)

watch(
  () => storeStore.listStore,
  newStores => {
    if (newStores.length) {
      selectedStore.value = newStores[0] as Store
    } else {
      selectedStore.value = undefined
    }
  },
  { immediate: true }
)
watch(
  () => storeStore.listCollection,
  newCollections => {
    if (newCollections.length && !selectedCollection.value) {
      selectedCollection.value = newCollections[0] as Collection
    }
  },
  { immediate: true }
)
watch(selectedStore, async (val: Store | undefined) => {
  if (val && val.id) {
    collectionPage.value = 1
    productPage.value = 1
    allProductsFetched.value = false
    allCollectionsFetched.value = false
    const res = await storeStore.getListCollection(val.id, {
      page: 1,
      limit: COLLECTION_PAGE_LIMIT,
    })
    selectedCollection.value = storeStore.listCollection[0]
    if (res.success && res.data.data.length < COLLECTION_PAGE_LIMIT) {
      allCollectionsFetched.value = true
    }
  }
})
watch(selectedCollection, async (val: Collection | undefined) => {
  productPage.value = 1
  allProductsFetched.value = false
  if (val?.handle && selectedStore.value?.id) {
    const res = await storeStore.getCollectionProducts({
      storeId: selectedStore.value.id,
      collectionId: val.id,
      handle: val.handle,
      page: 1,
      limit: PRODUCT_PAGE_LIMIT,
    })
    if (res.success && res.data.products.length < PRODUCT_PAGE_LIMIT) {
      allProductsFetched.value = true
    }
  } else if (val?.title === 'All' && selectedStore.value) {
    const res = await storeStore.getProducts({
      storeIds: selectedStore.value.id,
      page: 1,
      limit: PRODUCT_PAGE_LIMIT,
    })
    if (res.success && res.data.products.length < PRODUCT_PAGE_LIMIT) {
      allProductsFetched.value = true
    }
  }
})

onMounted(async () => {
  storeStore.getListTag()
  storeStore.getListStore()
})

const onLoadMoreProducts = async () => {
  if (allProductsFetched.value) return
  if (selectedCollection.value?.handle && selectedStore.value?.id) {
    const page = productPage.value + 1
    const res = await storeStore.getCollectionProducts({
      storeId: selectedStore.value.id,
      collectionId: selectedCollection.value.id,
      handle: selectedCollection.value.handle,
      page,
      limit: PRODUCT_PAGE_LIMIT,
    })
    if (res.success && res.data.products.length < PRODUCT_PAGE_LIMIT) {
      allProductsFetched.value = true
    }
  } else if (selectedStore.value) {
    const page = productPage.value + 1
    const res = await storeStore.getProducts({
      storeIds: selectedStore.value.id,
      page,
      limit: PRODUCT_PAGE_LIMIT,
    })
    if (res.success && res.data.products.length < PRODUCT_PAGE_LIMIT) {
      allProductsFetched.value = true
    }
  }
}
const updateCollectionPage = (page: number) => {
  collectionPage.value = page
}
</script>
<style lang="scss" scoped>
.home-page {
  > .filters {
    position: sticky;
    top: 16px;
    z-index: 1;
    margin-bottom: 16px;
  }
}
</style>
